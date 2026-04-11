This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FilterCrossSheetControl

A control from a filter that is scoped across more than one sheet. This represents your filter control on a sheet

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CascadingControlConfiguration" : CascadingControlConfiguration,
  "FilterControlId" : String,
  "SourceFilterId" : String
}

```

### YAML

```yaml

  CascadingControlConfiguration:
    CascadingControlConfiguration
  FilterControlId: String
  SourceFilterId: String

```

## Properties

`CascadingControlConfiguration`

The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.

_Required_: No

_Type_: [CascadingControlConfiguration](aws-properties-quicksight-template-cascadingcontrolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControlId`

The ID of the `FilterCrossSheetControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFilterId`

The source filter ID of the `FilterCrossSheetControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterControl

FilterDateTimePickerControl

All content copied from https://docs.aws.amazon.com/.
