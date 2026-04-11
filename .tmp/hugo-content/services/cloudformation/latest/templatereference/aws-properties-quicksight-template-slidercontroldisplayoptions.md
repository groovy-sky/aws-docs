This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SliderControlDisplayOptions

The display options of a control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfoIconLabelOptions" : SheetControlInfoIconLabelOptions,
  "TitleOptions" : LabelOptions
}

```

### YAML

```yaml

  InfoIconLabelOptions:
    SheetControlInfoIconLabelOptions
  TitleOptions:
    LabelOptions

```

## Properties

`InfoIconLabelOptions`

The configuration of info icon label options.

_Required_: No

_Type_: [SheetControlInfoIconLabelOptions](aws-properties-quicksight-template-sheetcontrolinfoiconlabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TitleOptions`

The options to configure the title visibility, name, and font size.

_Required_: No

_Type_: [LabelOptions](aws-properties-quicksight-template-labeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SingleAxisOptions

SmallMultiplesAxisProperties

All content copied from https://docs.aws.amazon.com/.
