package helper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	cronowriter "github.com/utahta/go-cronowriter"
	"golang.org/x/crypto/bcrypt"
)

// LogInfo write info
func LogInfo(log string) {
	w1 := cronowriter.MustNew("./tmp/example.log.%Y%m%d")
	w1.Write([]byte(log + "\n"))
}

// LogErr write info
func LogErr(log string) {
	w2 := cronowriter.MustNew("./tmp/internal_error.log.%Y%m%d")
	w2.Write([]byte(log + "\n"))
}

// HashPassword create pass
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare pass
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// StringToint64 helper
func StringToint64(num string) int64 {
	n, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

// int64ToString helper
func int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// ErrResponse struct
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render func
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest func
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound const
var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

//ErrForbiden const
var ErrForbiden = &ErrResponse{HTTPStatusCode: http.StatusForbidden, StatusText: "You are not permitted"}
