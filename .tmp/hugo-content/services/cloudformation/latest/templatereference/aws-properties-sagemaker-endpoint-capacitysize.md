This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint CapacitySize

Specifies the type and size of the endpoint capacity to activate for a blue/green
deployment, a rolling deployment, or a rollback strategy. You can specify your batches
as either instance count or the overall percentage or your fleet.

For a rollback strategy, if you don't specify the fields in this object, or if you set
the `Value` to 100%, then SageMaker uses a blue/green rollback strategy and rolls
all traffic back to the blue fleet.

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

Specifies the endpoint capacity type.

- `INSTANCE_COUNT`: The endpoint activates based on the number of
instances.

- `CAPACITY_PERCENT`: The endpoint activates based on the specified
percentage of capacity.

_Required_: Yes

_Type_: String

_Allowed values_: `INSTANCE_COUNT | CAPACITY_PERCENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Defines the capacity size, either as a number of instances or a capacity
percentage.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlueGreenUpdatePolicy

DeploymentConfig

All content copied from https://docs.aws.amazon.com/.
