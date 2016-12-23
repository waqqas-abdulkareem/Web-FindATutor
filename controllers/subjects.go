package controller

import(
	"io/ioutil"
	"encoding/json"
	_"log"
	"net/http"
	"strings"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/waqqas-abdulkareem/find_a_tutor/app"
	"github.com/waqqas-abdulkareem/find_a_tutor/models"
)

func GetSubject(c *app.Context, w http.ResponseWriter, r * http.Request) *app.Error{
	params := mux.Vars(r)
	id := params["id"]

	var subject model.Subject
	err := c.DB.Find(&subject,id).Error
	
	if err != nil{
		return app.NotFound(err)
	}
	app.SendDataEncoded(http.StatusOK, &subject,w,r)
	return nil
}

func GetSubjects(c *app.Context, w http.ResponseWriter, r * http.Request) *app.Error{
	
	limit := 10
	offset := 0
	var err error
	
	if _limit := r.FormValue("limit"); len(_limit) > 0{
		limit,err = strconv.Atoi(_limit)
	}
	if _offset := r.FormValue("offset"); len(_offset) > 0{
		offset,err = strconv.Atoi(_offset)
	}
	if err != nil{
		return app.NewInternalServerError(err)
	}

	var subjects []model.Subject
	err = c.DB.Limit(limit).Offset(offset).Find(&subjects).Error
	if err != nil{
		return app.NewInternalServerError(err)
	}
	app.SendDataEncoded(http.StatusOK,&subjects,w,r)
	return nil
}

func PostSubject(c *app.Context, w http.ResponseWriter, r * http.Request) *app.Error{
	subject,appErr := subjectFromBody(r)
	if appErr != nil{
		return appErr;
	}
	subject.Name = strings.Title(subject.Name)
	err := c.DB.Create(&subject).Error
	if err != nil{
		return app.NewInternalServerError(err)
	}
	
	app.SendDataEncoded(http.StatusOK,&subject,w,r)	
	return nil
}

func subjectFromBody(r * http.Request) (*model.Subject,*app.Error){
	
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return nil,app.NewBadRequest(err)
	}
	
	var subject model.Subject
	err = json.Unmarshal(bytes,&subject)
	if err != nil{
		return nil,app.NewBadRequest(err)
	}

	return &subject,nil
}
