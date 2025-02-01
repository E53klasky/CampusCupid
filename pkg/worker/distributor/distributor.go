package worker

import (
	"context"

	"github.com/achintya-7/dating-app/pkg/worker"
	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	// SendMatchedEmailTask enqueues a task to send an email to the users who have matched.
	SendMatchedEmailTask(
		ctx context.Context,
		payload *worker.SendMatchEmailTaskPayload,
	) error

	// CalculateUserAttractivenessTask enqueues a task to calculate the attractiveness of a user on a Yes swipe.
	CalculateUserAttractivenessTask(
		ctx context.Context,
		payload *worker.CalculateUserAttractivenessTaskPayload,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
