module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/v2 v2.10.1
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.8.0
	go.uber.org/zap v1.21.0
	k8s.io/api v0.24.4
	k8s.io/apimachinery v0.24.4
	k8s.io/client-go v0.24.4
	knative.dev/eventing v0.34.1-0.20220928014117-bd37622589a4
	knative.dev/hack v0.0.0-20220928014218-68872175ea9c
	knative.dev/pkg v0.0.0-20220928014015-6d718681cd8c
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
