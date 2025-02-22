package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
)

type processFootmarkService struct {
	repo repository.ProcessFootmarkRepository
}

func (p processFootmarkService) GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark string, claim int) *v1.ProcessFootmark {
	return p.repo.GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark, claim)
}

func (p processFootmarkService) GetByProcessIdAndStepAndClaim(processId, step string, claim int) []v1.ProcessFootmark {
	return p.repo.GetByProcessIdAndStepAndClaim(processId, step, claim)
}

func (p processFootmarkService) Store(processFootmark v1.ProcessFootmark) {
	p.repo.Store(processFootmark)
}

func (p processFootmarkService) GetByProcessId(processId string) []v1.ProcessFootmark {
	return p.repo.GetByProcessId(processId)
}

// NewProcessFootmarkService returns Process footmark service
func NewProcessFootmarkService(repo repository.ProcessFootmarkRepository) service.ProcessFootmark {
	return &processFootmarkService{
		repo: repo,
	}
}
