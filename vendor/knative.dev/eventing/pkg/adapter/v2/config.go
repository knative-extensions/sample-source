/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package adapter

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"k8s.io/apimachinery/pkg/types"

	"go.uber.org/zap"

	"knative.dev/eventing/pkg/observability"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	kle "knative.dev/pkg/leaderelection"
	"knative.dev/pkg/logging"
)

type EnvConfigConstructor func() EnvConfigAccessor

const (
	EnvConfigComponent            = "K_COMPONENT"
	EnvConfigNamespace            = "NAMESPACE"
	EnvConfigName                 = "NAME"
	EnvConfigResourceGroup        = "K_RESOURCE_GROUP"
	EnvConfigSink                 = "K_SINK"
	EnvConfigAudience             = "K_AUDIENCE"
	EnvConfigOIDCServiceAccount   = "K_OIDC_SERVICE_ACCOUNT"
	EnvConfigCACert               = "K_CA_CERTS"
	EnvConfigCEOverrides          = "K_CE_OVERRIDES"
	EnvConfigLoggingConfig        = "K_LOGGING_CONFIG"
	EnvConfigObservabilityConfig  = "K_OBSERVABILITY_CONFIG"
	EnvConfigLeaderElectionConfig = "K_LEADER_ELECTION_CONFIG"
	EnvSinkTimeout                = "K_SINK_TIMEOUT"
)

// EnvConfig is the minimal set of configuration parameters
// source adapters should support.
type EnvConfig struct {
	// Component is the kind of this adapter.
	Component string `envconfig:"K_COMPONENT"`

	// Environment variable containing the namespace of the adapter.
	Namespace string `envconfig:"NAMESPACE"`

	// Environment variable containing the name of the adapter.
	Name string `envconfig:"NAME" default:"adapter"`

	// Environment variable containing the resource group of the adapter for metrics.
	ResourceGroup string `envconfig:"K_RESOURCE_GROUP" default:"adapter.sources.knative.dev"`

	// Sink is the URI messages will be sent.
	Sink string `envconfig:"K_SINK"`

	// Audience is the audience of the target sink.
	Audience *string `envconfig:"K_AUDIENCE"`

	// OIDCServiceAccount Name is the name of the service account to use for the adapter.
	OIDCServiceAccountName *string `envconfig:"K_OIDC_SERVICE_ACCOUNT"`

	// CACerts are the Certification Authority (CA) certificates in PEM format
	// according to https://www.rfc-editor.org/rfc/rfc7468.
	// +optional
	CACerts *string `envconfig:"K_CA_CERTS"`

	// CEOverrides are the CloudEvents overrides to be applied to the outbound event.
	CEOverrides string `envconfig:"K_CE_OVERRIDES"`

	// LoggingConfigJson is a json string of logging.Config.
	// This is used to configure the logging config, the config is stored in
	// a config map inside the controllers namespace and copied here.
	LoggingConfigJson string `envconfig:"K_LOGGING_CONFIG" default:"{}"`

	// ObservabilityConfigJson is a json string of observability.Config.
	// This is used to configure the observability config, the config is stored in
	// a config map inside the controllers namespace and copied here.
	ObservabilityConfigJson string `envconfig:"K_OBSERVABILITY_CONFIG" default:"{}"`

	// LeaderElectionConfigJson is the leader election component configuration.
	LeaderElectionConfigJson string `envconfig:"K_LEADER_ELECTION_CONFIG"`

	// Time in seconds to wait for sink to respond
	EnvSinkTimeout string `envconfig:"K_SINK_TIMEOUT"`

	// cached zap logger
	logger *zap.SugaredLogger
}

// EnvConfigAccessor defines accessors for the minimal
// set of source adapter configuration parameters.
type EnvConfigAccessor interface {
	// Set the component name.
	SetComponent(string)

	// Get the URI where messages will be forwarded to.
	GetSink() string

	// GetCACerts gets the CACerts of the Sink.
	GetCACerts() *string

	// GetAudience gets the audience of the target sink.
	GetAudience() *string

	// GetOIDCServiceAccountName gets the service account name to use for the adapter.
	GetOIDCServiceAccountName() *types.NamespacedName

	// Get the namespace of the adapter.
	GetNamespace() string

	// Get the name of the adapter.
	GetName() string

	// Get the parsed logger.
	GetLogger() *zap.SugaredLogger

	GetCloudEventOverrides() (*duckv1.CloudEventOverrides, error)

	GetObservabilityConfig() (*observability.Config, error)

	// GetLeaderElectionConfig returns leader election configuration.
	GetLeaderElectionConfig() (*kle.ComponentConfig, error)

	// Get the timeout to apply on a request to a sink
	GetSinktimeout() int
}

