This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationReferenceDataSource S3ReferenceDataSource

Identifies the S3 bucket and object that contains the reference data. Also identifies
the IAM role Amazon Kinesis Analytics can assume to read this object on your
behalf.

An Amazon Kinesis Analytics application loads reference data only once. If the data
changes, you call the [UpdateApplication](../../../kinesisanalytics/latest/dev/api-updateapplication.md) operation to trigger reloading of data into your
application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketARN" : String,
  "FileKey" : String,
  "ReferenceRoleARN" : String
}

```

### YAML

```yaml

  BucketARN: String
  FileKey: String
  ReferenceRoleARN: String

```

## Properties

`BucketARN`

Amazon Resource Name (ARN) of the S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileKey`

Object key name containing reference data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReferenceRoleARN`

ARN of the IAM role that the service can assume to read data on your behalf. This role
must have permission for the `s3:GetObject` action on the object and trust
policy that allows Amazon Kinesis Analytics service principal to assume this
role.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceSchema

Next

All content copied from https://docs.aws.amazon.com/.
