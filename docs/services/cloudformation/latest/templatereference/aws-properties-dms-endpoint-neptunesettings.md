This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint NeptuneSettings

Provides information that defines an Amazon Neptune endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about the available settings, see
[Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorRetryDuration" : Integer,
  "IamAuthEnabled" : Boolean,
  "MaxFileSize" : Integer,
  "MaxRetryCount" : Integer,
  "S3BucketFolder" : String,
  "S3BucketName" : String,
  "ServiceAccessRoleArn" : String
}

```

### YAML

```yaml

  ErrorRetryDuration: Integer
  IamAuthEnabled: Boolean
  MaxFileSize: Integer
  MaxRetryCount: Integer
  S3BucketFolder: String
  S3BucketName: String
  ServiceAccessRoleArn: String

```

## Properties

`ErrorRetryDuration`

The number of milliseconds for AWS DMS to wait to retry a bulk-load of migrated graph
data to the Neptune target database before raising an error. The default is 250.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamAuthEnabled`

If you want IAM authorization enabled for this
endpoint, set this parameter to `true`. Then attach the appropriate IAM policy
document to your service role specified by `ServiceAccessRoleArn`. The default
is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFileSize`

The maximum size in kilobytes of migrated graph data stored in a .csv file before AWS DMS
bulk-loads the data to the Neptune target database. The default is 1,048,576 KB. If the
bulk load is successful, AWS DMS clears the bucket, ready to store the next batch of
migrated graph data.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRetryCount`

The number of times for AWS DMS to retry a bulk load of migrated graph data to the
Neptune target database before raising an error. The default is 5.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketFolder`

A folder path where you want AWS DMS to store migrated graph data in the S3 bucket
specified by `S3BucketName`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

The name of the Amazon S3 bucket where AWS DMS can temporarily store migrated graph data
in .csv files before bulk-loading it to the Neptune target database. AWS DMS maps the SQL
source data to graph data before storing it in these .csv files.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccessRoleArn`

The Amazon Resource Name (ARN) of the service role that you created for the Neptune
target endpoint. The role must allow the `iam:PassRole` action.

For more information, see
[Creating an IAM Service Role for Accessing Amazon Neptune as a Target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.ServiceRole)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MySqlSettings

OracleSettings
