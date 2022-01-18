package v1

import (
	"github.com/lkeme/QSearch/app/controller/v1/question"
	"github.com/lkeme/QSearch/app/controller/v1/system"
)

type ApiGroup struct {
	QuestionApiGroup question.ApiGroup
	SystemApiGroup   system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
