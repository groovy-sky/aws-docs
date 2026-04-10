This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime S3Location

The Amazon S3 location for storing data. This structure defines where in Amazon S3 data is stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Prefix" : String,
  "VersionId" : String
}

```

### YAML

```yaml

  Bucket: String
  Prefix: String
  VersionId: String

```

## Properties

`Bucket`

The name of the Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix for objects in the Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionId`

The version ID of the Amazon Amazon S3 object. If not specified, the latest version of the object is used.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestHeaderConfiguration

SessionStorageConfiguration

All content copied from https://docs.aws.amazon.com/.
