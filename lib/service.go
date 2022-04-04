package speed

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type SpeedTester struct {
	providers map[string]SpeedInterface
}

func (r *SpeedTester) getProvider(provider string) SpeedInterface {
	if _, ok := r.providers[provider]; !ok {
		logrus.Error("Provider not found: ", provider)
		panic("Provider not found: " + provider)
	}
	return r.providers[provider]
}

func (r *SpeedTester) speedTestHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: speedtest")

	download, upload := r.getProvider("speedtest").GetStats()
	logrus.Info("Download: ", download, "Upload: ", upload)

	// send response
	fmt.Fprintf(w, "Download: %.2f, Upload: %.2f", download, upload)

}

func (r *SpeedTester) fastHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Endpoint Hit: fast")

	download, upload := r.getProvider("fast").GetStats()
	logrus.Info("Download: ", download, "Upload: ", upload)

	// send response
	fmt.Fprintf(w, "Download: %.2f, Upload: %.2f", download, upload)
}

// initializes the handlers for the speed test
func (r *SpeedTester) InitHandlers() {
	http.HandleFunc("/stats/speedtest", r.speedTestHandler)
	http.HandleFunc("/stats/fast", r.fastHandler)
}

// creates an instance of the speed tester
func NewSpeedTest() *SpeedTester {
	// adding the providers
	providers := make(map[string]SpeedInterface)
	providers["speedtest"] = &SpeedTestProvider{}
	providers["fast"] = &FastProvider{}

	return &SpeedTester{providers}
}
