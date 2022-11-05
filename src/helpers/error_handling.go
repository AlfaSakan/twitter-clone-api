package helpers

import "fmt"

type errKind int

const (
	_ errKind = iota
	noHeader
	cantReadHeader
	cannotFoundUser
	cannotFoundUsername
	cannotFoundSession
	wrongPassword
	failedCreatedUser
	duplicateUser
)

type WaveError struct {
	kind  errKind
	value string
	err   error
}

func (e WaveError) Error() string {
	switch e.kind {
	case noHeader:
		return "no header (file too short?)"
	case cannotFoundUser:
		return fmt.Sprintf("can't found user with id : %s", e.value)
	case cannotFoundUsername:
		return fmt.Sprintf("can't found username : %s", e.value)
	case cannotFoundSession:
		return fmt.Sprintf("can't found session with id : %s", e.value)
	case wrongPassword:
		return "wrong input password!"
	case failedCreatedUser:
		return fmt.Sprintf("failed to created user. (%s)", e.err.Error())
	case duplicateUser:
		return fmt.Sprintf("this value already exist %s. (%s)", e.value, e.err.Error())
	default:
		return "bad request"
	}
}

func (e WaveError) With(val string) WaveError {
	e1 := e
	e1.value = val
	return e1
}

func (e WaveError) From(pos string, err error) WaveError {
	e1 := e
	e1.value = pos
	e1.err = err
	return e1
}

var (
	HeaderMissing     = WaveError{kind: noHeader}
	UserNotFound      = WaveError{kind: cannotFoundUser}
	UsernameNotFound  = WaveError{kind: cannotFoundUsername}
	SessionNotFound   = WaveError{kind: cannotFoundSession}
	WrongPassword     = WaveError{kind: wrongPassword}
	DuplicateUser     = WaveError{kind: duplicateUser}
	FailedCreatedUser = WaveError{kind: failedCreatedUser}
)
