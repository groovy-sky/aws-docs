This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::Workspace

The `AWS::WorkSpaces::Workspace` resource specifies a WorkSpace.

Updates are not supported for the `BundleId`,
`RootVolumeEncryptionEnabled`, `UserVolumeEncryptionEnabled`, or
`VolumeEncryptionKey` properties. To update these properties, you must also
update a property that triggers a replacement, such as the `UserName`
property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpaces::Workspace",
  "Properties" : {
      "BundleId" : String,
      "DirectoryId" : String,
      "RootVolumeEncryptionEnabled" : Boolean,
      "Tags" : [ Tag, ... ],
      "UserName" : String,
      "UserVolumeEncryptionEnabled" : Boolean,
      "VolumeEncryptionKey" : String,
      "WorkspaceProperties" : WorkspaceProperties
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpaces::Workspace
Properties:
  BundleId: String
  DirectoryId: String
  RootVolumeEncryptionEnabled: Boolean
  Tags:
    - Tag
  UserName: String
  UserVolumeEncryptionEnabled: Boolean
  VolumeEncryptionKey: String
  WorkspaceProperties:
    WorkspaceProperties

```

## Properties

`BundleId`

The identifier of the bundle for the WorkSpace.

_Required_: Yes

_Type_: String

_Pattern_: `^wsb-[0-9a-z]{8,63}$`

_Update requires_: Updates are not supported.

`DirectoryId`

The identifier of the Directory Service directory for the WorkSpace.

_Required_: Yes

_Type_: String

_Pattern_: `^(d-[0-9a-f]{8,63}$)|(wsd-[0-9a-z]{8,63}$)`

_Minimum_: `10`

_Maximum_: `65`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootVolumeEncryptionEnabled`

Indicates whether the data stored on the root volume is encrypted.

_Required_: No

_Type_: Boolean

_Update requires_: Updates are not supported.

`Tags`

The tags for the WorkSpace.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspaces-workspace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

The user name of the user for the WorkSpace. This user name must exist in the Directory Service directory for the WorkSpace.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserVolumeEncryptionEnabled`

Indicates whether the data stored on the user volume is encrypted.

_Required_: No

_Type_: Boolean

_Update requires_: Updates are not supported.

`VolumeEncryptionKey`

The symmetric AWS KMS key used to encrypt data stored on your WorkSpace.
Amazon WorkSpaces does not support asymmetric KMS keys.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`WorkspaceProperties`

The WorkSpace properties.

_Required_: No

_Type_: [WorkspaceProperties](aws-properties-workspaces-workspace-workspaceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The identifier of the WorkSpace, returned as a string.

`WorkspaceId`

The identifier of the WorkSpace.

## See also

- [CreateWorkspaces](../../../workspaces/latest/api/api-createworkspaces.md) in the _Amazon WorkSpaces API_
_Reference_

- [Launch a\
Virtual Desktop Using Amazon WorkSpaces](../../../workspaces/latest/adminguide/launch-workspaces-tutorials.md) in the _Amazon_
_WorkSpaces Administration Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
