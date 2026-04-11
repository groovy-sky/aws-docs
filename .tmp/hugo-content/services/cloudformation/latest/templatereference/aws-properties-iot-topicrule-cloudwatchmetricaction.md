This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule CloudwatchMetricAction

Describes an action that captures a CloudWatch metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricName" : String,
  "MetricNamespace" : String,
  "MetricTimestamp" : String,
  "MetricUnit" : String,
  "MetricValue" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  MetricName: String
  MetricNamespace: String
  MetricTimestamp: String
  MetricUnit: String
  MetricValue: String
  RoleArn: String

```

## Properties

`MetricName`

The CloudWatch metric name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricNamespace`

The CloudWatch metric namespace name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricTimestamp`

An optional [Unix timestamp](../../../amazoncloudwatch/latest/developerguide/cloudwatch-concepts.md#about_timestamp).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricUnit`

The [metric\
unit](../../../amazoncloudwatch/latest/developerguide/cloudwatch-concepts.md#Unit) supported by CloudWatch.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricValue`

The CloudWatch metric value.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that allows access to the CloudWatch metric.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudwatchLogsAction

DynamoDBAction

All content copied from https://docs.aws.amazon.com/.
