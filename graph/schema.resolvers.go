package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mahmudinashar/go/api/helpers"
	"github.com/mahmudinashar/go/graph/generated"
	"github.com/mahmudinashar/go/graph/model"
)

func (r *mutationResolver) InputUsers(ctx context.Context, input model.UserCreateInputParam) (*model.Users, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	hashedPassword, err := helpers.Hash(input.Password)
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

func (r *queryResolver) Users(ctx context.Context, nickname *string, email *string, role *int) ([]*model.Users, error) {
	users := []*model.Users{}
	var baseQuery string
	var inputQuery, firstQuery bool = false, true

	if nickname != nil {
		nicknameQuery := fmt.Sprintf("nickname LIKE '%s'", *nickname)
		if strings.Contains(*nickname, "%") {
			nicknameQuery = fmt.Sprintf("nickname LIKE '%s'", *nickname)
		}

		if firstQuery {
			baseQuery = nicknameQuery

		} else {
			baseQuery += " AND " + nicknameQuery
		}

		inputQuery = true
		firstQuery = false
	}

	if email != nil {
		emailQuery := fmt.Sprintf("email = '%s'", *email)
		if strings.Contains(*nickname, "%") {
			emailQuery = fmt.Sprintf("email LIKE '%s'", *email)
		}

		if firstQuery {
			baseQuery = emailQuery

		} else {
			baseQuery += " AND " + emailQuery
		}

		inputQuery = true
		firstQuery = false
	}

	if role != nil {
		roleQuery := fmt.Sprintf("role = '%d'", *role)

		if firstQuery {
			baseQuery = roleQuery

		} else {
			baseQuery += " AND " + roleQuery
		}

		inputQuery = true
		firstQuery = false
	}

	if inputQuery {
		r.DB.Debug().Where(baseQuery).Find(&users)

	} else {
		r.DB.Debug().Find(&users)
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
