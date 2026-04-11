This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule OperandTwo

Represents the right hand operand in the condition. In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BooleanValue" : Boolean,
  "DoubleValue" : Number,
  "EmptyValue" : Json,
  "StringValue" : String
}

```

### YAML

```yaml

  BooleanValue:
    Boolean
  DoubleValue: Number
  EmptyValue: Json
  StringValue:
    String

```

## Properties

`BooleanValue`

Boolean value type.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DoubleValue`

Double value type.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmptyValue`

Represents an empty operand value. In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

String value type.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OperandOne

RequiredCaseRule

All content copied from https://docs.aws.amazon.com/.
