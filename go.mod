module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
	k8s.io/api v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/eventing v0.26.1-0.20211028061727-2eb07cbddd07
	knative.dev/hack v0.0.0-20211027200727-f1228dd5e3e6
	knative.dev/pkg v0.0.0-20211027171921-f7b70f5ce303
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
