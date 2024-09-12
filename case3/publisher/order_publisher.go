package publisher

import (
	"context"
	"encoding/json"

	redislib "github.com/elangreza14/qbit/case3/lib/redis"
	"github.com/google/uuid"
)

const OrderChannel = "order-channel"

type orderPublisher struct {
	redisDB redislib.IRedis
}

func NewOrderPublisher(redisDB redislib.IRedis) *orderPublisher {
	return &orderPublisher{
		redisDB: redisDB,
	}
}

type PublishOrderPayload struct {
	OrderID uuid.UUID
	UserID  uuid.UUID
	CartIDs []int
}

func (op *orderPublisher) PublishOrder(ctx context.Context, userID, orderID uuid.UUID, cartIDs []int) error {
	message, err := json.Marshal(PublishOrderPayload{
		OrderID: orderID,
		UserID:  userID,
		CartIDs: cartIDs,
	})

	if err != nil {
		return err
	}

	return op.redisDB.Publish(ctx, OrderChannel, message).Err()
}
