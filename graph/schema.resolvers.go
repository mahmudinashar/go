package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/mahmudinashar/go/graph/generated"
	"github.com/mahmudinashar/go/graph/model"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (r *mutationResolver) InputUsers(ctx context.Context, input model.UserCreateInputParam) (*model.Users, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	hashedPassword, err := Hash(input.Password)
	if err != nil {
		return nil, err
	}

	createUser := model.Users{
		Nickname:  input.Nickname,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Role:      input.Role,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
	r.DB.Create(&createUser)

	return &createUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.Users, error) {
	users := []*model.Users{}
	r.DB.Find(&users)
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
