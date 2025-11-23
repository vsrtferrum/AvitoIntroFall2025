package aplication

func (g *GitModel) GetStat(pr StatReviewRequest) ([]StatReviewResponce, error) {
	return g.storage.GetReviwerStats(pr)
}
