package api

type ErrorResponseErrorCode string

const (
	NOCANDIDATE ErrorResponseErrorCode = "NO_CANDIDATE"
	NOTASSIGNED ErrorResponseErrorCode = "NOT_ASSIGNED"
	NOTFOUND    ErrorResponseErrorCode = "NOT_FOUND"
	PREXISTS    ErrorResponseErrorCode = "PR_EXISTS"
	PRMERGED    ErrorResponseErrorCode = "PR_MERGED"
	TEAMEXISTS  ErrorResponseErrorCode = "TEAM_EXISTS"
)

type PullRequestStatus string

const (
	PullRequestStatusMERGED PullRequestStatus = "MERGED"
	PullRequestStatusOPEN   PullRequestStatus = "OPEN"
)

type PullRequestShortStatus string

const (
	PullRequestShortStatusMERGED PullRequestShortStatus = "MERGED"
	PullRequestShortStatusOPEN   PullRequestShortStatus = "OPEN"
)

const (
	USER_NOT_FOUND     = "USER_NOT_FOUND"
	TEAM_NOT_FOUND     = "TEAM_NOT_FOUND"
	PR_CREATION_FAILED = "PR_CREATION_FAILED"
	PR_MERGE_FAILED    = "PR_MERGE_FAILED"
	PR_REASSIGN_FAILED = "PR_REASSIGN_FAILED"
	TEAM_ADD_FAILED    = "TEAM_ADD_FAILED"
	DEACTIVATE_FAILED  = "DEACTIVATE_FAILED"
	INTERNAL_ERROR     = "INTERNAL_ERROR"
	INVALID_REQUEST    = "INVALID_REQUEST"
)
