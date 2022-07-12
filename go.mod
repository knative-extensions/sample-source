module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.10.1
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.21.0
	k8s.io/api v0.23.8
	k8s.io/apimachinery v0.23.8
	k8s.io/client-go v0.23.8
	knative.dev/eventing v0.32.1-0.20220712162113-0eb566ae2067
	knative.dev/hack v0.0.0-20220701014203-65c463ac8c98
	knative.dev/pkg v0.0.0-20220705130606-e60d250dc637
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
