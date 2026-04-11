This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ListControlDisplayOptions

The display options of a control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfoIconLabelOptions" : SheetControlInfoIconLabelOptions,
  "SearchOptions" : ListControlSearchOptions,
  "SelectAllOptions" : ListControlSelectAllOptions,
  "TitleOptions" : LabelOptions
}

```

### YAML

```yaml

  InfoIconLabelOptions:
    SheetControlInfoIconLabelOptions
  SearchOptions:
    ListControlSearchOptions
  SelectAllOptions:
    ListControlSelectAllOptions
  TitleOptions:
    LabelOptions

```

## Properties

`InfoIconLabelOptions`

The configuration of info icon label options.

_Required_: No

_Type_: [SheetControlInfoIconLabelOptions](aws-properties-quicksight-dashboard-sheetcontrolinfoiconlabeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SearchOptions`

The configuration of the search options in a list control.

_Required_: No

_Type_: [ListControlSearchOptions](aws-properties-quicksight-dashboard-listcontrolsearchoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectAllOptions`

The configuration of the `Select all` options in a list control.

_Required_: No

_Type_: [ListControlSelectAllOptions](aws-properties-quicksight-dashboard-listcontrolselectalloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TitleOptions`

The options to configure the title visibility, name, and font size.

_Required_: No

_Type_: [LabelOptions](aws-properties-quicksight-dashboard-labeloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LinkSharingConfiguration

ListControlSearchOptions

All content copied from https://docs.aws.amazon.com/.
