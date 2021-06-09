package responses

import "errors"

var (
	Success           = errors.New("success")
	UnSuccess         = errors.New("unsuccessful")
	ErrUnknown        = errors.New("unknown")
	BadRequest        = errors.New("bad request")
	ErrNotInt         = errors.New("err not int")
	SessionExpired    = errors.New("session expired")
	VerifyCodeExpired = errors.New("verify code expired")
	Existed           = errors.New("existed")
	NotAdmin          = errors.New("not admin")
	NotPermission     = errors.New("not permission")
	UnAuthorized      = errors.New("unauthorized")
	ErrLogin          = errors.New("login error")
	ErrSystem         = errors.New("system error")
	NotExisted        = errors.New("data not existed")
	ErrChangePass     = errors.New("change password error")
	NotMore           = errors.New("no more data")
	CannotEmpty       = errors.New("cannot empty")
	ErrSig            = errors.New("sig error")
	MapDescription = map[error]string{
		Success:                         "Success!",
		ErrUnknown:                      "Unknown error!",
		UnAuthorized:                    "UnAuthorized",
		BadRequest:                      "Bad Request!",
		ErrNotInt:                       "Param not int!",
		UnSuccess:                       "UnSuccess!",
		VerifyCodeExpired:               "Verify Code Expired",
		SessionExpired:                  "SessionExpired!",
		Existed:                         "Existed !",
		NotAdmin:                        "Not Admin !",
		NotPermission:                   "Not Permission !",
		ErrLogin:                        "Wrong username, password. ",
		ErrSystem:                       "The system is having problems.",
		NotExisted:                      "Data Not Existed!",
		ErrChangePass:                   "Wrong username, password.",
		NotMore:                         "No more data",
		ErrSig:                          "Sig error !",
		CannotEmpty:                     "Data empty !",
	}
	MapErrorCode = map[error]int64{
		Success:                         200,
		UnSuccess:                       201,
		ErrNotInt:                       302,
		SessionExpired:                  303,
		NotExisted:                      304,
		Existed:                         305,
		ErrChangePass:                   306,
		NotAdmin:                        307,
		NotPermission:                   308,
		NotMore:                         309,
		CannotEmpty:                     311,
		ErrSig:                          316,
		BadRequest:                      400,
		ErrUnknown:                      401,
		ErrLogin:                        402,
		ErrSystem:                       403,
		UnAuthorized:                    405,
		VerifyCodeExpired:               406,

	}
)

// Returns a error.
// swagger:response Err
type Err struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message string `json:"message"`
}

type ValidateErr struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message map[string]string `json:"message"`
}

func NewErr(err error) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: MapDescription[err],
	}
}

func NewErrByText(err error, text string) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: text,
	}
}
