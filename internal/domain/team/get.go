package team

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (t *Team) Get() []vo.UserId {
	return t.users
}
