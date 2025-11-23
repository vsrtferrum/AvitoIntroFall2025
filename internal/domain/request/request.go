package request

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

type Request struct {
	Logger
	vo.PullRequestId
	vo.PullRequestName
	Author   vo.UserId
	Status   vo.Status
	Reviewer []vo.UserId
}

func NewRequest(pullRequestId vo.PullRequestId, pullRequestName vo.PullRequestName, author vo.UserId, logger Logger) *Request {
	return &Request{
		PullRequestId:   pullRequestId,
		PullRequestName: pullRequestName,
		Author:          author,
		Status:          vo.StatusOPEN,
		Logger:          logger,
		Reviewer:        make([]vo.UserId, 0, 2),
	}
}
