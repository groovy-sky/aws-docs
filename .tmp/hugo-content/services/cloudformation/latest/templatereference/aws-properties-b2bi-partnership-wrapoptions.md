This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership WrapOptions

Contains options for wrapping (line folding) in X12 EDI files. Wrapping controls how
long lines are handled in the EDI output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LineLength" : Number,
  "LineTerminator" : String,
  "WrapBy" : String
}

```

### YAML

```yaml

  LineLength: Number
  LineTerminator: String
  WrapBy: String

```

## Properties

`LineLength`

Specifies the maximum length of a line before wrapping occurs. This value is used when
`wrapBy` is set to `LINE_LENGTH`.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineTerminator`

Specifies the character sequence used to terminate lines when wrapping. Valid
values:

- `CRLF`: carriage return and line feed

- `LF`: line feed)

- `CR`: carriage return

_Required_: No

_Type_: String

_Allowed values_: `CRLF | LF | CR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WrapBy`

Specifies the method used for wrapping lines in the EDI output. Valid values:

- `SEGMENT`: Wraps by segment.

- `ONE_LINE`: Indicates that the entire content is on a single
line.

###### Note

When you specify `ONE_LINE`, do not provide either the line length
nor the line terminator value.

- `LINE_LENGTH`: Wraps by character count, as specified by
`lineLength` value.

_Required_: No

_Type_: String

_Allowed values_: `SEGMENT | ONE_LINE | LINE_LENGTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

X12AcknowledgmentOptions

All content copied from https://docs.aws.amazon.com/.
