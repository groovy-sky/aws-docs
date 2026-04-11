This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint PostgreSqlSettings

Provides information that defines a PostgreSQL endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For information about other available settings, see
[Extra connection attributes when using PostgreSQL as a source for AWS DMS](../../../dms/latest/userguide/chap-source-postgresql.md#CHAP_Source.PostgreSQL.ConnectionAttrib) and
[Extra connection attributes when using PostgreSQL as a target for AWS DMS](../../../dms/latest/userguide/chap-target-postgresql.md#CHAP_Target.PostgreSQL.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AfterConnectScript" : String,
  "BabelfishDatabaseName" : String,
  "CaptureDdls" : Boolean,
  "DatabaseMode" : String,
  "DdlArtifactsSchema" : String,
  "ExecuteTimeout" : Integer,
  "FailTasksOnLobTruncation" : Boolean,
  "HeartbeatEnable" : Boolean,
  "HeartbeatFrequency" : Integer,
  "HeartbeatSchema" : String,
  "MapBooleanAsBoolean" : Boolean,
  "MaxFileSize" : Integer,
  "PluginName" : String,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerSecretId" : String,
  "SlotName" : String
}

```

### YAML

```yaml

  AfterConnectScript: String
  BabelfishDatabaseName: String
  CaptureDdls: Boolean
  DatabaseMode: String
  DdlArtifactsSchema: String
  ExecuteTimeout: Integer
  FailTasksOnLobTruncation: Boolean
  HeartbeatEnable: Boolean
  HeartbeatFrequency: Integer
  HeartbeatSchema: String
  MapBooleanAsBoolean:
    Boolean
  MaxFileSize: Integer
  PluginName: String
  SecretsManagerAccessRoleArn: String
  SecretsManagerSecretId: String
  SlotName: String

```

## Properties

`AfterConnectScript`

For use with change data capture (CDC) only, this attribute has AWS DMS bypass foreign
keys and user triggers to reduce the time it takes to bulk load data.

Example: `afterConnectScript=SET session_replication_role='replica'`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BabelfishDatabaseName`

The Babelfish for Aurora PostgreSQL database name for the endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptureDdls`

To capture DDL events, AWS DMS creates various artifacts in
the PostgreSQL database when the task starts. You can later
remove these artifacts.

If this value is set to `True`, you don't have to create tables or
triggers on the source database.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseMode`

Specifies the default behavior of the replication's handling of PostgreSQL- compatible
endpoints that require some additional configuration, such as Babelfish endpoints.

_Required_: No

_Type_: String

_Allowed values_: `default | babelfish`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DdlArtifactsSchema`

The schema in which the operational DDL database artifacts are created.

The default value is `public`.

Example: `ddlArtifactsSchema=xyzddlschema;`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecuteTimeout`

Sets the client statement timeout for the PostgreSQL instance, in seconds. The default
value is 60 seconds.

Example: `executeTimeout=100;`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailTasksOnLobTruncation`

When set to `true`, this value causes a task to fail if the actual size of a
LOB column is greater than the specified `LobMaxSize`.

The default value is `false`.

If task is set to Limited LOB mode and this option is set to true, the task fails
instead of truncating the LOB data.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeartbeatEnable`

The write-ahead log (WAL) heartbeat feature mimics a dummy transaction. By doing this,
it prevents idle logical replication slots from holding onto old WAL logs, which can result
in storage full situations on the source. This heartbeat keeps `restart_lsn`
moving and prevents storage full scenarios.

The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeartbeatFrequency`

Sets the WAL heartbeat frequency (in minutes).

The default value is 5 minutes.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeartbeatSchema`

Sets the schema in which the heartbeat artifacts are created.

The default value is `public`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapBooleanAsBoolean`

When true, lets PostgreSQL migrate the boolean type as boolean. By default, PostgreSQL
migrates booleans as `varchar(5)`. You must set this setting on both the source
and target endpoints for it to take effect.

The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFileSize`

Specifies the maximum size (in KB) of any .csv file used to transfer data to
PostgreSQL.

The default value is 32,768 KB (32 MB).

Example: `maxFileSize=512`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PluginName`

Specifies the plugin to use to create a replication slot.

The default value is `pglogical`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value of the AWS Secrets Manager secret
that allows access to the PostgreSQL endpoint.

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

The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the PostgreSQL endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotName`

Sets the name of a previously created logical replication slot for a change data capture
(CDC) load of the PostgreSQL source instance.

When used with the `CdcStartPosition` request parameter for the AWS DMS API ,
this attribute also makes it possible to use native CDC start points. DMS verifies that the
specified logical replication slot exists before starting the CDC load task. It also
verifies that the task was created with a valid setting of `CdcStartPosition`.
If the specified slot doesn't exist or the task doesn't have a valid
`CdcStartPosition` setting, DMS raises an error.

For more information about setting the `CdcStartPosition` request parameter,
see [Determining a CDC native start point](../../../dms/latest/userguide/chap-task-cdc.md#CHAP_Task.CDC.StartPoint.Native) in the _AWS Database Migration Service User_
_Guide_. For more information about using `CdcStartPosition`, see
[CreateReplicationTask](../../../../reference/dms/latest/apireference/api-createreplicationtask.md), [StartReplicationTask](../../../../reference/dms/latest/apireference/api-startreplicationtask.md), and [ModifyReplicationTask](../../../../reference/dms/latest/apireference/api-modifyreplicationtask.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OracleSettings

RedisSettings

All content copied from https://docs.aws.amazon.com/.
