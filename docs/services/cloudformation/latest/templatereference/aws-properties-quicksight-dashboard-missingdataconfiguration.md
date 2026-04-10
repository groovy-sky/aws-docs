This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard MissingDataConfiguration

The configuration options that determine how missing data is treated during the rendering of a line chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TreatmentOption" : String
}

```

### YAML

```yaml

  TreatmentOption: String

```

## Properties

`TreatmentOption`

The treatment option that determines how missing data should be rendered. Choose
from the following options:

- `INTERPOLATE`: Interpolate missing values between the prior and the next known value.

- `SHOW_AS_ZERO`: Show missing values as the value `0`.

- `SHOW_AS_BLANK`: Display a blank space when rendering missing data.

_Required_: No

_Type_: String

_Allowed values_: `INTERPOLATE | SHOW_AS_ZERO | SHOW_AS_BLANK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MinimumLabelType

NegativeValueConfiguration

All content copied from https://docs.aws.amazon.com/.
