// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	"go.opentelemetry.io/collector/component"
	loggingexporter "go.opentelemetry.io/collector/exporter/loggingexporter"
	otlpexporter "go.opentelemetry.io/collector/exporter/otlpexporter"
	otlphttpexporter "go.opentelemetry.io/collector/exporter/otlphttpexporter"
	jaegerexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/jaegerexporter"
	zpagesextension "go.opentelemetry.io/collector/extension/zpagesextension"
	ballastextension "go.opentelemetry.io/collector/extension/ballastextension"
	healthcheckextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	pprofextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	oauth2clientauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"
	bearertokenauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
	basicauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	memorylimiterprocessor "go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	attributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	resourceprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	spanprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	k8sattributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	resourcedetectionprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	filterprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	otlpreceiver "go.opentelemetry.io/collector/receiver/otlpreceiver"
	jaegerreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	opencensusreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver"
	zipkinreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
)

func components() (component.Factories, error) {
	var err error
	factories := component.Factories{}

	factories.Extensions, err = component.MakeExtensionFactoryMap(
		zpagesextension.NewFactory(),
		ballastextension.NewFactory(),
		healthcheckextension.NewFactory(),
		pprofextension.NewFactory(),
		oauth2clientauthextension.NewFactory(),
		bearertokenauthextension.NewFactory(),
		basicauthextension.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	factories.Receivers, err = component.MakeReceiverFactoryMap(
		otlpreceiver.NewFactory(),
		jaegerreceiver.NewFactory(),
		opencensusreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	factories.Exporters, err = component.MakeExporterFactoryMap(
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		jaegerexporter.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	factories.Processors, err = component.MakeProcessorFactoryMap(
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		attributesprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		k8sattributesprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
		filterprocessor.NewFactory(),
	)
	if err != nil {
		return component.Factories{}, err
	}

	return factories, nil
}
