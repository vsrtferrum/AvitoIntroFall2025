package aplication

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/request"

func (g *GitModel) GetReview(pr GetReviewRequest) (*GetReviewResponce, error) {
	g.usersMtx.Lock()
	defer g.usersMtx.Unlock()
	defer g.requestsMtx.Unlock()
	prs, ok := g.users[pr.UserId]
	if !ok {
		return nil, ErrUserNotFound
	}
	req := make([]request.Request, 0, len(prs.PrForReview))
	g.requestsMtx.Lock()
	for _, pr := range prs.PrForReview {
		req = append(req, g.requests[pr])
	}
	return &GetReviewResponce{
		PullRequests: req,
	}, nil
}
