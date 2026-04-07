This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::Repository ImageScanningConfiguration

The image scanning configuration for a repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScanOnPush" : Boolean
}

```

### YAML

```yaml

  ScanOnPush: Boolean

```

## Properties

`ScanOnPush`

The setting that determines whether images are scanned after being pushed to a
repository. If set to `true`, images will be scanned after being pushed. If
this parameter is not specified, it will default to `false` and images will
not be scanned unless a scan is manually started.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EncryptionConfiguration

ImageTagMutabilityExclusionFilter
