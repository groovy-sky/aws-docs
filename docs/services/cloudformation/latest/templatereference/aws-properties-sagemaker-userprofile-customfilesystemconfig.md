This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile CustomFileSystemConfig

The settings for assigning a custom file system to a user profile or space for an Amazon SageMaker AI Domain. Permitted users can access this file system in Amazon SageMaker AI
Studio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EFSFileSystemConfig" : EFSFileSystemConfig,
  "FSxLustreFileSystemConfig" : FSxLustreFileSystemConfig,
  "S3FileSystemConfig" : S3FileSystemConfig
}

```

### YAML

```yaml

  EFSFileSystemConfig:
    EFSFileSystemConfig
  FSxLustreFileSystemConfig:
    FSxLustreFileSystemConfig
  S3FileSystemConfig:
    S3FileSystemConfig

```

## Properties

`EFSFileSystemConfig`

The settings for a custom Amazon EFS file system.

_Required_: No

_Type_: [EFSFileSystemConfig](aws-properties-sagemaker-userprofile-efsfilesystemconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FSxLustreFileSystemConfig`

The settings for a custom Amazon FSx for Lustre file system.

_Required_: No

_Type_: [FSxLustreFileSystemConfig](aws-properties-sagemaker-userprofile-fsxlustrefilesystemconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3FileSystemConfig`

Configuration settings for a custom Amazon S3 file system.

_Required_: No

_Type_: [S3FileSystemConfig](aws-properties-sagemaker-userprofile-s3filesystemconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeRepository

CustomImage

All content copied from https://docs.aws.amazon.com/.
