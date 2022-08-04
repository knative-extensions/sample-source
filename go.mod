module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.10.1
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.21.0
	k8s.io/api v0.23.9
	k8s.io/apimachinery v0.23.9
	k8s.io/client-go v0.23.9
	knative.dev/eventing v0.33.1-0.20220804105621-840fce727af9
	knative.dev/hack v0.0.0-20220728013938-9dabf7cf62e3
	knative.dev/pkg v0.0.0-20220802185824-a01dfedb0486
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
