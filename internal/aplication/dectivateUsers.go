package aplication

func (g *GitModel) DeactivateUsers(req DeactivateUsersRequest) error {
	prs, err := g.storage.DeactivateUsers(req)
	if err != nil {
		g.logger.WriteError(err)
		return ErrFailedToDeactivateUser
	}
	g.requestsMtx.Lock()
	g.teamsMtx.Lock()
	
	defer g.requestsMtx.Unlock()
	defer g.teamsMtx.Unlock()
	
	for _, pr := range prs {
		oldReq, ok := g.requests[pr.PullRequestId]
		if !ok {
			g.logger.WriteError(ErrRequestNotFound(oldReq.PullRequestId))
			continue
		}
		oldReq.Reviewer = pr.Reviewer
		g.requests[pr.PullRequestId] = oldReq
	}
	team, ok := g.teams[req.TeamName]
	if !ok {
		return ErrTeamNotFound(req.TeamName)
	}
	for _, user := range team.Get() {
		delete(g.users, user)
	}
	delete(g.teams, req.TeamName)
	return nil
}
