package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/4saito5/live-streaming-recorder/graph/generated"
	"github.com/4saito5/live-streaming-recorder/graph/model"
)

func (r *mutationResolver) CreateCheckList(ctx context.Context, input model.EditCheckList) (*model.CheckList, error) {
	checkList := &model.CheckList{
		Group:    input.Group,
		Name:     input.Name,
		Site:     input.Site,
		Key:      input.Key,
		URL:      input.URL,
		IsRecord: input.IsRecord,
		OnLive:   input.OnLive,
	}

	result := r.DB.Create(&checkList)
	if result.Error != nil {
		panic(result.Error)
	}
	return checkList, nil
}

func (r *mutationResolver) DeleteCheckList(ctx context.Context, input model.DeleteCheckListKey) (*model.Response, error) {
	result := r.DB.Where("id = ?", input.ID).Delete(&model.CheckList{})
	if result.Error != nil {
		panic(result.Error)
	}

	response := &model.Response{
		Code:    200,
		Message: "success.",
	}
	return response, nil
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
