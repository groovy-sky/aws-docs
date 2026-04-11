This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket AccelerateConfiguration

Configures the transfer acceleration state for an Amazon S3 bucket. For more information, see [Amazon S3 Transfer\
Acceleration](../../../s3/latest/dev/transfer-acceleration.md) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccelerationStatus" : String
}

```

### YAML

```yaml

  AccelerationStatus: String

```

## Properties

`AccelerationStatus`

Specifies the transfer acceleration status of the bucket.

_Required_: Yes

_Type_: String

_Allowed values_: `Enabled | Suspended`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- AWS::S3::Bucket [Examples](../userguide/aws-properties-s3-bucket.md#aws-properties-s3-bucket--examples)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AbortIncompleteMultipartUpload

AccessControlTranslation

All content copied from https://docs.aws.amazon.com/.
