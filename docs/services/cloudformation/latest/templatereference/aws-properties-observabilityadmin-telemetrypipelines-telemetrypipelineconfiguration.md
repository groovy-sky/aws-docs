This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryPipelines TelemetryPipelineConfiguration

Defines the configuration for a pipeline, including how data flows from sources through
processors to destinations. The configuration is specified in YAML format and must include a
valid pipeline definition with required source and sink components. This pipeline enables
end-to-end telemetry data collection, transformation, and delivery while supporting optional
processing steps and extensions for enhanced functionality.

The primary pipeline configuration section are:

- **Source:** Defines where log data originates from (S3
buckets, CloudWatch Logs, third-party APIs). Each pipeline must have exactly one
source.

- **Processors (optional):** Transform, parse, and enrich
log data as it flows through the pipeline. Processors are applied sequentially in the
order they are defined.

- **Sink:** Defines the destination where processed log
data is sent. Each pipeline must have exactly one sink.

- **Extensions (optional):** Provide additional
functionality such as AWS Secrets Manager integration for credential management.

For more details on each configuration section see [CloudWatch pipelines User\
Guide](../../../amazoncloudwatch/latest/monitoring/cloudwatch-pipelines.md). Additional comprehensive configuration examples can be found in the [CreateTelemetryPipeline API docs](https://docs.aws.amazon.com/cloudwatch/latest/observabilityadmin/API_CreateTelemetryPipeline.html#API_CreateTelemetryPipeline_Examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Body" : String
}

```

### YAML

```yaml

  Body: String

```

## Properties

`Body`

The pipeline configuration body that defines the data processing rules and
transformations.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `24000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TelemetryPipeline

TelemetryPipelineStatusReason
