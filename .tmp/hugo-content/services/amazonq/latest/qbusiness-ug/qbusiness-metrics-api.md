# Amazon Q Business API operation metrics

The following table shows the API operation metrics that Amazon Q Business sends to CloudWatch.

Metric nameUnitDescription

`success`

Count

The number of successful API operation calls. This metric is emitted for each successful API operation execution.

Valid dimensions: `MethodType`, `ApplicationId`

`failure`

Count

The number of failed API operation calls. This metric is emitted for each failed API operation execution.

Valid dimensions: `MethodType`, `ApplicationId`

`latency`

Milliseconds

The time taken to complete an API operation call. This metric measures the response time for individual API operations.

Valid dimensions: `MethodType`, `ApplicationId`

The `MethodType` dimension can include values such as:

- `ListPlugins`

- (Additional method types may be available depending on API usage)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Q Business chat metrics

Amazon Q Business index metrics

All content copied from https://docs.aws.amazon.com/.
