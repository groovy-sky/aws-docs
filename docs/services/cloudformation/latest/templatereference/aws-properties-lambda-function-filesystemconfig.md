This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function FileSystemConfig

Details about the connection between a Lambda function and an [Amazon EFS file system](../../../lambda/latest/dg/configuration-filesystem.md) or an [Amazon S3 Files file system](../../../lambda/latest/dg/configuration-filesystem.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "LocalMountPath" : String
}

```

### YAML

```yaml

  Arn: String
  LocalMountPath: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the Amazon EFS or Amazon S3 Files access point that provides access to the file
system.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:elasticfilesystem:(eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:\d{12}:access-point/fsap-[a-f0-9]{17}$|^arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}$`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalMountPath`

The path where the function can access the file system, starting with `/mnt/`.

_Required_: Yes

_Type_: String

_Pattern_: `^/mnt/[a-zA-Z0-9-_.]+$`

_Maximum_: `160`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EphemeralStorage

FunctionScalingConfig

All content copied from https://docs.aws.amazon.com/.
