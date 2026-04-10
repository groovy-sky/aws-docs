This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::Workspace WorkspaceProperties

Information about a WorkSpace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputeTypeName" : String,
  "RootVolumeSizeGib" : Integer,
  "RunningMode" : String,
  "RunningModeAutoStopTimeoutInMinutes" : Integer,
  "UserVolumeSizeGib" : Integer
}

```

### YAML

```yaml

  ComputeTypeName: String
  RootVolumeSizeGib: Integer
  RunningMode: String
  RunningModeAutoStopTimeoutInMinutes: Integer
  UserVolumeSizeGib: Integer

```

## Properties

`ComputeTypeName`

The compute type. For more information, see [Amazon WorkSpaces\
Bundles](https://aws.amazon.com/workspaces/details).

_Required_: No

_Type_: String

_Allowed values_: `VALUE | STANDARD | PERFORMANCE | POWER | GRAPHICS | POWERPRO | GENERALPURPOSE_4XLARGE | GENERALPURPOSE_8XLARGE | GRAPHICSPRO | GRAPHICS_G4DN | GRAPHICSPRO_G4DN | GRAPHICS_G6_XLARGE | GRAPHICS_G6_2XLARGE | GRAPHICS_G6_4XLARGE | GRAPHICS_G6_8XLARGE | GRAPHICS_G6_16XLARGE | GRAPHICS_GR6_4XLARGE | GRAPHICS_GR6_8XLARGE | GRAPHICS_G6F_LARGE | GRAPHICS_G6F_XLARGE | GRAPHICS_G6F_2XLARGE | GRAPHICS_G6F_4XLARGE | GRAPHICS_GR6F_4XLARGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootVolumeSizeGib`

The size of the root volume. For important information about how to modify the size of
the root and user volumes, see [Modify a\
WorkSpace](../../../workspaces/latest/adminguide/modify-workspaces.md).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunningMode`

The running mode. For more information, see [Manage the WorkSpace Running\
Mode](../../../workspaces/latest/adminguide/running-mode.md).

_Required_: No

_Type_: String

_Allowed values_: `AUTO_STOP | ALWAYS_ON | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunningModeAutoStopTimeoutInMinutes`

The time after a user logs off when WorkSpaces are automatically stopped. Configured in
60-minute intervals.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserVolumeSizeGib`

The size of the user storage. For important information about how to modify the size of
the root and user volumes, see [Modify a\
WorkSpace](../../../workspaces/latest/adminguide/modify-workspaces.md).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [WorkspaceProperties](../../../workspaces/latest/api/api-workspaceproperties.md) in the _Amazon WorkSpaces API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::WorkSpaces::WorkspacesPool

All content copied from https://docs.aws.amazon.com/.
