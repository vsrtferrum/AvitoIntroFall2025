package aplication

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (g *GitModel) CreatePr(pr CreatePullRequestRequest) (*CreatePullRequestResponce, error) {
	resp, err := g.storage.CreatePr(pr)
	if err != nil {
		g.logger.WriteError(err)
		return nil, ErrFailedToCreatePr
	}
	
	g.usersMtx.Lock()
	defer g.usersMtx.Unlock()

	for _, user := range resp.AssignedReviewers {
		temp, ok := g.users[user]
		if !ok {
			return nil, ErrUserNotFound
		}
		temp.PrForReview = append(temp.PrForReview, pr.PullRequestId)
		g.users[user] = temp

	}
	resp.PullRequestId = pr.PullRequestId
	resp.PullRequestName = pr.PullRequestName
	resp.AuthorId = pr.AuthorId
	resp.Status = vo.StatusOPEN
	return resp, nil
}
