This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::AppImageConfig FileSystemConfig

The Amazon Elastic File System storage configuration for a SageMaker AI image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultGid" : Integer,
  "DefaultUid" : Integer,
  "MountPath" : String
}

```

### YAML

```yaml

  DefaultGid: Integer
  DefaultUid: Integer
  MountPath: String

```

## Properties

`DefaultGid`

The default POSIX group ID (GID). If not specified, defaults to `100`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultUid`

The default POSIX user ID (UID). If not specified, defaults to `1000`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MountPath`

The path within the image to mount the user's EFS home directory. The directory
should be empty. If not specified, defaults to _/home/sagemaker-user_.

_Required_: No

_Type_: String

_Pattern_: `^/.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomImageContainerEnvironmentVariable

JupyterLabAppImageConfig

All content copied from https://docs.aws.amazon.com/.
