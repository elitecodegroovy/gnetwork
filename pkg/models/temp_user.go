package models

import (
	"errors"
	"time"
)

// Typed errors
var (
	ErrTempUserNotFound = errors.New("User not found")
)

type TempUserStatus string

const (
	TmpUserSignUpStarted TempUserStatus = "SignUpStarted"
	TmpUserInvitePending TempUserStatus = "InvitePending"
	TmpUserCompleted     TempUserStatus = "Completed"
	TmpUserRevoked       TempUserStatus = "Revoked"
)

// TempUser holds data for org invites and unconfirmed sign ups
type TempUser struct {
	Id              int64
	OrgId           int64
	Version         int
	Email           string
	Name            string
	InvitedByUserId int64
	Status          TempUserStatus

	EmailSent   bool
	EmailSentOn time.Time
	Code        string
	RemoteAddr  string

	Created time.Time
	Updated time.Time
}

// ---------------------
// COMMANDS

type CreateTempUserCommand struct {
	Email           string
	Name            string
	OrgId           int64
	InvitedByUserId int64
	Status          TempUserStatus
	Code            string
	RemoteAddr      string

	Result *TempUser
}

type UpdateTempUserStatusCommand struct {
	Code   string
	Status TempUserStatus
}

type UpdateTempUserWithEmailSentCommand struct {
	Code string
}

type GetTempUsersQuery struct {
	OrgId  int64
	Email  string
	Status TempUserStatus

	Result []*TempUserDTO
}

type GetTempUserByCodeQuery struct {
	Code string

	Result *TempUserDTO
}

type TempUserDTO struct {
	Id             int64          `json:"id"`
	OrgId          int64          `json:"orgId"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	InvitedByLogin string         `json:"invitedByLogin"`
	InvitedByEmail string         `json:"invitedByEmail"`
	InvitedByName  string         `json:"invitedByName"`
	Code           string         `json:"code"`
	Status         TempUserStatus `json:"status"`
	Url            string         `json:"url"`
	EmailSent      bool           `json:"emailSent"`
	EmailSentOn    time.Time      `json:"emailSentOn"`
	Created        time.Time      `json:"createdOn"`
}
