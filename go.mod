module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	knative.dev/eventing v0.31.1-0.20220516052256-d7a8a95792bd
	knative.dev/hack v0.0.0-20220512014059-f4972b4daff9
	knative.dev/pkg v0.0.0-20220512013937-2d8305b2e59a
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
