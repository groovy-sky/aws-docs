This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space CustomFileSystem

A file system, created by you, that you assign to a user profile or space for an Amazon SageMaker AI Domain. Permitted users can access this file system in Amazon SageMaker AI
Studio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EFSFileSystem" : EFSFileSystem,
  "FSxLustreFileSystem" : FSxLustreFileSystem,
  "S3FileSystem" : S3FileSystem
}

```

### YAML

```yaml

  EFSFileSystem:
    EFSFileSystem
  FSxLustreFileSystem:
    FSxLustreFileSystem
  S3FileSystem:
    S3FileSystem

```

## Properties

`EFSFileSystem`

A custom file system in Amazon EFS.

_Required_: No

_Type_: [EFSFileSystem](aws-properties-sagemaker-space-efsfilesystem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FSxLustreFileSystem`

A custom file system in Amazon FSx for Lustre.

_Required_: No

_Type_: [FSxLustreFileSystem](aws-properties-sagemaker-space-fsxlustrefilesystem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3FileSystem`

A custom file system in Amazon S3. This is only supported in Amazon SageMaker Unified Studio.

_Required_: No

_Type_: [S3FileSystem](aws-properties-sagemaker-space-s3filesystem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeRepository

CustomImage

All content copied from https://docs.aws.amazon.com/.
