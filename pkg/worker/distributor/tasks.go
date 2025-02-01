package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/achintya-7/dating-app/logger"
	"github.com/achintya-7/dating-app/pkg/worker"
	"github.com/hibiken/asynq"
)

// Sends a task to the worker to send an email to the user
func (distributor *RedisTaskDistributor) SendMatchedEmailTask(
	ctx context.Context,
	payload *worker.SendMatchEmailTaskPayload,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(worker.QUEUE_CRITICAL),
	}

	task := asynq.NewTask(worker.TASK_SEND_MATCHED_EMAIL, jsonPayload, opts...)
	_, err = distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}

	logger.Info(ctx, "send email task enqueued")

	return nil
}

// Sends a task to the worker to calculate the user's attractiveness
func (rtd *RedisTaskDistributor) CalculateUserAttractivenessTask(
	ctx context.Context,
	payload *worker.CalculateUserAttractivenessTaskPayload,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(worker.QUEUE_DEFAULT),
	}

	task := asynq.NewTask(worker.TASK_CALCULATE_USER_ATTRACTIVENESS, jsonPayload, opts...)
	_, err = rtd.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}

	logger.Info(ctx, "calculate user attractiveness task enqueued")

	return nil
}
