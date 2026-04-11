This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile FSxLustreFileSystemConfig

The settings for assigning a custom Amazon FSx for Lustre file system to a user profile or space for an
Amazon SageMaker Domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemId" : String,
  "FileSystemPath" : String
}

```

### YAML

```yaml

  FileSystemId: String
  FileSystemPath: String

```

## Properties

`FileSystemId`

The globally unique, 17-digit, ID of the file system, assigned by Amazon FSx for Lustre.

_Required_: Yes

_Type_: String

_Pattern_: `^(fs-[0-9a-f]{8,})$`

_Minimum_: `11`

_Maximum_: `21`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemPath`

The path to the file system directory that is accessible in Amazon SageMaker Studio. Permitted
users can access only this directory and below.

_Required_: No

_Type_: String

_Pattern_: `^\/\S*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EFSFileSystemConfig

HiddenSageMakerImage

All content copied from https://docs.aws.amazon.com/.
