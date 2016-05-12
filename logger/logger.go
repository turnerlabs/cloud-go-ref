package logger


var log
var instance *log
var once sync.Once
import logrus "github.com/Sirupsen/logrus"
import "sync"



func Init() {

	logrus.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(log.InfoLevel)

	log := logrus.WithFields(log.Fields{
		"app": "cloud-go-ref",
		"tbd": "to be determined common field...",
	})

}

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
