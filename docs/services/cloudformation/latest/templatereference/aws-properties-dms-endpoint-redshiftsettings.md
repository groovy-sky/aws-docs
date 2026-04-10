This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint RedshiftSettings

Provides information that defines an Amazon Redshift endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about other available settings, see
[Extra connection attributes when using Amazon Redshift as a target for AWS DMS](../../../dms/latest/userguide/chap-target-redshift.md#CHAP_Target.Redshift.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcceptAnyDate" : Boolean,
  "AfterConnectScript" : String,
  "BucketFolder" : String,
  "BucketName" : String,
  "CaseSensitiveNames" : Boolean,
  "CompUpdate" : Boolean,
  "ConnectionTimeout" : Integer,
  "DateFormat" : String,
  "EmptyAsNull" : Boolean,
  "EncryptionMode" : String,
  "ExplicitIds" : Boolean,
  "FileTransferUploadStreams" : Integer,
  "LoadTimeout" : Integer,
  "MapBooleanAsBoolean" : Boolean,
  "MaxFileSize" : Integer,
  "RemoveQuotes" : Boolean,
  "ReplaceChars" : String,
  "ReplaceInvalidChars" : String,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String,
  "ServerSideEncryptionKmsKeyId" : String,
  "ServiceAccessRoleArn" : String,
  "TimeFormat" : String,
  "TrimBlanks" : Boolean,
  "TruncateColumns" : Boolean,
  "WriteBufferSize" : Integer
}

```

### YAML

```yaml

  AcceptAnyDate: Boolean
  AfterConnectScript: String
  BucketFolder: String
  BucketName: String
  CaseSensitiveNames: Boolean
  CompUpdate: Boolean
  ConnectionTimeout: Integer
  DateFormat: String
  EmptyAsNull: Boolean
  EncryptionMode: String
  ExplicitIds: Boolean
  FileTransferUploadStreams: Integer
  LoadTimeout: Integer
  MapBooleanAsBoolean:
    Boolean
  MaxFileSize: Integer
  RemoveQuotes: Boolean
  ReplaceChars: String
  ReplaceInvalidChars: String
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String
  ServerSideEncryptionKmsKeyId: String
  ServiceAccessRoleArn: String
  TimeFormat: String
  TrimBlanks: Boolean
  TruncateColumns: Boolean
  WriteBufferSize: Integer

```

## Properties

`AcceptAnyDate`

A value that indicates to allow any date format, including invalid formats such as
00/00/00 00:00:00, to be loaded without generating an error. You can choose
`true` or `false` (the default).

This parameter applies only to TIMESTAMP and DATE columns. Always use ACCEPTANYDATE with
the DATEFORMAT parameter. If the date format for the data doesn't match the DATEFORMAT
specification, Amazon Redshift inserts a NULL value into that field.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AfterConnectScript`

Code to run after connecting. This parameter should contain the code itself, not the
name of a file containing the code.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketFolder`

An S3 folder where the comma-separated-value (.csv) files are stored before being
uploaded to the target Redshift cluster.

For full load mode, AWS DMS converts source records into .csv files and loads them to the
_BucketFolder/TableID_ path. AWS DMS uses the Redshift
`COPY` command to upload the .csv files to the target table. The files are
deleted once the `COPY` operation has finished. For more information, see [COPY](../../../redshift/latest/dg/r-copy.md) in the
_Amazon Redshift Database Developer Guide_.

For change-data-capture (CDC) mode, AWS DMS creates a _NetChanges_
table, and loads the .csv files to this _BucketFolder/NetChangesTableID_
path.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

The name of the intermediate S3 bucket used to store .csv files before uploading data to
Redshift.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaseSensitiveNames`

If Amazon Redshift is configured to support case sensitive schema names, set
`CaseSensitiveNames` to `true`. The default is
`false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompUpdate`

If you set `CompUpdate` to `true` Amazon Redshift applies
automatic compression if the table is empty. This applies even if the table columns already
have encodings other than `RAW`. If you set `CompUpdate` to
`false`, automatic compression is disabled and existing column encodings
aren't changed. The default is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionTimeout`

A value that sets the amount of time to wait (in milliseconds) before timing out,
beginning from when you initially establish a connection.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateFormat`

The date format that you are using. Valid values are `auto` (case-sensitive),
your date format string enclosed in quotes, or NULL. If this parameter is left unset
(NULL), it defaults to a format of 'YYYY-MM-DD'. Using `auto` recognizes most
strings, even some that aren't supported when you use a date format string.

If your date and time values use formats different from each other, set this to
`auto`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmptyAsNull`

