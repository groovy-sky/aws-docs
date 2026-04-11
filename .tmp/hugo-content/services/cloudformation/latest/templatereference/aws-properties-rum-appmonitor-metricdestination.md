This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor MetricDestination

Creates or updates a destination to receive extended metrics from CloudWatch RUM. You can send
extended metrics to CloudWatch or to a CloudWatch Evidently experiment.

For more information about extended metrics, see
[Extended metrics that you can send to CloudWatch and CloudWatch Evidently](../../../amazoncloudwatch/latest/monitoring/cloudwatch-rum-vended-metrics.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "DestinationArn" : String,
  "IamRoleArn" : String,
  "MetricDefinitions" : [ MetricDefinition, ... ]
}

```

### YAML

```yaml

  Destination: String
  DestinationArn: String
  IamRoleArn: String
  MetricDefinitions:
    - MetricDefinition

```

## Properties

`Destination`

Defines the destination to send the metrics to. Valid values are `CloudWatch` and
`Evidently`. If
you specify `Evidently`, you must also specify the ARN of the
CloudWatchEvidently experiment that is to
be the destination and an IAM role that has permission to write to the experiment.

_Required_: Yes

_Type_: String

_Allowed values_: `CloudWatch | Evidently`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationArn`

Use this parameter only if `Destination` is `Evidently`. This parameter specifies
the ARN of the Evidently experiment that will receive the extended metrics.

_Required_: No

_Type_: String

_Pattern_: `arn:[^:]*:[^:]*:[^:]*:[^:]*:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

This parameter is required if `Destination` is `Evidently`. If `Destination` is
`CloudWatch`, do not use this parameter.

This parameter specifies
the ARN of an IAM role that RUM will assume to write to the Evidently
experiment that you are sending metrics to. This role must have permission to write to that experiment.

_Required_: No

_Type_: String

_Pattern_: `arn:[^:]*:[^:]*:[^:]*:[^:]*:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricDefinitions`

An array of structures which define the metrics that you want to send.

_Required_: No

_Type_: Array of [MetricDefinition](aws-properties-rum-appmonitor-metricdefinition.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDefinition

ResourcePolicy

All content copied from https://docs.aws.amazon.com/.
