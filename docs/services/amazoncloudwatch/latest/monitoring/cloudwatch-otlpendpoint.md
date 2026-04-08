# OTLP Endpoints

OpenTelemetry Protocol (OTLP) is a general-purpose telemetry data delivery protocol
designed for the OpenTelemetry. CloudWatch OpenTelemetry endpoints are HTTP 1.1 endpoints. You
need to configure your OpenTelemetry collector to start sending open telemetry data to
CloudWatch. For more information, see [Getting started](cloudwatch-otlpgettingstarted.md).

## Logs endpoint

The logs endpoint follows the pattern
`https://logs.AWS Region.amazonaws.com/v1/logs`.
For example, for the `US West (Oregon) (us-west-2) Region`, the endpoint
will be `https://logs.us-west-2.amazonaws.com/v1/logs`. You can use the
above endpoint to forward logs to an existing `LogGroup` and
`LogStream`. For more information on setting up `LogGroup`
to ingest log data, see [Amazon CloudWatch\
Logs concepts](../logs/cloudwatchlogsconcepts.md).

You must configure `LogGroup` and `LogStream` when you
invoke CloudWatch Logs OpenTelemetry endpoint by setting `x-aws-log-group` and
`x-aws-log-stream` HTTP headers to `LogGroup` and
`LogStream` name respectively. For more information, see [Getting started](cloudwatch-otlpgettingstarted.md). The endpoint authenticates callers using Signature 4 authentication. For more
information, see [AWS Signature Version 4 for API\
requests](../../../iam/latest/userguide/reference-sigv.md).

When log event size exceed 1MB, CloudWatch Logs automatically truncates up to 10 fields,
starting with the largest fields. Each field is truncated as needed to keep the
total event size as close to 1MB as possible. The excess portions are stored as
Large Log Objects (LLOs) and LLO reference system fields are added. Optionally, you
can specify the field paths that need to be truncated by setting the
`x-aws-truncatable-fields ` HTTP header. The LLOs can be retrieved and
streamed back using the `GetLogObject` API. For more information, see
[GetLogObject](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogobject.md). Support of log events larger than 1MB and LLO experience
is available in US East (N. Virginia) US West (Oregon), Europe (Frankfurt),
Asia Pacific (Sydney), Asia Pacific (Mumbai), US East (Ohio),
Europe (Ireland), Asia Pacific (Tokyo), and Asia Pacific (Singapore).

## Metrics endpoint

The metrics endpoint follows the pattern `https://monitoring.AWS
                        Region.amazonaws.com/v1/metrics`. For example, for the
US West (Oregon) (us-west-2) Region, the endpoint will be
`https://monitoring.us-west-2.amazonaws.com/v1/metrics`.

You need to configure your OpenTelemetry collector to start sending metrics to
CloudWatch. The endpoint authenticates callers using Signature 4 authentication. For more
information, see [AWS Signature Version 4 for API\
requests](../../../iam/latest/userguide/reference-sigv.md).

## Traces endpoint

The traces endpoint follows the pattern `https://xray.AWS
                        Region.amazonaws.com/v1/traces`. For example, for the
US West (Oregon) (us-west-2) Region, the endpoint will be
`https://xray.us-west-2.amazonaws.com/v1/traces`.

You need to configure your OpenTelemetry collector to start sending traces to
CloudWatch. The endpoint authenticates callers using Signature 4 authentication. For more
information, see [AWS Signature Version 4 for API\
requests](../../../iam/latest/userguide/reference-sigv.md).

## RUM endpoint

The RUM endpoint follows the pattern `https://dataplane.rum.{AWS
                    Region}.amazonaws.com/v1/rum`. For example, for the US West (Oregon) Region, the
endpoint is `https://dataplane.rum.us-west-2.amazonaws.com/v1/rum`. This
endpoint handles _client-side telemetry data_ (only traces and log records with
`eventName`) for CloudWatch RUM applications.

To use this endpoint, you must create a [RUM app monitor](cloudwatch-rum-get-started-create-app-monitor.md)
with Mobile platform (Android/ iOS) and use the code snippet generated to instrument
your applications. The snippet will pull the RUM Mobile SDKs which are configured
with this endpoint. You can configure the SDKs further for RUM to collect telemetry
accordingly.

The endpoint supports both authenticated and unauthenticated requests. You can
use AWS Signature Version 4 (SigV4) for authenticated requests, or resource-based
policies to allow unauthenticated access from mobile applications.

To learn more about the authentication models as defined in their
SDKs, see the following:

