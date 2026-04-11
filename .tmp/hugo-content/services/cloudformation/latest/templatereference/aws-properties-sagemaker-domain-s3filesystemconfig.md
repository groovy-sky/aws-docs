This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain S3FileSystemConfig

Configuration for the custom Amazon S3 file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MountPath" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  MountPath: String
  S3Uri: String

```

## Properties

`MountPath`

The file system path where the Amazon S3 storage location will be mounted within the Amazon SageMaker
Studio environment.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Uri`

The Amazon S3 URI of the S3 file system configuration.

_Required_: No

_Type_: String

_Pattern_: `(s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RStudioServerProDomainSettings

SharingSettings

All content copied from https://docs.aws.amazon.com/.
