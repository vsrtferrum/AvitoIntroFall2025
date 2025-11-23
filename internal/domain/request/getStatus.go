package request

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (r *Request) GetStatus() vo.Status {
	return r.Status
}
