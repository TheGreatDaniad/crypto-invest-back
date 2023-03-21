package customErrors

import (
	"github.com/thegreatdaniad/crypto-invest/utils"
)

type Error struct {
	Err  error
	Code string
}

func (e Error) Error() string {
	return e.Err.Error()
}

func New(err error, code string) Error {
	return Error{
		Err:  err,
		Code: code,
	}
}
func (e Error) IsCommonHttpError() bool {
	return utils.SliceIncludes(commonHttpErrors, e.Code)
}

var commonHttpErrors []string = []string{
	InvalidInputs,
	InternalServerError,
	ProccessCanceledByParent,
	DatabaseConnectionError,
	BadRequest,
}

const (
	UserAlreadyExists        = "USER_ALREADY_EXISTS"
	UserAlreadyVerified      = "USER_ALREADY_VERIFIED"
	InvalidInputs            = "INVALID_INPUTS"
	InternalServerError      = "INTERNAL_SERVER_ERROR"
	DatabaseErrors           = "DATABASE_ERRORS"
	DatabaseConnectionError  = "DATABASE_CONNECTION_ERROR"
	ProccessCanceledByParent = "PROCESS_CANCELED_BY_PARENT"
	NoError                  = "NO_ERROR"
	NotFound                 = "NOT_FOUND"
	BadRequest               = "BAD_REQUEST"
	UserNotFound             = "USER_NOT_FOUND"
	AuthenticationFailed     = "AUTHENTICATION_FAILED"
	PasswordIsNotCorrect     = "PASSWORD_IS_NOT_CORRECT"
	UnknownError             = "UNKNOWN_ERROR"
	InvalidUserId            = "INVALID_USER_ID"
	InvalidCredentials       = "INVALID_CREDENTIALS"
	Unauthorized             = "UNAUTHORIZED"
	EmailNotSent             = "EMAIL_NOT_SENT"
	CannotParseJwt           = "CANNOT_PARSE_JWT"
	CannotParseToken         = "CANNOT_PARSE_TOKEN"
	ReceiverNotFound         = "RECEIVER_NOT_FOUND"
	UploadFileError          = "UPLOAD_FILE_ERROR"
	InvalidUsecase           = "INVALID_USECASE"
	InvalidMessageId         = "INVALID_MESSAGE_ID"
	MessageNotFound          = "MESSAGE_NOT_FOUND"
	DeliveryNotFound         = "DELIVERY_NOT_FOUND"
	NotificationNotFound     = "NOTIFICATION_NOT_FOUND"
	TicketNotFound           = "TICKET_NOT_FOUND"
	AnswerNotFound           = "ANSWER_NOT_FOUND"
	InvalidTicketId          = "INVALID_TICKET_ID"
	InvalidAnswerId          = "INVALID_ANSWER_ID"
	PostNotFound             = "POST_NOT_FOUND"
	InvalidPostId            = "INVALID_POST_ID"
	InvalidPostType          = "INVALID_POST_TYPE"
	FileNotFound             = "FILE_NOT_FOUND"
	InvalidFileId            = "INVALID_FILE_ID"
	InvalidId                = "INVALID_ID"
	MessageOrUserNotFound    = "MESSAGE_OR_USER_NOT_FOUND"
	DownloadFileError        = "DOWNLOAD_FILE_ERROR"
	FacebookError            = "FACEBOOK_API_ERROR"
)
