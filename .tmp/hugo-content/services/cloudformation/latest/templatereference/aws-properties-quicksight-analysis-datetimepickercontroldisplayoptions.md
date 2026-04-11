This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DateTimePickerControlDisplayOptions

The display options of a control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateIconVisibility" : String,
  "DateTimeFormat" : String,
  "HelperTextVisibility" : String,
  "InfoIconLabelOptions" : SheetControlInfoIconLabelOptions,
  "TitleOptions" : LabelOptions
}

```

### YAML

```yaml

  DateIconVisibility: String
  DateTimeFormat: String
  HelperTextVisibility: String
  InfoIconLabelOptions:
    SheetControlInfoIconLabelOptions
  TitleOptions:
    LabelOptions

```

## Properties

`DateIconVisibility`

The date icon visibility of the `DateTimePickerControlDisplayOptions`.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateTimeFormat`

Customize how dates are formatted in controls.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HelperTextVisibility`

The helper text visibility of the `DateTimePickerControlDisplayOptions`.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

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

DateTimeParameterDeclaration

DateTimeValueWhenUnsetConfiguration

All content copied from https://docs.aws.amazon.com/.
