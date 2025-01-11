package logger

import (
	"log"
	"os"
)

// var ok *log.Logger

// func SetLogger(byy *log.Logger){
// 	ok=byy
// }

var Ok *log.Logger

func GetLogger(logger_name string) (*log.Logger, error) {
	Regfile, err := os.OpenFile("../logger/"+logger_name+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return nil, err
	}
	// defer Regfile.Close()

	regerror := log.New(Regfile, "Register Error", log.Ldate|log.Ltime)
	return regerror, nil
}
