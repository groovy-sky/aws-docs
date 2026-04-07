This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Script S3Location

The location in Amazon S3 where build or script files can be stored for access by
Amazon GameLift.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Key" : String,
  "ObjectVersion" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  Bucket: String
  Key: String
  ObjectVersion: String
  RoleArn: String

```

## Properties

`Bucket`

An Amazon S3 bucket identifier. Thename of the S3 bucket.

###### Note

Amazon GameLift Servers doesn't support uploading from Amazon S3 buckets with names that contain a dot
(.).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The name of the zip file that contains the build files or script files.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectVersion`

The version of the file, if object versioning is turned on for the bucket. Amazon GameLift Servers uses
this information when retrieving files from an S3 bucket that you own. Use this
parameter to specify a specific version of the file. If not set, the latest version of
the file is retrieved.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) for an IAM role that
allows Amazon GameLift Servers to access the S3 bucket.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GameLift::Script

Tag
