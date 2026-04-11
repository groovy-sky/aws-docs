This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template WordCloudOptions

The word cloud options for a word cloud visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudLayout" : String,
  "MaximumStringLength" : Number,
  "WordCasing" : String,
  "WordOrientation" : String,
  "WordPadding" : String,
  "WordScaling" : String
}

```

### YAML

```yaml

  CloudLayout: String
  MaximumStringLength: Number
  WordCasing: String
  WordOrientation: String
  WordPadding: String
  WordScaling: String

```

## Properties

`CloudLayout`

The cloud layout options (fluid, normal) of a word cloud.

_Required_: No

_Type_: String

_Allowed values_: `FLUID | NORMAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumStringLength`

The length limit of each word from 1-100.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordCasing`

The word casing options (lower\_case, existing\_case) for the words in a word cloud.

_Required_: No

_Type_: String

_Allowed values_: `LOWER_CASE | EXISTING_CASE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordOrientation`

The word orientation options (horizontal, horizontal\_and\_vertical) for the words in a word cloud.

_Required_: No

_Type_: String

_Allowed values_: `HORIZONTAL | HORIZONTAL_AND_VERTICAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordPadding`

The word padding options (none, small, medium, large) for the words in a word cloud.

_Required_: No

_Type_: String

_Allowed values_: `NONE | SMALL | MEDIUM | LARGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WordScaling`

The word scaling options (emphasize, normal) for the words in a word cloud.

_Required_: No

_Type_: String

_Allowed values_: `EMPHASIZE | NORMAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WordCloudFieldWells

WordCloudSortConfiguration

All content copied from https://docs.aws.amazon.com/.
