This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDB::DBCluster ServerlessV2ScalingConfiguration

Sets the scaling configuration of an Amazon DocumentDB Serverless cluster.

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

The maximum number of Amazon DocumentDB capacity units (DCUs) for an instance in an Amazon DocumentDB Serverless cluster.
You can specify DCU values in half-step increments, such as 32, 32.5, 33, and so on.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The minimum number of Amazon DocumentDB capacity units (DCUs) for an instance in an Amazon DocumentDB Serverless cluster.
You can specify DCU values in half-step increments, such as 8, 8.5, 9, and so on.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DocDB::DBCluster

Tag

All content copied from https://docs.aws.amazon.com/.
