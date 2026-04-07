# Transformation metrics and errors

CloudWatch Logs publishes transformation metrics to CloudWatch. These metrics include
`TransformedLogEvents`, `TransformedBytes`, and
`TransformationErrors`. For more information, see [Log transformer metrics and dimensions](cloudwatch-logs-monitoring-cloudwatch-metrics.md#CloudWatchLogs-Transformer-Metrics).

Whenever CloudWatch Logs tries and fails to transform a log event, it adds a
`@transformationError` system field to that log event. When you run a
CloudWatch Logs Insights query, you will see this field in all log events that had a
transformation failure. You can query for this field with a query such as `filter
                ispresent(@transformationError)` to find all the failed transformation
events.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Datatype converter processors

Analyze with Amazon OpenSearch Service
