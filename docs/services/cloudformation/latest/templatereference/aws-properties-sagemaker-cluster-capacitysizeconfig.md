This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster CapacitySizeConfig

The configuration of the size measurements of the AMI update. Using this configuration,
you can specify whether SageMaker should update your instance group by an amount or percentage
of instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Type: String
  Value: Integer

```

## Properties

`Type`

Specifies whether SageMaker should process the update by amount or percentage of
instances.

_Required_: Yes

_Type_: String

_Pattern_: `INSTANCE_COUNT|CAPACITY_PERCENTAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Specifies the amount or percentage of instances SageMaker updates at a time.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmDetails

ClusterAutoScalingConfig

All content copied from https://docs.aws.amazon.com/.
