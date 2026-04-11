This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity DataType

The entity data type.

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

The allowed values.

_Required_: No

_Type_: Array of [DataValue](aws-properties-iottwinmaker-entity-datavalue.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NestedType`

The nested type.

_Required_: No

_Type_: [DataType](aws-properties-iottwinmaker-entity-datatype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Relationship`

The relationship.

_Required_: No

_Type_: [Relationship](aws-properties-iottwinmaker-entity-relationship.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The entity type.

_Required_: No

_Type_: String

_Allowed values_: `RELATIONSHIP | STRING | LONG | BOOLEAN | INTEGER | DOUBLE | LIST | MAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnitOfMeasure`

The unit of measure.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompositeComponent

DataValue

All content copied from https://docs.aws.amazon.com/.
