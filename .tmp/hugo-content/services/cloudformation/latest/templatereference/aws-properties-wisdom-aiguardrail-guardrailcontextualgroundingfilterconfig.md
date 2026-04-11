This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrail GuardrailContextualGroundingFilterConfig

A configuration for grounding filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Threshold" : Number,
  "Type" : String
}

```

### YAML

```yaml

  Threshold: Number
  Type: String

```

## Properties

`Threshold`

The threshold for this filter.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of this filter.

_Required_: Yes

_Type_: String

_Allowed values_: `GROUNDING | RELEVANCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GuardrailContentFilterConfig

GuardrailManagedWordsConfig

All content copied from https://docs.aws.amazon.com/.
