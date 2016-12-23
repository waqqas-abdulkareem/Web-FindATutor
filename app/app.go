package app

import(
	"net/http"
	"github.com/jinzhu/gorm"
)

type Context struct{
	DB	 *gorm.DB
	Port string
	Host string
	Environment string
}

type Error struct{
	Message string
	Code int
}

func NewBadRequest(err error) *Error{
	return &Error{Code: http.StatusBadRequest, Message: err.Error()}
}

func NotFound(err error) *Error{
	return &Error{Code: http.StatusNotFound, Message: err.Error()}
}

func NewInternalServerError(err error) *Error{
	return &Error{Code: http.StatusInternalServerError, Message: err.Error()}	
}

func (e Error) Error() string{
	return e.Message
}

type Handler struct{
	*Context
	HandleFunc func(*Context, http.ResponseWriter, *http.Request) *Error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if err := h.HandleFunc(h.Context,w,r); err != nil{
		SendAppErrorEncoded(err,w,r)
	}
}

