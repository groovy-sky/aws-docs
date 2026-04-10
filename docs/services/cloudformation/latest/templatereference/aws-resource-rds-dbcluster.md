This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBCluster

The `AWS::RDS::DBCluster` resource creates an Amazon Aurora DB cluster or Multi-AZ DB cluster.

For more information about creating an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](../../../amazonrds/latest/aurorauserguide/aurora-createinstance.md) in the _Amazon Aurora User Guide_.

For more information about creating a Multi-AZ DB cluster, see
[Creating a Multi-AZ DB cluster](../../../amazonrds/latest/userguide/create-multi-az-db-cluster.md)
in the _Amazon RDS User Guide_.

###### Note

You can only create this resource in AWS Regions where Amazon Aurora or Multi-AZ DB clusters are
supported.

**Updating DB clusters**

When properties labeled " _Update requires:_ [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)" are updated, AWS CloudFormation first creates a replacement DB
cluster, then changes references from other dependent resources to point to the
replacement DB cluster, and finally deletes the old DB cluster.

###### Important

We highly recommend that you take a snapshot of the database before updating the
stack. If you don't, you lose the data when AWS CloudFormation replaces your DB
cluster. To preserve your data, perform the following procedure:

1. Deactivate any applications that are using the DB cluster so that there's
    no activity on the DB instance.

2. Create a snapshot of the DB cluster. For more information, see [Creating a DB cluster snapshot](../../../amazonrds/latest/aurorauserguide/user-createsnapshotcluster.md).

3. If you want to restore your DB cluster using a DB cluster snapshot, modify
    the updated template with your DB cluster changes and add the
    `SnapshotIdentifier` property with the ID of the DB cluster
    snapshot that you want to use.

After you restore a DB cluster with a `SnapshotIdentifier`
    property, you must specify the same `SnapshotIdentifier` property
    for any future updates to the DB cluster. When you specify this property for
    an update, the DB cluster is not restored from the DB cluster snapshot
    again, and the data in the database is not changed. However, if you don't
    specify the `SnapshotIdentifier` property, an empty DB cluster is
    created, and the original DB cluster is deleted. If you specify a property
    that is different from the previous snapshot restore property, a new DB
    cluster is restored from the specified `SnapshotIdentifier`
    property, and the original DB cluster is deleted.

4. Update the stack.

Currently, when you are updating the stack for an Aurora Serverless DB cluster, you can't include changes to
any other properties when you specify one of the following properties: `PreferredBackupWindow`,
`PreferredMaintenanceWindow`, and `Port`. This limitation doesn't apply to provisioned
DB clusters.

For more information about updating other properties of this resource, see `
                    ModifyDBCluster
                `. For more information about updating stacks, see
[AWS \
CloudFormation Stacks Updates](../userguide/using-cfn-updating-stacks.md).

**Deleting DB clusters**

