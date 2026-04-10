This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task ManifestConfigSourceS3

The `ManifestConfigSourceS3` property type specifies Property description not available. for an [AWS::DataSync::Task](aws-resource-datasync-task.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketAccessRoleArn" : String,
  "ManifestObjectPath" : String,
  "ManifestObjectVersionId" : String,
  "S3BucketArn" : String
}

```

### YAML

```yaml

  BucketAccessRoleArn: String
  ManifestObjectPath: String
  ManifestObjectVersionId: String
  S3BucketArn: String

```

## Properties

`BucketAccessRoleArn`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestObjectPath`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}\p{C}]*$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestObjectVersionId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketArn`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):(s3|s3-outposts):[a-z\-0-9]*:[0-9]*:.*$`

_Maximum_: `156`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManifestConfig

Options

All content copied from https://docs.aws.amazon.com/.
