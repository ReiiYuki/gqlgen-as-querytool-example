package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-as-querytool/graph/generated"
	"gqlgen-as-querytool/graph/model"
)

var CurrentId = 1

var InMemList []*model.Todo

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	user := model.User{
		ID:   input.UserID,
		Name: input.UserID,
	}

	todo := model.Todo{
		User: &user,
		ID:   string(CurrentId),
		Text: input.Text,
		Done: false,
	}

	InMemList = append(InMemList, &todo)

	CurrentId += 1

	return &todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return InMemList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
