package aplication

import (
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/request"
)

type Storage interface {
	AddTeam(AddTeamRequest) error
	SetUserActive(SetIsAtivateRequest) error
	CreatePr(CreatePullRequestRequest) (*CreatePullRequestResponce, error)
	MergePr(MergePullRequestRequest) error
	PrReassign(ReassignPullRequestRequest) (*ReassignPullRequestResponce, error)
	GetReviwerStats(StatReviewRequest) ([]StatReviewResponce, error)
	DeactivateUsers(DeactivateUsersRequest) ([]request.Request, error)
}

type Logger interface {
	WriteError(error)
	WriteStatus(string)
	WriteDebugStatus(string)
}
