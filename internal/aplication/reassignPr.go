package aplication

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (g *GitModel) ReassignPr(pr ReassignPullRequestRequest) (*ReassignPullRequestResponce, error) {
	g.requestsMtx.Lock()
	defer g.requestsMtx.Unlock()
	oldPr := g.requests[pr.PullRequestId]
	if oldPr.GetStatus() == vo.StatusMERGED {
		return nil, ErrFailedToReassignPr
	}

	resp, err := g.storage.PrReassign(pr)
	if err != nil {
		g.logger.WriteError(err)
		return nil, ErrFailedToReassignPr
	}
	for i, val := range oldPr.Reviewer {
		if val == pr.OldReviewerId {
			oldPr.Reviewer[i] = resp.AssignedReviewers[0]
		}
	}
	resp.PullRequestId = pr.PullRequestId
	resp.PullRequestName = g.requests[pr.PullRequestId].PullRequestName
	resp.AuthorId = g.requests[pr.PullRequestId].Author
	resp.Status = vo.StatusOPEN
	resp.AssignedReviewers = oldPr.Reviewer
	return resp, nil
}
