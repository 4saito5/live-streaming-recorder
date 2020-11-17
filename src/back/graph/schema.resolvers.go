package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/4saito5/live-streaming-recorder/graph/generated"
	"github.com/4saito5/live-streaming-recorder/graph/model"
)

func (r *mutationResolver) CreateCheckList(ctx context.Context, input model.NewCheckList) (*model.CheckList, error) {
	// panic(fmt.Errorf("not implemented"))

	// isRecord := 0
	// if input.IsRecord {
	// 	isRecord := 1
	// }
	checkList := &model.CheckList{
		// ID:     	fmt.Sprintf("T%d", rand.Int()),
		Group:    input.Group,
		Name:     input.Name,
		Site:     input.Site,
		URL:      input.URL,
		IsRecord: input.IsRecord,
	}

	result := r.DB.Create(&checkList)
	if result.Error != nil {
		panic(result.Error)
	}

	// r.checklists = append(r.checklists, result.RowsAffected)
	// return result.RowsAffected, nil
	r.checklists = append(r.checklists, checkList)
	return checkList, nil
}

func (r *queryResolver) Checklists(ctx context.Context) ([]*model.CheckList, error) {
	// panic(fmt.Errorf("not implemented"))

	r.DB.Find(&r.checklists)
	return r.checklists, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
