package storage

import "errors"

var (
	ErrConnectionTimeLimit       = errors.New("connection time limit")
	ErrCreateConnection          = errors.New("failed to create connection")
	ErrCreateConfig              = errors.New("failed to create config")
	ErrCloseConnection           = errors.New("error while closing connection")
	ErrSendQuery                 = errors.New("error sending query or error responce")
	ErrConvertResponce           = errors.New("error converting database responce")
	ErrResultQuery               = errors.New("error nil result")
	ErrNonDeterministicUsers     = errors.New("error non unique user")
	ErrCreateTransaction         = errors.New("error create transaction")
	ErrExecTransaction           = errors.New("error exec transaction")
	ErrCommitTransaction         = errors.New("error commit transaction")
	ErrPrepareQuery              = errors.New("error failed to prepare query")
	ErrNonPointerValue           = errors.New("dest must be a pointer to struct")
	ErrFailedToCreateTransaction = errors.New("failed to create transaction")
)
