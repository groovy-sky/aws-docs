This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationS3 S3Config

Specifies the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that DataSync uses to access your S3 bucket.

For more information, see [Providing\
DataSync access to S3 buckets](../../../datasync/latest/userguide/create-s3-location.md#create-s3-location-access).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketAccessRoleArn" : String
}

```

### YAML

```yaml

  BucketAccessRoleArn: String

```

## Properties

`BucketAccessRoleArn`

Specifies the ARN of the IAM role that DataSync uses to access
your S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataSync::LocationS3

Tag

All content copied from https://docs.aws.amazon.com/.
