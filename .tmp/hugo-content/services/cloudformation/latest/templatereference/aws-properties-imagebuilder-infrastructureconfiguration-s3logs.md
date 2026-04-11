This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::InfrastructureConfiguration S3Logs

Amazon S3 logging configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3BucketName" : String,
  "S3KeyPrefix" : String
}

```

### YAML

```yaml

  S3BucketName: String
  S3KeyPrefix: String

```

## Properties

`S3BucketName`

The S3 bucket in which to store the logs.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KeyPrefix`

The Amazon S3 path to the bucket where the logs are stored.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Placement

AWS::ImageBuilder::LifecyclePolicy

All content copied from https://docs.aws.amazon.com/.
