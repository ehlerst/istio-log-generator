package generate
import (
	"fmt"
	"github.com/ehlerst/istio-log-generator/models"
	"time"
	"math/rand"
	"net"
)

var (
	codes = []int{200, 401, 404, 500, 503}
)

func GenerateIstioLog() models.IstioLog {
	log := models.IstioLog{
		ResponseFlags:                  "OK",
		UpstreamLocalAddress:           GenerateRandomIP(),
		UserAgent:                      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		RequestedServerName:            "example.com",
		ResponseCodeDetails:            "details",
		Path:                           "/api/v1/data",
		XDatadogParentID:               "parent_id",
		RequestID:                      "request_id",
		Authority:                      fmt.Sprintf("%x", rand.Int63()),
		BytesReceived:                  rand.Intn(8096) + 512,
		DownstreamRemoteAddress:        GenerateRandomIP(),
		BytesSent:                      rand.Intn(8096) + 512,
		XEnvoyAttemptCount:             "1",
		Protocol:                       "HTTP/1.1",
		Method:                         "GET",
		XForwardedFor:                  GenerateRandomIP(),
		DownstreamLocalAddress:         GenerateRandomIP(),
		StartTime:                      time.Now(),
		TraceID:                        fmt.Sprintf("%x", rand.Int63()),
		XDatadogTraceID:                fmt.Sprintf("%x", rand.Int63()),
		ResponseCode:	                GenerateRandomCode(),
		Duration:                       rand.Intn(500) + 100,
		UpstreamServiceTime:            fmt.Sprintf("%d", rand.Intn(500) + 100),
		UpstreamCluster: fmt.Sprintf("%x", rand.Int63()),
	}
	return log

}

func GenerateRandomCode() int {
    randomIndex := rand.Intn(len(codes))
    return codes[randomIndex]
}

func GenerateRandomIP() string {
	ip := net.IP(make([]byte, 4))
	rand.Read(ip)
	return ip.String()

}