package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/guibes/graphql-simple-api/database"
	"github.com/guibes/graphql-simple-api/graph/generated"
	"github.com/guibes/graphql-simple-api/graph/middleware"
	"github.com/guibes/graphql-simple-api/graph/model"
	"github.com/guibes/graphql-simple-api/graph/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	return db.Save(input), nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Token, error) {
	user, err := db.FindUserByEmail(email)

	if err != nil || user == nil {
		return nil, errors.New("User not found")
	}

	if !utils.ComparePassword(password, user.Password) {
		return nil, errors.New("Passwords doesn't match")
	}

	expiredAt := int(time.Now().Add(time.Hour * 1).Unix())
	obj := &model.Token{
		Token:     utils.GenerateJwt(user.ID, int64(expiredAt)),
		ExpiredAt: expiredAt,
	}

	return obj, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	fmt.Println(ctx)
	userAuth := middleware.GetAuthFromContext(ctx)
	fmt.Println(userAuth)
	if userAuth == nil {
		return nil, errors.New("Access denied")
	}
	return db.FindByID(id), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	fmt.Println(ctx)
	userAuth := middleware.GetAuthFromContext(ctx)
	fmt.Println(userAuth)
	if userAuth == nil {
		return nil, errors.New("Access denied")
	}
	return db.All(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) FirstName(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *userResolver) LastName(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

type userResolver struct{ *Resolver }

var db = database.Connect()