The default `DeletionPolicy` for `AWS::RDS::DBCluster` resources
is `Snapshot`. For more information about how AWS CloudFormation deletes
resources, see [DeletionPolicy Attribute](../userguide/aws-attribute-deletionpolicy.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBCluster",
  "Properties" : {
      "AllocatedStorage" : Integer,
      "AssociatedRoles" : [ DBClusterRole, ... ],
      "AutoMinorVersionUpgrade" : Boolean,
      "AvailabilityZones" : [ String, ... ],
      "BacktrackWindow" : Integer,
      "BackupRetentionPeriod" : Integer,
      "ClusterScalabilityType" : String,
      "CopyTagsToSnapshot" : Boolean,
      "DatabaseInsightsMode" : String,
      "DatabaseName" : String,
      "DBClusterIdentifier" : String,
      "DBClusterInstanceClass" : String,
      "DBClusterParameterGroupName" : String,
      "DBInstanceParameterGroupName" : String,
      "DBSubnetGroupName" : String,
      "DBSystemId" : String,
      "DeleteAutomatedBackups" : Boolean,
      "DeletionProtection" : Boolean,
      "Domain" : String,
      "DomainIAMRoleName" : String,
      "EnableCloudwatchLogsExports" : [ String, ... ],
      "EnableGlobalWriteForwarding" : Boolean,
      "EnableHttpEndpoint" : Boolean,
      "EnableIAMDatabaseAuthentication" : Boolean,
      "EnableLocalWriteForwarding" : Boolean,
      "Engine" : String,
      "EngineLifecycleSupport" : String,
      "EngineMode" : String,
      "EngineVersion" : String,
      "GlobalClusterIdentifier" : String,
      "Iops" : Integer,
      "KmsKeyId" : String,
      "ManageMasterUserPassword" : Boolean,
      "MasterUserAuthenticationType" : String,
      "MasterUsername" : String,
      "MasterUserPassword" : String,
      "MasterUserSecret" : MasterUserSecret,
      "MonitoringInterval" : Integer,
      "MonitoringRoleArn" : String,
      "NetworkType" : String,
      "PerformanceInsightsEnabled" : Boolean,
      "PerformanceInsightsKmsKeyId" : String,
      "PerformanceInsightsRetentionPeriod" : Integer,
      "Port" : Integer,
      "PreferredBackupWindow" : String,
      "PreferredMaintenanceWindow" : String,
      "PubliclyAccessible" : Boolean,
      "ReplicationSourceIdentifier" : String,
      "RestoreToTime" : String,
      "RestoreType" : String,
      "ScalingConfiguration" : ScalingConfiguration,
      "ServerlessV2ScalingConfiguration" : ServerlessV2ScalingConfiguration,
      "SnapshotIdentifier" : String,
      "SourceDBClusterIdentifier" : String,
      "SourceDbClusterResourceId" : String,
      "SourceRegion" : String,
      "StorageEncrypted" : Boolean,
      "StorageType" : String,
      "Tags" : [ Tag, ... ],
      "UseLatestRestorableTime" : Boolean,
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBCluster
Properties:
  AllocatedStorage: Integer
  AssociatedRoles:
    - DBClusterRole
  AutoMinorVersionUpgrade: Boolean
  AvailabilityZones:
    - String
  BacktrackWindow: Integer
  BackupRetentionPeriod: Integer
  ClusterScalabilityType: String
  CopyTagsToSnapshot: Boolean
  DatabaseInsightsMode: String
  DatabaseName: String
  DBClusterIdentifier: String
  DBClusterInstanceClass: String
  DBClusterParameterGroupName: String
  DBInstanceParameterGroupName: String
  DBSubnetGroupName: String
  DBSystemId: String
  DeleteAutomatedBackups: Boolean
  DeletionProtection: Boolean
  Domain: String
  DomainIAMRoleName: String
  EnableCloudwatchLogsExports:
    - String
  EnableGlobalWriteForwarding: Boolean
  EnableHttpEndpoint: Boolean
  EnableIAMDatabaseAuthentication: Boolean
  EnableLocalWriteForwarding: Boolean
  Engine: String
  EngineLifecycleSupport: String
  EngineMode: String
  EngineVersion: String
  GlobalClusterIdentifier: String
  Iops: Integer
  KmsKeyId: String
  ManageMasterUserPassword: Boolean
  MasterUserAuthenticationType: String
  MasterUsername: String
  MasterUserPassword: String
  MasterUserSecret:
    MasterUserSecret
  MonitoringInterval: Integer
  MonitoringRoleArn: String
  NetworkType: String
  PerformanceInsightsEnabled: Boolean
  PerformanceInsightsKmsKeyId: String
  PerformanceInsightsRetentionPeriod: Integer
  Port: Integer
  PreferredBackupWindow: String
  PreferredMaintenanceWindow: String
  PubliclyAccessible: Boolean
  ReplicationSourceIdentifier: String
  RestoreToTime: String
  RestoreType: String
  ScalingConfiguration:
    ScalingConfiguration
  ServerlessV2ScalingConfiguration:
    ServerlessV2ScalingConfiguration
  SnapshotIdentifier: String
  SourceDBClusterIdentifier: String
  SourceDbClusterResourceId: String
  SourceRegion: String
  StorageEncrypted: Boolean
  StorageType: String
  Tags:
    - Tag
  UseLatestRestorableTime: Boolean
  VpcSecurityGroupIds:
    - String

```

## Properties

`AllocatedStorage`

The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.

Valid for Cluster Type: Multi-AZ DB clusters only

This setting is required to create a Multi-AZ DB cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociatedRoles`

Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services
on your behalf.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Array of [DBClusterRole](aws-properties-rds-dbcluster-dbclusterrole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoMinorVersionUpgrade`

Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.
By default, minor engine upgrades are applied automatically.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB cluster.

For more information about automatic minor version upgrades, see [Automatically upgrading the minor engine version](../../../amazonrds/latest/userguide/user-upgradedbinstance-upgrading.md#USER_UpgradeDBInstance.Upgrading.AutoMinorVersionUpgrades).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZones`

A list of Availability Zones (AZs) where instances in the DB cluster can be created. For information on
AWS Regions and Availability Zones, see
[Choosing the Regions and \
Availability Zones](../../../amazonrds/latest/aurorauserguide/concepts-regionsandavailabilityzones.md) in the _Amazon Aurora User Guide_.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BacktrackWindow`

The target backtrack window, in seconds. To disable backtracking, set this value to
`0`.

Valid for Cluster Type: Aurora MySQL DB clusters only

Default: `0`

Constraints:

- If specified, this value must be set to a number from 0 to 259,200 (72 hours).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BackupRetentionPeriod`

The number of days for which automated backups are retained.

Default: 1

Constraints:

- Must be a value from 1 to 35

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ClusterScalabilityType`

Specifies the scalability mode of the Aurora DB cluster. When set to `limitless`, the cluster operates as an Aurora Limitless Database,
allowing you to create a DB shard group for horizontal scaling (sharding) capabilities. When set to `standard` (the default), the cluster uses normal DB instance creation.

**Important:** Automated backup retention isn't supported with Aurora Limitless Database clusters. If you set this property to `limitless`, you cannot set `DeleteAutomatedBackups` to `false`. To create a backup, use manual snapshots instead.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CopyTagsToSnapshot`

A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
The default is not to copy them.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseInsightsMode`

The mode of Database Insights to enable for the DB cluster.

If you set this value to `advanced`, you must also set the `PerformanceInsightsEnabled`
parameter to `true` and the `PerformanceInsightsRetentionPeriod` parameter to 465.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Allowed values_: `standard | advanced`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of your database. If you don't provide a name, then Amazon RDS won't create a
database in this DB cluster. For naming constraints, see [Naming\
Constraints](../../../amazonrds/latest/aurorauserguide/chap-limits.md#RDS_Limits.Constraints) in the _Amazon Aurora User Guide_.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBClusterIdentifier`

The DB cluster identifier. This parameter is stored as a lowercase string.

Constraints:

- Must contain from 1 to 63 letters, numbers, or hyphens.

- First character must be a letter.

- Can't end with a hyphen or contain two consecutive hyphens.

Example: `my-cluster1`

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBClusterInstanceClass`

The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example `db.m6gd.xlarge`.
Not all DB instance classes are available in all AWS Regions, or for all database engines.

For the full list of DB instance classes and availability for your engine, see [DB instance class](../../../amazonrds/latest/userguide/concepts-dbinstanceclass.md) in the _Amazon RDS User Guide_.

This setting is required to create a Multi-AZ DB cluster.

Valid for Cluster Type: Multi-AZ DB clusters only

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterParameterGroupName`

The name of the DB cluster parameter group to associate with this DB cluster.

###### Important

If you apply a parameter group to an existing DB cluster, then its DB instances
might need to reboot. This can result in an outage while the DB instances are
rebooting.

If you apply a change to parameter group associated with a stopped DB cluster,
then the update stack waits until the DB cluster is started.

To list all of the available DB cluster parameter group names, use the following
command:

`aws rds describe-db-cluster-parameter-groups --query
            "DBClusterParameterGroups[].DBClusterParameterGroupName" --output text`

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBInstanceParameterGroupName`

The name of the DB parameter group to apply to all instances of the DB cluster.

###### Note

When you apply a parameter group using the `DBInstanceParameterGroupName` parameter, the DB
cluster isn't rebooted automatically. Also, parameter changes are applied immediately rather than
during the next maintenance window.

Valid for Cluster Type: Aurora DB clusters only

Default: The existing name setting

Constraints:

- The DB parameter group must be in the same DB parameter group family as this DB cluster.

- The `DBInstanceParameterGroupName` parameter is valid in combination with the
`AllowMajorVersionUpgrade` parameter for a major version upgrade only.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBSubnetGroupName`

A DB subnet group that you want to associate with this DB cluster.

If you are restoring a DB cluster to a point in time with `RestoreType` set to `copy-on-write`, and don't
specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBSystemId`

Reserved for future use.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeleteAutomatedBackups`

Specifies whether to remove automated backups immediately after the DB
cluster is deleted. This parameter isn't case-sensitive. The default is to remove
automated backups immediately after the DB cluster is deleted, unless the AWS Backup
policy specifies a point-in-time restore rule.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtection`

A value that indicates whether the DB cluster has deletion protection enabled.
The database can't be deleted when deletion protection is enabled. By default,
deletion protection is disabled.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

Indicates the directory ID of the Active Directory to create the DB cluster.

For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.

For more information, see [Kerberos authentication](../../../amazonrds/latest/aurorauserguide/kerberos-authentication.md)
in the _Amazon Aurora User Guide_.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIAMRoleName`

Specifies the name of the IAM role to use when making API calls to the Directory Service.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableCloudwatchLogsExports`

The list of log types that need to be enabled for exporting to CloudWatch Logs. The values
in the list depend on the DB engine being used. For more information, see
[Publishing Database Logs to Amazon CloudWatch Logs](../../../amazonrds/latest/aurorauserguide/user-logaccess.md#USER_LogAccess.Procedural.UploadtoCloudWatch) in the _Amazon Aurora User Guide_.

**Aurora MySQL**

Valid values: `audit`, `error`, `general`, `slowquery`

**Aurora PostgreSQL**

Valid values: `postgresql`

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableGlobalWriteForwarding`

Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster
(Aurora global database). By default, write operations are not allowed on Aurora DB clusters that
are secondary clusters in an Aurora global database.

You can set this value only on Aurora DB clusters that are members of an Aurora global database. With this parameter
enabled, a secondary cluster can forward writes to the current primary cluster, and the resulting changes are replicated back to
this cluster. For the primary DB cluster of an Aurora global database, this value is used immediately if the
primary is demoted by a global cluster API operation, but it does nothing until then.

Valid for Cluster Type: Aurora DB clusters only

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableHttpEndpoint`

Specifies whether to enable the HTTP endpoint for the DB cluster. By default, the HTTP endpoint
isn't enabled.

When enabled, the HTTP endpoint provides a connectionless web service API (RDS Data API) for running
SQL queries on the DB cluster. You can also query your database
from inside the RDS console with the RDS query editor.

For more information, see [Using RDS Data API](../../../amazonrds/latest/aurorauserguide/data-api.md) in the
_Amazon Aurora User Guide_.

Valid for Cluster Type: Aurora DB clusters only

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableIAMDatabaseAuthentication`

A value that indicates whether to enable mapping of AWS Identity and Access
Management (IAM) accounts to database accounts. By default, mapping is disabled.

For more information, see
[IAM Database Authentication](../../../amazonrds/latest/aurorauserguide/usingwithrds-iamdbauth.md) in the _Amazon Aurora User Guide._

Valid for: Aurora DB clusters only

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableLocalWriteForwarding`

Specifies whether read replicas can forward write operations to the writer DB instance in the DB cluster. By
default, write operations aren't allowed on reader DB instances.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The name of the database engine to be used for this DB cluster.

Valid Values:

- `aurora-mysql`

- `aurora-postgresql`

- `mysql`

- `postgres`

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EngineLifecycleSupport`

The life cycle type for this DB cluster.

###### Note

By default, this value is set to `open-source-rds-extended-support`, which enrolls your DB cluster into Amazon RDS Extended Support.
At the end of standard support, you can avoid charges for Extended Support by setting the value to `open-source-rds-extended-support-disabled`. In this case,
creating the DB cluster will fail if the DB major version is past its end of standard support date.

You can use this setting to enroll your DB cluster into Amazon RDS Extended Support. With RDS Extended Support,
you can run the selected major engine version on your DB cluster past the end of standard support for that engine version. For more information, see the following sections:

- Amazon Aurora - [Amazon RDS Extended Support with Amazon Aurora](../../../amazonrds/latest/aurorauserguide/extended-support.md) in the _Amazon Aurora User Guide_

- Amazon RDS - [Amazon RDS Extended Support with Amazon RDS](../../../amazonrds/latest/userguide/extended-support.md) in the _Amazon RDS User Guide_

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

Valid Values: `open-source-rds-extended-support | open-source-rds-extended-support-disabled`

Default: `open-source-rds-extended-support`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EngineMode`

The DB engine mode of the DB cluster, either `provisioned` or `serverless`.

The `serverless` engine mode only applies for Aurora Serverless v1 DB clusters. Aurora Serverless v2 DB clusters use the
`provisioned` engine mode.

For information about limitations and requirements for Serverless DB clusters, see the
following sections in the _Amazon Aurora User Guide_:

- [Limitations of Aurora\
Serverless v1](../../../amazonrds/latest/aurorauserguide/aurora-serverless.md#aurora-serverless.limitations)

- [Requirements\
for Aurora Serverless v2](../../../amazonrds/latest/aurorauserguide/aurora-serverless-v2-requirements.md)

Valid for Cluster Type: Aurora DB clusters only

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The version number of the database engine to use.

###### Important

Don't use this property if your DB cluster is a member of a global database cluster. Instead, specify the `EngineVersion` property
on the [AWS::RDS::GlobalCluster](../userguide/aws-resource-rds-globalcluster.md) resource.
Major version upgrades aren't supported for individual members of a global cluster. Use `ModifyGlobalCluster` to upgrade all members of the global cluster.

To list all of the available engine versions for Aurora MySQL version 2 (5.7-compatible)
and version 3 (8.0-compatible), use the following command:

`aws rds describe-db-engine-versions --engine aurora-mysql --query
            "DBEngineVersions[].EngineVersion"`

You can supply either `5.7` or `8.0` to use the default engine version for Aurora MySQL version 2 or
version 3, respectively.

To list all of the available engine versions for Aurora PostgreSQL, use
the following command:

`aws rds describe-db-engine-versions --engine aurora-postgresql --query
            "DBEngineVersions[].EngineVersion"`

To list all of the available engine versions for RDS for MySQL, use the following command:

`aws rds describe-db-engine-versions --engine mysql --query "DBEngineVersions[].EngineVersion"`

To list all of the available engine versions for RDS for PostgreSQL, use the following command:

`aws rds describe-db-engine-versions --engine postgres --query "DBEngineVersions[].EngineVersion"`

**Aurora MySQL**

For information, see [Database engine updates for Amazon Aurora MySQL](../../../amazonrds/latest/aurorauserguide/auroramysql-updates.md) in the
_Amazon Aurora User Guide_.

**Aurora PostgreSQL**

For information, see [Amazon Aurora PostgreSQL releases and engine versions](../../../amazonrds/latest/aurorauserguide/aurorapostgresql-updates-20180305.md) in the
_Amazon Aurora User Guide_.

**MySQL**

For information, see [Amazon RDS for MySQL](../../../amazonrds/latest/userguide/chap-mysql.md#MySQL.Concepts.VersionMgmt) in the _Amazon RDS User Guide_.

**PostgreSQL**

For information, see [Amazon RDS for PostgreSQL](../../../amazonrds/latest/userguide/chap-postgresql.md#PostgreSQL.Concepts) in the _Amazon RDS User Guide_.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`GlobalClusterIdentifier`

If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database
cluster, specify the global cluster ID of the global database cluster. To define the primary database cluster of the global cluster,
use the [AWS::RDS::GlobalCluster](../userguide/aws-resource-rds-globalcluster.md)
resource.

If you aren't configuring a global database cluster, don't specify this property.

###### Note

To remove the DB cluster from a global database cluster, specify an empty value for the `GlobalClusterIdentifier` property.

For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](../../../amazonrds/latest/aurorauserguide/aurora-global-database.md) in the _Amazon Aurora User Guide_.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: String

_Pattern_: `^$|^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Iops`

The amount of Provisioned IOPS (input/output operations per second) to be initially allocated
for each DB instance in the Multi-AZ DB cluster.

For information about valid IOPS values, see [Provisioned IOPS storage](../../../amazonrds/latest/userguide/chap-storage.md#USER_PIOPS) in the _Amazon RDS_
_User Guide_.

This setting is required to create a Multi-AZ DB cluster.

Valid for Cluster Type: Multi-AZ DB clusters only

Constraints:

- Must be a multiple between .5 and 50 of the storage amount for the DB cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The Amazon Resource Name (ARN) of the AWS KMS key that is
used to encrypt the database instances in the DB cluster, such as
`arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef`.
If you enable the `StorageEncrypted` property but don't specify this
property, the default KMS key is used. If you specify this property, you must set the
`StorageEncrypted` property to `true`.

If you specify the `SnapshotIdentifier` property, the `StorageEncrypted` property
value is inherited from the snapshot, and if the DB cluster is encrypted, the specified `KmsKeyId`
property is used.

If you create a read replica of an encrypted DB cluster in another AWS Region, make sure to set `KmsKeyId` to a KMS key identifier that is valid in the destination AWS Region.
This KMS key is used to encrypt the read replica in that AWS Region.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManageMasterUserPassword`

Specifies whether to manage the master user password with AWS Secrets Manager.

For more information, see [Password management with AWS Secrets Manager](../../../amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide_ and [Password management with AWS Secrets Manager](../../../amazonrds/latest/aurorauserguide/rds-secrets-manager.md)
in the _Amazon Aurora User Guide._

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

Constraints:

- Can't manage the master user password with AWS Secrets Manager if `MasterUserPassword`
is specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserAuthenticationType`

Specifies the authentication type for the master user. With IAM master user authentication, you can configure the master DB user with IAM database authentication when you create a DB cluster.

You can specify one of the following values:

- `password` \- Use standard database authentication with a password.

- `iam-db-auth` \- Use IAM database authentication for the master user.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

This option is only valid for RDS for MySQL, RDS for MariaDB, RDS for PostgreSQL, Aurora MySQL, and Aurora PostgreSQL engines.

_Required_: No

_Type_: String

_Allowed values_: `password | iam-db-auth`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUsername`

The name of the master user for the DB cluster.

###### Note

If you specify the `SourceDBClusterIdentifier`, `SnapshotIdentifier`, or `GlobalClusterIdentifier`
property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB
cluster for the global database cluster, respectively.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z]{1}[a-zA-Z0-9_]*$`

_Minimum_: `1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MasterUserPassword`

The master password for the DB instance.

###### Note

If you specify the `SourceDBClusterIdentifier`, `SnapshotIdentifier`, or `GlobalClusterIdentifier`
property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB
cluster for the global database cluster, respectively.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserSecret`

The secret managed by RDS in AWS Secrets Manager for the master user password.

###### Note

When you restore a DB cluster from a snapshot, Amazon RDS generates a new secret instead of
reusing the secret specified in the `SecretArn` property. This ensures that
the restored DB cluster is securely managed with a dedicated secret. To maintain
consistent integration with your application, you might need to update resource
configurations to reference the newly created secret.

For more information, see [Password management with AWS Secrets Manager](../../../amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide_ and [Password management with AWS Secrets Manager](../../../amazonrds/latest/aurorauserguide/rds-secrets-manager.md)
in the _Amazon Aurora User Guide._

_Required_: No

_Type_: [MasterUserSecret](aws-properties-rds-dbcluster-masterusersecret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringInterval`

The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off
collecting Enhanced Monitoring metrics, specify `0`.

If `MonitoringRoleArn` is specified, also set `MonitoringInterval`
to a value other than `0`.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

Valid Values: `0 | 1 | 5 | 10 | 15 | 30 | 60`

Default: `0`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringRoleArn`

The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.
An example is `arn:aws:iam:123456789012:role/emaccess`. For information on creating a monitoring role,
see [Setting \
up and enabling Enhanced Monitoring](../../../amazonrds/latest/userguide/user-monitoring-os.md#USER_Monitoring.OS.Enabling) in the _Amazon RDS User Guide_.

If `MonitoringInterval` is set to a value other than `0`, supply a `MonitoringRoleArn` value.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

The network type of the DB cluster.

Valid values:

- `IPV4`

- `DUAL`

The network type is determined by the `DBSubnetGroup` specified for the DB cluster.
A `DBSubnetGroup` can support only the IPv4 protocol or the IPv4 and IPv6
protocols ( `DUAL`).

For more information, see [Working with a DB instance in a VPC](../../../amazonrds/latest/aurorauserguide/user-vpc-workingwithrdsinstanceinavpc.md) in the
_Amazon Aurora User Guide._

Valid for: Aurora DB clusters only

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceInsightsEnabled`

Specifies whether to turn on Performance Insights for the DB cluster.

For more information, see [Using Amazon Performance Insights](../../../amazonrds/latest/userguide/user-perfinsights.md) in the _Amazon RDS User Guide_.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceInsightsKmsKeyId`

The AWS KMS key identifier for encryption of Performance Insights data.

The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.

If you don't specify a value for `PerformanceInsightsKMSKeyId`, then Amazon RDS
uses your default KMS key. There is a default KMS key for your AWS account.
Your AWS account has a different default KMS key for each AWS Region.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceInsightsRetentionPeriod`

The number of days to retain Performance Insights data. When creating a DB cluster without enabling Performance Insights, you can't specify the parameter `PerformanceInsightsRetentionPeriod`.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

Valid Values:

- `7`

- _month_ \\* 31, where _month_ is a number of months from 1-23.
Examples: `93` (3 months \* 31), `341` (11 months \* 31), `589` (19 months \* 31)

- `731`

Default: `7` days

If you specify a retention period that isn't valid, such as `94`, Amazon RDS issues an error.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number on which the DB instances in the DB cluster accept connections.

Default:

- RDS for MySQL and Aurora MySQL - `3306`

- RDS for PostgreSQL and Aurora PostgreSQL - `5432`

###### Important

The `No interruption` on update behavior only applies to DB clusters. If you are updating a
DB instance, see [Port](../userguide/aws-properties-rds-database-instance.md#cfn-rds-dbinstance-port)
for the AWS::RDS::DBInstance resource.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredBackupWindow`

The daily time range during which automated backups are created. For more information, see [Backup Window](../../../amazonrds/latest/aurorauserguide/aurora-managing-backups.md#Aurora.Managing.Backups.BackupWindow) in the _Amazon Aurora User Guide._

Constraints:

- Must be in the format `hh24:mi-hh24:mi`.

- Must be in Universal Coordinated Time (UTC).

- Must not conflict with the preferred maintenance window.

- Must be at least 30 minutes.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).

Format: `ddd:hh24:mi-ddd:hh24:mi`

The default is a 30-minute window selected at random from an 8-hour block of time for
each AWS Region, occurring on a random day of the week. To see the time
blocks available, see [Maintaining an Amazon Aurora DB cluster](../../../amazonrds/latest/aurorauserguide/user-upgradedbinstance-maintenance.md#AdjustingTheMaintenanceWindow.Aurora) in the _Amazon Aurora User_
_Guide._

Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.

Constraints: Minimum 30-minute window.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Specifies whether the DB cluster is publicly accessible.

Valid for Cluster Type: Multi-AZ DB clusters only

When the DB cluster is publicly accessible and you connect from outside of the DB cluster's virtual private cloud (VPC), its domain name system (DNS) endpoint resolves to the public IP address. When you connect from within the same VPC as the DB cluster, the endpoint resolves to the private IP address. Access to the DB cluster is controlled by its security group settings.

When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.

The default behavior when `PubliclyAccessible` is not specified depends on whether a `DBSubnetGroup` is specified.

If `DBSubnetGroup` isn't specified, `PubliclyAccessible` defaults to `true`.

If `DBSubnetGroup` is specified, `PubliclyAccessible` defaults to `false` unless the value of `DBSubnetGroup` is `default`, in which case `PubliclyAccessible` defaults to `true`.

If `PubliclyAccessible` is true and the VPC that the `DBSubnetGroup` is in doesn't have an internet gateway attached to it, Amazon RDS returns an error.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationSourceIdentifier`

The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB
cluster is created as a read replica.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreToTime`

The date and time to restore the DB cluster to.

Valid Values: Value must be a time in Universal Coordinated Time (UTC) format

Constraints:

- Must be before the latest restorable time for the DB instance

- Must be specified if `UseLatestRestorableTime` parameter isn't provided

- Can't be specified if the `UseLatestRestorableTime` parameter is enabled

- Can't be specified if the `RestoreType` parameter is `copy-on-write`

This property must be used with `SourceDBClusterIdentifier` property. The resulting cluster will have the identifier
that matches the value of the `DBclusterIdentifier` property.

Example: `2015-03-07T23:45:00Z`

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestoreType`

The type of restore to be performed. You can specify one of the following values:

- `full-copy` \- The new DB cluster is restored as a full copy of the
source DB cluster.

- `copy-on-write` \- The new DB cluster is restored as a clone of the
source DB cluster.

If you don't specify a `RestoreType` value, then the new DB cluster is
restored as a full copy of the source DB cluster.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScalingConfiguration`

The scaling
configuration of an Aurora Serverless v1 DB cluster.

This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, Use the `ServerlessV2ScalingConfiguration` property.

Valid for: Aurora Serverless v1 DB clusters only

_Required_: No

_Type_: [ScalingConfiguration](aws-properties-rds-dbcluster-scalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerlessV2ScalingConfiguration`

The scaling configuration of an Aurora Serverless V2 DB cluster.

This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, Use the `ScalingConfiguration` property.

Valid for: Aurora Serverless v2 DB clusters only

_Required_: No

_Type_: [ServerlessV2ScalingConfiguration](aws-properties-rds-dbcluster-serverlessv2scalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotIdentifier`

The identifier for the DB snapshot or DB cluster snapshot to restore
from.

You can use either the name or the Amazon Resource Name (ARN) to specify a DB
cluster snapshot. However, you can use only the ARN to specify a DB snapshot.

After you restore a DB cluster with a `SnapshotIdentifier` property, you
must specify the same `SnapshotIdentifier` property for any future updates to
the DB cluster. When you specify this property for an update, the DB cluster is not
restored from the snapshot again, and the data in the database is not changed. However,
if you don't specify the `SnapshotIdentifier` property, an empty DB cluster
is created, and the original DB cluster is deleted. If you specify a property that is
different from the previous snapshot restore property, a new DB cluster is restored from
the specified `SnapshotIdentifier` property, and the original DB cluster is
deleted.

If you specify the `SnapshotIdentifier` property to restore a DB cluster (as opposed to specifying it for DB cluster updates),
then don't specify the following properties:

- `GlobalClusterIdentifier`

- `MasterUsername`

- `MasterUserPassword`

- `ReplicationSourceIdentifier`

- `RestoreType`

- `SourceDBClusterIdentifier`

- `SourceRegion`

- `StorageEncrypted` (for an encrypted snapshot)

- `UseLatestRestorableTime`

Constraints:

- Must match the identifier of an existing Snapshot.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceDBClusterIdentifier`

When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.

Constraints:

- Must match the identifier of an existing DBCluster.

- Cannot be specified if `SourceDbClusterResourceId` is specified. You must specify either `SourceDBClusterIdentifier` or `SourceDbClusterResourceId`, but not both.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceDbClusterResourceId`

The resource ID of the source DB cluster from which to restore.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceRegion`

The AWS Region which contains the source DB cluster when replicating a DB cluster. For
example, `us-east-1`.

Valid for: Aurora DB clusters only

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageEncrypted`

Indicates whether the DB cluster is encrypted.

If you specify the `KmsKeyId` property, then you must enable encryption.

If you specify the `SourceDBClusterIdentifier` property, don't specify this property. The
value is inherited from the source DB cluster, and if the DB cluster is encrypted, the specified
`KmsKeyId` property is used.

If you specify the `SnapshotIdentifier` and the specified snapshot is encrypted,
don't specify this property. The value is inherited from the snapshot, and the specified `KmsKeyId`
property is used.

If you specify the `SnapshotIdentifier` and the specified snapshot isn't encrypted, you can use this property
to specify that the restored DB cluster is encrypted. Specify the `KmsKeyId` property for the KMS key
to use for encryption. If you don't want the restored DB cluster to be encrypted, then don't set this property
or set it to `false`.

###### Note

If you specify both the `StorageEncrypted` and
`SnapshotIdentifier` properties without specifying the
`KmsKeyId` property, then the restored DB cluster inherits the encryption
settings from the DB snapshot that provide.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageType`

The storage type to associate with the DB cluster.

For information on storage types for Aurora DB clusters, see [Storage configurations for Amazon Aurora DB clusters](../../../amazonrds/latest/aurorauserguide/aurora-overview-storagereliability.md#aurora-storage-type). For information on storage types for Multi-AZ DB
clusters, see [Settings for creating Multi-AZ DB clusters](../../../amazonrds/latest/userguide/create-multi-az-db-cluster.md#create-multi-az-db-cluster-settings).

This setting is required to create a Multi-AZ DB cluster.

When specified for a Multi-AZ DB cluster, a value for the `Iops` parameter is required.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

Valid Values:

- Aurora DB clusters - `aurora | aurora-iopt1`

- Multi-AZ DB clusters - `io1 | io2 | gp3`

Default:

- Aurora DB clusters - `aurora`

- Multi-AZ DB clusters - `io1`

###### Note

When you create an Aurora DB cluster with the storage type set to `aurora-iopt1`, the storage type is returned
in the response. The storage type isn't returned when you set it to `aurora`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to assign to the DB cluster.

Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Array of [Tag](aws-properties-rds-dbcluster-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLatestRestorableTime`

A value that indicates whether to restore the DB cluster to the latest restorable
backup time. By default, the DB cluster is not restored to the latest restorable backup
time.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSecurityGroupIds`

A list of EC2 VPC security groups to associate with this DB cluster.

If you plan to update the resource, don't specify VPC security groups in a shared VPC.

Valid for: Aurora DB clusters and Multi-AZ DB clusters

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name ( `DBClusterIdentifier`) of the DB cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DBClusterArn`

The Amazon Resource Name (ARN) for the DB cluster.

`DBClusterResourceId`

The AWS Region-unique, immutable identifier for the DB cluster. This identifier is found in AWS CloudTrail log entries whenever
the KMS key for the DB cluster is accessed.

`Endpoint.Address`

The connection endpoint for the DB cluster. For example:
`mystack-mydbcluster-123456789012.us-east-2.rds.amazonaws.com`

`Endpoint.Port`

The port number that will accept connections on this DB cluster. For example:
`3306`

`MasterUserSecret.SecretArn`

The Amazon Resource Name (ARN) of the secret.

`ReadEndpoint.Address`

The reader endpoint for the DB cluster. For example:
`mystack-mydbcluster-ro-123456789012.us-east-2.rds.amazonaws.com`

`StorageEncryptionType`

The type of encryption used to protect data at rest in the DB cluster. Possible values:

- `none` \- The DB cluster is not encrypted.

- `sse-rds` \- The DB cluster is encrypted using an AWS owned KMS key.

- `sse-kms` \- The DB cluster is encrypted using a customer managed KMS key or AWS managed KMS key.

`StorageThroughput`

The storage throughput for the DB cluster. The throughput is automatically set based on the IOPS that you provision, and is not configurable.

This setting is only for non-Aurora Multi-AZ DB clusters.

## Remarks

_DB cluster creation modes_

This section outlines the modes for creating RDS DB cluters—empty
database, read replica, restore from snapshot, and point-in-time recovery. In the
console, API, or CLI, each mode corresponds to a distinct API or command. In
CloudFormation, however, the mode is determined by the following parameter
combinations.

- **Empty database** \- Creates a new, empty DB cluster.

To create an empty database, _don't_ specify any of the
following properties:

- `SnapshotIdentifier`

- `RestoreToTime`

- `UseLatestRestorableTime`

- `SourceDBClusterIdentifier`

- `ReplicationSourceIdentifier`

- **Read replica** \- Creates a replica of an existing
DB cluster to replicate data asynchronously.

To create a read replica, specify the `ReplicationSourceIdentifier`
property _without_ the following properties:

- `RestoreToTime`

- `UseLatestRestorableTime`

- `SnapshotIdentifier`

- `SourceDBClusterIdentifier`

- **Restore from snapshot** \- Restores the DB cluster
from the specified snapshot. Cannot be used with a point-in-time recovery.

To restore a DB cluster from a snapshot, specify the
`SnapshotIdentifier` property _without_ the
following properties:

- `RestoreToTime`

- `UseLatestRestorableTime`

- `SourceDBClusterIdentifier`

- `ReplicationSourceIdentifier`

- **Point-in-time recovery** \- Recovers the DB cluster
to a specific point in time. Requires identifying the source and the recovery
time. Cannot be used with a snapshot.

To restore a DB cluster to a specified time, include the
`SourceDBClusterIdentifier` property or `SourceDbClusterResourceId` property with
_either_ the `RestoreToTime` property or the
`UseLatestRestorableTime` property.

## Examples

The following examples create DB clusters.

###### Note

For an example that adds a scaling policy for a DB cluster with Application
Auto Scaling, see [Declaring a scaling policy for an Aurora DB cluster](../userguide/quickref-autoscaling.md#w2ab1c19c22c15c21c11).

- [Creating an Amazon Aurora DB cluster with two DB instances](#aws-resource-rds-dbcluster--examples--Creating_an_Amazon_Aurora_DB_cluster_with_two_DB_instances)

- [Creating an Amazon Aurora DB cluster that exports logs to Amazon CloudWatch Logs](#aws-resource-rds-dbcluster--examples--Creating_an_Amazon_Aurora_DB_cluster_that_exports_logs_to_Amazon_CloudWatch_Logs)

- [Creating a Multi-AZ DB cluster](#aws-resource-rds-dbcluster--examples--Creating_a_Multi-AZ_DB_cluster)

- [Creating an Amazon Aurora Serverless v2 DB cluster](#aws-resource-rds-dbcluster--examples--Creating_an_Amazon_Aurora_Serverless_v2_DB_cluster)

- [Creating an Amazon Aurora Serverless v1 DB cluster](#aws-resource-rds-dbcluster--examples--Creating_an_Amazon_Aurora_Serverless_v1_DB_cluster)

- [Create a Secrets Manager secret for a master password](#aws-resource-rds-dbcluster--examples--Create_a_Secrets_Manager_secret_for_a_master_password)

- [Restoring an Amazon Aurora DB cluster](#aws-resource-rds-dbcluster--examples--Restoring_an_Amazon_Aurora_DB_cluster)

### Creating an Amazon Aurora DB cluster with two DB instances

The following example creates an Amazon Aurora DB cluster and adds two DB
instances to it. Because Amazon RDS automatically assigns a writer and reader DB
instances in the cluster, use the cluster endpoint to read and write data, not
the individual DB instance endpoints.

###### Note

The example uses the `time_zone` Aurora MySQL parameter. For
Aurora PostgreSQL, use the `timezone` parameter instead.

#### JSON

```json

{
    "Parameters": {
        "Username": {
            "NoEcho": "true",
            "Description": "Username for Aurora MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "bevelvoerder",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "Password": {
            "NoEcho": "true",
            "Description": "Password for Aurora MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Passw0rd",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "MasterUsername": {
                    "Ref": "Username"
                },
                "MasterUserPassword": {
                    "Ref": "Password"
                },
                "Engine": "aurora-mysql",
                "DBClusterParameterGroupName": {
                    "Ref": "RDSDBClusterParameterGroup"
                }
            }
        },
        "RDSDBInstance1": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBParameterGroupName": {
                    "Ref": "RDSDBParameterGroup"
                },
                "Engine": "aurora-mysql",
                "DBClusterIdentifier": {
                    "Ref": "RDSCluster"
                },
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r3.xlarge"
            }
        },
        "RDSDBInstance2": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBParameterGroupName": {
                    "Ref": "RDSDBParameterGroup"
                },
                "Engine": "aurora-mysql",
                "DBClusterIdentifier": {
                    "Ref": "RDSCluster"
                },
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r3.xlarge"
            }
        },
        "RDSDBClusterParameterGroup": {
            "Type": "AWS::RDS::DBClusterParameterGroup",
            "Properties": {
                "Description": "CloudFormation Sample Aurora Cluster Parameter Group",
                "Family": "aurora5.6",
                "Parameters": {
                    "time_zone": "US/Eastern"
                }
            }
        },
        "RDSDBParameterGroup": {
            "Type": "AWS::RDS::DBParameterGroup",
            "Properties": {
                "Description": "CloudFormation Sample Aurora Parameter Group",
                "Family": "aurora5.6",
                "Parameters": {
                    "sql_mode": "IGNORE_SPACE",
                    "max_allowed_packet": 1024,
                    "innodb_buffer_pool_size": "{DBInstanceClassMemory*3/4}"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  Username:
    NoEcho: 'true'
    Description: Username for Aurora MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: "bevelvoerder"
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  Password:
    NoEcho: 'true'
    Description: Password for Aurora MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: "Passw0rd"
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername:
        Ref: Username
      MasterUserPassword:
        Ref: Password
      Engine: aurora-mysql
      DBClusterParameterGroupName:
        Ref: RDSDBClusterParameterGroup
  RDSDBInstance1:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBParameterGroupName:
        Ref: RDSDBParameterGroup
      Engine: aurora-mysql
      DBClusterIdentifier:
        Ref: RDSCluster
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r3.xlarge
  RDSDBInstance2:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBParameterGroupName:
        Ref: RDSDBParameterGroup
      Engine: aurora-mysql
      DBClusterIdentifier:
        Ref: RDSCluster
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r3.xlarge
  RDSDBClusterParameterGroup:
    Type: 'AWS::RDS::DBClusterParameterGroup'
    Properties:
      Description: CloudFormation Sample Aurora Cluster Parameter Group
      Family: aurora5.6
      Parameters:
        time_zone: US/Eastern
  RDSDBParameterGroup:
    Type: 'AWS::RDS::DBParameterGroup'
    Properties:
      Description: CloudFormation Sample Aurora Parameter Group
      Family: aurora5.6
      Parameters:
        sql_mode: IGNORE_SPACE
        max_allowed_packet: 1024
        innodb_buffer_pool_size: '{DBInstanceClassMemory*3/4}'
```

### Creating an Amazon Aurora DB cluster that exports logs to Amazon CloudWatch Logs

The following example creates an Amazon Aurora PostgreSQL DB cluster that
exports logs to Amazon CloudWatch Logs. For more information about exporting
Aurora DB cluster logs to Amazon CloudWatch Logs, see [Publishing Database Logs to Amazon CloudWatch Logs](../../../amazonrds/latest/aurorauserguide/user-logaccess.md#USER_LogAccess.Procedural.UploadtoCloudWatch) in the
_Amazon Aurora User Guide_.

#### JSON

```json

{
  "AWSTemplateFormatVersion" : "2010-09-09",

  "Description" : "AWS CloudFormation Sample Template for sending Aurora DB cluster logs to CloudWatch Logs: Sample template showing how to create an Aurora PostgreSQL DB cluster that exports logs to CloudWatch Logs. **WARNING** This template enables log exports to CloudWatch Logs. You will be billed for the AWS resources used if you create a stack from this template.",

  "Parameters" : {
      "DBUsername" : {
        "NoEcho" : "true",
        "Description" : "Username for PostgreSQL database access",

        "Type" : "String",
        "MinLength" : "1",
        "MaxLength" : "16",
        "AllowedPattern" : "[a-zA-Z][a-zA-Z0-9]*",
        "ConstraintDescription" : "must begin with a letter and contain only alphanumeric characters."
      },
      "DBPassword" : {
        "NoEcho" : "true",
        "Description" : "Password for PostgreSQL database access",

        "Type" : "String",
        "MinLength" : "8",

        "MaxLength" : "41",
        "AllowedPattern" : "[a-zA-Z0-9]*",
        "ConstraintDescription" : "must contain only alphanumeric characters."
      }
  },

  "Resources" : {
      "RDSCluster" : {
          "Type": "AWS::RDS::DBCluster",
          "Properties" : {
              "MasterUsername" : {
                  "Ref" : "DBUsername"
              },
              "MasterUserPassword" : {
                  "Ref" : "DBPassword"
              },
              "DBClusterIdentifier" : "aurora-postgresql-cluster",
              "Engine" : "aurora-postgresql",
              "EngineVersion" : "10.7",
              "DBClusterParameterGroupName" : "default.aurora-postgresql10",
              "EnableCloudwatchLogsExports" : ["postgresql"]
          }
      },
  "RDSDBInstance1": {
        "Type" : "AWS::RDS::DBInstance",
        "Properties" : {
            "DBInstanceIdentifier" : "aurora-postgresql-instance1",
            "Engine" : "aurora-postgresql",
            "DBClusterIdentifier" : {
                "Ref" : "RDSCluster"
            },
            "PubliclyAccessible" : "true",
            "DBInstanceClass" : "db.r4.large"
        }
    },
  "RDSDBInstance2": {
        "Type" : "AWS::RDS::DBInstance",
        "Properties" : {
            "DBInstanceIdentifier" : "aurora-postgresql-instance2",
            "Engine" : "aurora-postgresql",
            "DBClusterIdentifier" : {
                "Ref" : "RDSCluster"
            },
            "PubliclyAccessible" : "true",
            "DBInstanceClass" : "db.r4.large"
        }
    },
  }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  AWS CloudFormation Sample Template for sending Aurora DB cluster logs to
  CloudWatch Logs: Sample template showing how to create an Aurora PostgreSQL DB
  cluster that exports logs to CloudWatch Logs. **WARNING** This template
  enables log exports to CloudWatch Logs. You will be billed for the AWS
  resources used if you create a stack from this template.
Parameters:
  DBUsername:
    NoEcho: 'true'
    Description: Username for PostgreSQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  DBPassword:
    NoEcho: 'true'
    Description: Password for PostgreSQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername: !Ref DBUsername
      MasterUserPassword: !Ref DBPassword
      DBClusterIdentifier: aurora-postgresql-cluster
      Engine: aurora-postgresql
      EngineVersion: '10.7'
      DBClusterParameterGroupName: default.aurora-postgresql10
      EnableCloudwatchLogsExports:
        - postgresql
  RDSDBInstance1:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBInstanceIdentifier: aurora-postgresql-instance1
      Engine: aurora-postgresql
      DBClusterIdentifier: !Ref RDSCluster
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r4.large
  RDSDBInstance2:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBInstanceIdentifier: aurora-postgresql-instance2
      Engine: aurora-postgresql
      DBClusterIdentifier: !Ref RDSCluster
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r4.large
```

### Creating a Multi-AZ DB cluster

The following example creates a Multi-AZ DB cluster. A Multi-AZ DB cluster
provides high availability by maintaining standby replicas in different Availability
Zones. This template creates a DB cluster and automatically provisions the required
reader and writer DB instances. For more information, see [Multi-AZ DB\
cluster deployments](../../../amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "CloudFormation template to create a MySQL Multi-AZ DB cluster with encryption, CloudWatch logs, deletion protection, and allocated storage.",
    "Parameters": {
      "MasterPassword": {
        "Type": "String",
        "NoEcho": true,
        "Description": "The master password for the DB cluster."
      }
    },
    "Resources": {
      "MyDBCluster": {
        "Type": "AWS::RDS::DBCluster",
        "Properties": {
          "Engine": "mysql",
          "EngineVersion": "8.0.32",
          "DatabaseName": "MyDatabase",
          "MasterUsername": "admin",
          "MasterUserPassword": { "Ref": "MasterPassword" },
          "StorageEncrypted": true,
          "EnableCloudwatchLogsExports": [
            "error",
            "general",
            "slowquery"
          ],
          "DeletionProtection": true,
          "AllocatedStorage": 20,
          "DBClusterInstanceClass": "db.m5d.xlarge",
          "StorageType": "gp3"
        }
      }
    }
  }
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: "CloudFormation template to create a MySQL Multi-AZ DB cluster with encryption, CloudWatch logs, deletion protection, and allocated storage."
Parameters:
  MasterPassword:
    Type: String
    NoEcho: true
    Description: "The master password for the DB cluster."
Resources:
  MyDBCluster:
    Type: AWS::RDS::DBCluster
    Properties:
      Engine: mysql
      EngineVersion: "8.0.32"
      DatabaseName: MyDatabase
      MasterUsername: admin
      MasterUserPassword: !Ref MasterPassword
      StorageEncrypted: true
      EnableCloudwatchLogsExports:
        - "error"
        - "general"
        - "slowquery"
      DeletionProtection: true
      AllocatedStorage: 20
      DBClusterInstanceClass: db.m5d.xlarge
      StorageType: gp3
```

### Creating an Amazon Aurora Serverless v2 DB cluster

The following example creates an Amazon Aurora Serverless v2 DB cluster.
Aurora Serverless v2 automates the processes of monitoring the workload and adjusting the capacity for your databases.
It adjusts the capacity in increments to provide the necessary database resources to meet an application's requirements.
For more information,
see [Using Amazon Aurora Serverless v2](../../../amazonrds/latest/aurorauserguide/aurora-serverless-v2.md) in the
_Amazon Aurora User Guide_.

You must provide a supported engine version when you create an Amazon Aurora Serverless v2 DB cluster.
For more information, see [Aurora Serverless v2](../../../amazonrds/latest/aurorauserguide/concepts-aurora-fea-regions-db-eng-feature-serverlessv2.md) in the _Amazon Aurora User Guide_.

###### Note

This example creates an Aurora MySQL Serverless v2 DB cluster by setting
`Engine` to `aurora-mysql`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "ServerlessV2 Cluster",
    "Parameters": {
        "MasterUsername": {
            "Type": "String"
        },
        "MasterUserPassword": {
            "Type": "String",
            "NoEcho": true
        },
        "DBClusterIdentifier": {
            "Type": "String"
        },
        "EngineVersion": {
            "Type": "String"
        },
        "MinCapacity": {
            "Type": "String"
        },
        "MaxCapacity": {
            "Type": "String"
        }
    },
    "Resources": {
        "RDSDBCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "Engine": "aurora-mysql",
                "DBClusterIdentifier": {
                    "Ref": "DBClusterIdentifier"
                },
                "EngineVersion": {
                    "Ref": "EngineVersion"
                },
                "MasterUsername": {
                    "Ref": "MasterUsername"
                },
                "MasterUserPassword": {
                    "Ref": "MasterUserPassword"
                },
                "ServerlessV2ScalingConfiguration": {
                    "MinCapacity": {
                        "Ref": "MinCapacity"
                    },
                    "MaxCapacity": {
                        "Ref": "MaxCapacity"
                    }
                }
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "Engine": "aurora-mysql",
                "DBInstanceClass": "db.serverless",
                "DBClusterIdentifier": {
                    "Ref": "RDSDBCluster"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: ServerlessV2 Cluster
Parameters:
  MasterUsername:
    Type: String
  MasterUserPassword:
    Type: String
    NoEcho: true
  DBClusterIdentifier:
    Type: String
  EngineVersion:
    Type: String
  MinCapacity:
    Type: String
  MaxCapacity:
    Type: String
Resources:
  RDSDBCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      Engine: aurora-mysql
      DBClusterIdentifier: !Ref DBClusterIdentifier
      EngineVersion: !Ref EngineVersion
      MasterUsername: !Ref MasterUsername
      MasterUserPassword: !Ref MasterUserPassword
      ServerlessV2ScalingConfiguration:
        MinCapacity: !Ref MinCapacity
        MaxCapacity: !Ref MaxCapacity
   RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      Engine: aurora-mysql
      DBInstanceClass: db.serverless
      DBClusterIdentifier: !Ref RDSDBCluster

```

### Creating an Amazon Aurora Serverless v1 DB cluster

The following example creates an Amazon Aurora Serverless v1 DB cluster. An
Aurora Serverless v1 DB cluster is a DB cluster that automatically starts up, shuts
down, and scales up or down its compute capacity based on your application's
needs. For more information about Aurora Serverless v1 DB clusters, see [Using Amazon\
Aurora Serverless](../../../amazonrds/latest/aurorauserguide/aurora-serverless.md) in the _Amazon Aurora User_
_Guide_.

###### Note

This example creates an Aurora MySQL Serverless DB cluster by setting
`Engine` to `aurora` and
`EngineVersion` to `5.6.10a`. To create an Aurora
PostgreSQL Serverless DB cluster, set `Engine` to
`aurora-postgresql` and `EngineVersion` to
`10.7`.

#### JSON

```json

{
  "AWSTemplateFormatVersion" : "2010-09-09",

  "Description" : "AWS CloudFormation Sample Template AuroraServerlessDBCluster: Sample template showing how to create an Amazon Aurora Serverless DB cluster. **WARNING** This template creates an Amazon Aurora DB cluster. You will be billed for the AWS resources used if you create a stack from this template.",

  "Parameters" : {
      "DBUsername" : {
        "NoEcho" : "true",
        "Description" : "Username for MySQL database access",
        "Type" : "String",
        "MinLength" : "1",
        "MaxLength" : "16",
        "AllowedPattern" : "[a-zA-Z][a-zA-Z0-9]*",
        "ConstraintDescription" : "must begin with a letter and contain only alphanumeric characters."
      },
      "DBPassword" : {
        "NoEcho" : "true",
        "Description" : "Password MySQL database access",
        "Type" : "String",
        "MinLength" : "8",
        "MaxLength" : "41",
        "AllowedPattern" : "[a-zA-Z0-9]*",
        "ConstraintDescription" : "must contain only alphanumeric characters."
      }
  },
  "Resources" : {
      "RDSCluster" : {
          "Type": "AWS::RDS::DBCluster",
          "Properties" : {
              "MasterUsername" : {
                  "Ref": "DBUsername"
              },
              "MasterUserPassword" : {
                  "Ref": "DBPassword"
              },
              "DBClusterIdentifier" : "my-serverless-cluster",
              "Engine" : "aurora",
              "EngineVersion" : "5.6.10a",
              "EngineMode" : "serverless",
              "ScalingConfiguration" : {
                  "AutoPause" : true,
                  "MinCapacity" : 4,
                  "MaxCapacity" : 32,
                  "SecondsUntilAutoPause" : 1000
              }
          }
      }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  AWS CloudFormation Sample Template AuroraServerlessDBCluster: Sample template
  showing how to create an Amazon Aurora Serverless DB cluster. **WARNING** This
  template creates an Amazon Aurora DB cluster. You will be billed for the AWS
  resources used if you create a stack from this template.
Parameters:
  DBUsername:
    NoEcho: 'true'
    Description: Username for MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  DBPassword:
    NoEcho: 'true'
    Description: Password MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername: !Ref DBUsername
      MasterUserPassword: !Ref DBPassword
      DBClusterIdentifier: my-serverless-cluster
      Engine: aurora
      EngineVersion: 5.6.10a
      EngineMode: serverless
      ScalingConfiguration:
        AutoPause: true
        MinCapacity: 4
        MaxCapacity: 32
        SecondsUntilAutoPause: 1000
```

### Create a Secrets Manager secret for a master password

The following example creates a AWS CloudFormation stack with the AWS::RDS::DBInstance resource with managed master user password feature.
The secret which is used to authenticate the DB instance is shown in the `Secret` stack output.
This output can be removed if showing this value to other CloudFormation stack isn't required.

For more information about managing password with Secrets Manager,
see [Password management with AWS Secrets Manager](../../../amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide_
and [Password management with AWS Secrets Manager](../../../amazonrds/latest/aurorauserguide/rds-secrets-manager.md)
in the _Amazon Aurora User Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "A sample template",
    "Resources": {
        "MyCluster": {
            "Properties": {
                "Engine": "aurora-mysql",
                "MasterUsername": "masteruser",
                "ManageMasterUserPassword": true,
                "MasterUserSecret": {
                    "KmsKeyId": {
                        "Ref": "KMSKey"
                    }
                }
            },
            "Type": "AWS::RDS::DBCluster"
        },
        "MyInstance": {
            "Properties": {
                "DBClusterIdentifier": {
                    "Ref": "MyCluster"
                },
                "DBInstanceClass": "db.t2.small",
                "Engine": "aurora-mysql",
                "PubliclyAccessible": "true"
            },
            "Type": "AWS::RDS::DBInstance"
        },
        "KMSKey": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "Description": "Manual test KMS key",
                "EnableKeyRotation": true,
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Id": {
                        "Ref": "AWS::StackName"
                    },
                    "Statement": [
                        {
                            "Sid": "Allow administration of the key",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Sub": "arn:${AWS::Partition}:iam::${AWS::AccountId}:root"
                                }
                            },
                            "Action": [
                                "kms:*"
                            ],
                            "Resource": "*"
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "Secret": {
            "Value": {
                "Fn::GetAtt": "MyCluster.MasterUserSecret.SecretArn"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: A sample template
Resources:
  MyCluster:
    Properties:
      Engine: aurora-mysql
      MasterUsername: masteruser
      ManageMasterUserPassword: true
      MasterUserSecret:
        KmsKeyId: !Ref KMSKey
    Type: "AWS::RDS::DBCluster"
  MyInstance:
    Properties:
      DBClusterIdentifier:
        Ref: MyCluster
      DBInstanceClass: db.t2.small
      Engine: aurora-mysql
      PubliclyAccessible: "true"
    Type: "AWS::RDS::DBInstance"
  KMSKey:
    Type: 'AWS::KMS::Key'
    Properties:
      Description: Manual test KMS key
      EnableKeyRotation: True
      KeyPolicy:
        Version: "2012-10-17"
        Id: !Ref "AWS::StackName"
        Statement:
          - Sid: "Allow administration of the key"
            Effect: "Allow"
            Principal:
              AWS:
                Fn::Sub: 'arn:${AWS::Partition}:iam::${AWS::AccountId}:root'
            Action:
              - "kms:*"
            Resource: "*"
Outputs:
  Secret:
    Value:
      Fn::GetAtt: MyCluster.MasterUserSecret.SecretArn

```

### Restoring an Amazon Aurora DB cluster

The following example restores an Amazon Aurora DB cluster. The example assumes that `SourceDBCluster` property already exists.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation Sample Template To Restore AuroraDBCluster: Sample template showing how to restore an Amazon Aurora DB cluster. **WARNING** This template creates an Amazon Aurora DB cluster. You will be billed for the AWS resources used if you create a stack from this template.",
    "Resources": {
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "DBClusterIdentifier": "my-multi-az-cluster-pit",
                "SourceDBClusterIdentifier": "my-multi-az-cluster",
                "RestoreToTime": "2023-05-29T23:50:00Z",
                "DBClusterInstanceClass": "db.r6gd.large",
                "StorageType": "io1",
                "Iops": 1000,
                "PubliclyAccessible": true
            }
        }
    }
}
```

#### YAML

```yaml

Description: >-
  AWS CloudFormation Sample Template To Restore AuroraDBCluster: Sample template
  showing how to restore an Amazon Aurora DB cluster. **WARNING** This
  template creates an Amazon Aurora DB cluster. You will be billed for the AWS
  resources used if you create a stack from this template.

Resources:
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      DBClusterIdentifier: my-multi-az-cluster-pit
      SourceDBClusterIdentifier: my-multi-az-cluster
      RestoreToTime: 2023-05-29T23:50:00Z
      DBClusterInstanceClass: db.r6gd.large
      StorageType: io1
      Iops: 1000
      PubliclyAccessible: true
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DBClusterRole

All content copied from https://docs.aws.amazon.com/.
