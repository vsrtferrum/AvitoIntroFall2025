package request

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (r *Request) MergedCheck() bool {
	return r.Status == vo.StatusOPEN
}
