package models

import "time"

type IstioLog struct {
	ResponseFlags                  string    `json:"response_flags"`
	UpstreamLocalAddress           string    `json:"upstream_local_address"`
	UserAgent                      string    `json:"user_agent"`
	RequestedServerName            string    `json:"requested_server_name"`
	ResponseCodeDetails            string    `json:"response_code_details"`
	Path                           string    `json:"path"`
	XDatadogParentID               string    `json:"x_datadog_parent_id"`
	RequestID                      string    `json:"request_id"`
	Authority                      string    `json:"authority"`
	BytesReceived                  int       `json:"bytes_received"`
	DownstreamRemoteAddress        string    `json:"downstream_remote_address"`
	BytesSent                      int       `json:"bytes_sent"`
	XEnvoyAttemptCount             string    `json:"x_envoy_attempt_count"`
	Protocol                       string    `json:"protocol"`
	Method                         string    `json:"method"`
	XForwardedFor                  string    `json:"x_forwarded_for"`
	DownstreamLocalAddress         string    `json:"downstream_local_address"`
	StartTime                      time.Time `json:"start_time"`
	TraceID                        string    `json:"trace_id"`
	XDatadogTraceID                string    `json:"x_datadog_trace_id"`
	UpstreamCluster                string    `json:"upstream_cluster"`
	RouteName                      any       `json:"route_name"`
	Duration                       int       `json:"duration"`
	UpstreamTransportFailureReason any       `json:"upstream_transport_failure_reason"`
	UpstreamHost                   string    `json:"upstream_host"`
	ConnectionTerminationDetails   any       `json:"connection_termination_details"`
	UpstreamServiceTime            string    `json:"upstream_service_time"`
	ResponseCode                   int       `json:"response_code"`
}