package user

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

type User struct {
	Id          vo.UserId
	Name        vo.UserName
	PrForReview []vo.PullRequestId
	IsActive    bool
}

func NewUser(UserId vo.UserId, UserName vo.UserName, IsActive bool) *User {
	return &User{
		Id:          UserId,
		Name:        UserName,
		IsActive:    IsActive,
		PrForReview: []vo.PullRequestId{},
	}
}
