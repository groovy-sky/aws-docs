# Publishing logs with the embedded metric format

You can generate embedded metric format logs using the following methods:

- Generate and send the logs by using the [open-sourced client libraries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Libraries.html).

- Manually generate the logs using the [embedded metric format specification](cloudwatch-embedded-metric-format-specification.md), and then use the [CloudWatch agent](cloudwatch-embedded-metric-format-generation-cloudwatch-agent.md) or the [PutLogEvents API](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putlogevents.md) to send the logs.

The following topics provide more information about embedded metrics.

###### Topics

- [Creating logs in embedded metric format using the client libraries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Libraries.html)

- [Specification: Embedded metric format](cloudwatch-embedded-metric-format-specification.md)

- [Using the PutLogEvents API to send manually-created embedded metric format logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Generation_PutLogEvents.html)

- [Using the CloudWatch agent to send embedded metric format logs](cloudwatch-embedded-metric-format-generation-cloudwatch-agent.md)

- [Using the embedded metric format with AWS Distro for OpenTelemetry](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_OpenTelemetry.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Embedding metrics within logs

Using the client libraries
