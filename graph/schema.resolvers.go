package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/scorpionknifes/pts-backend/graph/generated"
	"github.com/scorpionknifes/pts-backend/graph/model"
	"github.com/scorpionknifes/pts-backend/internal/stories"
	"github.com/scorpionknifes/pts-backend/internal/turns"
)

func (r *mutationResolver) CreateUser(ctx context.Context) (*model.User, error) {
	user := model.User{
		Name: fmt.Sprintf("Anonymous%d", rand.Intn(9999)),
	}
	tx := r.DB.Create(&user)
	return &user, tx.Error
}

func (r *mutationResolver) CreateStory(ctx context.Context, input model.StoryInput) (*model.Story, error) {
	story := model.Story{
		Name: input.Name,
		Tags: input.Tags,
	}
	tx := r.DB.Create(&story)
	stories.Update(story)
	return &story, tx.Error
}

func (r *mutationResolver) CreateTurn(ctx context.Context, input model.TurnInput) (*model.Turn, error) {
	turn := model.Turn{
		UserID:  input.UserID,
		StoryID: input.StoryID,
		Value:   input.Value,
	}

	r.DB.Create(&turn)
	turns.Update(input.StoryID, turn)
	var story *model.Story
	type Result struct {
		Name string
		Age  int
	}

	var result model.Count
	r.DB.Raw("SELECT COUNT(DISTINCT user_id) as people, COUNT(user_id) as count FROM  [dbo].[turns] WHERE story_id = ?", input.StoryID).Scan(&result)

	r.DB.Model(&story).Where("id = ?", input.StoryID).Update("Count", result.Count).Update("People", result.People)
	return &turn, nil
}

func (r *queryResolver) Stories(ctx context.Context) ([]*model.Story, error) {
	var stories []*model.Story
	tx := r.DB.Preload("Turns").Find(&stories)
	return stories, tx.Error
}

func (r *queryResolver) Story(ctx context.Context, id int) (*model.Story, error) {
	story := model.Story{}
	tx := r.DB.Preload("Turns").First(&story, id)
	log.Println(story)
	return &story, tx.Error
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	tx := r.DB.Preload("Turns").Find(&users)
	return users, tx.Error
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user := model.User{}
	tx := r.DB.First(&user, id)
	return &user, tx.Error
}

func (r *subscriptionResolver) Turns(ctx context.Context, story int) (<-chan *model.Turn, error) {
	return turns.Add(story), nil
}

func (r *subscriptionResolver) Stories(ctx context.Context) (<-chan *model.Story, error) {
	return stories.Add(), nil
}

func (r *turnResolver) User(ctx context.Context, obj *model.Turn) (*model.User, error) {
	user := model.User{}
	tx := r.DB.First(&user, obj.UserID)
	return &user, tx.Error
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// Turn returns generated.TurnResolver implementation.
func (r *Resolver) Turn() generated.TurnResolver { return &turnResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type turnResolver struct{ *Resolver }
