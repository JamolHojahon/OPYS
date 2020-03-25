package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	ao "github.com/OPYS/internal/app/authentication"

	"github.com/OPYS/internal/pkg/databaseinit"
	log "github.com/sirupsen/logrus"
)

func main() {

	initLogging()

	log.Debug("Trying to initializa db connection!")
	ConVar, err := databaseinit.InitDataBase()
	defer databaseinit.Disconnect()
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	cRepo := ao.InitRepository(ConVar)
	cSrv := ao.InitService(cRepo)
	cContrl := ao.InitControllers(cSrv)

	router := http.NewServeMux()

	/* JSON request for signup
		{
	 		"Email":"myemail@gmail.com",
	  		"Password":"123456789",
	  		"ConfirmPassword":"123456789",
	  		"Claims": [
	    	{
	      		"type":"firstname",
		      	"value":"Jane"
	    	},
	    	{
	      		"type":"lastname",
	      		"value":"Doe"
	    	},
	    	{
	      		"type":"birthdate",
	      		"value":"1999.12.31"
	    	}]
		}
	*/
	router.HandleFunc("/signup/", cContrl.SignUp())

	/*JSON req for /signin/
	{
		"Email":"myemail@example.com",
	  	"Password":"123456789"
	}
	*/
	router.HandleFunc("/signin/", cContrl.SignIn())
	// router.HandleFunc("userpage/")

	log.Info("Starting http server...")
	http.ListenAndServe(":80", router)

}

func initLogging() {

	var file, err = os.OpenFile("../../logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetLevel(log.DebugLevel)

	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})

}
