package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodeyeen/chatsy/internal/auth"
	"github.com/kodeyeen/chatsy/internal/chat"
	"github.com/kodeyeen/chatsy/internal/config"
	"github.com/kodeyeen/chatsy/internal/ticket"
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

	ticketRepo := ticket.NewInmemoryRepository(ctx, cfg.TicketTTL)

	userRepo := user.NewPostgresRepository(dbpool)

	authSvc := auth.NewDefaultService(cfg.Secret, cfg.TokenTTL, userRepo, ticketRepo)
	authHlr := auth.NewHandler(authSvc)

	chatRepo := chat.NewPostgresRepository(dbpool)
	chatSvc := chat.NewDefaultService(chatRepo)
	chatHlr := chat.NewWsHandler(chatSvc)

	wsConnMng := websocket.NewConnManager(ticketRepo, userRepo)
	wsConnMng.Handle(websocket.EventFetchChats, chatHlr.FetchChats)

	checkJWT := auth.NewCheckJWTMiddleware(cfg.Secret)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/auth/register", authHlr.Register)
	serveMux.HandleFunc("/auth/login", authHlr.Login)
	serveMux.HandleFunc("/auth/logout", authHlr.Logout)
	serveMux.HandleFunc("/auth/me", checkJWT(authHlr.Me))
	serveMux.HandleFunc("/auth/ticket", checkJWT(authHlr.CreateTicket))
	serveMux.HandleFunc("/ws", wsConnMng.ServeWS)

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
