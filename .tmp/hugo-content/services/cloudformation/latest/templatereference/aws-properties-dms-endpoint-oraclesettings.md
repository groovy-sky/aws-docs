This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint OracleSettings

Provides information that defines an Oracle endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For information about other available settings, see
[Extra connection attributes when using Oracle as a source for AWS DMS](../../../dms/latest/userguide/chap-source-oracle.md#CHAP_Source.Oracle.ConnectionAttrib) and
[Extra connection attributes when using Oracle as a target for AWS DMS](../../../dms/latest/userguide/chap-target-oracle.md#CHAP_Target.Oracle.ConnectionAttrib)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessAlternateDirectly" : Boolean,
  "AdditionalArchivedLogDestId" : Integer,
  "AddSupplementalLogging" : Boolean,
  "AllowSelectNestedTables" : Boolean,
  "ArchivedLogDestId" : Integer,
  "ArchivedLogsOnly" : Boolean,
  "AsmPassword" : String,
  "AsmServer" : String,
  "AsmUser" : String,
  "CharLengthSemantics" : String,
  "DirectPathNoLog" : Boolean,
  "DirectPathParallelLoad" : Boolean,
  "EnableHomogenousTablespace" : Boolean,
  "ExtraArchivedLogDestIds" : [ Integer, ... ],
  "FailTasksOnLobTruncation" : Boolean,
  "NumberDatatypeScale" : Integer,
  "OraclePathPrefix" : String,
  "ParallelAsmReadThreads" : Integer,
  "ReadAheadBlocks" : Integer,
  "ReadTableSpaceName" : Boolean,
  "ReplacePathPrefix" : Boolean,
  "RetryInterval" : Integer,
  "SecretsManagerAccessRoleArn" : String,
  "SecretsManagerOracleAsmAccessRoleArn" : String,
  "SecretsManagerOracleAsmSecretId" : String,
  "SecretsManagerSecretId" : String,
  "SecurityDbEncryption" : String,
  "SecurityDbEncryptionName" : String,
  "SpatialDataOptionToGeoJsonFunctionName" : String,
  "StandbyDelayTime" : Integer,
  "UseAlternateFolderForOnline" : Boolean,
  "UseBFile" : Boolean,
  "UseDirectPathFullLoad" : Boolean,
  "UseLogminerReader" : Boolean,
  "UsePathPrefix" : String
}

```

### YAML

```yaml

  AccessAlternateDirectly: Boolean
  AdditionalArchivedLogDestId: Integer
  AddSupplementalLogging: Boolean
  AllowSelectNestedTables: Boolean
  ArchivedLogDestId: Integer
  ArchivedLogsOnly: Boolean
  AsmPassword: String
  AsmServer: String
  AsmUser: String
  CharLengthSemantics: String
  DirectPathNoLog: Boolean
  DirectPathParallelLoad: Boolean
  EnableHomogenousTablespace: Boolean
  ExtraArchivedLogDestIds:
    - Integer
  FailTasksOnLobTruncation: Boolean
  NumberDatatypeScale: Integer
  OraclePathPrefix: String
  ParallelAsmReadThreads: Integer
  ReadAheadBlocks: Integer
  ReadTableSpaceName: Boolean
  ReplacePathPrefix: Boolean
  RetryInterval: Integer
  SecretsManagerAccessRoleArn: String
  SecretsManagerOracleAsmAccessRoleArn: String
  SecretsManagerOracleAsmSecretId: String
  SecretsManagerSecretId: String
  SecurityDbEncryption: String
  SecurityDbEncryptionName: String
  SpatialDataOptionToGeoJsonFunctionName: String
  StandbyDelayTime: Integer
  UseAlternateFolderForOnline: Boolean
  UseBFile: Boolean
  UseDirectPathFullLoad: Boolean
  UseLogminerReader: Boolean
  UsePathPrefix: String

```

## Properties

`AccessAlternateDirectly`

Set this attribute to `false` in order to use the Binary Reader to capture
change data for an Amazon RDS for Oracle as the source. This tells the DMS instance to not
access redo logs through any specified path prefix replacement using direct file
access.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalArchivedLogDestId`

Set this attribute with `ArchivedLogDestId` in a primary/ standby setup. This
attribute is useful in the case of a switchover. In this case, AWS DMS needs to know which
destination to get archive redo logs from to read changes. This need arises because the
previous primary instance is now a standby instance after switchover.

Although AWS DMS supports the use of the Oracle `RESETLOGS` option to open the
database, never use `RESETLOGS` unless necessary. For additional information
about `RESETLOGS`, see [RMAN Data Repair Concepts](https://docs.oracle.com/en/database/oracle/oracle-database/19/bradv/rman-data-repair-concepts.html) in the _Oracle Database Backup and Recovery_
_User's Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AddSupplementalLogging`

Set this attribute to set up table-level supplemental logging for the Oracle database.
This attribute enables PRIMARY KEY supplemental logging on all tables selected for a
migration task.

If you use this option, you still need to enable database-level supplemental
logging.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowSelectNestedTables`

Set this attribute to `true` to enable replication of Oracle tables
containing columns that are nested tables or defined types.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ArchivedLogDestId`

Specifies the ID of the destination for the archived redo logs. This value should be the
same as a number in the dest\_id column of the v$archived\_log view. If you work with an
additional redo log destination, use the `AdditionalArchivedLogDestId` option to
specify the additional destination ID. Doing this improves performance by ensuring that the
correct logs are accessed from the outset.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ArchivedLogsOnly`

When this field is set to `True`, AWS DMS only accesses the archived redo logs.
If the archived redo logs are stored on Automatic Storage Management (ASM) only, the AWS DMS
user account needs to be granted ASM privileges.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AsmPassword`

For an Oracle source endpoint, your Oracle Automatic Storage Management (ASM) password.
You can set this value from the `
                            asm_user_password
                        ` value.
You set this value as part of the comma-separated value that you set to the
`Password` request parameter when you create the endpoint to access
transaction logs using Binary Reader. For more information, see [Configuration for change data capture (CDC) on an Oracle source\
database](../../../dms/latest/userguide/chap-source-oracle.md#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AsmServer`

For an Oracle source endpoint, your ASM server address. You can set this value from the
`asm_server` value. You set `asm_server` as part of the extra
connection attribute string to access an Oracle server with Binary Reader that uses ASM.
For more information, see [Configuration for change data capture (CDC) on an Oracle source\
database](../../../dms/latest/userguide/chap-source-oracle.md#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AsmUser`

For an Oracle source endpoint, your ASM user name. You can set this value from the
`asm_user` value. You set `asm_user` as part of the extra
connection attribute string to access an Oracle server with Binary Reader that uses ASM.
For more information, see [Configuration for change data capture (CDC) on an Oracle source\
database](../../../dms/latest/userguide/chap-source-oracle.md#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CharLengthSemantics`

Specifies whether the length of a character column is in bytes or in characters. To
indicate that the character column length is in characters, set this attribute to
`CHAR`. Otherwise, the character column length is in bytes.

Example: `charLengthSemantics=CHAR;`

_Required_: No

_Type_: String

_Allowed values_: `default | char | byte`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectPathNoLog`

When set to `true`, this attribute helps to increase the commit rate on the
Oracle target database by writing directly to tables and not writing a trail to database
logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectPathParallelLoad`

When set to `true`, this attribute specifies a parallel load when
`useDirectPathFullLoad` is set to `Y`. This attribute also only
applies when you use the AWS DMS parallel load feature. Note that the target table cannot
have any constraints or indexes.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableHomogenousTablespace`

Set this attribute to enable homogenous tablespace replication and create existing
tables or indexes under the same tablespace on the target.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtraArchivedLogDestIds`

Specifies the IDs of one more destinations for one or more archived redo logs. These IDs
are the values of the `dest_id` column in the `v$archived_log` view.
Use this setting with the `archivedLogDestId` extra connection attribute in a
primary-to-single setup or a primary-to-multiple-standby setup.

This setting is useful in a switchover when you use an Oracle Data Guard database as a
source. In this case, AWS DMS needs information about what destination to get archive redo
logs from to read changes. AWS DMS needs this because after the switchover the previous
primary is a standby instance. For example, in a primary-to-single standby setup you might
apply the following settings.

`archivedLogDestId=1; ExtraArchivedLogDestIds=[2]`

In a primary-to-multiple-standby setup, you might apply the following settings.

`archivedLogDestId=1; ExtraArchivedLogDestIds=[2,3,4]`

Although AWS DMS supports the use of the Oracle `RESETLOGS` option to open the
database, never use `RESETLOGS` unless it's necessary. For more information
about `RESETLOGS`, see [RMAN Data Repair Concepts](https://docs.oracle.com/en/database/oracle/oracle-database/19/bradv/rman-data-repair-concepts.html) in the _Oracle Database Backup and Recovery_
_User's Guide_.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailTasksOnLobTruncation`

When set to `true`, this attribute causes a task to fail if the actual size
of an LOB column is greater than the specified `LobMaxSize`.

If a task is set to limited LOB mode and this option is set to `true`, the
task fails instead of truncating the LOB data.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberDatatypeScale`

Specifies the number scale. You can select a scale up to 38, or you can select FLOAT. By
default, the NUMBER data type is converted to precision 38, scale 10.

Example: `numberDataTypeScale=12`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OraclePathPrefix`

Set this string attribute to the required value in order to use the Binary Reader to
capture change data for an Amazon RDS for Oracle as the source. This value specifies the
default Oracle root used to access the redo logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelAsmReadThreads`

Set this attribute to change the number of threads that DMS configures to perform a
change data capture (CDC) load using Oracle Automatic Storage Management (ASM). You can
specify an integer value between 2 (the default) and 8 (the maximum). Use this attribute
together with the `readAheadBlocks` attribute.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadAheadBlocks`

Set this attribute to change the number of read-ahead blocks that DMS configures to
perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).
You can specify an integer value between 1000 (the default) and 200,000 (the
maximum).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadTableSpaceName`

When set to `true`, this attribute supports tablespace replication.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplacePathPrefix`

Set this attribute to true in order to use the Binary Reader to capture change data for
an Amazon RDS for Oracle as the source. This setting tells DMS instance to replace the default
Oracle root with the specified `usePathPrefix` setting to access the redo
logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryInterval`

Specifies the number of seconds that the system waits before resending a query.

Example: `retryInterval=6;`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerAccessRoleArn`

The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the
trusted entity and grants the required permissions to access the value in
`SecretsManagerSecret`. The role must allow the `iam:PassRole` action.
`SecretsManagerSecret` has the value of the AWS Secrets Manager secret
that allows access to the Oracle endpoint.

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

`SecretsManagerOracleAsmAccessRoleArn`

Required only if your Oracle endpoint uses Advanced Storage Manager (ASM). The full ARN
of the IAM role that specifies AWS DMS as the trusted entity and grants the required
permissions to access the `SecretsManagerOracleAsmSecret`. This
`SecretsManagerOracleAsmSecret` has the secret value that allows access to
the Oracle ASM of the endpoint.

###### Note

You can specify one of two sets of values for these permissions. You can specify the
values for this setting and `SecretsManagerOracleAsmSecretId`. Or you can
specify clear-text values for `AsmUser`, `AsmPassword`, and
`AsmServerName`. You can't specify both.

For more information on creating this `SecretsManagerOracleAsmSecret`, the corresponding
`SecretsManagerOracleAsmAccessRoleArn`, and the `SecretsManagerOracleAsmSecretId`
that is required to access it, see
[Using secrets to access AWS Database Migration Service resources](../../../dms/latest/userguide/chap-security.md#security-iam-secretsmanager)
in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerOracleAsmSecretId`

Required only if your Oracle endpoint uses Advanced Storage Manager (ASM). The full ARN, partial ARN, or display name of the
`SecretsManagerOracleAsmSecret` that contains the Oracle ASM connection details for the Oracle endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManagerSecretId`

The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains
the Oracle endpoint connection details.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityDbEncryption`

For an Oracle source endpoint, the transparent data encryption (TDE) password required
by AWM DMS to access Oracle redo logs encrypted by TDE using Binary Reader. It is also the
`
                            TDE_Password
                        ` part of the comma-separated value you
set to the `Password` request parameter when you create the endpoint. The
`SecurityDbEncryptian` setting is related to this
`SecurityDbEncryptionName` setting. For more information, see [Supported encryption methods for using Oracle as a source for\
AWS DMS](../../../dms/latest/userguide/chap-source-oracle.md#CHAP_Source.Oracle.Encryption) in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityDbEncryptionName`

For an Oracle source endpoint, the name of a key used for the transparent data
encryption (TDE) of the columns and tablespaces in an Oracle source database that is
encrypted using TDE. The key value is the value of the `SecurityDbEncryption`
setting. For more information on setting the key name value of
`SecurityDbEncryptionName`, see the information and example for setting the
`securityDbEncryptionName` extra connection attribute in [Supported encryption methods for using Oracle as a source for\
AWS DMS](../../../dms/latest/userguide/chap-source-oracle.md#CHAP_Source.Oracle.Encryption) in the _AWS Database Migration Service User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpatialDataOptionToGeoJsonFunctionName`

Use this attribute to convert `SDO_GEOMETRY` to `GEOJSON` format.
By default, DMS calls the `SDO2GEOJSON` custom function if present and
accessible. Or you can create your own custom function that mimics the operation of
`SDOGEOJSON` and set `SpatialDataOptionToGeoJsonFunctionName` to
call it instead.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandbyDelayTime`

Use this attribute to specify a time in minutes for the delay in standby sync. If the
source is an Oracle Active Data Guard standby database, use this attribute to specify the
time lag between primary and standby databases.

In AWS DMS, you can create an Oracle CDC task that uses an Active Data Guard standby
instance as a source for replicating ongoing changes. Doing this eliminates the need to
connect to an active database that might be in production.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAlternateFolderForOnline`

Set this attribute to `true` in order to use the Binary Reader to capture
change data for an Amazon RDS for Oracle as the source. This tells the DMS instance to use any
specified prefix replacement to access all online redo logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBFile`

Set this attribute to True to capture change data using the Binary Reader utility. Set
`UseLogminerReader` to False to set this attribute to True. To use Binary
Reader with Amazon RDS for Oracle as the source, you set additional attributes. For more
information about using this setting with Oracle Automatic Storage Management (ASM), see
[Using Oracle LogMiner or AWS DMS Binary Reader for\
CDC](../../../dms/latest/userguide/chap-source-oracle.md#CHAP_Source.Oracle.CDC).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseDirectPathFullLoad`

Set this attribute to True to have AWS DMS use a direct path full load. Specify this value
to use the direct path protocol in the Oracle Call Interface (OCI). By using this OCI
protocol, you can bulk-load Oracle target tables during a full load.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLogminerReader`

Set this attribute to True to capture change data using the Oracle LogMiner utility (the
default). Set this attribute to False if you want to access the redo logs as a binary file.
When you set `UseLogminerReader` to False, also set `UseBfile` to
True. For more information on this setting and using Oracle ASM, see [Using Oracle LogMiner or AWS DMS Binary Reader for CDC](../../../dms/latest/userguide/chap-source-oracle.md#CHAP_Source.Oracle.CDC) in
the _AWS DMS User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsePathPrefix`

Set this string attribute to the required value in order to use the Binary Reader to
capture change data for an Amazon RDS for Oracle as the source. This value specifies the
path prefix used to replace the default Oracle root to access the redo logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NeptuneSettings

PostgreSqlSettings

All content copied from https://docs.aws.amazon.com/.
