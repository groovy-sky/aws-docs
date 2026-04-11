This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe S3LogDestination

Represents the Amazon S3 logging configuration settings for the pipe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketOwner" : String,
  "OutputFormat" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketOwner: String
  OutputFormat: String
  Prefix: String

```

## Properties

`BucketName`

The name of the Amazon S3 bucket to which EventBridge delivers the log
records for the pipe.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketOwner`

The AWS account that owns the Amazon S3 bucket to which EventBridge delivers the log records for the pipe.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputFormat`

The format EventBridge uses for the log records.

EventBridge currently only supports `json` formatting.

_Required_: No

_Type_: String

_Allowed values_: `json | plain | w3c`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix text with which to begin Amazon S3 log object names.

For more information, see [Organizing objects using\
prefixes](../../../s3/latest/userguide/using-prefixes.md) in the _Amazon Simple Storage Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PlacementStrategy

SageMakerPipelineParameter

All content copied from https://docs.aws.amazon.com/.
