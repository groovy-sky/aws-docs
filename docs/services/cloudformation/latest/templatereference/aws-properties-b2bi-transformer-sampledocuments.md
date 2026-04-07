This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer SampleDocuments

Describes a structure that contains the Amazon S3 bucket and an array of the corresponding keys used to identify the location for your sample documents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "Keys" : [ SampleDocumentKeys, ... ]
}

```

### YAML

```yaml

  BucketName: String
  Keys:
    - SampleDocumentKeys

```

## Properties

`BucketName`

Contains the Amazon S3 bucket that is used to hold your sample documents.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Keys`

Contains an array of the Amazon S3 keys used to identify the location for your
sample documents.

_Required_: Yes

_Type_: Array of [SampleDocumentKeys](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-sampledocumentkeys.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SampleDocumentKeys

Tag
