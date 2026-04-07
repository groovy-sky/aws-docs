This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint MicrosoftSqlServerSettings

Provides information that defines a Microsoft SQL Server endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For information about other available settings, see
[Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and
[Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BcpPacketSize" : Integer,
  "ControlTablesFileGroup" : String,
  "DatabaseName" : String,
  "ForceLobLookup" : Boolean,
  "Password" : String,
  "Port" : Integer,
  "QuerySingleAlwaysOnNode" : Boolean,
  "ReadBackupOnly" : Boolean,
  "SafeguardPolicy" : String,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String,
  "ServerName" : String,
  "TlogAccessMode" : String,
  "TrimSpaceInChar" : Boolean,
  "UseBcpFullLoad" : Boolean,
  "Username" : String,
  "UseThirdPartyBackupDevice" : Boolean
}

```

### YAML

```yaml

  BcpPacketSize: Integer
  ControlTablesFileGroup: String
  DatabaseName: String
  ForceLobLookup: Boolean
  Password: String
  Port: Integer
  QuerySingleAlwaysOnNode: Boolean
  ReadBackupOnly: Boolean
  SafeguardPolicy: String
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String
  ServerName: String
  TlogAccessMode: String
  TrimSpaceInChar: Boolean
  UseBcpFullLoad: Boolean
  Username: String
  UseThirdPartyBackupDevice: Boolean

```

## Properties

`BcpPacketSize`

The maximum size of the packets (in bytes) used to transfer data using BCP.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlTablesFileGroup`

Specifies a file group for the AWS DMS internal tables. When the replication task starts,
all the internal AWS DMS control tables (awsdms\_ apply\_exception, awsdms\_apply,
awsdms\_changes) are created for the specified file group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

Database name for the endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForceLobLookup`

Forces LOB lookup on inline LOB.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

Endpoint connection password.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

Endpoint TCP port.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuerySingleAlwaysOnNode`

Cleans and recreates table metadata information on the replication instance when a
mismatch occurs. An example is a situation where running an alter DDL statement on a table
might result in different information about the table cached in the replication
instance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadBackupOnly`

When this attribute is set to `Y`, AWS DMS only reads changes from transaction
log backups and doesn't read from the active transaction log file during ongoing
replication. Setting this parameter to `Y` enables you to control active
transaction log file growth during full load and ongoing replication tasks. However, it can
add some source latency to ongoing replication.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SafeguardPolicy`

Use this attribute to minimize the need to access the backup log and enable AWS DMS to
prevent truncation using one of the following two methods.

_Start transactions in the database:_ This is the default method.
When this method is used, AWS DMS prevents TLOG truncation by mimicking a transaction in the
database. As long as such a transaction is open, changes that appear after the transaction
started aren't truncated. If you need Microsoft Replication to be enabled in your database,
then you must choose this method.

_Exclusively use sp\_repldone within a single task_: When this method
is used, AWS DMS reads the changes and then uses sp\_repldone to mark the TLOG transactions as
ready for truncation. Although this method doesn't involve any transactional activities, it
can only be used when Microsoft Replication isn't running. Also, when using this method,
only one AWS DMS task can access the database at any given time. Therefore, if you need to
run parallel AWS DMS tasks against the same database, use the default method.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value of the AWS Secrets Manager
secret that allows access to the SQL Server endpoint.

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

The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MicrosoftSQLServer endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

Fully qualified domain name of the endpoint. For an Amazon RDS SQL Server instance, this is
the output of [DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md), in the `Endpoint.Address` field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlogAccessMode`

Indicates the mode used to fetch CDC data.

_Required_: No

_Type_: String

_Allowed values_: `BackupOnly | PreferBackup | PreferTlog | TlogOnly`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrimSpaceInChar`

Use the `TrimSpaceInChar` source endpoint setting to right-trim data on CHAR
and NCHAR data types during migration. Setting `TrimSpaceInChar` does not
left-trim data. The default value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBcpFullLoad`

Use this to attribute to transfer data for full-load operations using BCP. When the
target table contains an identity column that does not exist in the source table, you must
disable the use BCP for loading table option.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

Endpoint connection user name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseThirdPartyBackupDevice`

When this attribute is set to `Y`, DMS processes third-party transaction log
backups if they are created in native format.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KinesisSettings

MongoDbSettings