- iOS applications - [AWS Distro\
for OpenTelemetry (ADOT) iOS SDK](https://github.com/aws-observability/aws-otel-swift).

- Android applications - [AWS Distro\
for OpenTelemetry (ADOT) Android SDK](https://github.com/aws-observability/aws-otel-android).

## Endpoint limits and restrictions

The table lists the common endpoint limits and restrictions for metrics, traces, and
logs.

LimitEndpointAdditional information

Required collector extension

[sigv4authextension](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/extension/sigv4authextension)

To send metrics or traces to OTLP endpoint you must use
sigv4authextension

Supported protocol

HTTP

The endpoint supports only HTTP and doesn't support
gRPC

Supported OTLP versions

OTLP 1.x

Payload format

binary, json

The endpoint accepts requests using binary and json
formats

Compression methods

gzip, none

The endpoint only supports gzip and none compression
methods

The table lists the endpoint limits and restrictions for traces.

LimitTraces endpointAdditional information

Maximum uncompressed bytes / request

5 MB

The OTLP endpoint will reject requests larger than 5MB when
payload is uncompressed.

Maximum events / request

10,000 spans

The maximum number of spans in a batch is 10,000. Exceeding
this limit will cause rejection of the API call.

Single resource and scope size

16 KB

Each unique resource and corresponding scope should not exceed
16 KB of size. Exceeding this limit for any resource will cause
rejection of the entire API call.

Single span maximum size

200 KB

Spans more than 200KB will be rejected by the endpoint.

Span created timestamps

2 hours in the future and 14 days in the past

None of the spans in the batch can be more than two hours in
the future or more than 14 days in the past.

Maximum time gap in events / request

24 hours

The table lists the endpoint limits and restrictions for metrics.

LimitMetrics endpointAdditional informationError code

Maximum TPS

500

The maximum number of requests per second allowed per account.

429

Maximum new series creation rate

1,000,000

Maximum number of new series that can be created in a 10-minute window. This limit applies only to creating new metric names, not to ingesting data points for existing metrics.

429 - all metrics in the requests are throttled; 200 - metrics in the requests are partially throttled

Maximum request size

1 MB

The maximum uncompressed size of the request.

400

Maximum datapoint count

1,000

The maximum number of datapoints sent in a single request. This number is a sum across ResourceMetrics, ScopeMetrics and Metrics.

400

Maximum metadata size

40 KB

The maximum combined size of all labels and label values accepted for a series per datapoint.

400 - all metrics in the requests are invalid; 200 - metrics in the requests are partially invalid

Maximum label count

150

Maximum number of labels across Resource/Scope/Datapoint attributes per datapoint.

400 - all metrics in the requests are invalid; 200 - metrics in the requests are partially invalid

The table lists the endpoint limits and restrictions for logs.

LimitLogs endpointAdditional information

Maximum uncompressed bytes / request

1 MB

The OTLP endpoint will reject requests larger than 1MB when
payload is uncompressed.

The maximum request size is 1,048,576 bytes after
decompression and deserialization of binary data serialized by
Protocol buffers. This size is calculated as the sum of all
event messages in UTF-8, plus 26 bytes for each log
record.

20 MB

Available only in US East (N. Virginia) US West (Oregon),
Europe (Frankfurt), Asia Pacific (Sydney),
Asia Pacific (Mumbai), US East (Ohio), Europe (Ireland),
Asia Pacific (Tokyo), and Asia Pacific (Singapore).

The maximum request size is 20MB (20,971,520 bytes) after the
OTLP payload is decompressed and decoded from JSON
format.

For logs up to 1MB – These logs have full access to all
CloudWatch Logs features including queries and live tail.

For logs larger than 1MB – The excess portion is
processed as Large Log Objects (LLOs).

Request per second

5000

5000 transactions per second per account per Region You can
request an increase to the per-second throttling quota by using
the Service Quotas service.

Single resource and scope size

16 KB

Each unique resource and corresponding scope should not exceed
16 KB of size. Exceeding this limit for any resource will cause
rejection of the entire API call.

Single LogEvent size

1 MB

LogEvent size is calculated as sum of sizes for each
LogRecord, Scope and Resource. This quota can't be
changed.

Logs created timestamps

2 hours in future and 14 days old

The log records in the batch does not have to be in a
chronological order. However, the log records in the batch
cannot be more than 2 hours in the future and cannot be more
than 14 days in the past. Also, none of the log records can be
earlier than the retention period of the log
group.

Maximum time gap in events / request

24 hours

Maximum events / request

10,000 logs

The maximum number of log events in a batch is 10,000.
Exceeding this limit will cause rejection of the API
call.

Maximum large log objects / request

1 log record

Available in US East (N. Virginia) US West (Oregon),
Europe (Frankfurt), Asia Pacific (Sydney),
Asia Pacific (Mumbai), US East (Ohio), Europe (Ireland),
Asia Pacific (Tokyo), and Asia Pacific (Singapore).

For content exceeding 1MB in a log event, the excess content
is stored as LLOs. Limited to 1 log record per request.

Maximum large log objects / record

10 LLOs

Available in US East (N. Virginia) US West (Oregon),
Europe (Frankfurt), Asia Pacific (Sydney),
Asia Pacific (Mumbai), US East (Ohio), Europe (Ireland),
Asia Pacific (Tokyo), and Asia Pacific (Singapore).

A single log record can contain up to 10 LLOs.

###### Note

The account limits for Logs are shared across the SDK and the new Logs
endpoint.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenTelemetry

Getting started

All content copied from https://docs.aws.amazon.com/.
