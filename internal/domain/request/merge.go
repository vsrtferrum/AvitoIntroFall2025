package request

import "github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"

func (r *Request) Merge() {
	r.Status = vo.StatusMERGED
}
