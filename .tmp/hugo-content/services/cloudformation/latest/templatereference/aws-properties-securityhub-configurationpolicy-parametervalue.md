This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConfigurationPolicy ParameterValue

An object that includes the data type of a security control parameter and its current value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Boolean" : Boolean,
  "Double" : Number,
  "Enum" : String,
  "EnumList" : [ String, ... ],
  "Integer" : Integer,
  "IntegerList" : [ Integer, ... ],
  "String" : String,
  "StringList" : [ String, ... ]
}

```

### YAML

```yaml

  Boolean:
    Boolean
  Double: Number
  Enum: String
  EnumList:
    - String
  Integer:
    Integer
  IntegerList:
    - Integer
  String:
    String
  StringList:
    - String

```

## Properties

`Boolean`

A control parameter that is a boolean.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Double`

A control parameter that is a double.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enum`

A control parameter that is an enum.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnumList`

A control parameter that is a list of enums.

_Required_: No

_Type_: Array of String

_Maximum_: `2048 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Integer`

A control parameter that is an integer.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerList`

A control parameter that is a list of integers.

_Required_: No

_Type_: Array of Integer

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`String`

A control parameter that is a string.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringList`

A control parameter that is a list of strings.

_Required_: No

_Type_: Array of String

_Maximum_: `2048 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterConfiguration

Policy

All content copied from https://docs.aws.amazon.com/.
