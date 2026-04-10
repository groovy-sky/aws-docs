This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBCluster ServerlessScalingConfiguration

Contains the scaling configuration of a Neptune Serverless DB cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacity" : Number,
  "MinCapacity" : Number
}

```

### YAML

```yaml

  MaxCapacity: Number
  MinCapacity: Number

```

## Properties

`MaxCapacity`

The maximum number of Neptune capacity units (NCUs) for a DB instance in a Neptune Serverless cluster.
You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on.

_Required_: Yes

_Type_: Number

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The minimum number of Neptune capacity units (NCUs) for a DB instance
in a Neptune Serverless cluster. You can specify NCU values in half-step
increments, such as 8, 8.5, 9, and so on.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBClusterRole

Tag

All content copied from https://docs.aws.amazon.com/.
