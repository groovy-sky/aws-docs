This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Capability S3Location

Specifies the details for the Amazon S3 file location that is being used with AWS B2B Data Interchange. File
locations in Amazon S3 are identified using a combination of the bucket and key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "Key" : String
}

```

### YAML

```yaml

  BucketName: String
  Key: String

```

## Properties

`BucketName`

Specifies the name of the Amazon S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

Specifies the Amazon S3 key for the file location.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EdiType

Tag
