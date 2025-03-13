package generate
import (
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
		UpstreamLocalAddress:           "192.168.1.1:5678",
		UserAgent:                      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		RequestedServerName:            "example.com",
		ResponseCodeDetails:            "details",
		Path:                           "/api/v1/data",
		XDatadogParentID:               "parent_id",
		RequestID:                      "request_id",
		Authority:                      "authority",
		BytesReceived:                  1024,
		DownstreamRemoteAddress:        "192.168.1.2:8080",
		BytesSent:                      512,
		XEnvoyAttemptCount:             "1",
		Protocol:                       "HTTP/1.1",
		Method:                         "GET",
		XForwardedFor:                  "192.168.1.3",
		DownstreamLocalAddress:         "192.168.1.4:80",
		StartTime:                      time.Now(),
		TraceID:                        "trace_id",
		XDatadogTraceID:                "datadog_trace_id",
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