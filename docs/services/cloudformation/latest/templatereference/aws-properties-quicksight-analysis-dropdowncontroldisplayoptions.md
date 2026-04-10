This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DropDownControlDisplayOptions

The display options of a control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfoIconLabelOptions" : SheetControlInfoIconLabelOptions,
  "SelectAllOptions" : ListControlSelectAllOptions,
  "TitleOptions" : LabelOptions
}

```

### YAML

```yaml

  InfoIconLabelOptions:
    SheetControlInfoIconLabelOptions
  SelectAllOptions:
    ListControlSelectAllOptions
  TitleOptions:
    LabelOptions

```

## Properties

`InfoIconLabelOptions`

The configuration of info icon label options.

_Required_: No

_Type_: [SheetControlInfoIconLabelOptions](aws-properties-quicksight-analysis-sheetcontrolinfoiconlabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectAllOptions`

The configuration of the `Select all` options in a
dropdown control.

_Required_: No

_Type_: [ListControlSelectAllOptions](aws-properties-quicksight-analysis-listcontrolselectalloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TitleOptions`

The options to configure the title visibility, name, and font size.

_Required_: No

_Type_: [LabelOptions](aws-properties-quicksight-analysis-labeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DrillDownFilter

DynamicDefaultValue

All content copied from https://docs.aws.amazon.com/.
