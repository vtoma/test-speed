package speed

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SpeedTestProviderMocked struct {
	mock.Mock
}

func (r *SpeedTestProviderMocked) GetStats() (float64, float64) {
	return 1111, 2222
}

func CreateNewSpeedTest() *SpeedTester {
	providers := make(map[string]SpeedInterface)
	providers["speedtest"] = &SpeedTestProviderMocked{}
	providers["fast"] = &SpeedTestProviderMocked{}

	return &SpeedTester{providers}
}

func TestNonExistingProvider(t *testing.T) {
	speedTester := CreateNewSpeedTest()

	provider := "nonExistingProvider"
	assert.PanicsWithValue(t, "Provider not found: "+provider, func() { speedTester.getProvider(provider) })
}

func TestExistingProvider(t *testing.T) {
	speedTester := CreateNewSpeedTest()

	provider := "fast"
	assert.NotPanics(t, func() { speedTester.getProvider(provider) })
}

func Test_speedTestHandler(t *testing.T) {
	speedTester := CreateNewSpeedTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	speedTester.speedTestHandler(res, req)
	assert.Equal(t, "Download: 1111.00, Upload: 2222.00", res.Body.String())
}

func Test_FastHandler(t *testing.T) {
	speedTester := CreateNewSpeedTest()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	speedTester.fastHandler(res, req)
	assert.Equal(t, "Download: 1111.00, Upload: 2222.00", res.Body.String())
}

func TestNewSpeedTest(t *testing.T) {
	speedTester := NewSpeedTest()
	assert.NotNil(t, speedTester)
}

var downloadCache, uploadCache float64

func BenchmarkSpeedTest(b *testing.B) {
	speedTestObj := NewSpeedTest()
	for n := 0; n < b.N; n++ {
		downloadCache, uploadCache = speedTestObj.getProvider("speedtest").GetStats()
	}
}

func BenchmarkFast(b *testing.B) {
	speedTestObj := NewSpeedTest()
	for n := 0; n < b.N; n++ {
		downloadCache, uploadCache = speedTestObj.getProvider("fast").GetStats()
	}
}
