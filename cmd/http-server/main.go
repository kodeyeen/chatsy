package main

import (
	"context"
	"log"

	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodeyeen/chatsy/internal/auth"
	"github.com/kodeyeen/chatsy/internal/chat"
	"github.com/kodeyeen/chatsy/internal/config"
	"github.com/kodeyeen/chatsy/internal/delivery/http/httpmw"
	httpdelivery "github.com/kodeyeen/chatsy/internal/delivery/http/v1"
	websocketdelivery "github.com/kodeyeen/chatsy/internal/delivery/websocket"
	"github.com/kodeyeen/chatsy/internal/message"
	"github.com/kodeyeen/chatsy/internal/persistence"
	postgrespersistence "github.com/kodeyeen/chatsy/internal/persistence/postgres"
	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	connString := persistence.NewConnString(
		"postgres",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Name,
	)

	dbpool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal("failed to create db pool")
	}
	defer dbpool.Close()

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatalf("failed to ping db: %s", err.Error())
	}

	userRepo := postgrespersistence.NewUserRepository(dbpool)

	authSvc := auth.NewService(cfg.Secret, cfg.TokenTTL, cfg.TicketTTL, userRepo)
	authClr := httpdelivery.NewAuthController(authSvc)

	chatRepo := postgrespersistence.NewChatRepository(dbpool)
	chatSvc := chat.NewService(chatRepo)
	chatClr := httpdelivery.NewChatController(chatSvc)

	msgRepo := postgrespersistence.NewMessageRepository(dbpool)
	msgSvc := message.NewService(msgRepo, userRepo)
	msgClr := httpdelivery.NewMessageController(msgSvc)

	eventHlr := websocketdelivery.NewEventHandler(chatSvc, msgSvc)
	wsMgr := websocketdelivery.NewManager(eventHlr)

	checkJWT := httpmw.NewJWTChecker(cfg.Secret)
	checkTicket := httpmw.NewTicketChecker(cfg.Secret)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("POST /auth/register", authClr.Register)
	serveMux.HandleFunc("POST /auth/login", authClr.Login)
	serveMux.HandleFunc("POST /auth/logout", authClr.Logout)
	serveMux.HandleFunc("GET /auth/me", checkJWT(authClr.Me))
	serveMux.HandleFunc("POST /auth/ticket", checkJWT(authClr.CreateTicket))

	serveMux.HandleFunc("GET /ws", checkTicket(wsMgr.ServeHTTP))

	serveMux.HandleFunc("GET /chats/mine", checkJWT(chatClr.GetMine))
	serveMux.HandleFunc("GET /chats/{chatId}/messages", checkJWT(msgClr.GetByChatID))

	c := cors.New(cors.Options{
		AllowedOrigins:   cfg.Cors.AllowedOrigins,
		AllowCredentials: true,
		Debug:            cfg.Cors.Debug,
	})
	handler := c.Handler(serveMux)

	server := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      handler,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	log.Fatal(server.ListenAndServe())
}
