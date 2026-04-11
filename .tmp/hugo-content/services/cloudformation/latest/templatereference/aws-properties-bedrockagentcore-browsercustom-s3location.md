This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::BrowserCustom S3Location

The Amazon S3 location for storing data. This structure defines where in Amazon S3 data is stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  Bucket: String
  Prefix: String

```

## Properties

`Bucket`

The name of the Amazon S3 bucket. This bucket contains the stored data.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Prefix`

The prefix for objects in the Amazon S3 bucket. This prefix is added to the object keys to organize the data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordingConfig

VpcConfig

All content copied from https://docs.aws.amazon.com/.
