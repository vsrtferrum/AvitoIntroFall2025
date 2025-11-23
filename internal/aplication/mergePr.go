package aplication

import (
	"time"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

func (g *GitModel) MergePr(pr MergePullRequestRequest) (*MergePullRequestResponce, error) {
	data, ok := g.requests[pr.PullRequestId]
	if !ok {
		return nil, ErrFailedToMergePr
	}
	resp := &MergePullRequestResponce{
		MergedAt:          time.Now(),
		PullRequestId:     data.PullRequestId,
		PullRequestName:   data.PullRequestName,
		AuthorId:          data.Author,
		Status:            vo.StatusMERGED,
		AssignedReviewers: data.Reviewer,
	}
	
	err := g.storage.MergePr(pr)
	if err != nil {
		g.logger.WriteError(err)
		return nil, ErrFailedToMergePr
	}

	g.requestsMtx.Lock()
	defer g.requestsMtx.Unlock()
	delete(g.requests, pr.PullRequestId)
	return resp, nil
}
