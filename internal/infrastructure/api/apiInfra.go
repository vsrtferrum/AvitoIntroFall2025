package api

import (
	"net/http"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/team"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/user"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
	httpGen "github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/http"
)

func writeError(w http.ResponseWriter, status int, code, message string) error {
	errorResp := httpGen.ErrorResponse{
		Error: struct {
			Code    httpGen.ErrorResponseErrorCode `json:"code"`
			Message string                         `json:"message"`
		}{
			Code:    httpGen.ErrorResponseErrorCode(code),
			Message: message,
		},
	}
	return writeJSONResponse(w, status, errorResp)
}
func handleAppError(w http.ResponseWriter, err error) {
	switch err {
	case aplication.ErrUserNotFound:
		writeError(w, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
	case aplication.ErrNoTeam:
		writeError(w, http.StatusNotFound, "TEAM_NOT_FOUND", "Team not found")
	case aplication.ErrFailedToCreatePr:
		writeError(w, http.StatusBadRequest, "PR_CREATION_FAILED", "Failed to create pull request")
	case aplication.ErrFailedToMergePr:
		writeError(w, http.StatusBadRequest, "PR_MERGE_FAILED", "Failed to merge pull request")
	case aplication.ErrFailedToReassignPr:
		writeError(w, http.StatusBadRequest, "PR_REASSIGN_FAILED", "Failed to reassign pull request")
	case aplication.ErrFailedToAddTeam:
		writeError(w, http.StatusBadRequest, "TEAM_ADD_FAILED", "Failed to add team")
	case aplication.ErrFailedToDeactivateUser:
		writeError(w, http.StatusBadRequest, "DEACTIVATE_FAILED", "Failed to deactivate users")
	default:
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error")
	}
}

func convertUserIdsToStrings(userIds []vo.UserId) []string {
	result := make([]string, len(userIds))
	for i, id := range userIds {
		result[i] = string(id)
	}
	return result
}

func convertUsersToTeamMembers(users []user.User) []httpGen.TeamMember {
	members := make([]httpGen.TeamMember, len(users))
	for i, user := range users {
		members[i] = httpGen.TeamMember{
			UserId:   string(user.Id),
			Username: string(user.Name),
			IsActive: user.IsActive,
		}
	}
	return members
}

func convertTeamToTeamMembers(t *team.Team) []httpGen.TeamMember {
	userIds := t.Get()
	members := make([]httpGen.TeamMember, len(userIds))

	for i, userId := range userIds {
		members[i] = httpGen.TeamMember{
			UserId:   string(userId),
			IsActive: true,
		}
	}
	return members
}
