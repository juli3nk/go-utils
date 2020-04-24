package tracing

import (
	"fmt"
	"io"

	"github.com/juli3nk/go-utils"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegerprom "github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewTracer(logger *log.Entry) (opentracing.Tracer, io.Closer, error) {
	serviceName, err := utils.GetEnv("JAEGER_SERVICE_NAME")
	if err != nil {
		return nil, nil, err
	}
	agentHost := utils.GetEnvDefault("JAEGER_AGENT_HOST", "localhost")
	agentPort := utils.GetEnvDefault("JAEGER_AGENT_PORT", "6831")

	agentURI := fmt.Sprintf("%s:%s", agentHost, agentPort)

	factory := jaegerprom.New()
	metrics := jaeger.NewMetrics(factory, map[string]string{"lib": "jaeger"})

	transport, err := jaeger.NewUDPTransport(agentURI, 0)
	if err != nil {
		return nil, nil, err
	}

	logAdapt := NewLogger(logger)

	reporter := jaeger.NewCompositeReporter(
		jaeger.NewLoggingReporter(logAdapt),
		jaeger.NewRemoteReporter(
			transport,
			jaeger.ReporterOptions.Metrics(metrics),
			jaeger.ReporterOptions.Logger(logAdapt),
		),
	)

	sampler := jaeger.NewConstSampler(true)

	tracer, closer := jaeger.NewTracer(serviceName,
		sampler,
		reporter,
		jaeger.TracerOptions.Metrics(metrics),
	)

	return tracer, closer, nil
}
