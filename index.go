package main

import(
	"flag"
	"io/ioutil"
	"encoding/json"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/waqqas-abdulkareem/find_a_tutor/controllers"
	"github.com/waqqas-abdulkareem/find_a_tutor/models"
	"github.com/waqqas-abdulkareem/find_a_tutor/app"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var config *app.Configuration
var ConfigurationFileName string

func init(){
	initConfig()
	initDB();
}

func initConfig(){
	ConfigurationFileName = *flag.String(
		"Configuration",
		"Configuration.json",
		"path to json file containing configuration",
	);
	
	var err error
	config, err = ReadConfig()
	if err != nil{
		log.Fatal(err)
	}
}

func ReadConfig() (*app.Configuration,error){
	bytes, err := ioutil.ReadFile(ConfigurationFileName)
	if err != nil{
		return nil,err
	}

	var config app.Configuration
	err = json.Unmarshal(bytes,&config)
	if err != nil{
		return nil,err
	}
	return &config,nil;
}

func initDB(){
	var err error;
	db, err = gorm.Open(
		config.DBConfig.DriverName, 
		config.DBConfig.DriverSourceName,
	)
	if err != nil{
		log.Fatal(err)
	}

	MigrateDB();

}

func MigrateDB(){
	db.AutoMigrate(&model.Subject{})	
}

func main(){
	defer db.Close()

	router := mux.NewRouter();	
	context := &app.Context{
		DB:	  db,
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	};
	apiRouter := router.PathPrefix("/api").Subrouter();	
	
	//subjects
	apiRouter.Handle("/subjects/",app.Handler{
		Context: context,
		HandleFunc: controller.GetSubjects,
	}).Methods("GET")

	apiRouter.Handle("/subjects/{id}/",app.Handler{
		Context: context,
		HandleFunc: controller.GetSubject,
	}).Methods("GET")
	
	apiRouter.Handle("/subjects/",app.Handler{
		Context: context,
		HandleFunc: controller.PostSubject,
	}).Methods("POST")

	log.Printf("Listening...%s:%s",context.Host,context.Port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}


