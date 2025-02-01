package worker

import (
	"github.com/achintya-7/dating-app/pkg/mail"
	db "github.com/achintya-7/dating-app/pkg/sql/sqlc"
	"github.com/achintya-7/dating-app/pkg/worker"
	"github.com/hibiken/asynq"
)

type RedisTaskProcessor struct {
	server *asynq.Server
	store  *db.Store
	mailer mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store *db.Store, mailer mail.EmailSender) *RedisTaskProcessor {
	server := asynq.NewServer(
		redisOpt, asynq.Config{
			Queues: map[string]int{
				worker.QUEUE_CRITICAL: 10,
				worker.QUEUE_DEFAULT:  5,
			},
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
		mailer: mailer,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(worker.TASK_SEND_MATCHED_EMAIL, processor.SendMatchedEmailProcessor)
	mux.HandleFunc(worker.TASK_CALCULATE_USER_ATTRACTIVENESS, processor.CalculateUserAttractivenessProcessor)

	return processor.server.Start(mux)
}

func (processor *RedisTaskProcessor) Shutdown() {
	processor.server.Shutdown()
}
