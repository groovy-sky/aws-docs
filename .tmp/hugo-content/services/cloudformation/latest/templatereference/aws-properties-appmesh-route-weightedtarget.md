This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route WeightedTarget

An object that represents a target and its relative weight. Traffic is distributed
across targets according to their relative weight. For example, a weighted target with a
relative weight of 50 receives five times as much traffic as one with a relative weight of
10\. The total weight for all targets combined must be less than or equal to 100.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : Integer,
  "VirtualNode" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  Port: Integer
  VirtualNode: String
  Weight: Integer

```

## Properties

`Port`

The targeted port of the weighted object.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualNode`

The virtual node to associate with the weighted target.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The relative weight of the weighted target.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TcpTimeout

AWS::AppMesh::VirtualGateway

All content copied from https://docs.aws.amazon.com/.
