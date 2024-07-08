package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodeyeen/chatsy/internal/auth"
	"github.com/kodeyeen/chatsy/internal/chat"
	"github.com/kodeyeen/chatsy/internal/config"
	"github.com/kodeyeen/chatsy/internal/database"
	"github.com/kodeyeen/chatsy/internal/http"
	"github.com/kodeyeen/chatsy/internal/message"
	"github.com/kodeyeen/chatsy/internal/postgres"
	"github.com/kodeyeen/chatsy/internal/websocket"
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

	userRepo := postgres.NewUserRepository(dbpool)

	authSvc := auth.NewDefaultService(cfg.Secret, cfg.TokenTTL, cfg.TicketTTL, userRepo)
	_ = http.NewAuthController(authSvc)

	chatRepo := postgres.NewChatRepository(dbpool)
	chatSvc := chat.NewDefaultService(chatRepo)

	msgRepo := postgres.NewMessageRepository(dbpool)
	msgSvc := message.NewDefaultService(msgRepo, userRepo)

	eventHandler := websocket.NewEventHandler(chatSvc, msgSvc)
	_ = websocket.NewManager(eventHandler)

	_ = http.NewCheckJWTMiddleware(cfg.Secret)
	_ = http.NewCheckTicketMiddleware(cfg.Secret)

	// serveMux := http.NewServeMux()
	// serveMux.HandleFunc("/auth/register", authClr.Register)
	// serveMux.HandleFunc("/auth/login", authClr.Login)
	// serveMux.HandleFunc("/auth/logout", authClr.Logout)
	// serveMux.HandleFunc("/auth/me", checkJWT(authClr.Me))
	// serveMux.HandleFunc("/auth/ticket", checkJWT(authClr.CreateTicket))
	// serveMux.HandleFunc("/ws", checkTicket(wsMng.ServeHTTP))

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   cfg.Cors.AllowedOrigins,
	// 	AllowCredentials: true,
	// 	Debug:            cfg.Cors.Debug,
	// })
	// handler := c.Handler(serveMux)

	// server := &http.Server{
	// 	Addr:         cfg.HTTPServer.Address,
	// 	Handler:      handler,
	// 	ReadTimeout:  cfg.HTTPServer.Timeout,
	// 	WriteTimeout: cfg.HTTPServer.Timeout,
	// 	IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	// }
	// log.Fatal(server.ListenAndServe())
}
