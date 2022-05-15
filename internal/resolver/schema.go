package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mkdirJava/golangAwsPipelinePlayground/graph/generated"
	"github.com/mkdirJava/golangAwsPipelinePlayground/graph/generated/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "1",
		Text: "THIS IS A TEST",
		Done: true,
		User: &model.User{ID: "1", Name: "BOB"},
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	models := []*model.Todo{
		{ID: "1",
			Text: "THIS IS A TEST",
			Done: true,
			User: &model.User{ID: "1", Name: "BOB"}},
	}
	return models, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
