This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard CustomParameterValues

The customized parameter values.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeValues" : [ String, ... ],
  "DecimalValues" : [ Number, ... ],
  "IntegerValues" : [ Number, ... ],
  "StringValues" : [ String, ... ]
}

```

### YAML

```yaml

  DateTimeValues:
    - String
  DecimalValues:
    - Number
  IntegerValues:
    - Number
  StringValues:
    - String

```

## Properties

`DateTimeValues`

A list of datetime-type parameter values.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DecimalValues`

A list of decimal-type parameter values.

_Required_: No

_Type_: Array of Number

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerValues`

A list of integer-type parameter values.

_Required_: No

_Type_: Array of Number

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValues`

A list of string-type parameter values.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomNarrativeOptions

CustomValuesConfiguration

All content copied from https://docs.aws.amazon.com/.
