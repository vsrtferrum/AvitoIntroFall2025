package api

import (
	"encoding/json"
	"net/http"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/user"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
	httpGen "github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/http"
)

type APIHandler struct {
	Logger
	app *aplication.GitModel
}

func NewAPIHandler(app *aplication.GitModel, logger Logger) *APIHandler {
	return &APIHandler{app: app, Logger: logger}
}

func (h *APIHandler) PostPullRequestCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody httpGen.PostPullRequestCreateJSONRequestBody
	if err := readJSONBody(r, &reqBody); err != nil {
		err = writeError(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body")
		if err != nil {
			h.WriteError(err)
		}
		return
	}

	appReq := aplication.CreatePullRequestRequest{
		PullRequestId:   vo.PullRequestId(reqBody.PullRequestId),
		PullRequestName: vo.PullRequestName(reqBody.PullRequestName),
		AuthorId:        vo.UserId(reqBody.AuthorId),
	}

	resp, err := h.app.CreatePr(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := httpGen.PullRequest{
		PullRequestId:     string(resp.PullRequestId),
		PullRequestName:   string(resp.PullRequestName),
		AuthorId:          string(resp.AuthorId),
		Status:            httpGen.PullRequestStatus(resp.Status),
		AssignedReviewers: convertUserIdsToStrings(resp.AssignedReviewers),
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) PostPullRequestMerge(w http.ResponseWriter, r *http.Request) {
	var reqBody httpGen.PostPullRequestMergeJSONRequestBody
	if err := readJSONBody(r, &reqBody); err != nil {
		err := writeError(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body")
		if err != nil {
			h.WriteError(err)
		}
		return
	}

	appReq := aplication.MergePullRequestRequest{
		PullRequestId: vo.PullRequestId(reqBody.PullRequestId),
	}

	resp, err := h.app.MergePr(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := httpGen.PullRequest{
		PullRequestId:     string(resp.PullRequestId),
		PullRequestName:   string(resp.PullRequestName),
		AuthorId:          string(resp.AuthorId),
		Status:            httpGen.PullRequestStatus(resp.Status),
		AssignedReviewers: convertUserIdsToStrings(resp.AssignedReviewers),
		MergedAt:          &resp.MergedAt,
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) PostPullRequestReassign(w http.ResponseWriter, r *http.Request) {
	var reqBody httpGen.PostPullRequestReassignJSONRequestBody
	if err := readJSONBody(r, &reqBody); err != nil {
		err := writeError(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body")
		if err != nil {
			h.WriteError(err)
		}
		return
	}

	appReq := aplication.ReassignPullRequestRequest{
		PullRequestId: vo.PullRequestId(reqBody.PullRequestId),
		OldReviewerId: vo.UserId(reqBody.OldUserId),
	}

	resp, err := h.app.ReassignPr(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := httpGen.PullRequest{
		PullRequestId:     string(resp.PullRequestId),
		PullRequestName:   string(resp.PullRequestName),
		AuthorId:          string(resp.AuthorId),
		Status:            httpGen.PullRequestStatus(resp.Status),
		AssignedReviewers: convertUserIdsToStrings(resp.AssignedReviewers),
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) PostTeamAdd(w http.ResponseWriter, r *http.Request) {
	var reqBody httpGen.Team
	if err := readJSONBody(r, &reqBody); err != nil {
		err := writeError(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body")
		if err != nil {
			h.WriteError(err)
		}
		return
	}

	users := make([]user.User, len(reqBody.Members))
	for i, member := range reqBody.Members {
		users[i] = user.User{
			Id:       vo.UserId(member.UserId),
			Name:     vo.UserName(member.Username),
			IsActive: member.IsActive,
		}
	}

	appReq := aplication.AddTeamRequest{
		TeamName: vo.TeamName(reqBody.TeamName),
		Users:    users,
	}

	resp, err := h.app.AddTeam(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := httpGen.Team{
		TeamName: string(resp.TeamName),
		Members:  convertUsersToTeamMembers(resp.Users),
	}

	err = writeJSONResponse(w, http.StatusCreated, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) PostTeamDeactivate(w http.ResponseWriter, r *http.Request) {
	var reqBody httpGen.PostTeamDeactivateJSONRequestBody
	if err := readJSONBody(r, &reqBody); err != nil {
		err = writeError(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body")
		if err != nil {
			h.WriteError(err)
		}
		return
	}

	appReq := aplication.DeactivateUsersRequest{
		TeamName: vo.TeamName(reqBody.TeamName),
	}

	err := h.app.DeactivateUsers(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *APIHandler) GetTeamGet(w http.ResponseWriter, r *http.Request, params httpGen.GetTeamGetParams) {
	team, err := h.app.GetTeam(vo.TeamName(params.TeamName))
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := httpGen.Team{
		TeamName: string(team.TeamName),
		Members:  convertTeamToTeamMembers(team),
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) GetUsersGetReview(w http.ResponseWriter, r *http.Request, params httpGen.GetUsersGetReviewParams) {
	appReq := aplication.GetReviewRequest{
		UserId: vo.UserId(params.UserId),
	}

	resp, err := h.app.GetReview(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := make([]httpGen.PullRequestShort, len(resp.PullRequests))
	for i, pr := range resp.PullRequests {
		apiResp[i] = httpGen.PullRequestShort{
			PullRequestId:   string(pr.PullRequestId),
			PullRequestName: string(pr.PullRequestName),
			AuthorId:        string(pr.Author),
			Status:          httpGen.PullRequestShortStatus(pr.Status),
		}
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) PostUsersSetIsActive(w http.ResponseWriter, r *http.Request) {
	var reqBody httpGen.PostUsersSetIsActiveJSONRequestBody
	if err := readJSONBody(r, &reqBody); err != nil {
		err = writeError(w, http.StatusBadRequest, INVALID_REQUEST, "Invalid request body")
		if err != nil {
			h.WriteError(err)
		}
		return
	}

	user, err := h.app.SetActivate(vo.UserId(reqBody.UserId), reqBody.IsActive)
	if err != nil {
		handleAppError(w, err)
		return
	}

	apiResp := httpGen.User{
		UserId:   string(user.UserId),
		Username: string(user.UserName),
		IsActive: user.IsActive,
		TeamName: string(user.TeamName),
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func (h *APIHandler) GetUsersStatsReviewedPRs(w http.ResponseWriter, r *http.Request, params httpGen.GetUsersStatsReviewedPRsParams) {
	limit := 50
	if params.Limit != nil && *params.Limit > 0 && *params.Limit <= 100 {
		limit = *params.Limit
	}

	var cursor vo.UserId
	if params.Cursor != nil {
		cursor = vo.UserId(*params.Cursor)
	}

	appReq := aplication.StatReviewRequest{
		Cursor: cursor,
		Limit:  limit,
	}

	stats, err := h.app.GetStat(appReq)
	if err != nil {
		handleAppError(w, err)
		return
	}

	type UserStats struct {
		UserId     string `json:"user_id"`
		ReviewedPr int    `json:"reviewed_pr"`
	}

	apiResp := make([]UserStats, len(stats))
	for i, stat := range stats {
		apiResp[i] = UserStats{
			UserId:     string(stat.UserId),
			ReviewedPr: stat.ReviewedPr,
		}
	}

	err = writeJSONResponse(w, http.StatusOK, apiResp)
	if err != nil {
		h.WriteError(err)
	}
}

func readJSONBody(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func writeJSONResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}
