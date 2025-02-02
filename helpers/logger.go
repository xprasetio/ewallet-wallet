package helpers

import "github.com/sirupsen/logrus"

var Loggger *logrus.Logger
func SetupLogger() {	

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.Info("logger initialized using logrus") 
	Loggger = log
}