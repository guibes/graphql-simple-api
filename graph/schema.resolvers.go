package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/guibes/graphql-simple-api/database"
	"github.com/guibes/graphql-simple-api/graph/generated"
	"github.com/guibes/graphql-simple-api/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return db.Save(input), nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
