This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Fleet S3Location

Describes the S3 location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : String,
  "S3Key" : String
}

```

### YAML

```yaml

  S3Bucket: String
  S3Key: String

```

## Properties

`S3Bucket`

The S3 bucket of the S3 object.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z\.\-]*(?<!\.)$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Key`

The S3 key of the S3 object.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DomainJoinInfo

Tag
