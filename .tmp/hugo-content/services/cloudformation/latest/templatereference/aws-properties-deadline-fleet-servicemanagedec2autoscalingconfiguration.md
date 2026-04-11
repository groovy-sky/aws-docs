This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet ServiceManagedEc2AutoScalingConfiguration

The `ServiceManagedEc2AutoScalingConfiguration` property type specifies Property description not available. for an [AWS::Deadline::Fleet](aws-resource-deadline-fleet.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScaleOutWorkersPerMinute" : Integer,
  "StandbyWorkerCount" : Integer,
  "WorkerIdleDurationSeconds" : Integer
}

```

### YAML

```yaml

  ScaleOutWorkersPerMinute: Integer
  StandbyWorkerCount: Integer
  WorkerIdleDurationSeconds: Integer

```

## Properties

`ScaleOutWorkersPerMinute`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandbyWorkerCount`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerIdleDurationSeconds`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemoryMiBRange

ServiceManagedEc2FleetConfiguration

All content copied from https://docs.aws.amazon.com/.
