package ticket

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type inmemoryRepository struct {
	tickets map[string]*Ticket
}

func NewInmemoryRepository(ctx context.Context, retentionPeriod time.Duration) *inmemoryRepository {
	r := &inmemoryRepository{
		tickets: make(map[string]*Ticket),
	}

	go r.checkTickets(ctx, retentionPeriod)

	return r
}

func (r *inmemoryRepository) Add(ctx context.Context, t *Ticket) string {
	t.ID = uuid.NewString()
	t.CreatedAt = time.Now()

	r.tickets[t.ID] = t

	return t.ID
}

func (r *inmemoryRepository) FindByID(ctx context.Context, id string) (*Ticket, error) {
	t, ok := r.tickets[id]
	if !ok {
		return &Ticket{}, ErrNotFound
	}

	delete(r.tickets, id)

	return t, nil
}

func (r *inmemoryRepository) checkTickets(ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			for _, t := range r.tickets {
				if t.CreatedAt.Add(retentionPeriod).Before(time.Now()) {
					delete(r.tickets, t.ID)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
