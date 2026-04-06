# Monitor WebSocket API execution with CloudWatch metrics

You can use [Amazon CloudWatch](../../../amazoncloudwatch/latest/monitoring/whatiscloudwatch.md) metrics to
monitor WebSocket APIs. The configuration is similar to that used for REST APIs. For more
information, see [Monitor REST API execution with Amazon CloudWatch metrics](https://docs.aws.amazon.com/apigateway/latest/developerguide/monitoring-cloudwatch.html).

The following metrics are supported for WebSocket APIs:

MetricDescriptionConnectCountThe number of messages sent to the `$connect` route
integration.MessageCountThe number of messages sent to the WebSocket API, either from or to the
client.IntegrationErrorThe number of requests that return a 4XX/5XX response from the
integration.ClientErrorThe number of requests that have a 4XX response returned by API Gateway before
the integration is invoked.ExecutionErrorErrors that occurred when calling the integration.IntegrationLatencyThe time difference between API Gateway sending the request to the integration
and API Gateway receiving the response from the integration. Suppressed for
callbacks and mock integrations.

You can use the dimensions in the following table to filter API Gateway metrics.

DimensionDescriptionApiIdFilters API Gateway metrics for an API with the specified API ID.ApiId, StageFilters API Gateway metrics for an API stage with the specified API ID and
stage ID.ApiId, Method, Resource, Stage

Filters API Gateway metrics for an API method with the specified API ID,
stage ID, resource path, and route ID.

API Gateway will not send these metrics unless you have explicitly enabled
detailed CloudWatch metrics. You can do this by calling the [UpdateStage](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-stages-stagename.html) action
of the API Gateway V2 REST API to update the `detailedMetricsEnabled` property to
`true`. Alternatively, you can call
the [update-stage](https://docs.aws.amazon.com/cli/latest/reference/apigatewayv2/update-stage.html) AWS CLI command to update the `DetailedMetricsEnabled` property to
`true`. Enabling such metrics will incur additional charges to your
account. For pricing information, see [Amazon CloudWatch\
Pricing](https://aws.amazon.com/cloudwatch/pricing).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Monitor

Logging
