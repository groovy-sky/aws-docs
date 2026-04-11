This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application LogPattern

The `AWS::ApplicationInsights::Application LogPattern` property type
specifies an object that defines the log patterns that belong to a
`LogPatternSet`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Pattern" : String,
  "PatternName" : String,
  "Rank" : Integer
}

```

### YAML

```yaml

  Pattern: String
  PatternName: String
  Rank: Integer

```

## Properties

`Pattern`

A regular expression that defines the log pattern. A log pattern can contain up to 50
characters, and it cannot be empty.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternName`

The name of the log pattern. A log pattern name can contain up to 50 characters, and
it cannot be empty. The characters can be Unicode letters, digits, or one of the
following symbols: period, dash, underscore.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9.-_]*`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rank`

The rank of the log pattern.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Log

LogPatternSet

All content copied from https://docs.aws.amazon.com/.
