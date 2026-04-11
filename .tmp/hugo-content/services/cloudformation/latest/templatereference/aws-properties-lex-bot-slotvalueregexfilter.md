This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotValueRegexFilter

Provides a regular expression used to validate the value of a
slot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Pattern" : String
}

```

### YAML

```yaml

  Pattern: String

```

## Properties

`Pattern`

A regular expression used to validate the value of a slot.

Use a standard regular expression. Amazon Lex supports the following
characters in the regular expression:

- A-Z, a-z

- 0-9

- Unicode characters ("\\⁠u<Unicode>")

Represent Unicode characters with four digits, for example "\\⁠u0041"
or "\\⁠u005A".

The following regular expression operators are not supported:

- Infinite repeaters: \*, +, or {x,} with no upper bound.

- Wild card (.)

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotValueOverrideMap

SlotValueSelectionSetting

All content copied from https://docs.aws.amazon.com/.
