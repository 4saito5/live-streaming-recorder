package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/4saito5/live-streaming-recorder/graph/generated"
	"github.com/4saito5/live-streaming-recorder/graph/model"
)

func (r *mutationResolver) CreateCheckList(ctx context.Context, input model.NewCheckList) (*model.CheckList, error) {
	checkList := &model.CheckList{
		Group:    input.Group,
		Name:     input.Name,
		Site:     input.Site,
		Key:      input.Key,
		URL:      input.URL,
		IsRecord: input.IsRecord,
	}

	result := r.DB.Create(&checkList)
	if result.Error != nil {
		panic(result.Error)
	}

	r.checklists = append(r.checklists, checkList)
	return checkList, nil
}

func (r *queryResolver) Checklists(ctx context.Context) ([]*model.CheckList, error) {
	r.DB.Find(&r.checklists)
	return r.checklists, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