A value that specifies whether AWS DMS should migrate empty CHAR and VARCHAR fields as
NULL. A value of `true` sets empty CHAR and VARCHAR fields to null. The default
is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionMode`

The type of server-side encryption that you want to use for your data. This encryption
type is part of the endpoint settings or the extra connections attributes for Amazon S3.
You can choose either `SSE_S3` (the default) or `SSE_KMS`.

###### Note

For the `ModifyEndpoint` operation, you can change the existing value of
the `EncryptionMode` parameter from `SSE_KMS` to
`SSE_S3`. But you can’t change the existing value from `SSE_S3`
to `SSE_KMS`.

To use `SSE_S3`, create an AWS Identity and Access Management (IAM) role with a policy that allows
`"arn:aws:s3:::*"` to use the following actions: `"s3:PutObject",
            "s3:ListBucket"`

_Required_: No

_Type_: String

_Allowed values_: `sse-s3 | sse-kms`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExplicitIds`

This setting is only valid for a full-load migration task. Set `ExplicitIds`
to `true` to have tables with `IDENTITY` columns override their
auto-generated values with explicit values loaded from the source data files used to
populate the tables. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileTransferUploadStreams`

The number of threads used to upload a single file. This parameter accepts a value from
1 through 64. It defaults to 10.

The number of parallel streams used to upload a single .csv file to an S3 bucket using
S3 Multipart Upload. For more information, see [Multipart upload\
overview](../../../s3/latest/dev/mpuoverview.md).

`FileTransferUploadStreams` accepts a value from 1 through 64. It defaults to
10.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadTimeout`

The amount of time to wait (in seconds) before timing out of operations performed
by AWS DMS on a Redshift cluster, such as Redshift COPY, INSERT, DELETE, and UPDATE.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapBooleanAsBoolean`

When true, lets Redshift migrate the boolean type as boolean. By default, Redshift
migrates booleans as `varchar(1)`. You must set this setting on both the source
and target endpoints for it to take effect.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFileSize`

The maximum size (in KB) of any .csv file used to load data on an S3 bucket and transfer
data to Amazon Redshift. It defaults to 1048576KB (1 GB).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveQuotes`

A value that specifies to remove surrounding quotation marks from strings in the
incoming data. All characters within the quotation marks, including delimiters, are
retained. Choose `true` to remove quotation marks. The default is
`false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceChars`

A value that specifies to replaces the invalid characters specified in
`ReplaceInvalidChars`, substituting the specified characters instead. The
default is `"?"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceInvalidChars`

A list of characters that you want to replace. Use with
`ReplaceChars`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value of the AWS Secrets Manager secret
that allows access to the Amazon Redshift endpoint.

###### Note

You can specify one of two sets of values for these permissions. You can specify the
values for this setting and `SecretsManagerSecretId`. Or you can specify
clear-text values for `UserName`, `Password`,
`ServerName`, and `Port`. You can't specify both.

For more information on creating this `SecretsManagerSecret`, the corresponding
`SecretsManagerAccessRoleArn`, and the `SecretsManagerSecretId`
that is required to access it, see
[Using secrets to access AWS Database Migration Service resources](../../../dms/latest/userguide/chap-security.md#security-iam-secretsmanager)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecretId`

The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the Amazon Redshift endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionKmsKeyId`

The AWS KMS key ID. If you are using `SSE_KMS` for the
`EncryptionMode`, provide this key ID. The key that you use needs an attached
policy that enables IAM user permissions and allows use of the key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccessRoleArn`

The Amazon Resource Name (ARN) of the IAM role that has access to the Amazon Redshift
service. The role must allow the `iam:PassRole` action.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeFormat`

The time format that you want to use. Valid values are `auto`
(case-sensitive), `'timeformat_string'`, `'epochsecs'`, or
`'epochmillisecs'`. It defaults to 10. Using `auto` recognizes
most strings, even some that aren't supported when you use a time format string.

If your date and time values use formats different from each other, set this parameter
to `auto`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrimBlanks`

A value that specifies to remove the trailing white space characters from a VARCHAR
string. This parameter applies only to columns with a VARCHAR data type. Choose
`true` to remove unneeded white space. The default is
`false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TruncateColumns`

A value that specifies to truncate data in columns to the appropriate number of
characters, so that the data fits in the column. This parameter applies only to columns
with a VARCHAR or CHAR data type, and rows with a size of 4 MB or less. Choose
`true` to truncate data. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteBufferSize`

The size (in KB) of the in-memory file write buffer used when generating .csv files on
the local disk at the DMS replication instance. The default value is 1000 (buffer size is
1000KB).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedisSettings

S3Settings

All content copied from https://docs.aws.amazon.com/.
