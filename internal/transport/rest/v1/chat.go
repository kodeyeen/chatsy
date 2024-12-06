package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kodeyeen/chatsy/v1"
)

type chatService interface {
	GetByUserID(ctx context.Context, userID int, limit, offset int) (*chatsy.PageResponse[*chatsy.GetChatResponse], error)
}

type ChatController struct {
	svc chatService
}

func NewChatController(svc chatService) *ChatController {
	return &ChatController{
		svc: svc,
	}
}

func (c *ChatController) GetMine(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Set("Content-Type", "application/json;charset=utf-8")
	headers.Set("X-Content-Type-Options", "nosniff")

	ctx := r.Context()
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := r.URL.Query()

	limit, err := strconv.ParseUint(query.Get("limit"), 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(chatsy.ErrorResponse{
			Message: "Invalid limit specified",
		})
		return
	}

	offset, err := strconv.ParseUint(query.Get("offset"), 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(chatsy.ErrorResponse{
			Message: "Invalid offset specified",
		})
		return
	}

	resp, err := c.svc.GetByUserID(ctx, userID, int(limit), int(offset))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(chatsy.ErrorResponse{
			Message: "Something went wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *ChatController) Login(w http.ResponseWriter, r *http.Request) {

}
