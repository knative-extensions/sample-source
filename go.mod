module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.23.4
	k8s.io/apimachinery v0.23.4
	k8s.io/client-go v0.23.4
	knative.dev/eventing v0.30.1-0.20220318182120-5df47e13e70b
	knative.dev/hack v0.0.0-20220318020218-14f832e506f8
	knative.dev/pkg v0.0.0-20220318185521-e6e3cf03d765
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
