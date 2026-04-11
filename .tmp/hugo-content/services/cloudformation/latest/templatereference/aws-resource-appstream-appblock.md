This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::AppBlock

This resource creates an app block. App blocks store details about the virtual hard disk that
contains the files for the application in an S3 bucket. It also stores the setup script
with details about how to mount the virtual hard disk. App blocks are only supported for
Elastic fleets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::AppBlock",
  "Properties" : {
      "Description" : String,
      "DisplayName" : String,
      "Name" : String,
      "PackagingType" : String,
      "PostSetupScriptDetails" : ScriptDetails,
      "SetupScriptDetails" : ScriptDetails,
      "SourceS3Location" : S3Location,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::AppBlock
Properties:
  Description: String
  DisplayName: String
  Name: String
  PackagingType: String
  PostSetupScriptDetails:
    ScriptDetails
  SetupScriptDetails:
    ScriptDetails
  SourceS3Location:
    S3Location
  Tags:
    - Tag

```

## Properties

`Description`

The description of the app block.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisplayName`

The display name of the app block.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the app block.

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PackagingType`

The packaging type of the app block.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOM | APPSTREAM2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PostSetupScriptDetails`

The post setup script details of the app block.

_Required_: No

_Type_: [ScriptDetails](aws-properties-appstream-appblock-scriptdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SetupScriptDetails`

The setup script details of the app block.

_Required_: No

_Type_: [ScriptDetails](aws-properties-appstream-appblock-scriptdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceS3Location`

The source S3 location of the app block.

_Required_: Yes

_Type_: [S3Location](aws-properties-appstream-appblock-s3location.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags of the app block.

_Required_: No

_Type_: Array of [Tag](aws-properties-appstream-appblock-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the `Arn` of the app block, such as
`arn:aws:appstream:us-west-2:123456789123:app-block/abcdefg`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the app block.

`CreatedTime`

The time when the app block was created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon WorkSpaces Applications

S3Location

All content copied from https://docs.aws.amazon.com/.
