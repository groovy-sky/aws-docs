This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings ToolbarConfiguration

The configuration of the toolbar. This allows administrators to select the toolbar type and visual mode, set maximum display resolution for sessions, and choose which items are visible to end users during their sessions. If administrators do not modify these settings, end users retain control over their toolbar preferences.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HiddenToolbarItems" : [ String, ... ],
  "MaxDisplayResolution" : String,
  "ToolbarType" : String,
  "VisualMode" : String
}

```

### YAML

```yaml

  HiddenToolbarItems:
    - String
  MaxDisplayResolution: String
  ToolbarType: String
  VisualMode: String

```

## Properties

`HiddenToolbarItems`

The list of toolbar items to be hidden.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxDisplayResolution`

The maximum display resolution that is allowed for the session.

_Required_: No

_Type_: String

_Allowed values_: `size4096X2160 | size3840X2160 | size3440X1440 | size2560X1440 | size1920X1080 | size1280X720 | size1024X768 | size800X600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolbarType`

The type of toolbar displayed during the session.

_Required_: No

_Type_: String

_Allowed values_: `Floating | Docked`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualMode`

The visual mode of the toolbar.

_Required_: No

_Type_: String

_Allowed values_: `Dark | Light`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
