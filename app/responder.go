package app

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/xml"
)

func SendErrorEncoded(code int,err error, w http.ResponseWriter, r * http.Request){
	SendAppErrorEncoded(&Error{Code: code, Message: err.Error()},w,r)		
}

func SendAppErrorEncoded(err *Error, w http.ResponseWriter, r * http.Request){
	SendDataEncoded(err.Code,struct{Message string}{err.Message},w,r)
}

func SendDataEncoded(code int, data interface{},w http.ResponseWriter, r * http.Request){
		
	log.Printf("Data: %v",data)
	encodingType := "application/json"

	var bytes []byte
	var err error
	switch encodingType{
		case "application/xml":
			bytes,err = xml.Marshal(data)  
		case "application/json":
			fallthrough
		default:
			bytes,err = json.Marshal(data)
	}

	log.Printf("encodingType = %s\n",encodingType)
	log.Printf("bytes = %s\v",string(bytes))
	log.Printf("err = %v\n",err)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,err.Error())
		return
	}

	w.Header().Set("Content-Type",encodingType)
	w.WriteHeader(code)
	fmt.Fprint(w,string(bytes))	
}
