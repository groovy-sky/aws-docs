This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket AbortIncompleteMultipartUpload

Specifies the days since the initiation of an incomplete multipart upload that Amazon S3
will wait before permanently removing all parts of the upload. For more information, see
[Stopping Incomplete Multipart Uploads Using a Bucket Lifecycle Policy](../../../s3/latest/dev/mpuoverview.md#mpu-abort-incomplete-mpu-lifecycle-config) in the
_Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DaysAfterInitiation" : Integer
}

```

### YAML

```yaml

  DaysAfterInitiation: Integer

```

## Properties

`DaysAfterInitiation`

Specifies the number of days after which Amazon S3 stops an incomplete multipart
upload.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3::Bucket

AccelerateConfiguration

All content copied from https://docs.aws.amazon.com/.
