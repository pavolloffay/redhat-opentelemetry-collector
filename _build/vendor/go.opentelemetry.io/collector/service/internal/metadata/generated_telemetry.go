// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("go.opentelemetry.io/collector/service")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("go.opentelemetry.io/collector/service")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                                    metric.Meter
	ProcessCPUSeconds                        metric.Float64ObservableCounter
	observeProcessCPUSeconds                 func(context.Context, metric.Observer) error
	ProcessMemoryRss                         metric.Int64ObservableGauge
	observeProcessMemoryRss                  func(context.Context, metric.Observer) error
	ProcessRuntimeHeapAllocBytes             metric.Int64ObservableGauge
	observeProcessRuntimeHeapAllocBytes      func(context.Context, metric.Observer) error
	ProcessRuntimeTotalAllocBytes            metric.Int64ObservableCounter
	observeProcessRuntimeTotalAllocBytes     func(context.Context, metric.Observer) error
	ProcessRuntimeTotalSysMemoryBytes        metric.Int64ObservableGauge
	observeProcessRuntimeTotalSysMemoryBytes func(context.Context, metric.Observer) error
	ProcessUptime                            metric.Float64ObservableCounter
	observeProcessUptime                     func(context.Context, metric.Observer) error
	level                                    configtelemetry.Level
}

// telemetryBuilderOption applies changes to default builder.
type telemetryBuilderOption func(*TelemetryBuilder)

// WithLevel sets the current telemetry level for the component.
func WithLevel(lvl configtelemetry.Level) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.level = lvl
	}
}

// WithProcessCPUSecondsCallback sets callback for observable ProcessCPUSeconds metric.
func WithProcessCPUSecondsCallback(cb func() float64, opts ...metric.ObserveOption) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.observeProcessCPUSeconds = func(_ context.Context, o metric.Observer) error {
			o.ObserveFloat64(builder.ProcessCPUSeconds, cb(), opts...)
			return nil
		}
	}
}

// WithProcessMemoryRssCallback sets callback for observable ProcessMemoryRss metric.
func WithProcessMemoryRssCallback(cb func() int64, opts ...metric.ObserveOption) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.observeProcessMemoryRss = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.ProcessMemoryRss, cb(), opts...)
			return nil
		}
	}
}

// WithProcessRuntimeHeapAllocBytesCallback sets callback for observable ProcessRuntimeHeapAllocBytes metric.
func WithProcessRuntimeHeapAllocBytesCallback(cb func() int64, opts ...metric.ObserveOption) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.observeProcessRuntimeHeapAllocBytes = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.ProcessRuntimeHeapAllocBytes, cb(), opts...)
			return nil
		}
	}
}

// WithProcessRuntimeTotalAllocBytesCallback sets callback for observable ProcessRuntimeTotalAllocBytes metric.
func WithProcessRuntimeTotalAllocBytesCallback(cb func() int64, opts ...metric.ObserveOption) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.observeProcessRuntimeTotalAllocBytes = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.ProcessRuntimeTotalAllocBytes, cb(), opts...)
			return nil
		}
	}
}

// WithProcessRuntimeTotalSysMemoryBytesCallback sets callback for observable ProcessRuntimeTotalSysMemoryBytes metric.
func WithProcessRuntimeTotalSysMemoryBytesCallback(cb func() int64, opts ...metric.ObserveOption) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.observeProcessRuntimeTotalSysMemoryBytes = func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(builder.ProcessRuntimeTotalSysMemoryBytes, cb(), opts...)
			return nil
		}
	}
}

// WithProcessUptimeCallback sets callback for observable ProcessUptime metric.
func WithProcessUptimeCallback(cb func() float64, opts ...metric.ObserveOption) telemetryBuilderOption {
	return func(builder *TelemetryBuilder) {
		builder.observeProcessUptime = func(_ context.Context, o metric.Observer) error {
			o.ObserveFloat64(builder.ProcessUptime, cb(), opts...)
			return nil
		}
	}
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...telemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{level: configtelemetry.LevelBasic}
	for _, op := range options {
		op(&builder)
	}
	var err, errs error
	if builder.level >= configtelemetry.LevelBasic {
		builder.meter = Meter(settings)
	} else {
		builder.meter = noop.Meter{}
	}
	builder.ProcessCPUSeconds, err = builder.meter.Float64ObservableCounter(
		"otelcol_process_cpu_seconds",
		metric.WithDescription("Total CPU user and system time in seconds"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meter.RegisterCallback(builder.observeProcessCPUSeconds, builder.ProcessCPUSeconds)
	errs = errors.Join(errs, err)
	builder.ProcessMemoryRss, err = builder.meter.Int64ObservableGauge(
		"otelcol_process_memory_rss",
		metric.WithDescription("Total physical memory (resident set size)"),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meter.RegisterCallback(builder.observeProcessMemoryRss, builder.ProcessMemoryRss)
	errs = errors.Join(errs, err)
	builder.ProcessRuntimeHeapAllocBytes, err = builder.meter.Int64ObservableGauge(
		"otelcol_process_runtime_heap_alloc_bytes",
		metric.WithDescription("Bytes of allocated heap objects (see 'go doc runtime.MemStats.HeapAlloc')"),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meter.RegisterCallback(builder.observeProcessRuntimeHeapAllocBytes, builder.ProcessRuntimeHeapAllocBytes)
	errs = errors.Join(errs, err)
	builder.ProcessRuntimeTotalAllocBytes, err = builder.meter.Int64ObservableCounter(
		"otelcol_process_runtime_total_alloc_bytes",
		metric.WithDescription("Cumulative bytes allocated for heap objects (see 'go doc runtime.MemStats.TotalAlloc')"),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meter.RegisterCallback(builder.observeProcessRuntimeTotalAllocBytes, builder.ProcessRuntimeTotalAllocBytes)
	errs = errors.Join(errs, err)
	builder.ProcessRuntimeTotalSysMemoryBytes, err = builder.meter.Int64ObservableGauge(
		"otelcol_process_runtime_total_sys_memory_bytes",
		metric.WithDescription("Total bytes of memory obtained from the OS (see 'go doc runtime.MemStats.Sys')"),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meter.RegisterCallback(builder.observeProcessRuntimeTotalSysMemoryBytes, builder.ProcessRuntimeTotalSysMemoryBytes)
	errs = errors.Join(errs, err)
	builder.ProcessUptime, err = builder.meter.Float64ObservableCounter(
		"otelcol_process_uptime",
		metric.WithDescription("Uptime of the process"),
		metric.WithUnit("s"),
	)
	errs = errors.Join(errs, err)
	_, err = builder.meter.RegisterCallback(builder.observeProcessUptime, builder.ProcessUptime)
	errs = errors.Join(errs, err)
	return &builder, errs
}
