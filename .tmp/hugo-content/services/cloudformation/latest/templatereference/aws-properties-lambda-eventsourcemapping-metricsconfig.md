This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping MetricsConfig

The metrics configuration for your event source. Use this configuration object to define which metrics you want your
event source mapping to produce.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Metrics" : [ String, ... ]
}

```

### YAML

```yaml

  Metrics:
    - String

```

## Properties

`Metrics`

The metrics you want your event source mapping to produce, including `EventCount`, `ErrorCount`, `KafkaMetrics`.

- `EventCount` to receive metrics related to the number of events processed by your event source mapping.

- `ErrorCount` (Amazon MSK and self-managed Apache Kafka) to receive metrics related to the number of errors in your event source mapping processing.

- `KafkaMetrics` (Amazon MSK and self-managed Apache Kafka) to receive metrics related to the Kafka consumers from your event source mapping.

For more information about these metrics,
see [Event source mapping metrics](../../../lambda/latest/dg/monitoring-metrics-types.md#event-source-mapping-metrics).

_Required_: No

_Type_: Array of String

_Allowed values_: `EventCount | ErrorCount | KafkaMetrics`

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfig

OnFailure

All content copied from https://docs.aws.amazon.com/.
