This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate S3Configuration

Specifies the configuration for experiment logging to Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  BucketName: String
  Prefix: String

```

## Properties

`BucketName`

The name of the destination bucket.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The bucket prefix.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `700`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Outputs

AWS::FIS::TargetAccountConfiguration

All content copied from https://docs.aws.amazon.com/.
