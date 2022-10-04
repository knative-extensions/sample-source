module knative.dev/sample-source

go 1.14

require (
	github.com/cloudevents/sdk-go/sql/v2 v2.12.0 // indirect
	github.com/cloudevents/sdk-go/v2 v2.12.0
	github.com/google/go-cmp v0.5.6
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.8.0
	go.uber.org/zap v1.21.0
	k8s.io/api v0.24.4
	k8s.io/apimachinery v0.24.4
	k8s.io/client-go v0.24.4
	knative.dev/eventing v0.34.1-0.20221004074129-91e8838ce1e1
	knative.dev/hack v0.0.0-20221003111221-b3c97904d3b5
	knative.dev/pkg v0.0.0-20221003153827-158538cc46ec
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
