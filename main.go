package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/maliaga/egrh_watcher/src"
	gonanoid "github.com/matoous/go-nanoid/v2"
	logger "github.com/sirupsen/logrus"
)

func init() {
	logger.SetFormatter(&logger.JSONFormatter{})
}
func main() {

	id, idErr := gonanoid.New()
	if idErr != nil {
		logger.Fatalf("Cant generate id")
	}
	resp, err := http.Get(os.Getenv("SERVICE_URL"))
	if err != nil {
		logger.WithField("id", id).Fatalf("Cant reach egrh service")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.WithField("id", id).Fatalf("cant parse service response")
	}
	data := &src.Response{}
	parsedErr := json.Unmarshal(body, data)
	if parsedErr != nil {
		logger.WithField("id", id).Fatalf("cant unmarshall body")
	}
	serviceTime := time.Unix(data.Egrh.CreatedTime, 0)
	logger.WithField("id", id).Infof("service time is %s", serviceTime.String())
	now := time.Now()
	timeDiff := math.Round(now.Sub(serviceTime).Minutes())
	logger.WithField("id", id).Infof("timediff is %d minutes", int(timeDiff))
	if timeDiff > 70 {
		logger.WithField("id", id).Fatal("Egrh is outdated")
	} else {
		logger.WithField("id", id).Infof("Egrh is ok")
	}
}
