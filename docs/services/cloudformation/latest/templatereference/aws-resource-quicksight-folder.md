This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Folder

Creates an empty shared folder.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::Folder",
  "Properties" : {
      "AwsAccountId" : String,
      "FolderId" : String,
      "FolderType" : String,
      "Name" : String,
      "ParentFolderArn" : String,
      "Permissions" : [ ResourcePermission, ... ],
      "SharingModel" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::Folder
Properties:
  AwsAccountId: String
  FolderId: String
  FolderType: String
  Name: String
  ParentFolderArn: String
  Permissions:
    - ResourcePermission
  SharingModel: String
  Tags:
    - Tag

```

## Properties

`AwsAccountId`

The ID for the AWS account where you want to create the folder.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FolderId`

The ID of the folder.

_Required_: No

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FolderType`

The type of folder it is.

_Required_: No

_Type_: String

_Allowed values_: `SHARED | RESTRICTED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A display name for the folder.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentFolderArn`

The Amazon Resource Name (ARN) for the folder.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

A structure that describes the principals and the resource-level permissions of a folder.

To specify no permissions, omit `Permissions`.

_Required_: No

_Type_: Array of [ResourcePermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-folder-resourcepermission.html)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SharingModel`

The sharing scope of the folder.

_Required_: No

_Type_: String

_Allowed values_: `ACCOUNT | NAMESPACE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags for the folders that you want to apply overrides to.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-folder-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) for the folder.

`CreatedTime`

The time that the folder was created.

`LastUpdatedTime`

The time that the folder was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConnectionProperties

ResourcePermission
