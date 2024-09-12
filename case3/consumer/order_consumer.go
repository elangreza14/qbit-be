package consumer

import (
	"context"
	"encoding/json"

	"github.com/elangreza14/qbit/case3/dto"
	redislib "github.com/elangreza14/qbit/case3/lib/redis"
	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

type (
	orderService interface {
		UpdateOrder(ctx context.Context, req dto.UpdateOrder) error
	}

	OrderConsumer struct {
		orderService orderService
		redisDB      redislib.IRedis
		logger       *zap.Logger
	}
)

func NewOrderConsumer(orderService orderService, redisDB redislib.IRedis, logger *zap.Logger) *OrderConsumer {
	return &OrderConsumer{
		orderService: orderService,
		redisDB:      redisDB,
		logger:       logger,
	}
}

func (cc *OrderConsumer) ConsumeOrder(ctx context.Context, msg *redis.Message, channel string) {
	if msg == nil {
		return
	}

	if msg.Channel != channel {
		return
	}

	order := dto.UpdateOrder{}
	err := json.Unmarshal([]byte(msg.Payload), &order)
	if err != nil {
		cc.logger.Error("error unmarshaling order", zap.Error(err))
		return
	}

	err = cc.orderService.UpdateOrder(ctx, order)
	if err != nil {
		cc.logger.Error("error update order", zap.Error(err))
	}
}
