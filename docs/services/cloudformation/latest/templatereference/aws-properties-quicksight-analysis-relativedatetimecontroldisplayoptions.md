This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis RelativeDateTimeControlDisplayOptions

The display options of a control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeFormat" : String,
  "InfoIconLabelOptions" : SheetControlInfoIconLabelOptions,
  "TitleOptions" : LabelOptions
}

```

### YAML

```yaml

  DateTimeFormat: String
  InfoIconLabelOptions:
    SheetControlInfoIconLabelOptions
  TitleOptions:
    LabelOptions

```

## Properties

`DateTimeFormat`

Customize how dates are formatted in controls.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InfoIconLabelOptions`

The configuration of info icon label options.

_Required_: No

_Type_: [SheetControlInfoIconLabelOptions](aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TitleOptions`

The options to configure the title visibility, name, and font size.

_Required_: No

_Type_: [LabelOptions](aws-properties-quicksight-analysis-labeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RelativeDatesFilter

ResourcePermission

All content copied from https://docs.aws.amazon.com/.
