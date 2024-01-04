// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// DeleteAppConfig invokes deleteAppConfig operation.
	//
	// Deletes the AppConfig with the requested Key.
	//
	// DELETE /app-configs/{key}
	DeleteAppConfig(ctx context.Context, params DeleteAppConfigParams) (DeleteAppConfigRes, error)
	// Health invokes Health operation.
	//
	// Simple endpoint to check if the server is up.
	//
	// GET /health
	Health(ctx context.Context) (HealthRes, error)
	// ListAppConfig invokes listAppConfig operation.
	//
	// GET /app-configs
	ListAppConfig(ctx context.Context, params ListAppConfigParams) (ListAppConfigRes, error)
	// ReadAppConfig invokes readAppConfig operation.
	//
	// Finds the AppConfig with the requested Key and returns it.
	//
	// GET /app-configs/{key}
	ReadAppConfig(ctx context.Context, params ReadAppConfigParams) (ReadAppConfigRes, error)
	// SetAppConfig invokes setAppConfig operation.
	//
	// Sets an AppConfig with the specified Key/Value.
	//
	// POST /app-configs
	SetAppConfig(ctx context.Context, request *SetAppConfigReq, params SetAppConfigParams) (SetAppConfigRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// DeleteAppConfig invokes deleteAppConfig operation.
//
// Deletes the AppConfig with the requested Key.
//
// DELETE /app-configs/{key}
func (c *Client) DeleteAppConfig(ctx context.Context, params DeleteAppConfigParams) (DeleteAppConfigRes, error) {
	res, err := c.sendDeleteAppConfig(ctx, params)
	return res, err
}

func (c *Client) sendDeleteAppConfig(ctx context.Context, params DeleteAppConfigParams) (res DeleteAppConfigRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteAppConfig"),
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/app-configs/{key}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "DeleteAppConfig",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/app-configs/"
	{
		// Encode "key" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "key",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Key))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "app_name" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "app_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AppName))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "environment" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "environment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Environment))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stack" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stack",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stack.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_profile" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_profile",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSProfile))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_region" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_region",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSRegion))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_namespace" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_namespace",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sNamespace))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_cluster_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_cluster_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sClusterID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Access-Key-Id",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSAccessKeyID))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Secret-Access-Key",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSecretAccessKey))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Session-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSessionToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteAppConfigResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Health invokes Health operation.
//
// Simple endpoint to check if the server is up.
//
// GET /health
func (c *Client) Health(ctx context.Context) (HealthRes, error) {
	res, err := c.sendHealth(ctx)
	return res, err
}

func (c *Client) sendHealth(ctx context.Context) (res HealthRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Health"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/health"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "Health",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/health"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeHealthResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListAppConfig invokes listAppConfig operation.
//
// GET /app-configs
func (c *Client) ListAppConfig(ctx context.Context, params ListAppConfigParams) (ListAppConfigRes, error) {
	res, err := c.sendListAppConfig(ctx, params)
	return res, err
}

func (c *Client) sendListAppConfig(ctx context.Context, params ListAppConfigParams) (res ListAppConfigRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listAppConfig"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/app-configs"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListAppConfig",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app-configs"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "itemsPerPage" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ItemsPerPage.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "app_name" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "app_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AppName))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "environment" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "environment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Environment))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stack" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stack",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stack.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_profile" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_profile",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSProfile))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_region" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_region",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSRegion))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_namespace" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_namespace",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sNamespace))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_cluster_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_cluster_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sClusterID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Access-Key-Id",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSAccessKeyID))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Secret-Access-Key",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSecretAccessKey))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Session-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSessionToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListAppConfigResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadAppConfig invokes readAppConfig operation.
//
// Finds the AppConfig with the requested Key and returns it.
//
// GET /app-configs/{key}
func (c *Client) ReadAppConfig(ctx context.Context, params ReadAppConfigParams) (ReadAppConfigRes, error) {
	res, err := c.sendReadAppConfig(ctx, params)
	return res, err
}

func (c *Client) sendReadAppConfig(ctx context.Context, params ReadAppConfigParams) (res ReadAppConfigRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readAppConfig"),
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/app-configs/{key}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ReadAppConfig",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/app-configs/"
	{
		// Encode "key" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "key",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Key))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "app_name" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "app_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AppName))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "environment" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "environment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Environment))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stack" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stack",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stack.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_profile" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_profile",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSProfile))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_region" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_region",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSRegion))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_namespace" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_namespace",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sNamespace))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_cluster_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_cluster_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sClusterID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Access-Key-Id",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSAccessKeyID))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Secret-Access-Key",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSecretAccessKey))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Session-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSessionToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeReadAppConfigResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SetAppConfig invokes setAppConfig operation.
//
// Sets an AppConfig with the specified Key/Value.
//
// POST /app-configs
func (c *Client) SetAppConfig(ctx context.Context, request *SetAppConfigReq, params SetAppConfigParams) (SetAppConfigRes, error) {
	res, err := c.sendSetAppConfig(ctx, request, params)
	return res, err
}

func (c *Client) sendSetAppConfig(ctx context.Context, request *SetAppConfigReq, params SetAppConfigParams) (res SetAppConfigRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setAppConfig"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/app-configs"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "SetAppConfig",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/app-configs"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "itemsPerPage" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "itemsPerPage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ItemsPerPage.Get(); ok {
				return e.EncodeValue(conv.IntToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "app_name" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "app_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AppName))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "environment" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "environment",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.Environment))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "stack" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "stack",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Stack.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_profile" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_profile",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSProfile))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "aws_region" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "aws_region",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.AWSRegion))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_namespace" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_namespace",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sNamespace))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "k8s_cluster_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "k8s_cluster_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.K8sClusterID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeSetAppConfigRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Access-Key-Id",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSAccessKeyID))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Secret-Access-Key",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSecretAccessKey))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Aws-Session-Token",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.StringToString(params.XAWSSessionToken))
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeSetAppConfigResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
