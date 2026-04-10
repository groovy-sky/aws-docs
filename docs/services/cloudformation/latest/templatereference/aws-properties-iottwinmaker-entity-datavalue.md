This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity DataValue

An object that specifies a value for a property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanValue" : Boolean,
  "DoubleValue" : Number,
  "Expression" : String,
  "IntegerValue" : Integer,
  "ListValue" : [ DataValue, ... ],
  "LongValue" : Number,
  "MapValue" : {Key: Value, ...},
  "RelationshipValue" : RelationshipValue,
  "StringValue" : String
}

```

### YAML

```yaml

  BooleanValue:
    Boolean
  DoubleValue: Number
  Expression: String
  IntegerValue:
    Integer
  ListValue:
    - DataValue
  LongValue: Number
  MapValue:
    Key: Value
  RelationshipValue:
    RelationshipValue
  StringValue:
    String

```

## Properties

`BooleanValue`

A boolean value.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DoubleValue`

A double value.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

An expression that produces the value.

_Required_: No

_Type_: String

_Pattern_: `(^\$\{Parameters\.[a-zA-z]+([a-zA-z_0-9]*)}$)`

_Minimum_: `1`

_Maximum_: `316`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerValue`

An integer value.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ListValue`

A list of multiple values.

_Required_: No

_Type_: Array of [DataValue](aws-properties-iottwinmaker-entity-datavalue.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LongValue`

A long value.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapValue`

An object that maps strings to multiple DataValue objects.

_Required_: No

_Type_: Object of [DataValue](aws-properties-iottwinmaker-entity-datavalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationshipValue`

A value that relates a component to another component.

_Required_: No

_Type_: [RelationshipValue](aws-properties-iottwinmaker-entity-relationshipvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

A string value.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataType

Definition

All content copied from https://docs.aws.amazon.com/.