var _ EnvConfigAccessor = (*EnvConfig)(nil)

func (e *EnvConfig) SetComponent(component string) {
	e.Component = component
}

func (e *EnvConfig) GetLogger() *zap.SugaredLogger {
	if e.logger == nil {
		loggingConfig, err := logging.JSONToConfig(e.LoggingConfigJson)
		if err != nil {
			// Use default logging config.
			if loggingConfig, err = logging.NewConfigFromMap(map[string]string{}); err != nil {
				// If this fails, there is no recovering.
				panic(err)
			}
		}

		logger, _ := logging.NewLoggerFromConfig(loggingConfig, e.Component)
		e.logger = logger
	}
	return e.logger
}

func (e *EnvConfig) GetSink() string {
	return e.Sink
}

func (e *EnvConfig) GetOIDCServiceAccountName() *types.NamespacedName {
	if e.OIDCServiceAccountName != nil {
		return &types.NamespacedName{
			Namespace: e.Namespace,
			Name:      *e.OIDCServiceAccountName,
		}
	}
	return nil
}

func (e *EnvConfig) GetCACerts() *string {
	return e.CACerts
}

func (e *EnvConfig) GetAudience() *string {
	return e.Audience
}

func (e *EnvConfig) GetNamespace() string {
	return e.Namespace
}

func (e *EnvConfig) GetName() string {
	return e.Name
}

func (e *EnvConfig) GetSinktimeout() int {
	if duration, err := strconv.Atoi(e.EnvSinkTimeout); err == nil {
		return duration
	}
	e.GetLogger().Warn("Sink timeout configuration is invalid, default to -1 (no timeout)")
	return -1
}

func (e *EnvConfig) GetObservabilityConfig() (*observability.Config, error) {
	cfg := &observability.Config{}
	err := json.Unmarshal([]byte(e.ObservabilityConfigJson), cfg)
	return cfg, err
}

func (e *EnvConfig) GetCloudEventOverrides() (*duckv1.CloudEventOverrides, error) {
	var ceOverrides duckv1.CloudEventOverrides
	if len(e.CEOverrides) > 0 {
		err := json.Unmarshal([]byte(e.CEOverrides), &ceOverrides)
		if err != nil {
			return nil, err
		}
	}
	return &ceOverrides, nil
}

func (e *EnvConfig) GetLeaderElectionConfig() (*kle.ComponentConfig, error) {
	if e.LeaderElectionConfigJson == "" {
		return e.defaultLeaderElectionConfig(), nil
	}

	var config kle.ComponentConfig
	if err := json.Unmarshal([]byte(e.LeaderElectionConfigJson), &config); err != nil {
		return e.defaultLeaderElectionConfig(), err
	}
	config.Component = e.Component
	return &config, nil
}

func (e *EnvConfig) defaultLeaderElectionConfig() *kle.ComponentConfig {
	return &kle.ComponentConfig{
		Component:     e.Component,
		Buckets:       1,
		LeaseDuration: 15 * time.Second,
		RenewDeadline: 10 * time.Second,
		RetryPeriod:   2 * time.Second,
	}
}

// LeaderElectionComponentConfigToJSON converts a ComponentConfig to a json string.
func LeaderElectionComponentConfigToJSON(cfg *kle.ComponentConfig) (string, error) {
	if cfg == nil {
		return "", nil
	}

	jsonCfg, err := json.Marshal(cfg)
	return string(jsonCfg), err
}

func GetSinkTimeout(logger *zap.SugaredLogger) int {
	str := os.Getenv(EnvSinkTimeout)
	if str != "" {
		var err error
		duration, err := strconv.Atoi(str)
		if err != nil || duration < 0 {
			if logger != nil {
				logger.Errorf("%s environment value is invalid. It must be a integer greater than zero. (got %s)", EnvSinkTimeout, str)
			}
			return -1
		}
		return duration
	}
	return -1
}
