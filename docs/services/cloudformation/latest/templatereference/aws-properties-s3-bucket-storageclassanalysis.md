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

_Type_: [DataExport](aws-properties-s3-bucket-dataexport.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SseKmsEncryptedObjects

Tag

All content copied from https://docs.aws.amazon.com/.
