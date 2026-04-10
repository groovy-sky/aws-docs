This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis LayerCustomAction

A layer custom action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionOperations" : [ LayerCustomActionOperation, ... ],
  "CustomActionId" : String,
  "Name" : String,
  "Status" : String,
  "Trigger" : String
}

```

### YAML

```yaml

  ActionOperations:
    - LayerCustomActionOperation
  CustomActionId: String
  Name: String
  Status: String
  Trigger: String

```

## Properties

`ActionOperations`

A list of `LayerCustomActionOperations`.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

_Required_: Yes

_Type_: Array of [LayerCustomActionOperation](aws-properties-quicksight-analysis-layercustomactionoperation.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomActionId`

The ID of the custom action.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the custom action.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the `LayerCustomAction`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trigger`

The trigger of the `LayerCustomAction`.

Valid values are defined as follows:

- `DATA_POINT_CLICK`: Initiates a custom action by a left pointer click on a data point.

- `DATA_POINT_MENU`: Initiates a custom action by right pointer click from the menu.

_Required_: Yes

_Type_: String

_Allowed values_: `DATA_POINT_CLICK | DATA_POINT_MENU`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelOptions

LayerCustomActionOperation

All content copied from https://docs.aws.amazon.com/.
