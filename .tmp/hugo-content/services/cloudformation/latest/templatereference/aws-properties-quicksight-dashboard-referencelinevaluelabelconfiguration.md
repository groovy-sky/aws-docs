This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ReferenceLineValueLabelConfiguration

The value label configuration of the label in a reference line.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FormatConfiguration" : NumericFormatConfiguration,
  "RelativePosition" : String
}

```

### YAML

```yaml

  FormatConfiguration:
    NumericFormatConfiguration
  RelativePosition: String

```

## Properties

`FormatConfiguration`

The format configuration of the value label.

_Required_: No

_Type_: [NumericFormatConfiguration](aws-properties-quicksight-dashboard-numericformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativePosition`

The relative position of the value label. Choose one of the following options:

- `BEFORE_CUSTOM_LABEL`

- `AFTER_CUSTOM_LABEL`

_Required_: No

_Type_: String

_Allowed values_: `BEFORE_CUSTOM_LABEL | AFTER_CUSTOM_LABEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceLineStyleConfiguration

RelativeDatesFilter

All content copied from https://docs.aws.amazon.com/.
