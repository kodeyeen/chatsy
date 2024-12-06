package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kodeyeen/chatsy/api/v1"
)

type messageService interface {
	GetByChatID(ctx context.Context, chatID int, limit, offset int) (*api.PageResponse[*api.GetMessageResponse], error)
}

type MessageController struct {
	svc messageService
}

func NewMessageController(svc messageService) *MessageController {
	return &MessageController{
		svc: svc,
	}
}

func (c *MessageController) GetByChatID(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json;charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()

	chatID, err := strconv.Atoi(r.PathValue("chatId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.ErrorResponse{
			Message: "Invalid chatId specified",
		})
		return
	}

	query := r.URL.Query()

	limit, err := strconv.ParseUint(query.Get("limit"), 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.ErrorResponse{
			Message: "Invalid limit specified",
		})
		return
	}

	offset, err := strconv.ParseUint(query.Get("offset"), 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.ErrorResponse{
			Message: "Invalid offset specified",
		})
		return
	}

	resp, err := c.svc.GetByChatID(ctx, chatID, int(limit), int(offset))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(api.ErrorResponse{
			Message: "Something went wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *MessageController) Login(w http.ResponseWriter, r *http.Request) {

}
