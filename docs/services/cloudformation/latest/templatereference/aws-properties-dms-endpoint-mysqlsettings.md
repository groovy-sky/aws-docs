This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint MySqlSettings

Provides information that defines a MySQL endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For information about other available settings, see
[Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and
[Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AfterConnectScript" : String,
  "CleanSourceMetadataOnMismatch" : Boolean,
  "EventsPollInterval" : Integer,
  "MaxFileSize" : Integer,
  "ParallelLoadThreads" : Integer,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String,
  "ServerTimezone" : String,
  "TargetDbType" : String
}

```

### YAML

```yaml

  AfterConnectScript: String
  CleanSourceMetadataOnMismatch: Boolean
  EventsPollInterval: Integer
  MaxFileSize: Integer
  ParallelLoadThreads: Integer
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String
  ServerTimezone: String
  TargetDbType: String

```

## Properties

`AfterConnectScript`

Specifies a script to run immediately after AWS DMS connects to the endpoint. The
migration task continues running regardless if the SQL statement succeeds or fails.

For this parameter, provide the code of the script itself, not the name of a file
containing the script.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CleanSourceMetadataOnMismatch`

Cleans and recreates table metadata information on the replication instance when a
mismatch occurs. For example, in a situation where running an alter DDL on the table could
result in different information about the table cached in the replication instance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventsPollInterval`

Specifies how often to check the binary log for new changes/events when the database is
idle. The default is five seconds.

Example: `eventsPollInterval=5;`

In the example, AWS DMS checks for changes in the binary logs every five seconds.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFileSize`

Specifies the maximum size (in KB) of any .csv file used to transfer data to a
MySQL-compatible database.

Example: `maxFileSize=512`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelLoadThreads`

Improves performance when loading data into the MySQL-compatible target database.
Specifies how many threads to use to load the data into the MySQL-compatible target
database. Setting a large number of threads can have an adverse effect on database
performance, because a separate connection is required for each thread. The default is
one.

Example: `parallelLoadThreads=1`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value of the AWS Secrets Manager secret
that allows access to the MySQL endpoint.

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

The full ARN, partial ARN, or display name of the `SecretsManagerSecret`
that contains the MySQL endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerTimezone`

Specifies the time zone for the source MySQL database.

Example: `serverTimezone=US/Pacific;`

Note: Do not enclose time zones in single quotes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDbType`

Specifies where to migrate source tables on the target, either to a single database or
multiple databases. If you specify `SPECIFIC_DATABASE`, specify the database
name using the `DatabaseName` parameter of the `Endpoint`
object.

Example: `targetDbType=MULTIPLE_DATABASES`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MongoDbSettings

NeptuneSettings
