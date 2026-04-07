This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint IbmDb2Settings

Provides information that defines an IBMDB2 endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about other available settings, see
[Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CurrentLsn" : String,
  "KeepCsvFiles" : Boolean,
  "LoadTimeout" : Integer,
  "MaxFileSize" : Integer,
  "MaxKBytesPerRead" : Integer,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String,
  "SetDataCaptureChanges" : Boolean,
  "WriteBufferSize" : Integer
}

```

### YAML

```yaml

  CurrentLsn: String
  KeepCsvFiles: Boolean
  LoadTimeout: Integer
  MaxFileSize: Integer
  MaxKBytesPerRead: Integer
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String
  SetDataCaptureChanges: Boolean
  WriteBufferSize: Integer

```

## Properties

`CurrentLsn`

For ongoing replication (CDC), use CurrentLSN to specify a log sequence number (LSN)
where you want the replication to start.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeepCsvFiles`

If true, AWS DMS saves any .csv files to the Db2 LUW target that were used to replicate
data. DMS uses these files for analysis and troubleshooting.

The default value is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadTimeout`

The amount of time (in milliseconds) before AWS DMS times out operations performed by DMS
on the Db2 target. The default value is 1200 (20 minutes).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFileSize`

Specifies the maximum size (in KB) of .csv files used to transfer data to Db2
LUW.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxKBytesPerRead`

Maximum number of bytes per read, as a NUMBER value. The default is 64 KB.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value ofthe AWS Secrets Manager secret
that allows access to the Db2 LUW endpoint.

###### Note

You can specify one of two sets of values for these permissions. You can specify the
values for this setting and `SecretsManagerSecretId`. Or you can specify
clear-text values for `UserName`, `Password`,
`ServerName`, and `Port`. You can't specify both.

For more information on creating this `SecretsManagerSecret`, the corresponding
`SecretsManagerAccessRoleArn`, and the `SecretsManagerSecretId`
that is required to access it, see
[Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecretId`

The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the IBMDB2 endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SetDataCaptureChanges`

Enables ongoing replication (CDC) as a BOOLEAN value. The default is true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteBufferSize`

The size (in KB) of the in-memory file write buffer used when generating .csv files on
the local disk on the DMS replication instance. The default value is 1024 (1 MB).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GcpMySQLSettings

KafkaSettings
