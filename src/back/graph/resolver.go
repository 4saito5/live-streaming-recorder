package graph

import (
  "gorm.io/gorm"
	"github.com/4saito5/live-streaming-recorder/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB					*gorm.DB
	checklists 	[]*model.CheckList
}
