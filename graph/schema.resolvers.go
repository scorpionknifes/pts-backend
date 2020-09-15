package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/scorpionknifes/pts-backend/graph/generated"
	"github.com/scorpionknifes/pts-backend/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateStory(ctx context.Context, input model.StoryInput) (*model.Story, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTurn(ctx context.Context, input model.TurnInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stories(ctx context.Context) ([]*model.Story, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Story(ctx context.Context, id int) (*model.Story, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Turns(ctx context.Context, story int) (<-chan *model.Turn, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Stories(ctx context.Context) (<-chan *model.Story, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
