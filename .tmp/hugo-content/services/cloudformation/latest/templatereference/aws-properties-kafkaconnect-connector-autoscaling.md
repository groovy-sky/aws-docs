This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector AutoScaling

Specifies how the connector scales.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxAutoscalingTaskCount" : Integer,
  "MaxWorkerCount" : Integer,
  "McuCount" : Integer,
  "MinWorkerCount" : Integer,
  "ScaleInPolicy" : ScaleInPolicy,
  "ScaleOutPolicy" : ScaleOutPolicy
}

```

### YAML

```yaml

  MaxAutoscalingTaskCount: Integer
  MaxWorkerCount: Integer
  McuCount: Integer
  MinWorkerCount: Integer
  ScaleInPolicy:
    ScaleInPolicy
  ScaleOutPolicy:
    ScaleOutPolicy

```

## Properties

`MaxAutoscalingTaskCount`

The maximum number of tasks allocated to the connector during autoscaling operations.
Must be at least equal to maxWorkerCount.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxWorkerCount`

The maximum number of workers allocated to the connector.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`McuCount`

The number of microcontroller units (MCUs) allocated to each connector worker. The valid
values are 1,2,4,8.

_Required_: Yes

_Type_: Integer

_Allowed values_: `1 | 2 | 4 | 8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinWorkerCount`

The minimum number of workers allocated to the connector.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleInPolicy`

The scale-in policy for the connector.

_Required_: Yes

_Type_: [ScaleInPolicy](aws-properties-kafkaconnect-connector-scaleinpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleOutPolicy`

The scale-out policy for the connector.

_Required_: Yes

_Type_: [ScaleOutPolicy](aws-properties-kafkaconnect-connector-scaleoutpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApacheKafkaCluster

Capacity

All content copied from https://docs.aws.amazon.com/.
