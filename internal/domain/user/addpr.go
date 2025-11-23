package user

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (u *User) Add(requests ...vo.PullRequestId) {
	u.PrForReview = append(u.PrForReview, requests...)
}
