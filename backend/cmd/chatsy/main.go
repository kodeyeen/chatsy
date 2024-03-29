package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodeyeen/chatsy/internal/auth"
	"github.com/kodeyeen/chatsy/internal/chat"
	"github.com/kodeyeen/chatsy/internal/config"
	"github.com/kodeyeen/chatsy/internal/message"
	"github.com/kodeyeen/chatsy/internal/user"
	"github.com/kodeyeen/chatsy/internal/websocket"
	"github.com/kodeyeen/chatsy/pkg/database"
	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()

	connString := database.NewConnString(
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

	userRepo := user.NewPostgresRepository(dbpool)

	authSvc := auth.NewDefaultService(cfg.Secret, cfg.TokenTTL, cfg.TicketTTL, userRepo)
	authHlr := auth.NewHTTPHandler(authSvc)

	chatRepo := chat.NewPostgresRepository(dbpool)
	chatSvc := chat.NewDefaultService(chatRepo)

	msgRepo := message.NewPostgresRepository(dbpool)
	msgSvc := message.NewDefaultService(msgRepo, userRepo)

	eventHandler := websocket.NewEventHandler(chatSvc, msgSvc)
	wsMng := websocket.NewManager(eventHandler)

	checkJWT := auth.NewCheckJWTMiddleware(cfg.Secret)
	checkTicket := auth.NewCheckTicketMiddleware(cfg.Secret)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/auth/register", authHlr.Register)
	serveMux.HandleFunc("/auth/login", authHlr.Login)
	serveMux.HandleFunc("/auth/logout", authHlr.Logout)
	serveMux.HandleFunc("/auth/me", checkJWT(authHlr.Me))
	serveMux.HandleFunc("/auth/ticket", checkJWT(authHlr.CreateTicket))
	serveMux.HandleFunc("/ws", checkTicket(wsMng.ServeWS))

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
