package db

import (
	"github.com/achintya-7/dating-app/internal/dto"
	"github.com/gin-gonic/gin"
)

type CreateUserTx struct {
	UserReq CreateUserParams
}

// CreateUserTx creates a new user and an entry in the ranking table
func (store *Store) CreateUserTx(ctx *gin.Context, arg CreateUserTx) (*dto.CreateUserResponse, error) {
	var user dto.CreateUserResponse

	err := store.ExecTx(ctx, func(q *Queries) error {
		// Create a user
		_, err := q.CreateUser(ctx, arg.UserReq)
		if err != nil {
			return err
		}

		// Create an entry in the ranking table
		_, err = q.CreateRanking(ctx, arg.UserReq.UserID)
		if err != nil {
			return err
		}

		user.ID = arg.UserReq.UserID
		user.Email = arg.UserReq.Email
		user.Name = arg.UserReq.Name
		user.Age = int8(arg.UserReq.Age)
		user.Gender = arg.UserReq.Gender

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}
