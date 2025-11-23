package aplication

import (
	"time"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/request"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/user"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

type AddTeamRequest struct {
	vo.TeamName
	Users []user.User
}

type AddTeamResponce struct {
	vo.TeamName
	Users []user.User
}

type GetTeamResponce struct {
	Users []user.User
}

type SetIsAtivateRequest struct {
	vo.UserId
	IsActive bool
}

type SetIsAtivateResponce struct {
	vo.UserId
	vo.UserName
	vo.TeamName
	IsActive bool
}

type CreatePullRequestRequest struct {
	vo.PullRequestId
	vo.PullRequestName
	AuthorId vo.UserId
}

type CreatePullRequestResponce struct {
	vo.PullRequestId
	vo.PullRequestName
	AuthorId vo.UserId
	vo.Status
	AssignedReviewers []vo.UserId
}

type MergePullRequestRequest struct {
	vo.PullRequestId
}

type MergePullRequestResponce struct {
	MergedAt time.Time
	vo.PullRequestId
	vo.PullRequestName
	AuthorId vo.UserId
	vo.Status
	AssignedReviewers []vo.UserId
}

type ReassignPullRequestRequest struct {
	vo.PullRequestId
	OldReviewerId vo.UserId
}

type ReassignPullRequestResponce struct {
	vo.PullRequestId
	vo.PullRequestName
	AuthorId          vo.UserId
	Status            vo.Status
	AssignedReviewers []vo.UserId
}

type GetReviewRequest struct {
	vo.UserId
}

type GetReviewResponce struct {
	PullRequests []request.Request
}

type StatReviewRequest struct {
	Cursor vo.UserId
	Limit  int
}

type StatReviewResponce struct {
	vo.UserId
	ReviewedPr int
}

type DeactivateUsersRequest struct {
	vo.TeamName
}
