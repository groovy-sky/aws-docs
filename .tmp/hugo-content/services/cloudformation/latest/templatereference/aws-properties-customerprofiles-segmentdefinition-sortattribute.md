This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition SortAttribute

The `SortAttribute` property type specifies Property description not available. for an [AWS::CustomerProfiles::SegmentDefinition](aws-resource-customerprofiles-segmentdefinition.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataType" : String,
  "Name" : String,
  "Order" : String,
  "Type" : String
}

```

### YAML

```yaml

  DataType: String
  Name: String
  Order: String
  Type: String

```

## Properties

`DataType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `STRING | NUMBER | DATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Order`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `PROFILE | CALCULATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SegmentSort

SourceSegment

All content copied from https://docs.aws.amazon.com/.
