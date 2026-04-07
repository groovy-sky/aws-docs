This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket StorageClassAnalysis

Specifies data related to access patterns to be collected and made available to analyze the
tradeoffs between different storage classes for an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataExport" : DataExport
}

```

### YAML

```yaml

  DataExport:
    DataExport

```

## Properties

`DataExport`

Specifies how data related to the storage class analysis for an Amazon S3 bucket should be
exported.

_Required_: No

_Type_: [DataExport](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-dataexport.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SseKmsEncryptedObjects

Tag
