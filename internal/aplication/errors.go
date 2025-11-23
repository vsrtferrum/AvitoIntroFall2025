package aplication

import (
	"errors"
	"fmt"

	"github.com/vsrtferrum/AvitoIntroFall2025/internal/domain/vo"
)

var (
	ErrNoTeam                 = errors.New("error no team")
	ErrFailedToAddTeam        = errors.New("failed to add team")
	ErrUserNotFound           = errors.New("no user found")
	ErrFailedToCreatePr       = errors.New("failed to create pull request")
	ErrFailedToMergePr        = errors.New("failed to merge pull request")
	ErrFailedToReassignPr     = errors.New("failed to reassign pull request")
	ErrFailedToDeactivateUser = errors.New("failed to deactivate users")
)

func ErrUsrNotFound(usr vo.UserId) error {
	return fmt.Errorf("error no user found %s", usr)
}

func ErrTeamNotFound(tm vo.TeamName) error {
	return fmt.Errorf("team no user found %s", tm)
}

func ErrRequestNotFound(req vo.PullRequestId) error {
	return fmt.Errorf("team no user found %s", req)
}
