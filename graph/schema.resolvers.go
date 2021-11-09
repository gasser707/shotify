package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gasser707/go-gql-server/auth"
	"github.com/gasser707/go-gql-server/graph/generated"
	"github.com/gasser707/go-gql-server/graph/model"
)

func (r *imageResolver) User(ctx context.Context, obj *model.Image) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadImages(ctx context.Context, input []*model.NewImageInput) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteImages(ctx context.Context, input []*model.DeleteImageInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateImage(ctx context.Context, input model.UpdateImageInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BuyImage(ctx context.Context, input *model.BuyImageInput) (*model.Sale, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (bool, error) {
	ok, err := auth.AuthService.Login(ctx)

	if ok {
		return ok, nil
	} else {
		return false, fmt.Errorf(err.Error())
	}
}

func (r *mutationResolver) Logout(ctx context.Context, input *bool) (bool, error) {
	ok, err := auth.AuthService.Logout(ctx)

	if ok {
		return ok, nil
	} else {
		return false, fmt.Errorf(err.Error())
	}
}
func (r *queryResolver) Images(ctx context.Context, input *model.ImageFilterInput) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, input *model.UserFilterInput) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Image returns generated.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type imageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
