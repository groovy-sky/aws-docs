This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::ComponentType DataType

An object that specifies the data type of a property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedValues" : [ DataValue, ... ],
  "NestedType" : DataType,
  "Relationship" : Relationship,
  "Type" : String,
  "UnitOfMeasure" : String
}

```

### YAML

```yaml

  AllowedValues:
    - DataValue
  NestedType:
    DataType
  Relationship:
    Relationship
  Type: String
  UnitOfMeasure: String

```

## Properties

`AllowedValues`

The allowed values for this data type.

_Required_: No

_Type_: Array of [DataValue](aws-properties-iottwinmaker-componenttype-datavalue.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NestedType`

The nested type in the data type.

_Required_: No

_Type_: [DataType](aws-properties-iottwinmaker-componenttype-datatype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Relationship`

A relationship that associates a component with another component.

_Required_: No

_Type_: [Relationship](aws-properties-iottwinmaker-componenttype-relationship.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The underlying type of the data type.

Valid Values: `RELATIONSHIP | STRING | LONG | BOOLEAN | INTEGER | DOUBLE | LIST | MAP`

_Required_: Yes

_Type_: String

_Allowed values_: `RELATIONSHIP | STRING | LONG | BOOLEAN | INTEGER | DOUBLE | LIST | MAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnitOfMeasure`

The unit of measure used in this data type.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataConnector

DataValue

All content copied from https://docs.aws.amazon.com/.
