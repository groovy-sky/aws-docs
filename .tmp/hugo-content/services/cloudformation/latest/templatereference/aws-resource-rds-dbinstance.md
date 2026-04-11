This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBInstance

The `AWS::RDS::DBInstance` resource creates an Amazon DB instance. The new DB instance can be an RDS DB instance, or it can be a DB instance
in an Aurora DB cluster.

For more information about creating an RDS DB instance, see [Creating an Amazon \
RDS DB instance](../../../amazonrds/latest/userguide/user-createdbinstance.md) in the _Amazon RDS User Guide_.

For more information about creating a DB instance in an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](../../../amazonrds/latest/aurorauserguide/aurora-createinstance.md) in the _Amazon Aurora User Guide_.

If you import an existing DB instance, and the template configuration doesn't match the actual configuration of the DB instance,
AWS CloudFormation applies the changes in the template during the import operation.

###### Important

If a DB instance is deleted or replaced during an update, AWS CloudFormation
deletes all automated snapshots. However, it retains manual DB snapshots. During an
update that requires replacement, you can apply a stack policy to prevent DB
instances from being replaced. For more information, see [Prevent\
Updates to Stack Resources](../userguide/protect-stack-resources.md).

**Updating DB instances**

When properties labeled " _Update requires:_ [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)"
are updated, AWS CloudFormation first creates a replacement DB
instance, then changes references from other dependent resources to point to the
replacement DB instance, and finally deletes the old DB instance.

###### Important

We highly recommend that you take a snapshot of the database before updating the
stack. If you don't, you lose the data when AWS CloudFormation replaces your DB
instance. To preserve your data, perform the following procedure:

1. Deactivate any applications that are using the DB instance so that there's
    no activity on the DB instance.

2. Create a snapshot of the DB instance. For more information, see [Creating a DB\
    Snapshot](../../../amazonrds/latest/userguide/user-createsnapshot.md).

3. If you want to restore your instance using a DB snapshot, modify the
    updated template with your DB instance changes and add the
    `DBSnapshotIdentifier` property with the ID of the DB
    snapshot that you want to use.

After you restore a DB instance with a `DBSnapshotIdentifier`
    property, you can delete the `DBSnapshotIdentifier`
    property. When you specify this property for an update, the DB instance is not restored from the DB snapshot
    again, and the data in the database is not changed. However, if you don't
    specify the `DBSnapshotIdentifier` property, an empty DB instance
    is created, and the original DB instance is deleted. If you specify a
    property that is different from the previous snapshot restore property, a
    new DB instance is restored from the specified
    `DBSnapshotIdentifier` property, and the original DB instance
    is deleted.

4. Update the stack.

For more information about updating other properties of this resource, see `
                    ModifyDBInstance
                `. For more information about updating stacks,
see [AWS \
CloudFormation Stacks Updates](../userguide/using-cfn-updating-stacks.md).

**Deleting DB instances**

For DB instances that are part of an Aurora DB cluster, you can set a deletion policy
for your DB instance to control how AWS CloudFormation handles the DB instance when the
stack is deleted. For Amazon RDS DB instances, you can choose to
_retain_ the DB instance, to _delete_ the DB
instance, or to _create a snapshot_ of the DB instance. The default
AWS CloudFormation behavior depends on the `DBClusterIdentifier`
property:

1. For `AWS::RDS::DBInstance` resources that don't specify the
    `DBClusterIdentifier` property, AWS CloudFormation saves a
    snapshot of the DB instance.

2. For `AWS::RDS::DBInstance` resources that do specify the
    `DBClusterIdentifier` property, AWS CloudFormation deletes the DB
    instance.

For more information, see [DeletionPolicy Attribute](../userguide/aws-attribute-deletionpolicy.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBInstance",
  "Properties" : {
      "AdditionalStorageVolumes" : [ AdditionalStorageVolume, ... ],
      "AllocatedStorage" : String,
      "AllowMajorVersionUpgrade" : Boolean,
      "ApplyImmediately" : Boolean,
      "AssociatedRoles" : [ DBInstanceRole, ... ],
      "AutomaticBackupReplicationKmsKeyId" : String,
      "AutomaticBackupReplicationRegion" : String,
      "AutomaticBackupReplicationRetentionPeriod" : Integer,
      "AutoMinorVersionUpgrade" : Boolean,
      "AvailabilityZone" : String,
      "BackupRetentionPeriod" : Integer,
      "BackupTarget" : String,
      "CACertificateIdentifier" : String,
      "CertificateRotationRestart" : Boolean,
      "CharacterSetName" : String,
      "CopyTagsToSnapshot" : Boolean,
      "CustomIAMInstanceProfile" : String,
      "DatabaseInsightsMode" : String,
      "DBClusterIdentifier" : String,
      "DBClusterSnapshotIdentifier" : String,
      "DBInstanceClass" : String,
      "DBInstanceIdentifier" : String,
      "DBName" : String,
      "DBParameterGroupName" : String,
      "DBSecurityGroups" : [ String, ... ],
      "DBSnapshotIdentifier" : String,
      "DBSubnetGroupName" : String,
      "DBSystemId" : String,
      "DedicatedLogVolume" : Boolean,
      "DeleteAutomatedBackups" : Boolean,
      "DeletionProtection" : Boolean,
      "Domain" : String,
      "DomainAuthSecretArn" : String,
      "DomainDnsIps" : [ String, ... ],
      "DomainFqdn" : String,
      "DomainIAMRoleName" : String,
      "DomainOu" : String,
      "EnableCloudwatchLogsExports" : [ String, ... ],
      "EnableIAMDatabaseAuthentication" : Boolean,
      "EnablePerformanceInsights" : Boolean,
      "Engine" : String,
      "EngineLifecycleSupport" : String,
      "EngineVersion" : String,
      "Iops" : Integer,
      "KmsKeyId" : String,
      "LicenseModel" : String,
      "ManageMasterUserPassword" : Boolean,
      "MasterUserAuthenticationType" : String,
      "MasterUsername" : String,
      "MasterUserPassword" : String,
      "MasterUserSecret" : MasterUserSecret,
      "MaxAllocatedStorage" : Integer,
      "MonitoringInterval" : Integer,
      "MonitoringRoleArn" : String,
      "MultiAZ" : Boolean,
      "NcharCharacterSetName" : String,
      "NetworkType" : String,
      "OptionGroupName" : String,
      "PerformanceInsightsKMSKeyId" : String,
      "PerformanceInsightsRetentionPeriod" : Integer,
      "Port" : String,
      "PreferredBackupWindow" : String,
      "PreferredMaintenanceWindow" : String,
      "ProcessorFeatures" : [ ProcessorFeature, ... ],
      "PromotionTier" : Integer,
      "PubliclyAccessible" : Boolean,
      "ReplicaMode" : String,
      "RestoreTime" : String,
      "SourceDBClusterIdentifier" : String,
      "SourceDBInstanceAutomatedBackupsArn" : String,
      "SourceDBInstanceIdentifier" : String,
      "SourceDbiResourceId" : String,
      "SourceRegion" : String,
      "StorageEncrypted" : Boolean,
      "StorageThroughput" : Integer,
      "StorageType" : String,
      "Tags" : [ Tag, ... ],
      "Timezone" : String,
      "UseDefaultProcessorFeatures" : Boolean,
      "UseLatestRestorableTime" : Boolean,
      "VPCSecurityGroups" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBInstance
Properties:
  AdditionalStorageVolumes:
    - AdditionalStorageVolume
  AllocatedStorage: String
  AllowMajorVersionUpgrade: Boolean
  ApplyImmediately: Boolean
  AssociatedRoles:
    - DBInstanceRole
  AutomaticBackupReplicationKmsKeyId: String
  AutomaticBackupReplicationRegion: String
  AutomaticBackupReplicationRetentionPeriod: Integer
  AutoMinorVersionUpgrade: Boolean
  AvailabilityZone: String
  BackupRetentionPeriod: Integer
  BackupTarget: String
  CACertificateIdentifier: String
  CertificateRotationRestart: Boolean
  CharacterSetName: String
  CopyTagsToSnapshot: Boolean
  CustomIAMInstanceProfile: String
  DatabaseInsightsMode: String
  DBClusterIdentifier: String
  DBClusterSnapshotIdentifier: String
  DBInstanceClass: String
  DBInstanceIdentifier: String
  DBName: String
  DBParameterGroupName: String
  DBSecurityGroups:
    - String
  DBSnapshotIdentifier: String
  DBSubnetGroupName: String
  DBSystemId: String
  DedicatedLogVolume: Boolean
  DeleteAutomatedBackups: Boolean
  DeletionProtection: Boolean
  Domain: String
  DomainAuthSecretArn: String
  DomainDnsIps:
    - String
  DomainFqdn: String
  DomainIAMRoleName: String
  DomainOu: String
  EnableCloudwatchLogsExports:
    - String
  EnableIAMDatabaseAuthentication: Boolean
  EnablePerformanceInsights: Boolean
  Engine: String
  EngineLifecycleSupport: String
  EngineVersion: String
  Iops: Integer
  KmsKeyId: String
  LicenseModel: String
  ManageMasterUserPassword: Boolean
  MasterUserAuthenticationType: String
  MasterUsername: String
  MasterUserPassword: String
  MasterUserSecret:
    MasterUserSecret
  MaxAllocatedStorage: Integer
  MonitoringInterval: Integer
  MonitoringRoleArn: String
  MultiAZ: Boolean
  NcharCharacterSetName: String
  NetworkType: String
  OptionGroupName: String
  PerformanceInsightsKMSKeyId: String
  PerformanceInsightsRetentionPeriod: Integer
  Port: String
  PreferredBackupWindow: String
  PreferredMaintenanceWindow: String
  ProcessorFeatures:
    - ProcessorFeature
  PromotionTier: Integer
  PubliclyAccessible: Boolean
  ReplicaMode: String
  RestoreTime: String
  SourceDBClusterIdentifier: String
  SourceDBInstanceAutomatedBackupsArn: String
  SourceDBInstanceIdentifier: String
  SourceDbiResourceId: String
  SourceRegion: String
  StorageEncrypted: Boolean
  StorageThroughput: Integer
  StorageType: String
  Tags:
    - Tag
  Timezone: String
  UseDefaultProcessorFeatures: Boolean
  UseLatestRestorableTime: Boolean
  VPCSecurityGroups:
    - String

```

## Properties

`AdditionalStorageVolumes`

The additional storage volumes associated with the DB instance. RDS supports
additional storage volumes for RDS for Oracle and RDS for SQL Server.

_Required_: No

_Type_: Array of [AdditionalStorageVolume](aws-properties-rds-dbinstance-additionalstoragevolume.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllocatedStorage`

The amount of storage in gibibytes (GiB) to be initially allocated for the database
instance.

###### Note

If any value is set in the `Iops` parameter,
`AllocatedStorage` must be at least 100 GiB, which corresponds to the
minimum Iops value of 1,000. If you increase the `Iops` value (in 1,000
IOPS increments), then you must also increase the `AllocatedStorage`
value (in 100-GiB increments).

**Amazon Aurora**

Not applicable. Aurora cluster volumes automatically grow as the amount of data in your
database increases, though you are only charged for the space that you use in an Aurora cluster volume.

**Db2**

Constraints to the amount of storage for each storage type are the following:

- General Purpose (SSD) storage (gp3): Must be an integer from 20 to 64000.

- Provisioned IOPS storage (io1): Must be an integer from 100 to 64000.

**MySQL**

Constraints to the amount of storage for each storage type are the following:

- General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.

- Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.

- Magnetic storage (standard): Must be an integer from 5 to 3072.

**MariaDB**

Constraints to the amount of storage for each storage type are the following:

- General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.

- Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.

- Magnetic storage (standard): Must be an integer from 5 to 3072.

**PostgreSQL**

Constraints to the amount of storage for each storage type are the following:

- General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.

- Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.

- Magnetic storage (standard): Must be an integer from 5 to 3072.

**Oracle**

Constraints to the amount of storage for each storage type are the following:

- General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.

- Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.

- Magnetic storage (standard): Must be an integer from 10 to 3072.

**SQL Server**

Constraints to the amount of storage for each storage type are the following:

- General Purpose (SSD) storage (gp2):

- Enterprise and Standard editions: Must be an integer from 20 to 16384.

- Web and Express editions: Must be an integer from 20 to 16384.

- Provisioned IOPS storage (io1):

- Enterprise and Standard editions: Must be an integer from 20 to 16384.

- Web and Express editions: Must be an integer from 20 to 16384.

- Magnetic storage (standard):

- Enterprise and Standard editions: Must be an integer from 20 to 1024.

- Web and Express editions: Must be an integer from 20 to 1024.

_Required_: Conditional

_Type_: String

_Pattern_: `^[0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowMajorVersionUpgrade`

A value that indicates whether major version upgrades are allowed. Changing this
parameter doesn't result in an outage and the change is asynchronously applied as soon
as possible.

Constraints: Major version upgrades must be allowed when specifying a value for the
`EngineVersion` parameter that is a different major version than the DB
instance's current version.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplyImmediately`

Specifies whether changes to the DB instance and any pending modifications are applied
immediately, regardless of the `PreferredMaintenanceWindow` setting. If set to
`false`, changes are applied during the next maintenance window. Until RDS
applies the changes, the DB instance remains in a drift state. As a result, the
configuration doesn't fully reflect the requested modifications and temporarily diverges
from the intended state.

In addition to the settings described in [Modifying a DB\
instance](../../../amazonrds/latest/userguide/overview-dbinstance-modifying.md), this property also determines whether the DB instance reboots when a
static parameter is modified in the associated DB parameter group.

Default: `true`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociatedRoles`

The AWS Identity and Access Management (IAM) roles associated with the DB instance.

**Amazon Aurora**

Not applicable. The associated roles are managed by the DB cluster.

_Required_: No

_Type_: Array of [DBInstanceRole](aws-properties-rds-dbinstance-dbinstancerole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticBackupReplicationKmsKeyId`

The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the
Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS Region, for example,
`arn:aws:kms:us-east-1:123456789012:key/AWS_ACCESS_KEY_ID_REDACTED`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticBackupReplicationRegion`

The AWS Region associated with the automated backup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticBackupReplicationRetentionPeriod`

The retention period for automated backups in a different AWS Region.
Use this parameter to set a unique retention period that only applies to cross-Region
automated backups. To enable automated backups in a different Region, specify a positive
value for the `AutomaticBackupReplicationRegion` parameter.

If not specified, this parameter defaults to the value of the `BackupRetentionPeriod` parameter. The maximum allowed value is 35.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoMinorVersionUpgrade`

A value that indicates whether minor engine upgrades are applied automatically to the
DB instance during the maintenance window. By default, minor engine upgrades are applied
automatically.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AvailabilityZone`

The Availability Zone (AZ) where the database will be created. For information on
AWS Regions and Availability Zones, see
[Regions\
and Availability Zones](../../../amazonrds/latest/userguide/concepts-regionsandavailabilityzones.md).

For Amazon Aurora, each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these
Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.

Default: A random, system-chosen Availability Zone in the endpoint's AWS Region.

Constraints:

- The `AvailabilityZone` parameter can't be specified if the DB instance is a Multi-AZ deployment.

- The specified Availability Zone must be in the same AWS Region as the current endpoint.

Example: `us-east-1d`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`BackupRetentionPeriod`

The number of days for which automated backups are retained. Setting this parameter to
a positive number enables backups. Setting this parameter to 0 disables automated
backups.

**Amazon Aurora**

Not applicable. The retention period for automated backups is managed by the DB
cluster.

Default: 1

Constraints:

- Must be a value from 0 to 35

- Can't be set to 0 if the DB instance is a source to read replicas

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`BackupTarget`

The location for storing automated backups and manual snapshots.

Valid Values:

- `local` (Dedicated Local Zone)

- `outposts` (AWS Outposts)

- `region` (AWS Region)

Default: `region`

For more information, see [Working \
with Amazon RDS on AWS Outposts](../../../amazonrds/latest/userguide/rds-on-outposts.md) in the _Amazon RDS User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CACertificateIdentifier`

The identifier of the CA certificate for this DB instance.

For more information, see [Using SSL/TLS to encrypt a connection to a DB \
instance](../../../amazonrds/latest/userguide/usingwithrds-ssl.md) in the _Amazon RDS User Guide_ and
[Using SSL/TLS to encrypt a connection to a DB cluster](../../../amazonrds/latest/aurorauserguide/usingwithrds-ssl.md) in the _Amazon Aurora_
_User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateRotationRestart`

Specifies whether the DB instance is restarted when you rotate your
SSL/TLS certificate.

By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate
is not updated until the DB instance is restarted.

###### Important

Set this parameter only if you are _not_ using SSL/TLS to connect to the DB instance.

If you are using SSL/TLS to connect to the DB instance, follow the appropriate instructions for your
DB engine to rotate your SSL/TLS certificate:

- For more information about rotating your SSL/TLS certificate for RDS DB engines, see
[Rotating Your SSL/TLS Certificate.](../../../amazonrds/latest/userguide/usingwithrds-ssl-certificate-rotation.md) in the _Amazon RDS User Guide._

- For more information about rotating your SSL/TLS certificate for Aurora DB engines, see
[Rotating Your SSL/TLS Certificate](../../../amazonrds/latest/aurorauserguide/usingwithrds-ssl-certificate-rotation.md) in the _Amazon Aurora User Guide_.

This setting doesn't apply to RDS Custom DB instances.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CharacterSetName`

For supported engines, indicates that the DB instance should be associated with the
specified character set.

**Amazon Aurora**

Not applicable. The character set is managed by the DB cluster. For more information,
see [AWS::RDS::DBCluster](../userguide/aws-resource-rds-dbcluster.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CopyTagsToSnapshot`

Specifies whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.

This setting doesn't apply to Amazon Aurora DB instances. Copying tags to snapshots is managed by the DB cluster. Setting this
value for an Aurora DB instance has no effect on the DB cluster setting.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomIAMInstanceProfile`

The instance profile associated with the underlying Amazon EC2 instance of an
RDS Custom DB instance.

This setting is required for RDS Custom.

Constraints:

- The profile must exist in your account.

- The profile must have an IAM role that Amazon EC2 has permissions to assume.

- The instance profile name and the associated IAM role name must start with the prefix `AWSRDSCustom`.

For the list of permissions required for the IAM role, see
[Configure IAM and your VPC](../../../amazonrds/latest/userguide/custom-setup-orcl.md#custom-setup-orcl.iam-vpc) in the _Amazon RDS User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseInsightsMode`

The mode of Database Insights to enable for the DB instance.

###### Note

Aurora DB instances inherit this value from the DB cluster, so you can't change this value.

_Required_: No

_Type_: String

_Allowed values_: `standard | advanced`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterIdentifier`

The identifier of the DB cluster that this DB instance will belong to.

This setting doesn't apply to RDS Custom DB instances.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBClusterSnapshotIdentifier`

The identifier for the Multi-AZ DB cluster snapshot to restore from.

For more information on Multi-AZ DB clusters, see [Multi-AZ DB\
cluster deployments](../../../amazonrds/latest/userguide/multi-az-db-clusters-concepts.md) in the _Amazon RDS User_
_Guide_.

Constraints:

- Must match the identifier of an existing Multi-AZ DB cluster snapshot.

- Can't be specified when `DBSnapshotIdentifier` is specified.

- Must be specified when `DBSnapshotIdentifier` isn't specified.

- If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the `DBClusterSnapshotIdentifier`
must be the ARN of the shared snapshot.

- Can't be the identifier of an Aurora DB cluster snapshot.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBInstanceClass`

The compute and memory capacity of the DB instance, for example `db.m5.large`.
Not all DB instance classes are available in all AWS Regions, or for all database engines.
For the full list of DB instance classes, and availability for your engine, see
[DB instance \
classes](../../../amazonrds/latest/userguide/concepts-dbinstanceclass.md) in the _Amazon RDS User Guide_ or
[Aurora \
DB instance classes](../../../amazonrds/latest/aurorauserguide/concepts-dbinstanceclass.md) in the _Amazon Aurora User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBInstanceIdentifier`

A name for the DB instance. If you specify a name, AWS CloudFormation converts it to
lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical
ID and uses that ID for the DB instance. For more information, see [Name\
Type](../userguide/aws-properties-name.md).

For information about constraints that apply to DB instance identifiers, see
[Naming constraints in \
Amazon RDS](../../../amazonrds/latest/userguide/chap-limits.md#RDS_Limits.Constraints) in the _Amazon RDS User Guide_.

###### Important

If you specify a name, you can't perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Pattern_: `^$|^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBName`

The meaning of this parameter differs according to the database engine you use.

###### Important

If you specify the `

                DBSnapshotIdentifier
                            ` property, this property only applies to RDS for Oracle.

**Amazon Aurora**

Not applicable. The database name is managed by the DB cluster.

**Db2**

The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.

Constraints:

- Must contain 1 to 64 letters or numbers.

- Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).

- Can't be a word reserved by the specified database engine.

**MySQL**

The name of the database to create when the DB instance is created. If this parameter
is not specified, no database is created in the DB instance.

Constraints:

- Must contain 1 to 64 letters or numbers.

- Can't be a word reserved by the specified database engine

**MariaDB**

The name of the database to create when the DB instance is created. If this parameter
is not specified, no database is created in the DB instance.

Constraints:

- Must contain 1 to 64 letters or numbers.

- Can't be a word reserved by the specified database engine

**PostgreSQL**

The name of the database to create when the DB instance is created. If this parameter
is not specified, the default `postgres` database is created in the DB instance.

Constraints:

- Must begin with a letter. Subsequent characters can be
letters, underscores, or digits (0-9).

- Must contain 1 to 63 characters.

- Can't be a word reserved by the specified database engine

**Oracle**

The Oracle System ID (SID) of the created DB instance. If you specify
`null`, the default value `ORCL` is used. You can't specify
the string NULL, or any other reserved word, for `DBName`.

Default: `ORCL`

Constraints:

- Can't be longer than 8 characters

**SQL Server**

Not applicable. Must be null.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBParameterGroupName`

The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](../userguide/aws-properties-rds-dbparametergroup.md) resource created in the template.

To list all of the available DB parameter group names, use the following
command:

`aws rds describe-db-parameter-groups --query
                "DBParameterGroups[].DBParameterGroupName" --output text`

###### Note

If any of the data members of the referenced parameter group are changed during an
update, the DB instance might need to be restarted, which causes some interruption.
If the parameter group contains static parameters, whether they were changed or not,
an update triggers a reboot.

If you don't specify a value for `DBParameterGroupName` property,
the default DB parameter group for the specified engine and engine version is used.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBSecurityGroups`

A list of the DB security groups to assign to the DB instance. The list can include
both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup
resources created in the template.

If you set DBSecurityGroups, you must not set VPCSecurityGroups, and vice versa.
Also, note that the DBSecurityGroups property exists only for backwards compatibility with older
regions and is no longer recommended for providing security information to an RDS DB
instance. Instead, use VPCSecurityGroups.

###### Important

If you specify this property, AWS CloudFormation sends only the following
properties (if specified) to Amazon RDS during create operations:

- `AllocatedStorage`

- `AutoMinorVersionUpgrade`

- `AvailabilityZone`

- `BackupRetentionPeriod`

- `CharacterSetName`

- `DBInstanceClass`

- `DBName`

- `DBParameterGroupName`

- `DBSecurityGroups`

- `DBSubnetGroupName`

- `Engine`

- `EngineVersion`

- `Iops`

- `LicenseModel`

- `MasterUsername`

- `MasterUserPassword`

- `MultiAZ`

- `OptionGroupName`

- `PreferredBackupWindow`

- `PreferredMaintenanceWindow`

All other properties are ignored. Specify a virtual private cloud (VPC) security
group if you want to submit other properties, such as `StorageType`,
`StorageEncrypted`, or `KmsKeyId`. If you're already using
the `DBSecurityGroups` property, you can't use these other properties by
updating your DB instance to use a VPC security group. You must recreate the DB
instance.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSnapshotIdentifier`

The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the
DB instance. If you're restoring from a shared manual DB snapshot, you must specify the
ARN of the snapshot.

By specifying this property, you can create a DB instance from the specified DB
snapshot. If the `DBSnapshotIdentifier` property is an empty string or the
`AWS::RDS::DBInstance` declaration has no
`DBSnapshotIdentifier` property, AWS CloudFormation creates a new
database. If the property contains a value (other than an empty string), AWS
CloudFormation creates a database from the specified snapshot. If a snapshot with the
specified name doesn't exist, AWS CloudFormation can't create the database and it rolls
back the stack.

Some DB instance properties aren't valid when you restore from a snapshot, such as the
`MasterUsername` and `MasterUserPassword` properties, and the
point-in-time recovery properties `RestoreTime` and
`UseLatestRestorableTime`. For information about the properties that you can
specify, see the [`RestoreDBInstanceFromDBSnapshot`](../../../../reference/amazonrds/latest/apireference/api-restoredbinstancefromdbsnapshot.md) action in the _Amazon_
_RDS API Reference_.

After you restore a DB instance with a `DBSnapshotIdentifier` property, you
must specify the same `DBSnapshotIdentifier` property for any future updates
to the DB instance. When you specify this property for an update, the DB instance is not
restored from the DB snapshot again, and the data in the database is not changed.
However, if you don't specify the `DBSnapshotIdentifier` property, an empty
DB instance is created, and the original DB instance is deleted. If you specify a
property that is different from the previous snapshot restore property, a new DB
instance is restored from the specified `DBSnapshotIdentifier` property, and
the original DB instance is deleted.

If you specify the `DBSnapshotIdentifier` property to restore a DB instance (as opposed to specifying it for DB instance updates),
then don't specify the following properties:

- `CharacterSetName`

- `DBClusterIdentifier`

- `DBName`

- `KmsKeyId`

- `MasterUsername`

- `MasterUserPassword`

- `PromotionTier`

- `SourceDBInstanceIdentifier`

- `SourceRegion`

- `StorageEncrypted` (for an unencrypted snapshot)

- `Timezone`

**Amazon Aurora**

Not applicable. Snapshot restore is managed by the DB cluster.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBSubnetGroupName`

A DB subnet group to associate with the DB instance. If you update this value, the new
subnet group must be a subnet group in a new VPC.

If you don't specify a DB subnet group, RDS uses the default DB subnet group if one
exists. If a default DB subnet group does not exist, and you don't specify a
`DBSubnetGroupName`, the DB instance fails to launch.

For more information about using Amazon RDS in a VPC, see [Amazon VPC and Amazon RDS](../../../amazonrds/latest/userguide/user-vpc.md)
in the _Amazon RDS User Guide_.

This setting doesn't apply to Amazon Aurora DB instances. The DB subnet group is managed by the DB cluster. If specified, the setting must match the DB cluster setting.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBSystemId`

The Oracle system identifier (SID), which is the name of the Oracle database instance that
manages your database files. In this context, the term "Oracle database instance" refers exclusively
to the system global area (SGA) and Oracle background processes. If you don't specify a SID,
the value defaults to `RDSCDB`. The Oracle SID is also the name of your CDB.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DedicatedLogVolume`

Indicates whether the DB instance has a dedicated log volume (DLV) enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteAutomatedBackups`

A value that indicates whether to remove automated backups immediately after the DB
instance is deleted. This parameter isn't case-sensitive. The default is to remove
automated backups immediately after the DB instance is deleted.

**Amazon Aurora**

Not applicable. When you delete a DB cluster, all automated backups for that DB cluster are deleted and can't be recovered.
Manual DB cluster snapshots of the DB cluster are not deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtection`

Specifies whether the DB instance has deletion protection enabled.
The database can't be deleted when deletion protection is enabled. By default,
deletion protection isn't enabled. For more information, see
[Deleting a DB Instance](../../../amazonrds/latest/userguide/user-deleteinstance.md).

This setting doesn't apply to Amazon Aurora DB instances. You can enable or disable deletion protection for the DB cluster.
For more information, see `CreateDBCluster`. DB instances in a DB
cluster can be deleted even when deletion protection is enabled for the DB cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The Active Directory directory ID to create the DB instance in. Currently, only Db2, MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.

For more information, see [Kerberos Authentication](../../../amazonrds/latest/userguide/kerberos-authentication.md) in the _Amazon RDS User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainAuthSecretArn`

The ARN for the Secrets Manager secret with the credentials for the user joining the domain.

Example: `arn:aws:secretsmanager:region:account-number:secret:myselfmanagedADtestsecret-123456`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DomainDnsIps`

The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.

Constraints:

- Two IP addresses must be provided. If there isn't a secondary domain controller, use the IP address of the primary domain controller for both entries in the list.

Example: `123.124.125.126,234.235.236.237`

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DomainFqdn`

The fully qualified domain name (FQDN) of an Active Directory domain.

Constraints:

- Can't be longer than 64 characters.

Example: `mymanagedADtest.mymanagedAD.mydomain`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DomainIAMRoleName`

The name of the IAM role to use when making API calls to the Directory Service.

This setting doesn't apply to the following DB instances:

- Amazon Aurora (The domain is managed by the DB cluster.)

- RDS Custom

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainOu`

The Active Directory organizational unit for your DB instance to join.

Constraints:

- Must be in the distinguished name format.

- Can't be longer than 64 characters.

Example: `OU=mymanagedADtestOU,DC=mymanagedADtest,DC=mymanagedAD,DC=mydomain`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EnableCloudwatchLogsExports`

The list of log types that need to be enabled for exporting to CloudWatch Logs. The values
in the list depend on the DB engine being used. For more information, see
[Publishing Database Logs to Amazon CloudWatch Logs](../../../amazonrds/latest/userguide/user-logaccess.md#USER_LogAccess.Procedural.UploadtoCloudWatch) in the _Amazon Relational Database_
_Service User Guide_.

**Amazon Aurora**

Not applicable. CloudWatch Logs exports are managed by the DB cluster.

**Db2**

Valid values: `diag.log`, `notify.log`

**MariaDB**

Valid values: `audit`, `error`, `general`, `slowquery`

**Microsoft SQL Server**

Valid values: `agent`, `error`

**MySQL**

Valid values: `audit`, `error`, `general`, `slowquery`

**Oracle**

Valid values: `alert`, `audit`, `listener`, `trace`, `oemagent`

**PostgreSQL**

Valid values: `postgresql`, `upgrade`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableIAMDatabaseAuthentication`

A value that indicates whether to enable mapping of AWS Identity and Access
Management (IAM) accounts to database accounts. By default, mapping is disabled.

This property is supported for RDS for MariaDB, RDS for MySQL, and RDS for PostgreSQL.
For more information, see [IAM Database Authentication for MariaDB, MySQL, and PostgreSQL](../../../amazonrds/latest/userguide/usingwithrds-iamdbauth.md) in the _Amazon RDS User Guide._

**Amazon Aurora**

Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePerformanceInsights`

Specifies whether to enable Performance Insights for the DB instance. For more information, see
[Using Amazon Performance Insights](../../../amazonrds/latest/userguide/user-perfinsights.md) in the _Amazon RDS User Guide_.

This setting doesn't apply to RDS Custom DB instances.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The name of the database engine to use for this DB instance. Not every database engine is available in every AWS Region.

This property is required when creating a DB instance.

###### Note

You can convert an Oracle database from the non-CDB architecture to the container
database (CDB) architecture by updating the `Engine` value in your templates
from `oracle-ee` to `oracle-ee-cdb` or from
`oracle-se2` to `oracle-se2-cdb`. Converting to the CDB
architecture requires an interruption.

Valid Values:

- `aurora-mysql` (for Aurora MySQL DB instances)

- `aurora-postgresql` (for Aurora PostgreSQL DB instances)

- `custom-oracle-ee` (for RDS Custom for Oracle DB instances)

- `custom-oracle-ee-cdb` (for RDS Custom for Oracle DB instances)

- `custom-sqlserver-ee` (for RDS Custom for SQL Server DB instances)

- `custom-sqlserver-se` (for RDS Custom for SQL Server DB instances)

- `custom-sqlserver-web` (for RDS Custom for SQL Server DB instances)

- `db2-ae`

- `db2-se`

- `mariadb`

- `mysql`

- `oracle-ee`

- `oracle-ee-cdb`

- `oracle-se2`

- `oracle-se2-cdb`

- `postgres`

- `sqlserver-ee`

- `sqlserver-se`

- `sqlserver-ex`

- `sqlserver-web`

_Required_: Conditional

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EngineLifecycleSupport`

The life cycle type for this DB instance.

###### Note

By default, this value is set to `open-source-rds-extended-support`, which enrolls your DB instance into Amazon RDS Extended Support.
At the end of standard support, you can avoid charges for Extended Support by setting the value to `open-source-rds-extended-support-disabled`. In this case,
creating the DB instance will fail if the DB major version is past its end of standard support date.

This setting applies only to RDS for MySQL and RDS for PostgreSQL. For Amazon Aurora DB instances, the life cycle type is managed by the DB cluster.

You can use this setting to enroll your DB instance into Amazon RDS Extended Support. With RDS Extended Support,
you can run the selected major engine version on your DB instance past the end of standard support for that engine version. For more information, see [Amazon RDS Extended Support with Amazon RDS](../../../amazonrds/latest/userguide/extended-support.md) in the _Amazon RDS User Guide_.

Valid Values: `open-source-rds-extended-support | open-source-rds-extended-support-disabled`

Default: `open-source-rds-extended-support`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EngineVersion`

The version number of the database engine to use.

For a list of valid engine versions, use the `DescribeDBEngineVersions` action.

The following are the database engines and links to information about the major and minor versions that are available with
Amazon RDS. Not every database engine is available for every AWS Region.

**Amazon Aurora**

Not applicable. The version number of the database engine to be used by the DB
instance is managed by the DB cluster.

**Db2**

See [Amazon RDS for Db2](../../../amazonrds/latest/userguide/chap-db2.md#Db2.Concepts.VersionMgmt) in the
_Amazon RDS User Guide._

**MariaDB**

See [MariaDB on Amazon RDS Versions](../../../amazonrds/latest/userguide/chap-mariadb.md#MariaDB.Concepts.VersionMgmt) in the
_Amazon RDS User Guide._

**Microsoft SQL Server**

See [Microsoft SQL Server Versions on Amazon RDS](../../../amazonrds/latest/userguide/chap-sqlserver.md#SQLServer.Concepts.General.VersionSupport) in the
_Amazon RDS User Guide._

**MySQL**

See [MySQL on Amazon RDS Versions](../../../amazonrds/latest/userguide/chap-mysql.md#MySQL.Concepts.VersionMgmt) in the
_Amazon RDS User Guide._

**Oracle**

See [Oracle Database Engine Release Notes](../../../amazonrds/latest/userguide/appendix-oracle-patchcomposition.md) in the
_Amazon RDS User Guide._

**PostgreSQL**

See [Supported PostgreSQL Database Versions](../../../amazonrds/latest/userguide/chap-postgresql.md#PostgreSQL.Concepts.General.DBVersions) in the
_Amazon RDS User Guide._

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Iops`

The number of I/O operations per second (IOPS) that the database provisions. The value
must be equal to or greater than 1000.

If you specify this property, you must follow the range of allowed ratios of your
requested IOPS rate to the amount of storage that you allocate (IOPS to allocated
storage). For example, you can provision an Oracle database instance with 1000 IOPS and
200 GiB of storage (a ratio of 5:1), or specify 2000 IOPS with 200 GiB of storage (a ratio
of 10:1). For more information, see [Amazon RDS\
Provisioned IOPS Storage to Improve Performance](../../../amazonrds/latest/developerguide/chap-storage.md#USER_PIOPS) in the _Amazon RDS_
_User Guide_.

###### Note

If you specify `io1` for the `StorageType` property, then
you must also specify the `Iops` property.

Constraints:

- For RDS for Db2, MariaDB, MySQL, Oracle, and PostgreSQL - Must be a multiple between .5 and 50 of the storage amount for the DB instance.

- For RDS for SQL Server - Must be a multiple between 1 and 50 of the storage amount for the DB instance.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ARN of the AWS KMS key that's used to encrypt
the DB instance, such as
`arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef`.
If you enable the StorageEncrypted property but don't specify this property, AWS
CloudFormation uses the default KMS key. If you specify this property, you must set
the StorageEncrypted property to true.

If you specify the `SourceDBInstanceIdentifier` or
`SourceDbiResourceId` property, don't specify this property. The value is
inherited from the source DB instance, and if the DB instance is encrypted, the specified
`KmsKeyId` property is used. However, if the source DB instance is in a
different AWS Region, you must specify a KMS key ID.

If you specify the `SourceDBInstanceAutomatedBackupsArn` property, don't
specify this property. The value is inherited from the source DB instance automated backup,
and if the automated backup is encrypted, the specified `KmsKeyId` property is
used.

If you create an encrypted read replica in a different AWS Region, then you must
specify a KMS key for the destination AWS Region. KMS encryption keys are specific to
the region that they're created in, and you can't use encryption keys from one region in
another region.

If you specify the `DBSnapshotIdentifier` property, don't specify this
property. The `StorageEncrypted` property value is inherited from the snapshot.
If the DB instance is encrypted, the specified `KmsKeyId` property is also
inherited from the snapshot.

If you specify `DBSecurityGroups`, AWS CloudFormation ignores this
property. To specify both a security group and this property, you must use a VPC
security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](../../../amazonrds/latest/userguide/user-vpc.md)
in the _Amazon RDS User Guide_.

**Amazon Aurora**

Not applicable. The KMS key identifier is managed by the DB cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LicenseModel`

License model information for this DB instance.

Valid Values:

- Aurora MySQL - `general-public-license`

- Aurora PostgreSQL - `postgresql-license`

- RDS for Db2 - `bring-your-own-license`. For more information about RDS for Db2 licensing, see
[https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-licensing.html](../../../amazonrds/latest/userguide/db2-licensing.md) in the
_Amazon RDS User Guide._

- RDS for MariaDB - `general-public-license`

- RDS for Microsoft SQL Server - `license-included`

- RDS for MySQL - `general-public-license`

- RDS for Oracle - `bring-your-own-license` or `license-included`

- RDS for PostgreSQL - `postgresql-license`

###### Note

If you've specified `DBSecurityGroups` and then you update the license
model, AWS CloudFormation replaces the underlying DB instance. This will incur some
interruptions to database availability.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManageMasterUserPassword`

Specifies whether to manage the master user password with AWS Secrets Manager.

For more information, see [Password management with AWS Secrets Manager](../../../amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide._

Constraints:

- Can't manage the master user password with AWS Secrets Manager if `MasterUserPassword`
is specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserAuthenticationType`

Specifies the authentication type for the master user. With IAM master user authentication, you can configure the master DB user with IAM database authentication when you create a DB instance.

You can specify one of the following values:

- `password` \- Use standard database authentication with a password.

- `iam-db-auth` \- Use IAM database authentication for the master user.

This option is only valid for RDS for MySQL, RDS for MariaDB, RDS for PostgreSQL, Aurora MySQL, and Aurora PostgreSQL engines.

_Required_: No

_Type_: String

_Allowed values_: `password | iam-db-auth`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUsername`

The master user name for the DB instance.

###### Note

If you specify the `SourceDBInstanceIdentifier` or
`DBSnapshotIdentifier` property, don't specify this property. The
value is inherited from the source DB instance or snapshot.

When migrating a self-managed Db2 database, we recommend that you use the same master username as your self-managed Db2 instance name.

**Amazon Aurora**

Not applicable. The name for the master user is managed by the DB cluster.

**RDS for Db2**

Constraints:

- Must be 1 to 16 letters or numbers.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

**RDS for MariaDB**

Constraints:

- Must be 1 to 16 letters or numbers.

- Can't be a reserved word for the chosen database engine.

**RDS for Microsoft SQL Server**

Constraints:

- Must be 1 to 128 letters or numbers.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

**RDS for MySQL**

Constraints:

- Must be 1 to 16 letters or numbers.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

**RDS for Oracle**

Constraints:

- Must be 1 to 30 letters or numbers.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

**RDS for PostgreSQL**

Constraints:

- Must be 1 to 63 letters or numbers.

- First character must be a letter.

- Can't be a reserved word for the chosen database engine.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,127}$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MasterUserPassword`

The password for the master user. The password can include any printable ASCII
character except "/", """, or "@".

**Amazon Aurora**

Not applicable. The password for the master user is managed by the DB cluster.

**RDS for Db2**

Must contain from 8 to 255 characters.

**RDS for MariaDB**

Constraints: Must contain from 8 to 41 characters.

**RDS for Microsoft SQL Server**

Constraints: Must contain from 8 to 128 characters.

**RDS for MySQL**

Constraints: Must contain from 8 to 41 characters.

**RDS for Oracle**

Constraints: Must contain from 8 to 30 characters.

**RDS for PostgreSQL**

Constraints: Must contain from 8 to 128 characters.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserSecret`

The secret managed by RDS in AWS Secrets Manager for the master user password.

For more information, see [Password management with AWS Secrets Manager](../../../amazonrds/latest/userguide/rds-secrets-manager.md)
in the _Amazon RDS User Guide._

_Required_: No

_Type_: [MasterUserSecret](aws-properties-rds-dbinstance-masterusersecret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAllocatedStorage`

The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.

For more information about this setting, including limitations that apply to it, see
[Managing capacity automatically with Amazon RDS storage autoscaling](../../../amazonrds/latest/userguide/user-piops-storagetypes.md#USER_PIOPS.Autoscaling)
in the _Amazon RDS User Guide_.

This setting doesn't apply to the following DB instances:

- Amazon Aurora (Storage is managed by the DB cluster.)

- RDS Custom

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringInterval`

The interval, in seconds, between points when Enhanced Monitoring metrics are collected for
the DB instance. To disable collection of Enhanced Monitoring metrics, specify `0`.

If `MonitoringRoleArn` is specified, then you must set `MonitoringInterval`
to a value other than `0`.

This setting doesn't apply to RDS Custom DB instances.

Valid Values: `0 | 1 | 5 | 10 | 15 | 30 | 60`

Default: `0`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringRoleArn`

The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For
example, `arn:aws:iam:123456789012:role/emaccess`. For information on creating a monitoring role,
see [Setting Up and Enabling Enhanced Monitoring](../../../amazonrds/latest/userguide/user-monitoring-os.md#USER_Monitoring.OS.Enabling)
in the _Amazon RDS User Guide_.

If `MonitoringInterval` is set to a value other than `0`, then you must supply a `MonitoringRoleArn` value.

This setting doesn't apply to RDS Custom DB instances.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiAZ`

Specifies whether the DB instance is a Multi-AZ deployment. You can't set
the `AvailabilityZone` parameter if the DB instance is a Multi-AZ deployment.

This setting doesn't apply to Amazon Aurora because the DB instance Availability Zones (AZs)
are managed by the DB cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`NcharCharacterSetName`

The name of the NCHAR character set for the Oracle DB instance.

This setting doesn't apply to RDS Custom DB instances.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

The network type of the DB instance.

Valid values:

- `IPV4`

- `DUAL`

The network type is determined by the `DBSubnetGroup` specified for the DB instance.
A `DBSubnetGroup` can support only the IPv4 protocol or the IPv4 and IPv6
protocols ( `DUAL`).

For more information, see [Working with a DB instance in a VPC](../../../amazonrds/latest/userguide/user-vpc-workingwithrdsinstanceinavpc.md) in the
_Amazon RDS User Guide._

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OptionGroupName`

Indicates that the DB instance should be associated with the specified option
group.

Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be
removed from an option group. Also, that option group can't be removed from a DB
instance once it is associated with a DB instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceInsightsKMSKeyId`

The AWS KMS key identifier for encryption of Performance Insights data.

The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.

If you do not specify a value for `PerformanceInsightsKMSKeyId`, then Amazon RDS
uses your default KMS key. There is a default KMS key for your AWS account.
Your AWS account has a different default KMS key for each AWS Region.

###### Important

**Update behavior:** Once Performance Insights is enabled with a KMS key,
you cannot change to a different physical KMS key without replacing the DB instance. However, the following
updates do not require replacement:

- Enabling or disabling Performance Insights using the `EnablePerformanceInsights` property

- Changing between different identifier formats (key ARN, key ID, alias ARN, alias name) of the
same physical KMS key

- Removing the `PerformanceInsightsKMSKeyId` property from your template

###### Note

**Drift behavior:** If you specify `PerformanceInsightsKMSKeyId`
while `EnablePerformanceInsights` is set to `false`, CloudFormation will report drift.
This occurs because the RDS API does not allow setting a KMS key when Performance Insights is disabled.
CloudFormation ignores the `PerformanceInsightsKMSKeyId` value during instance creation to avoid
API errors, resulting in a mismatch between your template and the actual instance configuration.

To avoid drift, omit both `EnablePerformanceInsights` and `PerformanceInsightsKMSKeyId`
during initial instance creation, then set both properties together when you're ready to enable Performance Insights.

For information about enabling Performance Insights, see [EnablePerformanceInsights](../userguide/aws-properties-rds-database-instance.md#cfn-rds-dbinstance-enableperformanceinsights).

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PerformanceInsightsRetentionPeriod`

The number of days to retain Performance Insights data. When creating a DB instance without enabling Performance Insights, you can't specify the parameter `PerformanceInsightsRetentionPeriod`.

This setting doesn't apply to RDS Custom DB instances.

Valid Values:

- `7`

- _month_ \\* 31, where _month_ is a number of months from 1-23.
Examples: `93` (3 months \* 31), `341` (11 months \* 31), `589` (19 months \* 31)

- `731`

Default: `7` days

If you specify a retention period that isn't valid, such as `94`, Amazon RDS returns an error.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number on which the database accepts connections.

This setting doesn't apply to Aurora DB instances. The port number is managed by the cluster.

Valid Values: `1150-65535`

Default:

- RDS for Db2 - `50000`

- RDS for MariaDB - `3306`

- RDS for Microsoft SQL Server - `1433`

- RDS for MySQL - `3306`

- RDS for Oracle - `1521`

- RDS for PostgreSQL - `5432`

Constraints:

- For RDS for Microsoft SQL Server, the value can't be `1234`, `1434`,
`3260`, `3343`, `3389`, `47001`, or
`49152-49156`.

_Required_: No

_Type_: String

_Pattern_: `^\d*$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PreferredBackupWindow`

The daily time range during which automated backups are created if automated backups are enabled,
using the `BackupRetentionPeriod` parameter. For more information, see [Backup Window](../../../amazonrds/latest/userguide/user-workingwithautomatedbackups.md#USER_WorkingWithAutomatedBackups.BackupWindow) in the _Amazon RDS User Guide._

Constraints:

- Must be in the format `hh24:mi-hh24:mi`.

- Must be in Universal Coordinated Time (UTC).

- Must not conflict with the preferred maintenance window.

- Must be at least 30 minutes.

**Amazon Aurora**

Not applicable. The daily time range for creating automated backups is managed by
the DB cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur, in Universal
Coordinated Time (UTC).

Format: `ddd:hh24:mi-ddd:hh24:mi`

The default is a 30-minute window selected at random from an 8-hour block of time for
each AWS Region, occurring on a random day of the week. To see the time blocks
available, see [Maintaining a DB instance](../../../amazonrds/latest/userguide/user-upgradedbinstance-maintenance.md#AdjustingTheMaintenanceWindow) in the
_Amazon RDS User Guide._

###### Note

This property applies when AWS CloudFormation initially creates the DB instance.
If you use AWS CloudFormation to update the DB instance, those updates are applied
immediately.

Constraints: Minimum 30-minute window.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ProcessorFeatures`

The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.

This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.

_Required_: No

_Type_: Array of [ProcessorFeature](aws-properties-rds-dbinstance-processorfeature.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromotionTier`

The order of priority in which an Aurora Replica is promoted to the primary instance
after a failure of the existing primary instance. For more information,
see [Fault Tolerance for an Aurora DB Cluster](../../../amazonrds/latest/aurorauserguide/concepts-aurorahighavailability.md#Aurora.Managing.FaultTolerance) in the _Amazon Aurora User Guide_.

This setting doesn't apply to RDS Custom DB instances.

Default: `1`

Valid Values: `0 - 15`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Indicates whether the DB instance is an internet-facing instance. If you specify true, AWS CloudFormation creates an instance with a publicly resolvable DNS name,
which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address.

The default behavior value depends on your VPC setup and the database subnet group. For more information, see the `PubliclyAccessible` parameter in
the [CreateDBInstance](../../../../reference/amazonrds/latest/apireference/api-createdbinstance.md)
in the _Amazon RDS API Reference_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicaMode`

The open mode of an Oracle read replica.
For more information, see [Working with Oracle Read Replicas for Amazon RDS](../../../amazonrds/latest/userguide/oracle-read-replicas.md)
in the _Amazon RDS User Guide_.

This setting is only supported in RDS for Oracle.

Default: `open-read-only`

Valid Values: `open-read-only` or `mounted`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreTime`

The date and time to restore from. This parameter applies to point-in-time recovery. For more
information, see [Restoring a DB instance to a specified time](../../../amazonrds/latest/userguide/user-pit.md)
in the in the _Amazon RDS User Guide_.

Constraints:

- Must be a time in Universal Coordinated Time (UTC) format.

- Must be before the latest restorable time for the DB instance.

- Can't be specified if the `UseLatestRestorableTime` parameter is enabled.

Example: `2009-09-07T23:45:00Z`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SourceDBClusterIdentifier`

The identifier of the Multi-AZ DB cluster that will act as the source for the read
replica. Each DB cluster can have up to 15 read replicas.

Constraints:

- Must be the identifier of an existing Multi-AZ DB cluster.

- Can't be specified if the `SourceDBInstanceIdentifier` parameter is
also specified.

- The specified DB cluster must have automatic backups enabled, that is, its
backup retention period must be greater than 0.

- The source DB cluster must be in the same AWS Region as the read replica.
Cross-Region replication isn't supported.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SourceDBInstanceAutomatedBackupsArn`

The Amazon Resource Name (ARN) of the replicated automated backups from which to restore, for example,
`arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE`.

This setting doesn't apply to RDS Custom.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SourceDBInstanceIdentifier`

If you want to create a read replica DB instance, specify the ID of the source DB
instance. Each DB instance can have a limited number of read replicas. For more
information, see [Working with Read\
Replicas](../../../amazonrds/latest/developerguide/user-readrepl.md) in the _Amazon RDS User Guide_.

For information about constraints that apply to DB instance identifiers, see
[Naming constraints in \
Amazon RDS](../../../amazonrds/latest/userguide/chap-limits.md#RDS_Limits.Constraints) in the _Amazon RDS User Guide_.

The `SourceDBInstanceIdentifier` property determines whether a DB instance is
a read replica. If you remove the `SourceDBInstanceIdentifier` property from
your template and then update your stack, AWS CloudFormation promotes the
read replica to a standalone DB instance.

If you specify the `UseLatestRestorableTime` or `RestoreTime`
properties in conjunction with the `SourceDBInstanceIdentifier` property, RDS
restores the DB instance to the requested point in time, thereby creating a new DB
instance.

###### Important

- If you specify a source DB instance that uses VPC security groups, we
recommend that you specify the `VPCSecurityGroups` property. If
you don't specify the property, the read replica inherits the value of the
`VPCSecurityGroups` property from the source DB when you
create the replica. However, if you update the stack, AWS CloudFormation
reverts the replica's `VPCSecurityGroups` property to the default
value because it's not defined in the stack's template. This change might
cause unexpected issues.

- Read replicas don't support deletion policies. AWS CloudFormation ignores
any deletion policy that's associated with a read replica.

- If you specify `SourceDBInstanceIdentifier`, don't specify the
`DBSnapshotIdentifier` property. You can't create a read
replica from a snapshot.

- Don't set the `BackupRetentionPeriod`, `DBName`,
`MasterUsername`, `MasterUserPassword`, and
`PreferredBackupWindow` properties. The database attributes
are inherited from the source DB instance, and backups are disabled for read
replicas.

- If the source DB instance is in a different region than the read replica,
specify the source region in `SourceRegion`, and specify an ARN
for a valid DB instance in `SourceDBInstanceIdentifier`. For more
information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](../../../amazonrds/latest/userguide/user-tagging.md#USER_Tagging.ARN) in the
_Amazon RDS User Guide_.

- For DB instances in Amazon Aurora clusters, don't specify this property.
Amazon RDS automatically assigns writer and reader DB instances.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SourceDbiResourceId`

The resource ID of the source DB instance from which to restore.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SourceRegion`

The ID of the region that contains the source DB instance for the read replica.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageEncrypted`

A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.

If you specify the `KmsKeyId` property, then you must enable encryption.

If you specify the `SourceDBInstanceIdentifier` or
`SourceDbiResourceId` property, don't specify this property. The value is
inherited from the source DB instance, and if the DB instance is encrypted, the specified
`KmsKeyId` property is used.

If you specify the `SourceDBInstanceAutomatedBackupsArn` property, don't
specify this property. The value is inherited from the source DB instance automated backup.

If you specify `DBSnapshotIdentifier` property, don't specify this property.
The value is inherited from the snapshot.

**Amazon Aurora**

Not applicable. The encryption for DB instances is managed by
the DB cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageThroughput`

Specifies the storage throughput value, in mebibyte per second (MiBps), for the DB instance. This setting applies only to the `gp3` storage type.

This setting doesn't apply to RDS Custom or Amazon Aurora.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageType`

The storage type to associate with the DB instance.

If you specify `io1`, `io2`, or `gp3`, you must also include a value for the
`Iops` parameter.

This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster.

Valid Values: `gp2 | gp3 | io1 | io2 | standard`

Default: `io1`, if the `Iops` parameter is specified. Otherwise,
`gp3`.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

Tags to assign to the DB instance.

_Required_: No

_Type_: Array of [Tag](aws-properties-rds-dbinstance-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timezone`

The time zone of the DB instance.
The time zone parameter is currently supported only by [RDS for Db2](../../../amazonrds/latest/userguide/db2-time-zone.md) and
[RDS for SQL Server](../../../amazonrds/latest/userguide/chap-sqlserver.md#SQLServer.Concepts.General.TimeZone).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UseDefaultProcessorFeatures`

Specifies whether the DB instance class of the DB instance uses its default
processor features.

This setting doesn't apply to RDS Custom DB instances.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLatestRestorableTime`

Specifies whether the DB instance is restored from the latest backup time. By default, the DB instance
isn't restored from the latest backup time. This parameter applies to point-in-time recovery. For more
information, see [Restoring a DB instance to a specified time](../../../amazonrds/latest/userguide/user-pit.md)
in the in the _Amazon RDS User Guide_.

Constraints:

- Can't be specified if the `RestoreTime` parameter is provided.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VPCSecurityGroups`

A list of the VPC security group IDs to assign to the DB instance. The list can
include both the physical IDs of existing VPC security groups and references to [AWS::EC2::SecurityGroup](../userguide/aws-properties-ec2-security-group.md) resources created in the template.

If you plan to update the resource, don't specify VPC security groups in a shared VPC.

If you set `VPCSecurityGroups`, you must not set [`DBSecurityGroups`](../userguide/aws-properties-rds-database-instance.md#cfn-rds-dbinstance-dbsecuritygroups), and vice versa.

###### Important

You can migrate a DB instance in your stack from an RDS DB security group to a VPC
security group, but keep the following in mind:

- You can't revert to using an RDS security group after you establish a VPC
security group membership.

- When you migrate your DB instance to VPC security groups, if your stack
update rolls back because the DB instance update fails or because an update
fails in another AWS CloudFormation resource, the rollback fails because it
can't revert to an RDS security group.

- To use the properties that are available when you use a VPC security
group, you must recreate the DB instance. If you don't, AWS CloudFormation
submits only the property values that are listed in the [`DBSecurityGroups`](../userguide/aws-properties-rds-database-instance.md#cfn-rds-dbinstance-dbsecuritygroups) property.

To avoid this situation, migrate your DB instance to using VPC security groups only
when that is the only change in your stack template.

**Amazon Aurora**

Not applicable. The associated list of EC2 VPC security groups is managed by
the DB cluster. If specified, the setting must match the DB cluster setting.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DB instance name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutomaticRestartTime`

The time when a stopped DB instance is restarted automatically.

`CertificateDetails.CAIdentifier`

The CA identifier of the CA certificate used for the DB instance's server certificate.

`CertificateDetails.ValidTill`

The expiration date of the DB instance’s server certificate.

`DBInstanceArn`

The Amazon Resource Name (ARN) for the DB instance.

`DBInstanceStatus`

The current state of this DB instance.

`DbiResourceId`

The AWS Region-unique, immutable identifier for the DB instance.
This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB instance is accessed.

`DBSystemId`

The Oracle system ID (Oracle SID) for a container database (CDB).
The Oracle SID is also the name of the CDB.

This setting is valid for RDS Custom only.

`Endpoint.Address`

The connection endpoint for the database. For example:
`mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`

For Aurora Serverless DB clusters, the connection endpoint only applies to the DB cluster.

`Endpoint.HostedZoneId`

The ID that Amazon Route 53 assigns when you create a hosted zone.

`Endpoint.Port`

The port number on which the database accepts connections. For example:
`3306`

`InstanceCreateTime`

The date and time when the DB instance was created.

`IsStorageConfigUpgradeAvailable`

Indicates whether an upgrade is recommended for the storage file system configuration on the DB instance.

`LatestRestorableTime`

The latest time to which a database in this DB instance can be restored with point-in-time restore.

`ListenerEndpoint.Address`

The connection endpoint for the database. For example: `mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`

`ListenerEndpoint.HostedZoneId`

The ID of the hosted zone in that is associated with the database endpoint.

`ListenerEndpoint.Port`

The port number on which the database accepts connections.

This setting doesn't apply to Aurora DB instances. The port number is managed by the cluster.

Valid Values: `1150-65535`

Default:

- RDS for Db2 - `50000`

- RDS for MariaDB - `3306`

- RDS for Microsoft SQL Server - `1433`

- RDS for MySQL - `3306`

- RDS for Oracle - `1521`

- RDS for PostgreSQL - `5432`

Constraints:

- For RDS for Microsoft SQL Server, the value can't be `1234`, `1434`,
`3260`, `3343`, `3389`, `47001`, or
`49152-49156`.

`MasterUserSecret.SecretArn`

The Amazon Resource Name (ARN) of the secret. This parameter is a return value that you can retrieve using the `Fn::GetAtt`
intrinsic function. For more information, see [Return values](../userguide/aws-resource-rds-dbinstance.md#aws-resource-rds-dbinstance-return-values).

`PercentProgress`

The progress of the storage optimization operation as a percentage.

`ReadReplicaDBClusterIdentifiers`

The identifiers of Aurora DB clusters to which the RDS DB instance is replicated as a read replica.

`ReadReplicaDBInstanceIdentifiers`

The identifiers of the read replicas associated with this DB instance.

`ResumeFullAutomationModeTime`

The number of minutes to pause the automation. When the time period ends, RDS Custom resumes full automation.
The minimum value is 60 (default). The maximum value is 1,440.

`SecondaryAvailabilityZone`

If present, specifies the name of the secondary Availability Zone for a DB instance with multi-AZ support.

`StatusInfos`

The status of a read replica. If the DB instance isn't a read replica, the value is blank.

## Remarks

_DB instance creation modes_

This section outlines the modes for creating RDS DB instances—empty
database, read replica, restore from snapshot, and point-in-time recovery. In the
console, API, or CLI, each mode corresponds to a distinct API or command. In
CloudFormation, however, the mode is determined by the following parameter
combinations.

- **Empty database** \- Creates a new, empty DB
instance.

To create an empty database, _don't_ specify any of the
following properties:

- `RestoreTime`

- `UseLatestRestorableTime`

- `SourceDBInstanceAutomatedBackupsArn`

- `SourceDbiResourceId`

- `SourceDBInstanceIdentifier`

- `SourceDBClusterIdentifier`

- `DBSnapshotIdentifier`

- **Read replica** \- Creates a replica of an
existing DB instance to replicate data asynchronously.

To create a read replica, specify the `SourceDBInstanceIdentifier`
property or the `SourceDBClusterIdentifier` property
_without_ the following properties:

- `RestoreTime`

- `UseLatestRestorableTime`

- `SourceDBInstanceAutomatedBackupsArn`

- `SourceDbiResourceId`

- `DBSnapshotIdentifier`

- **Restore from snapshot** \- Restores the DB
instance from the specified snapshot. Cannot be used with a point-in-time
recovery.

To restore a DB instance from a snapshot, specify the
`DBSnapshotIdentifier` property _without_ the
following properties:

- `RestoreTime`

- `UseLatestRestorableTime`

- `SourceDbiResourceId`

- `SourceDBInstanceAutomatedBackupsArn`

- `SourceDBInstanceIdentifier`

- `SourceDBClusterIdentifier`

- **Point-in-time recovery** \- Recovers the DB
instance to a specific point in time. Requires identifying the source and the
recovery time. Cannot be used with a snapshot.

To restore a DB instance to a specified time, include any of the following
properties:

- `SourceDBInstanceIdentifier`

- `SourceDbiResourceId`

- `SourceDBInstanceAutomatedBackupsArn`

with _either_ the `RestoreTime` property or the
`UseLatestRestorableTime` property.

## Examples

- [Creating a DB instance with Enhanced Monitoring enabled](#aws-resource-rds-dbinstance--examples--Creating_a_DB_instance_with_Enhanced_Monitoring_enabled)

- [Creating a cross-region encrypted read replica](#aws-resource-rds-dbinstance--examples--Creating_a_cross-region_encrypted_read_replica)

- [Setting a AWS region for the backup replication of a DB instance](#aws-resource-rds-dbinstance--examples--Setting_a_region_for_the_backup_replication_of_a_DB_instance)

- [Creating an Amazon RDS Custom DB instance](#aws-resource-rds-dbinstance--examples--Creating_an_Amazon_RDS_Custom_DB_instance)

- [Deploying RDS Custom for Oracle with single and multiple Availability Zones](#aws-resource-rds-dbinstance--examples--Deploying_RDS_Custom_for_Oracle_with_single_and_multiple_Availability_Zones)

- [Creating a Secrets Manager secret for a master password](#aws-resource-rds-dbinstance--examples--Creating_a_Secrets_Manager_secret_for_a_master_password)

- [Creating a DB instance with self-managed Active Directory domain](#aws-resource-rds-dbinstance--examples--Creating_a_DB_instance_with_self-managed_Active_Directory_domain)

### Creating a DB instance with Enhanced Monitoring enabled

The following example creates an Amazon RDS MySQL DB instance with Enhanced
Monitoring enabled. The IAM role for Enhanced Monitoring specified in
`MonitoringRoleArn` must exist before you run this example. For
more information, see [Enhanced\
Monitoring](../../../amazonrds/latest/userguide/user-monitoring-os.md) in the _Amazon RDS User Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation Sample Template for creating an Amazon RDS DB instance: Sample template showing how to create a DB instance with Enhanced Monitoring enabled. **WARNING** This template creates an RDS DB instance. You will be billed for the AWS resources used if you create a stack from this template.",
    "Parameters": {
        "DBInstanceID": {
            "Default": "mydbinstance",
            "Description": "My database instance",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "63",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "Must begin with a letter and must not end with a hyphen or contain two consecutive hyphens."
        },
        "DBName": {
            "Default": "mydb",
            "Description": "My database",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "64",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "Must begin with a letter and contain only alphanumeric characters."
        },
        "DBInstanceClass": {
            "Default": "db.m5.large",
            "Description": "DB instance class",
            "Type": "String",
            "ConstraintDescription": "Must select a valid DB instance type."
        },
        "DBAllocatedStorage": {
            "Default": "50",
            "Description": "The size of the database (GiB)",
            "Type": "Number",
            "MinValue": "20",
            "MaxValue": "65536",
            "ConstraintDescription": "must be between 20 and 65536 GiB."
        },
        "DBUsername": {
            "NoEcho": "true",
            "Description": "Username for MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "DBPassword": {
            "NoEcho": "true",
            "Description": "Password MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "MyDB": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBInstanceIdentifier": {
                    "Ref": "DBInstanceID"
                },
                "DBName": {
                    "Ref": "DBName"
                },
                "DBInstanceClass": {
                    "Ref": "DBInstanceClass"
                },
                "AllocatedStorage": {
                    "Ref": "DBAllocatedStorage"
                },
                "Engine": "MySQL",
                "EngineVersion": "8.0.33",
                "MasterUsername": {
                    "Ref": "DBUsername"
                },
                "MasterUserPassword": {
                    "Ref": "DBPassword"
                },
                "MonitoringInterval": 60,
                "MonitoringRoleArn": "arn:aws:iam::1233456789012:role/rds-monitoring-role"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  Description": "AWS CloudFormation Sample Template for creating an Amazon RDS DB instance:
  Sample template showing how to create a DB instance with Enhanced Monitoring enabled.
  **WARNING** This template creates an RDS DB instance. You will be billed for the AWS
  resources used if you create a stack from this template.
Parameters:
  DBInstanceID:
    Default: mydbinstance
    Description: My database instance
    Type: String
    MinLength: '1'
    MaxLength: '63'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: >-
      Must begin with a letter and must not end with a hyphen or contain two
      consecutive hyphens.
  DBName:
    Default: mydb
    Description: My database
    Type: String
    MinLength: '1'
    MaxLength: '64'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: Must begin with a letter and contain only alphanumeric characters.
  DBInstanceClass:
    Default: db.m5.large
    Description: DB instance class
    Type: String
    ConstraintDescription: Must select a valid DB instance type.
  DBAllocatedStorage:
    Default: '50'
    Description: The size of the database (GiB)
    Type: Number
    MinValue: '20'
    MaxValue: '65536'
    ConstraintDescription: must be between 20 and 65536 GiB.
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
  MyDB:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBInstanceIdentifier: !Ref DBInstanceID
      DBName: !Ref DBName
      DBInstanceClass: !Ref DBInstanceClass
      AllocatedStorage: !Ref DBAllocatedStorage
      Engine: MySQL
      EngineVersion: "8.0.33"
      MasterUsername: !Ref DBUsername
      MasterUserPassword: !Ref DBPassword
      MonitoringInterval: 60
      MonitoringRoleArn: 'arn:aws:iam::123456789012:role/rds-monitoring-role'
```

### Creating a cross-region encrypted read replica

The following example creates an encrypted read replica from a cross-region
source DB instance.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "RDS Storage Encrypted",
    "Parameters": {
        "SourceDBInstanceIdentifier": {
            "Type": "String"
        },
        "DBInstanceType": {
            "Type": "String"
        },
        "SourceRegion": {
            "Type": "String"
        }
    },
    "Resources": {
        "MyKey": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Id": "key-default-1",
                    "Statement": [
                        {
                            "Sid": "Enable IAM User Permissions",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        "",
                                        [
                                            "arn:aws:iam::",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            ":root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "kms:*",
                            "Resource": "*"
                        }
                    ]
                }
            }
        },
        "MyDBSmall": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBInstanceClass": {
                    "Ref": "DBInstanceType"
                },
                "SourceDBInstanceIdentifier": {
                    "Ref": "SourceDBInstanceIdentifier"
                },
                "SourceRegion": {
                    "Ref": "SourceRegion"
                },
                "KmsKeyId": {
                    "Ref": "MyKey"
                }
            }
        }
    },
    "Outputs": {
        "InstanceId": {
            "Description": "InstanceId of the newly created RDS Instance",
            "Value": {
                "Ref": "MyDBSmall"
            }
        }
    }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: 2010-09-09
Description: RDS Storage Encrypted
Parameters:
  SourceDBInstanceIdentifier:
    Type: String
  DBInstanceType:
    Type: String
  SourceRegion:
    Type: String
Resources:
  MyKey:
    Type: AWS::KMS::Key
    Properties:
      KeyPolicy:
        Version: 2012-10-17
        Id: key-default-1
        Statement:
          - Sid: Enable IAM User Permissions
            Effect: Allow
            Principal:
              AWS: !Join
                - ""
                - - "arn:aws:iam::"
                  - !Ref "AWS::AccountId"
                  - ":root"
            Action: "kms:*"
            Resource: "*"
  MyDBSmall:
    Type: AWS::RDS::DBInstance
    Properties:
      DBInstanceClass: !Ref DBInstanceType
      SourceDBInstanceIdentifier: !Ref SourceDBInstanceIdentifier
      SourceRegion: !Ref SourceRegion
      KmsKeyId: !Ref MyKey
Outputs:
  InstanceId:
    Description: InstanceId of the newly created RDS Instance
    Value: !Ref MyDBSmall

```

### Setting a AWS region for the backup replication of a DB instance

The following example defines the destination region for automated backup replication of a DB instance. In this example
the DB instance backup will be replicated to `ap-northeast-1` region.

#### JSON

```json

{
    "Resources": {
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "AllocatedStorage": "100",
                "AutomaticBackupReplicationRegion": "ap-northeast-1",
                "BackupRetentionPeriod": 1,
                "DBInstanceClass": "db.t3.micro",
                "Engine": "postgres",
                "Iops": 1000,
                "MasterUserPassword": "masterpassword",
                "MasterUsername": "masteruser",
                "StorageType": "IO1"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      AllocatedStorage: '100'
      AutomaticBackupReplicationRegion: ap-northeast-1
      BackupRetentionPeriod: 1
      DBInstanceClass: db.t3.micro
      Engine: postgres
      Iops: 1000
      MasterUserPassword: masterpassword
      MasterUsername: masteruser
      StorageType: IO1
```

### Creating an Amazon RDS Custom DB instance

This example creates an Amazon RDS Custom DB instance. Amazon RDS Custom allows you as a database administrator to access and customize your database environment
and operating system. With this access, you can configure settings, install patches, and enable native features to meet the dependent application's requirements.
With RDS Custom, you can run your database workload using the AWS Management Console or AWS CLI.
For more information, see [Working with Amazon RDS Custom](../../../amazonrds/latest/userguide/rds-custom.md) in the _Amazon RDS User_
_Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "AllocatedStorage": {
            "Type": "String"
        },
        "CustomIAMInstanceProfile": {
            "Type": "String"
        },
        "DBInstanceIdentifier": {
            "Type": "String"
        },
        "DBInstanceClass": {
            "Type": "String"
        },
        "DBName": {
            "Type": "String"
        },
        "DBSubnetGroupName": {
            "Type": "String"
        },
        "Engine": {
            "Type": "String"
        },
        "EngineVersion": {
            "Type": "String"
        },
        "KMSKeyId": {
            "Type": "String"
        },
        "LicenseModel": {
            "Type": "String"
        },
        "MasterUsername": {
            "Type": "String"
        },
        "MasterUserPassword": {
            "Type": "String",
            "NoEcho": true
        }
    },
    "Resources": {
        "RDSDBInstance": {
            "Properties": {
                "AllocatedStorage": {
                    "Ref": "AllocatedStorage"
                },
                "AutoMinorVersionUpgrade": false,
                "CustomIAMInstanceProfile": {
                    "Ref": "CustomIAMInstanceProfile"
                },
                "DBInstanceClass": {
                    "Ref": "DBInstanceClass"
                },
                "DBSubnetGroupName": {
                    "Ref": "DBSubnetGroupName"
                },
                "Engine": {
                    "Ref": "Engine"
                },
                "EngineVersion": {
                    "Ref": "EngineVersion"
                },
                "LicenseModel": {
                    "Ref": "LicenseModel"
                },
                "KmsKeyId": {
                    "Ref": "KMSKeyId"
                },
                "MasterUsername": {
                    "Ref": "MasterUsername"
                },
                "MasterUserPassword": {
                    "Ref": "MasterUserPassword"
                }
            },
            "Type": "AWS::RDS::DBInstance"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  AllocatedStorage:
    Type: String
  CustomIAMInstanceProfile:
    Type: String
  DBInstanceIdentifier:
    Type: String
  DBInstanceClass:
    Type: String
  DBName:
    Type: String
  DBSubnetGroupName:
    Type: String
  Engine:
    Type: String
  EngineVersion:
    Type: String
  KMSKeyId:
    Type: String
  LicenseModel:
    Type: String
  MasterUsername:
    Type: String
  MasterUserPassword:
    Type: String
    NoEcho: true
Resources:
  RDSDBInstance:
    Properties:
      AllocatedStorage: !Ref AllocatedStorage
      AutoMinorVersionUpgrade: false
      CustomIAMInstanceProfile: !Ref CustomIAMInstanceProfile
      DBInstanceClass: !Ref DBInstanceClass
      DBSubnetGroupName: !Ref DBSubnetGroupName
      Engine: !Ref Engine
      EngineVersion: !Ref EngineVersion
      LicenseModel: !Ref LicenseModel
      KmsKeyId: !Ref KMSKeyId
      MasterUsername: !Ref MasterUsername
      MasterUserPassword: !Ref MasterUserPassword
    Type: 'AWS::RDS::DBInstance'

```

### Deploying RDS Custom for Oracle with single and multiple Availability Zones

These examples deploy RDS Custom for Oracle with single and multi-AZ support.
The prerequisites template creates the required infrastructure, and the second
template creates Custom Engine Versions (CEVs) as well as both single-AZ and
multi-AZ RDS Custom for Oracle instances. For more information, see [Deploying RDS Custom for Oracle with CloudFormation](../../../amazonrds/latest/userguide/custom-oracle-multiaz-deployment.md) in the
_Amazon RDS User Guide_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "CloudFormation Template for the Pre-reqs of RDS Custom for Oracle",

    "Resources": {
        "RDSVPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "EnableDnsSupport": "true",
                "EnableDnsHostnames": "true",
                "CidrBlock": "10.20.0.0/16",
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "RDSVPC"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "Subnet1": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "CidrBlock": "10.20.1.0/24",
                "AvailabilityZone": {
                    "Fn::Select": [
                        0,
                        {
                            "Fn::GetAZs": ""
                        }
                    ]
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "Subnet1"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "Subnet2": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "CidrBlock": "10.20.2.0/24",
                "AvailabilityZone": {
                    "Fn::Select": [
                        1,
                        {
                            "Fn::GetAZs": ""
                        }
                    ]
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "Subnet2"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "Subnet3": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "CidrBlock": "10.20.3.0/24",
                "AvailabilityZone": {
                    "Fn::Select": [
                        2,
                        {
                            "Fn::GetAZs": ""
                        }
                    ]
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "Subnet3"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "Subnet4": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "CidrBlock": "10.20.4.0/24",
                "AvailabilityZone": {
                    "Fn::Select": [
                        2,
                        {
                            "Fn::GetAZs": ""
                        }
                    ]
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "Subnet4"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "InternetGateway": {
            "Type": "AWS::EC2::InternetGateway",
            "Properties": {
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "InternetGateway"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "AttachGateway": {
            "Type": "AWS::EC2::VPCGatewayAttachment",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "InternetGatewayId": {
                    "Ref": "InternetGateway"
                }
            }
        },
        "RouteTablePvt": {
            "Type": "AWS::EC2::RouteTable",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "RouteTablePvt"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "RouteTablePub": {
            "Type": "AWS::EC2::RouteTable",
            "Properties": {
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Join": [
                                "-",
                                [
                                    {
                                        "Ref": "AWS::StackName"
                                    },
                                    "RouteTablePub"
                                ]
                            ]
                        }
                    }
                ]
            }
        },
        "RoutePub": {
            "Type": "AWS::EC2::Route",
            "DependsOn": "AttachGateway",
            "Properties": {
                "RouteTableId": {
                    "Ref": "RouteTablePub"
                },
                "DestinationCidrBlock": "0.0.0.0/0",
                "GatewayId": {
                    "Ref": "InternetGateway"
                }
            }
        },
        "Subnet1RouteTableAssociation": {
            "Type": "AWS::EC2::SubnetRouteTableAssociation",
            "Properties": {
                "SubnetId": {
                    "Ref": "Subnet1"
                },
                "RouteTableId": {
                    "Ref": "RouteTablePvt"
                }
            }
        },
        "Subnet2RouteTableAssociation": {
            "Type": "AWS::EC2::SubnetRouteTableAssociation",
            "Properties": {
                "SubnetId": {
                    "Ref": "Subnet2"
                },
                "RouteTableId": {
                    "Ref": "RouteTablePvt"
                }
            }
        },
        "Subnet3RouteTableAssociation": {
            "Type": "AWS::EC2::SubnetRouteTableAssociation",
            "Properties": {
                "SubnetId": {
                    "Ref": "Subnet3"
                },
                "RouteTableId": {
                    "Ref": "RouteTablePvt"
                }
            }
        },
        "Subnet4RouteTableAssociation": {
            "Type": "AWS::EC2::SubnetRouteTableAssociation",
            "Properties": {
                "SubnetId": {
                    "Ref": "Subnet4"
                },
                "RouteTableId": {
                    "Ref": "RouteTablePub"
                }
            }
        },
        "RDSSG": {
            "Type": "AWS::EC2::SecurityGroup",
            "Properties": {
                "GroupDescription": "Security Group for RDS Oracle Custom instances",
                "GroupName": "rds-custom-db-sg",
                "SecurityGroupIngress": [
                    {
                        "IpProtocol": "tcp",
                        "FromPort": "0",
                        "ToPort": "65535",
                        "CidrIp": "10.20.0.0/16",
                        "Description": "Allows all traffic from within the same VPC"
                    }
                ],
                "SecurityGroupEgress": [
                    {
                        "IpProtocol": -1,
                        "CidrIp": "0.0.0.0/0",
                        "Description": "Allows outbound communication"
                    }
                ],
                "VpcId": {
                    "Ref": "RDSVPC"
                }
            }
        },
        "VPCESG": {
            "Type": "AWS::EC2::SecurityGroup",
            "Properties": {
                "GroupDescription": "Security Group for VPC Endpoints",
                "GroupName": "rds-custom-vpce-sg",
                "SecurityGroupIngress": [
                    {
                        "IpProtocol": "tcp",
                        "FromPort": "443",
                        "ToPort": "443",
                        "CidrIp": "10.20.0.0/16",
                        "Description": "Allows SSM access from the RDS Oracle DB instance"
                    }
                ],
                "SecurityGroupEgress": [
                    {
                        "IpProtocol": -1,
                        "CidrIp": "0.0.0.0/0",
                        "Description": "Allows outbound communication"
                    }
                ],
                "VpcId": {
                    "Ref": "RDSVPC"
                }
            }
        },
        "DBSubnetGroup": {
            "Type": "AWS::RDS::DBSubnetGroup",
            "Properties": {
                "DBSubnetGroupName": "rds-custom-private",
                "DBSubnetGroupDescription": "RDS Custom Private Network",
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ]
            }
        },
        "vpceS3": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "RouteTableIds": [
                    {
                        "Ref": "RouteTablePvt"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.s3"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                }
            }
        },
        "vpceEC2": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.ec2"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceEC2Messages": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.ec2messages"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceMonitoring": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.monitoring"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceSSM": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.ssm"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceSSMMessages": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.ssmmessages"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceLogs": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.logs"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceEvents": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.events"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceSecretsManager": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.secretsmanager"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "vpceSQS": {
            "Type": "AWS::EC2::VPCEndpoint",
            "Properties": {
                "VpcEndpointType": "Interface",
                "PrivateDnsEnabled": true,
                "SubnetIds": [
                    {
                        "Ref": "Subnet1"
                    },
                    {
                        "Ref": "Subnet2"
                    },
                    {
                        "Ref": "Subnet3"
                    }
                ],
                "ServiceName": {
                    "Fn::Sub": "com.amazonaws.${AWS::Region}.sqs"
                },
                "VpcId": {
                    "Ref": "RDSVPC"
                },
                "SecurityGroupIds": [
                    {
                        "Ref": "VPCESG"
                    }
                ]
            }
        },
        "KMSKey": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "Description": "Symmetric key for the CEV",
                "Enabled": true,
                "EnableKeyRotation": false,
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Id": "key-default-1",
                    "Statement": {
                        "Sid": "Enable IAM policies",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": {
                                "Fn::Join": [
                                    ":",
                                    [
                                        "arn",
                                        {
                                            "Ref": "AWS::Partition"
                                        },
                                        "iam:",
                                        {
                                            "Ref": "AWS::AccountId"
                                        },
                                        "root"
                                    ]
                                ]
                            }
                        },
                        "Action": "kms:*",
                        "Resource": "*"
                    }
                },
                "KeySpec": "SYMMETRIC_DEFAULT",
                "KeyUsage": "ENCRYPT_DECRYPT",
                "MultiRegion": false
            }
        },
        "KMSKeyAlias": {
            "Type": "AWS::KMS::Alias",
            "Properties": {
                "AliasName": "alias/rds-custom",
                "TargetKeyId": {
                    "Ref": "KMSKey"
                }
            }
        },
        "RDSCustomInstanceServiceRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "RoleName": {
                    "Fn::Sub": "AWSRDSCustomInstanceRole-${AWS::Region}"
                },
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Action": "sts:AssumeRole",
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "ec2.amazonaws.com"
                            }
                        }
                    ]
                },
                "Path": "/",
                "Policies": [
                    {
                        "PolicyName": "AWSRDSCustomEc2InstanceRolePolicy",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Sid": "1",
                                    "Effect": "Allow",
                                    "Action": [
                                        "ssm:DescribeAssociation",
                                        "ssm:GetDeployablePatchSnapshotForInstance",
                                        "ssm:GetDocument",
                                        "ssm:DescribeDocument",
                                        "ssm:GetManifest",
                                        "ssm:GetParameter",
                                        "ssm:GetParameters",
                                        "ssm:ListAssociations",
                                        "ssm:ListInstanceAssociations",
                                        "ssm:PutInventory",
                                        "ssm:PutComplianceItems",
                                        "ssm:PutConfigurePackageResult",
                                        "ssm:UpdateAssociationStatus",
                                        "ssm:UpdateInstanceAssociationStatus",
                                        "ssm:UpdateInstanceInformation",
                                        "ssm:GetConnectionStatus",
                                        "ssm:DescribeInstanceInformation",
                                        "ssmmessages:CreateControlChannel",
                                        "ssmmessages:CreateDataChannel",
                                        "ssmmessages:OpenControlChannel",
                                        "ssmmessages:OpenDataChannel"
                                    ],
                                    "Resource": [
                                        "*"
                                    ]
                                },
                                {
                                    "Sid": "2",
                                    "Effect": "Allow",
                                    "Action": [
                                        "ec2messages:AcknowledgeMessage",
                                        "ec2messages:DeleteMessage",
                                        "ec2messages:FailMessage",
                                        "ec2messages:GetEndpoint",
                                        "ec2messages:GetMessages",
                                        "ec2messages:SendReply"
                                    ],
                                    "Resource": [
                                        "*"
                                    ]
                                },
                                {
                                    "Sid": "3",
                                    "Effect": "Allow",
                                    "Action": [
                                        "logs:PutRetentionPolicy",
                                        "logs:PutLogEvents",
                                        "logs:DescribeLogStreams",
                                        "logs:DescribeLogGroups",
                                        "logs:CreateLogStream",
                                        "logs:CreateLogGroup"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:logs:${AWS::Region}:${AWS::AccountId}:log-group:rds-custom-instance*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "4",
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:putObject",
                                        "s3:getObject",
                                        "s3:getObjectVersion"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:s3:::do-not-delete-rds-custom-*/*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "5",
                                    "Effect": "Allow",
                                    "Action": [
                                        "cloudwatch:PutMetricData"
                                    ],
                                    "Resource": [
                                        "*"
                                    ],
                                    "Condition": {
                                        "StringEquals": {
                                            "cloudwatch:namespace": [
                                                "RDSCustomForOracle/Agent"
                                            ]
                                        }
                                    }
                                },
                                {
                                    "Sid": "6",
                                    "Effect": "Allow",
                                    "Action": [
                                        "events:PutEvents"
                                    ],
                                    "Resource": [
                                        "*"
                                    ]
                                },
                                {
                                    "Sid": "7",
                                    "Effect": "Allow",
                                    "Action": [
                                        "secretsmanager:GetSecretValue",
                                        "secretsmanager:DescribeSecret"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:do-not-delete-rds-custom-*"
                                        },
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:rds-custom!oracle-do-not-delete-*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "8",
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:ListBucketVersions"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:s3:::do-not-delete-rds-custom-*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "9",
                                    "Effect": "Allow",
                                    "Action": "ec2:CreateSnapshots",
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:ec2:*:*:instance/*"
                                        },
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:ec2:*:*:volume/*"
                                        }
                                    ],
                                    "Condition": {
                                        "StringEquals": {
                                            "ec2:ResourceTag/AWSRDSCustom": "custom-oracle"
                                        }
                                    }
                                },
                                {
                                    "Sid": "10",
                                    "Effect": "Allow",
                                    "Action": "ec2:CreateSnapshots",
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:ec2:*::snapshot/*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "11",
                                    "Effect": "Allow",
                                    "Action": [
                                        "kms:Decrypt",
                                        "kms:GenerateDataKey"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::GetAtt": [
                                                "KMSKey",
                                                "Arn"
                                            ]
                                        }
                                    ]
                                },
                                {
                                    "Sid": "12",
                                    "Effect": "Allow",
                                    "Action": "ec2:CreateTags",
                                    "Resource": "*",
                                    "Condition": {
                                        "StringLike": {
                                            "ec2:CreateAction": [
                                                "CreateSnapshots"
                                            ]
                                        }
                                    }
                                },
                                {
                                    "Sid": "13",
                                    "Effect": "Allow",
                                    "Action": [
                                        "sqs:SendMessage",
                                        "sqs:ReceiveMessage",
                                        "sqs:DeleteMessage",
                                        "sqs:GetQueueUrl"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:sqs:${AWS::Region}:${AWS::AccountId}:do-not-delete-rds-custom-*"
                                        }
                                    ],
                                    "Condition": {
                                        "StringLike": {
                                            "aws:ResourceTag/AWSRDSCustom": "custom-oracle"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "RDSCustomInstanceProfile": {
            "Type": "AWS::IAM::InstanceProfile",
            "Properties": {
                "InstanceProfileName": "AWSRDSCustomInstanceProfileForRdsCustomInstance",
                "Path": "/",
                "Roles": [
                    {
                        "Ref": "RDSCustomInstanceServiceRole"
                    }
                ]
            }
        },
        "PublicSubnetParameter": {
            "Type": "AWS::SSM::Parameter",
            "Properties": {
                "Description": "The subnet ID of the public subnet",
                "Name": "PublicSubnetParameter",
                "Type": "String",
                "Value": {
                    "Ref": "Subnet4"
                }
            }
        },
        "RDSSGParameter": {
            "Type": "AWS::SSM::Parameter",
            "Properties": {
                "Description": "The security group ID of the RDSSG",
                "Name": "RDSSGParameter",
                "Type": "String",
                "Value": {
                    "Fn::GetAtt": [
                        "RDSSG",
                        "GroupId"
                    ]
                }
            }
        },
        "RDSKMSParameter": {
            "Type": "AWS::SSM::Parameter",
            "Properties": {
                "Description": "KMS Key",
                "Name": "KMSParameter",
                "Type": "String",
                "Value": {
                    "Fn::GetAtt": [
                        "KMSKey",
                        "KeyId"
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

# AWSTemplateFormatVersion: 2010-09-09
# RDS Custom for AWS RDS Oracle
# 06/15/2025  created by Sharath Chandra Kampili

AWSTemplateFormatVersion: '2010-09-09'
Description: CloudFormation Template for the pre-reqs of RDS Custom for Oracle

Resources:
  RDSVPC:
    Type: AWS::EC2::VPC
    Properties:
      EnableDnsSupport: 'true'
      EnableDnsHostnames: 'true'
      CidrBlock: 10.20.0.0/16
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - RDSVPC
  Subnet1:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref RDSVPC
      CidrBlock: 10.20.1.0/24
      AvailabilityZone: !Select [0, !GetAZs ]
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - Subnet1
  Subnet2:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref RDSVPC
      CidrBlock: 10.20.2.0/24
      AvailabilityZone: !Select [1, !GetAZs ]
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - Subnet2
  Subnet3:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref RDSVPC
      CidrBlock: 10.20.3.0/24
      AvailabilityZone: !Select [2, !GetAZs ]
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - Subnet3
  Subnet4:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref RDSVPC
      CidrBlock: 10.20.4.0/24
      AvailabilityZone: !Select [2, !GetAZs ]
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - Subnet4

  InternetGateway:
    Type: AWS::EC2::InternetGateway
    Properties:
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - InternetGateway
  AttachGateway:
    Type: AWS::EC2::VPCGatewayAttachment
    Properties:
      VpcId: !Ref RDSVPC
      InternetGatewayId: !Ref InternetGateway

  RouteTablePvt:
    Type: AWS::EC2::RouteTable
    Properties:
      VpcId: !Ref RDSVPC
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - RouteTablePvt
  RouteTablePub:
    Type: AWS::EC2::RouteTable
    Properties:
      VpcId: !Ref RDSVPC
      Tags:
      - Key: Name
        Value:
          Fn::Join:
          - "-"
          - - !Ref AWS::StackName
            - RouteTablePub
  RoutePub:
    Type: AWS::EC2::Route
    DependsOn: AttachGateway
    Properties:
      RouteTableId: !Ref RouteTablePub
      DestinationCidrBlock: 0.0.0.0/0
      GatewayId: !Ref InternetGateway

  Subnet1RouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId: !Ref Subnet1
      RouteTableId: !Ref RouteTablePvt
  Subnet2RouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId: !Ref Subnet2
      RouteTableId: !Ref RouteTablePvt
  Subnet3RouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId: !Ref Subnet3
      RouteTableId: !Ref RouteTablePvt
  Subnet4RouteTableAssociation:
    Type: AWS::EC2::SubnetRouteTableAssociation
    Properties:
      SubnetId: !Ref Subnet4
      RouteTableId: !Ref RouteTablePub

  RDSSG:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Security Group for RDS Oracle Custom instances
      GroupName: rds-custom-db-sg
      SecurityGroupIngress:
        - IpProtocol: tcp
          FromPort: '0'
          ToPort: '65535'
          CidrIp: 10.20.0.0/16
          Description: Allows all traffic from within the same VPC
      SecurityGroupEgress:
        - IpProtocol: -1
          CidrIp: 0.0.0.0/0
          Description: Allows outbound communication
      VpcId: !Ref RDSVPC
  VPCESG:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: Security Group for VPC Endpoints
      GroupName: rds-custom-vpce-sg
      SecurityGroupIngress:
        - IpProtocol: tcp
          FromPort: '443'
          ToPort: '443'
          CidrIp: 10.20.0.0/16
          Description: Allows SSM access from the RDS Oracle DB instance
      SecurityGroupEgress:
        - IpProtocol: -1
          CidrIp: 0.0.0.0/0
          Description: Allows outbound communication
      VpcId: !Ref RDSVPC

  DBSubnetGroup:
    Type: AWS::RDS::DBSubnetGroup
    Properties:
      DBSubnetGroupName: rds-custom-private
      DBSubnetGroupDescription: RDS Custom Private Network
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
  vpceS3:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      RouteTableIds:
      - Ref: RouteTablePvt
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.s3
      VpcId: !Ref RDSVPC
  vpceEC2:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.ec2
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceEC2Messages:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.ec2messages
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceMonitoring:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.monitoring
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceSSM:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.ssm
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceSSMMessages:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.ssmmessages
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceLogs:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.logs
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceEvents:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.events
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceSecretsManager:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.secretsmanager
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  vpceSQS:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      VpcEndpointType: Interface
      PrivateDnsEnabled: true
      SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
      - !Ref Subnet3
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.sqs
      VpcId: !Ref RDSVPC
      SecurityGroupIds:
      - !Ref VPCESG
  KMSKey:
    Type: "AWS::KMS::Key"
    Properties:
      Description: "Symmetric key for the CEV"
      Enabled: true
      EnableKeyRotation: false
      KeyPolicy:
        Version: "2012-10-17"
        Id: "key-default-1"
        Statement:
          Sid: "Enable IAM policies"
          Effect: "Allow"
          Principal:
            AWS: !Join [ ':', [ 'arn', !Ref AWS::Partition, 'iam:', !Ref AWS::AccountId, 'root' ] ]
          Action: kms:*
          Resource: '*'
      KeySpec: "SYMMETRIC_DEFAULT"
      KeyUsage: "ENCRYPT_DECRYPT"
      MultiRegion: false
  KMSKeyAlias:
    Type: AWS::KMS::Alias
    Properties:
      AliasName: 'alias/rds-custom'
      TargetKeyId: !Ref KMSKey
  RDSCustomInstanceServiceRole:
    Type: AWS::IAM::Role
    Properties:
      RoleName:
        Fn::Sub: AWSRDSCustomInstanceRole-${AWS::Region}
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Action: sts:AssumeRole
            Effect: Allow
            Principal:
              Service: ec2.amazonaws.com
      Path: /
      Policies:
        - PolicyName: AWSRDSCustomEc2InstanceRolePolicy
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Sid: '1'
                Effect: Allow
                Action:
                  - ssm:DescribeAssociation
                  - ssm:GetDeployablePatchSnapshotForInstance
                  - ssm:GetDocument
                  - ssm:DescribeDocument
                  - ssm:GetManifest
                  - ssm:GetParameter
                  - ssm:GetParameters
                  - ssm:ListAssociations
                  - ssm:ListInstanceAssociations
                  - ssm:PutInventory
                  - ssm:PutComplianceItems
                  - ssm:PutConfigurePackageResult
                  - ssm:UpdateAssociationStatus
                  - ssm:UpdateInstanceAssociationStatus
                  - ssm:UpdateInstanceInformation
                  - ssm:GetConnectionStatus
                  - ssm:DescribeInstanceInformation
                  - ssmmessages:CreateControlChannel
                  - ssmmessages:CreateDataChannel
                  - ssmmessages:OpenControlChannel
                  - ssmmessages:OpenDataChannel
                Resource:
                  - '*'
              - Sid: '2'
                Effect: Allow
                Action:
                  - ec2messages:AcknowledgeMessage
                  - ec2messages:DeleteMessage
                  - ec2messages:FailMessage
                  - ec2messages:GetEndpoint
                  - ec2messages:GetMessages
                  - ec2messages:SendReply
                Resource:
                  - '*'
              - Sid: '3'
                Effect: Allow
                Action:
                  - logs:PutRetentionPolicy
                  - logs:PutLogEvents
                  - logs:DescribeLogStreams
                  - logs:DescribeLogGroups
                  - logs:CreateLogStream
                  - logs:CreateLogGroup
                Resource:
                  - Fn::Sub: >-
                      arn:${AWS::Partition}:logs:${AWS::Region}:${AWS::AccountId}:log-group:rds-custom-instance*
              - Sid: '4'
                Effect: Allow
                Action:
                  - s3:putObject
                  - s3:getObject
                  - s3:getObjectVersion
                Resource:
                  - Fn::Sub: >-
                      arn:${AWS::Partition}:s3:::do-not-delete-rds-custom-*/*
              - Sid: '5'
                Effect: Allow
                Action:
                  - cloudwatch:PutMetricData
                Resource:
                  - '*'
                Condition:
                  StringEquals:
                    cloudwatch:namespace:
                      - RDSCustomForOracle/Agent
              - Sid: '6'
                Effect: Allow
                Action:
                  - events:PutEvents
                Resource:
                  - '*'
              - Sid: '7'
                Effect: Allow
                Action:
                  - secretsmanager:GetSecretValue
                  - secretsmanager:DescribeSecret
                Resource:
                  - Fn::Sub: arn:${AWS::Partition}:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:do-not-delete-rds-custom-*
                  - Fn::Sub: arn:${AWS::Partition}:secretsmanager:${AWS::Region}:${AWS::AccountId}:secret:rds-custom!oracle-do-not-delete-*
              - Sid: '8'
                Effect: Allow
                Action:
                  - s3:ListBucketVersions
                Resource:
                  - Fn::Sub: >-
                      arn:${AWS::Partition}:s3:::do-not-delete-rds-custom-*
              - Sid: '9'
                Effect: Allow
                Action: ec2:CreateSnapshots
                Resource:
                  - Fn::Sub: arn:${AWS::Partition}:ec2:*:*:instance/*
                  - Fn::Sub: arn:${AWS::Partition}:ec2:*:*:volume/*
                Condition:
                  StringEquals:
                    ec2:ResourceTag/AWSRDSCustom: custom-oracle
              - Sid: '10'
                Effect: Allow
                Action: ec2:CreateSnapshots
                Resource:
                  - Fn::Sub: arn:${AWS::Partition}:ec2:*::snapshot/*
              - Sid: '11'
                Effect: Allow
                Action:
                  - kms:Decrypt
                  - kms:GenerateDataKey
                Resource:
                  - !GetAtt KMSKey.Arn
              - Sid: '12'
                Effect: Allow
                Action: ec2:CreateTags
                Resource: '*'
                Condition:
                  StringLike:
                    ec2:CreateAction:
                      - CreateSnapshots
              - Sid: '13'
                Effect: Allow
                Action:
                  - sqs:SendMessage
                  - sqs:ReceiveMessage
                  - sqs:DeleteMessage
                  - sqs:GetQueueUrl
                Resource:
                  - Fn::Sub: arn:${AWS::Partition}:sqs:${AWS::Region}:${AWS::AccountId}:do-not-delete-rds-custom-*
                Condition:
                  StringLike:
                    aws:ResourceTag/AWSRDSCustom: custom-oracle

  RDSCustomInstanceProfile:
    Type: AWS::IAM::InstanceProfile
    Properties:
      InstanceProfileName: 'AWSRDSCustomInstanceProfileForRdsCustomInstance'
      Path: /
      Roles:
        - !Ref RDSCustomInstanceServiceRole

  PublicSubnetParameter:
    Type: AWS::SSM::Parameter
    Properties:
      Description: The subnet ID of the public subnet
      Name: PublicSubnetParameter
      Type: String
      Value: !Ref Subnet4
  RDSSGParameter:
    Type: AWS::SSM::Parameter
    Properties:
      Description: The security group ID of the RDSSG
      Name: RDSSGParameter
      Type: String
      Value: !GetAtt RDSSG.GroupId
  RDSKMSParameter:
    Type: AWS::SSM::Parameter
    Properties:
      Description: KMS Key
      Name: KMSParameter
      Type: String
      Value: !GetAtt KMSKey.KeyId

```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "CloudFormation Template for the RDS Custom for Oracle",
    "Parameters": {
        "BucketName": {
            "Type": "String",
            "MinLength": 3,
            "MaxLength": 63,
            "AllowedPattern": "[a-zA-Z0-9.-]*",
            "Description": "S3 bucket name containing Oracle installation files"
        },
        "CEVS3Prefix": {
            "Type": "String",
            "Default": "oracle_cev"
        },
        "RDSSG": {
            "Type": "AWS::SSM::Parameter::Value<String>",
            "Default": "RDSSGParameter"
        },
        "KMSParameter": {
            "Type": "AWS::SSM::Parameter::Value<String>",
            "Default": "KMSParameter"
        },
        "CEVCreation": {
            "Description": "Please select the option \"yes\" for creating the CEV , if not leave it default for creating one for you",
            "Type": "String",
            "Default": "Yes"
        },
        "DBInstanceIdentifier": {
            "Description": "RDS Instance Name for Single AZ",
            "Type": "String",
            "Default": "rds-custom"
        }
    },
    "Conditions": {
        "ShouldAddCEV": {
            "Fn::Equals": [
                {
                    "Ref": "CEVCreation"
                },
                "Yes"
            ]
        }
    },
    "Resources": {
        "SSMExecutionRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "RoleName": {
                    "Fn::Join": [
                        "-",
                        [
                            "SSM-Execution-Role",
                            {
                                "Fn::Select": [
                                    4,
                                    {
                                        "Fn::Split": [
                                            "-",
                                            {
                                                "Fn::Select": [
                                                    2,
                                                    {
                                                        "Fn::Split": [
                                                            "/",
                                                            {
                                                                "Ref": "AWS::StackId"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "Fn::Sub": "${AWS::Region}"
                            }
                        ]
                    ]
                },
                "ManagedPolicyArns": [
                    {
                        "Fn::Sub": "arn:${AWS::Partition}:iam::aws:policy/service-role/AmazonSSMAutomationRole"
                    },
                    "arn:aws:iam::aws:policy/AmazonRDSFullAccess",
                    "arn:aws:iam::aws:policy/AmazonEC2FullAccess",
                    "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
                ],
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "ec2.amazonaws.com",
                                    "ssm.amazonaws.com"
                                ]
                            },
                            "Action": "sts:AssumeRole"
                        }
                    ]
                }
            }
        },
        "EC2Role": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "Path": "/",
                "RoleName": {
                    "Fn::Join": [
                        "-",
                        [
                            "EC2-Role",
                            {
                                "Fn::Select": [
                                    4,
                                    {
                                        "Fn::Split": [
                                            "-",
                                            {
                                                "Fn::Select": [
                                                    2,
                                                    {
                                                        "Fn::Split": [
                                                            "/",
                                                            {
                                                                "Ref": "AWS::StackId"
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                "Fn::Sub": "${AWS::Region}"
                            }
                        ]
                    ]
                },
                "AssumeRolePolicyDocument": "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"},\"Action\":\"sts:AssumeRole\"}]}",
                "MaxSessionDuration": 43200,
                "ManagedPolicyArns": [
                    "arn:aws:iam::aws:policy/AmazonRDSFullAccess",
                    "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
                    "arn:aws:iam::aws:policy/IAMReadOnlyAccess"
                ],
                "Description": "Allows Lambda functions to call AWS services on your behalf.",
                "Policies": [
                    {
                        "PolicyName": {
                            "Fn::Join": [
                                "-",
                                [
                                    "cev-policy-1",
                                    {
                                        "Fn::Select": [
                                            4,
                                            {
                                                "Fn::Split": [
                                                    "-",
                                                    {
                                                        "Fn::Select": [
                                                            2,
                                                            {
                                                                "Fn::Split": [
                                                                    "/",
                                                                    {
                                                                        "Ref": "AWS::StackId"
                                                                    }
                                                                ]
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "Fn::Sub": "${AWS::Region}"
                                    }
                                ]
                            ]
                        },
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Sid": "CreateKmsGrant",
                                    "Effect": "Allow",
                                    "Action": [
                                        "kms:DescribeKey",
                                        "kms:CreateGrant"
                                    ],
                                    "Resource": "*"
                                },
                                {
                                    "Sid": "CreateS3Bucket",
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:CreateBucket",
                                        "s3:PutBucketPolicy",
                                        "s3:PutBucketObjectLockConfiguration",
                                        "s3:PutBucketVersioning"
                                    ],
                                    "Resource": "arn:aws:s3:::do-not-delete-rds-custom-*"
                                },
                                {
                                    "Sid": "AccessToS3MediaBucket",
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:GetObjectAcl",
                                        "s3:GetObject",
                                        "s3:GetObjectTagging",
                                        "s3:ListBucket"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:s3:::${BucketName}"
                                        },
                                        {
                                           "Fn::Sub": "arn:${AWS::Partition}:s3:::${BucketName}/*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "PermissionForByom",
                                    "Effect": "Allow",
                                    "Action": [
                                        "mediaimport:CreateDatabaseBinarySnapshot"
                                    ],
                                    "Resource": "*"
                                },
                                {
                                    "Sid": "CreateServiceLinkedRole",
                                    "Effect": "Allow",
                                    "Action": [
                                        "iam:CreateServiceLinkedRole"
                                    ],
                                    "Resource": "arn:aws:iam::*:role/aws-service-role/custom.rds.amazonaws.com/AWSServiceRoleForRDSCustom"
                                }
                            ]
                        }
                    },
                    {
                        "PolicyName": {
                            "Fn::Join": [
                                "-",
                                [
                                    "cev-policy-2",
                                    {
                                        "Fn::Select": [
                                            4,
                                            {
                                                "Fn::Split": [
                                                    "-",
                                                    {
                                                        "Fn::Select": [
                                                            2,
                                                            {
                                                                "Fn::Split": [
                                                                    "/",
                                                                    {
                                                                        "Ref": "AWS::StackId"
                                                                    }
                                                                ]
                                                            }
                                                        ]
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        "Fn::Sub": "${AWS::Region}"
                                    }
                                ]
                            ]
                        },
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Sid": "ValidateIamRole",
                                    "Effect": "Allow",
                                    "Action": "iam:SimulatePrincipalPolicy",
                                    "Resource": "*"
                                },
                                {
                                    "Sid": "CreateCloudTrail",
                                    "Effect": "Allow",
                                    "Action": [
                                        "cloudtrail:CreateTrail",
                                        "cloudtrail:StartLogging"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::Sub": "arn:${AWS::Partition}:cloudtrail:*:*:trail/do-not-delete-rds-custom-*"
                                        }
                                    ]
                                },
                                {
                                    "Sid": "CreateCloudWatchLogStreamAndPushEvents",
                                    "Effect": "Allow",
                                    "Action": [
                                        "logs:CreateLogStream",
                                        "logs:CreateLogGroup",
                                        "logs:DescribeLogStreams",
                                        "logs:GetLogEvents",
                                        "logs:PutLogEvents"
                                    ],
                                    "Resource": "*"
                                },
                                {
                                    "Sid": "IAMPassRole",
                                    "Effect": "Allow",
                                    "Action": [
                                        "iam:PassRole"
                                    ],
                                    "Resource": "arn:aws:iam::*:role/AWSRDSCustomInstanceRole*"
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "CEVWaithandle": {
            "Condition": "ShouldAddCEV",
            "DependsOn": "CustomOracleJanPatch",
            "Type": "AWS::CloudFormation::WaitConditionHandle"
        },
        "WaitHandle": {
            "Type": "AWS::CloudFormation::WaitConditionHandle"
        },
        "WaitCondition": {
            "Type": "AWS::CloudFormation::WaitCondition",
            "Properties": {
                "Handle": {
                    "Fn::If": [
                        "ShouldAddCEV",
                        {
                            "Ref": "CEVWaithandle"
                        },
                        {
                            "Ref": "WaitHandle"
                        }
                    ]
                },
                "Timeout": "1",
                "Count": 0
            }
        },
        "RDSCustomSecret": {
            "Type": "AWS::SecretsManager::Secret",
            "Properties": {
                "Name": {
                    "Fn::Sub": "${AWS::StackName}-rds-custom-secret"
                },
                "Description": "RDS Custom Oracle Admin Password",
                "GenerateSecretString": {
                    "SecretStringTemplate": "{\"username\": \"oraadmin\"}",
                    "GenerateStringKey": "password",
                    "PasswordLength": 16,
                    "ExcludeCharacters": "\"@/\\"
                }
            }
        },
        "CustomOracleJanPatch": {
            "DependsOn": "SSMExecutionRole",
            "Condition": "ShouldAddCEV",
            "Type": "AWS::RDS::CustomDBEngineVersion",
            "Properties": {
                "DatabaseInstallationFilesS3BucketName": {
                    "Ref": "BucketName"
                },
                "DatabaseInstallationFilesS3Prefix": {
                    "Ref": "CEVS3Prefix"
                },
                "Description": "CEV for RDS custom for Oracle",
                "Engine": "custom-oracle-ee",
                "EngineVersion": "19.18-rdscustom-oracle-jan23psu-cev",
                "KMSKeyId": {
                    "Ref": "KMSParameter"
                },
                "Manifest": "{\"mediaImportTemplateVersion\":\"2020-08-14\",\"databaseInstallationFileNames\":[\"V982063-01.zip\"],\"opatchFileNames\":[\"p6880880_190000_Linux-x86-64.zip\"],\"psuRuPatchFileNames\": [\"p34765931_190000_Linux-x86-64.zip\",\"p34786990_190000_Linux-x86-64.zip\"],\"otherPatchFileNames\": [\"p35099667_190000_Linux-x86-64.zip\",\"p35099674_190000_Generic.zip\",\"p28730253_190000_Linux-x86-64.zip\",\"p29213893_1918000DBRU_Generic.zip\",\"p33125873_1918000DBRU_Linux-x86-64.zip\",\"p35012866_1918000DBRU_Linux-x86-64.zip\"]}"
            }
        },
        "CustomOracleAprCOHMTPatch": {
            "DependsOn": "SSMExecutionRole",
            "Condition": "ShouldAddCEV",
            "Type": "AWS::RDS::CustomDBEngineVersion",
            "Properties": {
                "DatabaseInstallationFilesS3BucketName": {
                    "Ref": "BucketName"
                },
                "DatabaseInstallationFilesS3Prefix": {
                    "Ref": "CEVS3Prefix"
                },
                "Description": "CEV for RDS custom for Oracle",
                "Engine": "custom-oracle-ee-cdb",
                "EngineVersion": "19.19-custom-oracle-home-mt",
                "KMSKeyId": {
                    "Ref": "KMSParameter"
                },
                "Manifest": "{\"mediaImportTemplateVersion\":\"2020-08-14\",\"databaseInstallationFileNames\":[\"V982063-01.zip\"],\"opatchFileNames\":[\"p6880880_190000_Linux-x86-64.zip\"],\"psuRuPatchFileNames\": [\"p35042068_190000_Linux-x86-64.zip\",\"p35050341_190000_Linux-x86-64.zip\"],\"otherPatchFileNames\": [\"p28730253_190000_Linux-x86-64.zip\",\"p29213893_1919000DBRU_Generic.zip\",\"p33125873_1919000DBRU_Linux-x86-64.zip\",\"p35220732_190000_Linux-x86-64.zip\",\"p35239280_190000_Generic.zip\"],\"installationParameters\": { \"unixGroupName\": \"dba\",\"unixGroupId\": 65789,\"unixUname\": \"oracle\",\"unixUid\": 6758,\"oracleHome\": \"/opt/oracle/product/19.19/dbhome_1\",\"oracleBase\": \"/opt/oracle\"}}"
            }
        },
        "CustomOracleAprPatch": {
            "DependsOn": "SSMExecutionRole",
            "Condition": "ShouldAddCEV",
            "Type": "AWS::RDS::CustomDBEngineVersion",
            "Properties": {
                "DatabaseInstallationFilesS3BucketName": {
                    "Ref": "BucketName"
                },
                "DatabaseInstallationFilesS3Prefix": {
                    "Ref": "CEVS3Prefix"
                },
                "Description": "CEV for RDS custom for Oracle",
                "Engine": "custom-oracle-ee",
                "EngineVersion": "19.19-rdscustom-oracle-apr23psu-cev",
                "KMSKeyId": {
                    "Ref": "KMSParameter"
                },
                "Manifest": "{\"mediaImportTemplateVersion\":\"2020-08-14\",\"databaseInstallationFileNames\":[\"V982063-01.zip\"],\"opatchFileNames\":[\"p6880880_190000_Linux-x86-64.zip\"],\"psuRuPatchFileNames\": [\"p35042068_190000_Linux-x86-64.zip\",\"p35050341_190000_Linux-x86-64.zip\"],\"otherPatchFileNames\": [\"p28730253_190000_Linux-x86-64.zip\",\"p29213893_1919000DBRU_Generic.zip\",\"p33125873_1919000DBRU_Linux-x86-64.zip\",\"p35220732_190000_Linux-x86-64.zip\",\"p35239280_190000_Generic.zip\"]}"
            }
        },
        "RDSDBInstance": {
            "DependsOn": "WaitCondition",
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "AllocatedStorage": 40,
                "DBInstanceIdentifier": {
                    "Fn::Sub": "${DBInstanceIdentifier}-saz"
                },
                "AutoMinorVersionUpgrade": false,
                "CustomIAMInstanceProfile": "AWSRDSCustomInstanceProfileForRdsCustomInstance",
                "DBInstanceClass": "db.m5.large",
                "DBSubnetGroupName": "rds-custom-private",
                "Engine": "custom-oracle-ee",
                "EngineVersion": "19.18-rdscustom-oracle-jan23psu-cev",
                "KmsKeyId": {
                    "Ref": "KMSParameter"
                },
                "MasterUsername": {
                    "Fn::Sub": "{{resolve:secretsmanager:${RDSCustomSecret}:SecretString:username}}"
                },
                "MasterUserPassword": {
                    "Fn::Sub": "{{resolve:secretsmanager:${RDSCustomSecret}:SecretString:password}}"
                },
                "StorageEncrypted": "true",
                "VPCSecurityGroups": [
                    {
                        "Ref": "RDSSG"
                    }
                ],
                "DBName": "ORCL",
                "PreferredBackupWindow": "03:00-04:00",
                "PreferredMaintenanceWindow": "Mon:05:00-Mon:06:00",
                "PubliclyAccessible": false
            }
        },
        "HARDSDBInstance": {
            "DependsOn": "WaitCondition",
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "AllocatedStorage": 40,
                "DBInstanceIdentifier": {
                    "Fn::Sub": "${DBInstanceIdentifier}-maz"
                },
                "AutoMinorVersionUpgrade": false,
                "CustomIAMInstanceProfile": "AWSRDSCustomInstanceProfileForRdsCustomInstance",
                "DBInstanceClass": "db.m5.large",
                "DBSubnetGroupName": "rds-custom-private",
                "Engine": "custom-oracle-ee",
                "EngineVersion": "19.19-rdscustom-oracle-apr23psu-cev",
                "KmsKeyId": {
                    "Ref": "KMSParameter"
                },
                "MasterUsername": {
                    "Fn::Sub": "{{resolve:secretsmanager:${RDSCustomSecret}:SecretString:username}}"
                },
                "MasterUserPassword": {
                    "Fn::Sub": "{{resolve:secretsmanager:${RDSCustomSecret}:SecretString:password}}"
                },
                "StorageEncrypted": "true",
                "MultiAZ": "true",
                "VPCSecurityGroups": [
                    {
                        "Ref": "RDSSG"
                    }
                ],
                "DBName": "ORCL",
                "PreferredBackupWindow": "03:00-04:00",
                "PreferredMaintenanceWindow": "Mon:05:00-Mon:06:00",
                "PubliclyAccessible": false
            }
        }
    }
}
```

#### YAML

```yaml

# AWSTemplateFormatVersion: 2010-09-09
# RDS Custom for AWS RDS Oracle
# 06/15/2025 created by Sharath Chandra Kampili

AWSTemplateFormatVersion: '2010-09-09'
Description: CloudFormation Template for the RDS Custom for Oracle

Parameters:
  BucketName:
    Type: "String"
    MinLength: 3
    MaxLength: 63
    AllowedPattern: ^[a-z0-9][a-z0-9.-]*[a-z0-9]$
    Description: "S3 bucket name (must be between 3 and 63 characters)"
  CEVS3Prefix:
    Type: "String"
    Default: "oracle_cev"
  RDSSG:
    Type: 'AWS::SSM::Parameter::Value<String>'
    Default: RDSSGParameter
  KMSParameter:
    Type: 'AWS::SSM::Parameter::Value<String>'
    Default: KMSParameter
  CEVCreation:
    Description: Please select the option "yes" for creating the CEV , if not leave it default for creating one for you
    Type: String
    Default: "Yes"
  DBInstanceIdentifier:
    Description: RDS Instance Name for Single AZ
    Type: String
    Default: "rds-custom"

Conditions:
  ShouldAddCEV: !Equals [ !Ref CEVCreation, "Yes" ]

Resources:

  SSMExecutionRole:
    Type : AWS::IAM::Role
    Properties:
      RoleName: !Join ['-', ['SSM-Execution-Role', !Select [4, !Split ['-', !Select [2, !Split ['/', !Ref AWS::StackId]]]],!Sub "${AWS::Region}"]]
      ManagedPolicyArns:
        - !Sub 'arn:${AWS::Partition}:iam::aws:policy/service-role/AmazonSSMAutomationRole'
        - arn:aws:iam::aws:policy/AmazonRDSFullAccess
        - arn:aws:iam::aws:policy/AmazonEC2FullAccess
        - arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
        - Effect: "Allow"
          Principal:
            Service:
            - "ec2.amazonaws.com"
            - "ssm.amazonaws.com"
          Action: "sts:AssumeRole"

  RDSSecret:
    Type: 'AWS::SecretsManager::Secret'
    Properties:
      Name: !Sub '${AWS::StackName}-rds-credentials'
      Description: 'RDS Custom Oracle database credentials'
      GenerateSecretString:
        SecretStringTemplate: '{"username": "oraadmin"}'
        GenerateStringKey: "password"
        PasswordLength: 16
        ExcludeCharacters: '"@/\\'

  EC2Role:
    Type: "AWS::IAM::Role"
    Properties:
      Path: "/"
      RoleName: !Join ['-', ['EC2-Role', !Select [4, !Split ['-', !Select [2, !Split ['/', !Ref AWS::StackId]]]],!Sub "${AWS::Region}"]]
      AssumeRolePolicyDocument: '{"Version": "2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"},"Action":"sts:AssumeRole"}]}'
      MaxSessionDuration: 43200
      ManagedPolicyArns:
        - "arn:aws:iam::aws:policy/AmazonRDSFullAccess"
        - "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
        - "arn:aws:iam::aws:policy/IAMReadOnlyAccess"
      Description: "Allows Lambda functions to call AWS services on your behalf."
      Policies:
        - PolicyName: !Join ['-', ['cev-policy-1', !Select [4, !Split ['-', !Select [2, !Split ['/', !Ref AWS::StackId]]]],!Sub "${AWS::Region}"]]
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
            - Sid: 'CreateKmsGrant'
              Effect: 'Allow'
              Action:
              - 'kms:DescribeKey'
              - 'kms:CreateGrant'
              Resource: '*'
            - Sid: 'CreateS3Bucket'
              Effect: 'Allow'
              Action:
              - 's3:CreateBucket'
              - 's3:PutBucketPolicy'
              - 's3:PutBucketObjectLockConfiguration'
              - 's3:PutBucketVersioning'
              Resource: "arn:aws:s3:::do-not-delete-rds-custom-*"
            - Sid: 'AccessToS3MediaBucket'
              Effect: 'Allow'
              Action:
              - 's3:GetObjectAcl'
              - 's3:GetObject'
              - 's3:GetObjectTagging'
              - 's3:ListBucket'
              Resource:
              - !Sub 'arn:${AWS::Partition}:s3:::${BucketName}'
              - !Sub 'arn:${AWS::Partition}:s3:::${BucketName}/*'
            - Sid: 'PermissionForByom'
              Effect: 'Allow'
              Action:
              - 'mediaimport:CreateDatabaseBinarySnapshot'
              Resource: "*"
            - Sid: 'CreateServiceLinkedRole'
              Effect: 'Allow'
              Action:
              - 'iam:CreateServiceLinkedRole'
              Resource: "arn:aws:iam::*:role/aws-service-role/custom.rds.amazonaws.com/AWSServiceRoleForRDSCustom"
        - PolicyName: !Join ['-', ['cev-policy-2', !Select [4, !Split ['-', !Select [2, !Split ['/', !Ref AWS::StackId]]]],!Sub "${AWS::Region}"]]
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
            - Sid: 'ValidateIamRole'
              Effect: 'Allow'
              Action: 'iam:SimulatePrincipalPolicy'
              Resource: '*'
            - Sid: 'CreateCloudTrail'
              Effect: 'Allow'
              Action:
                - 'cloudtrail:CreateTrail'
                - 'cloudtrail:StartLogging'
              Resource:
                  - Fn::Sub: >-
                      arn:${AWS::Partition}:cloudtrail:*:*:trail/do-not-delete-rds-custom-*
            - Sid: 'CreateCloudWatchLogStreamAndPushEvents'
              Effect: 'Allow'
              Action:
                - 'logs:CreateLogStream'
                - 'logs:CreateLogGroup'
                - 'logs:DescribeLogStreams'
                - 'logs:GetLogEvents'
                - 'logs:PutLogEvents'
              Resource: '*'
            - Sid: 'IAMPassRole'
              Effect: 'Allow'
              Action:
                - 'iam:PassRole'
              Resource: arn:aws:iam::*:role/AWSRDSCustomInstanceRole*
            - Sid: 'SecretsManagerAccess'
              Effect: 'Allow'
              Action:
                - 'secretsmanager:GetSecretValue'
                - 'secretsmanager:DescribeSecret'
              Resource: !Ref RDSSecret

  CEVWaithandle:
    Condition: ShouldAddCEV
    DependsOn: CustomOracleJanPatch
    Type: "AWS::CloudFormation::WaitConditionHandle"

  WaitHandle:
    Type: "AWS::CloudFormation::WaitConditionHandle"

  WaitCondition:
    Type: "AWS::CloudFormation::WaitCondition"
    Properties:
      Handle: !If [ShouldAddCEV, !Ref CEVWaithandle , !Ref WaitHandle ]
      Timeout: "1"
      Count: 0

  CustomOracleJanPatch:
    DependsOn: SSMExecutionRole
    Condition: ShouldAddCEV
    Type: AWS::RDS::CustomDBEngineVersion
    Properties:
      DatabaseInstallationFilesS3BucketName: !Ref BucketName
      DatabaseInstallationFilesS3Prefix: !Ref CEVS3Prefix
      Description: CEV for RDS custom for Oracle
      Engine: custom-oracle-ee
      EngineVersion: 19.18-rdscustom-oracle-jan23psu-cev
      KMSKeyId: !Ref KMSParameter
      Manifest: '{"mediaImportTemplateVersion":"2020-08-14","databaseInstallationFileNames":["V982063-01.zip"],"opatchFileNames":["p6880880_190000_Linux-x86-64.zip"],"psuRuPatchFileNames": ["p34765931_190000_Linux-x86-64.zip","p34786990_190000_Linux-x86-64.zip"],"otherPatchFileNames": ["p35099667_190000_Linux-x86-64.zip","p35099674_190000_Generic.zip","p28730253_190000_Linux-x86-64.zip","p29213893_1918000DBRU_Generic.zip","p33125873_1918000DBRU_Linux-x86-64.zip","p35012866_1918000DBRU_Linux-x86-64.zip"]}'

  CustomOracleAprCOHMTPatch:
    DependsOn: SSMExecutionRole
    Condition: ShouldAddCEV
    Type: AWS::RDS::CustomDBEngineVersion
    Properties:
      DatabaseInstallationFilesS3BucketName:  !Ref BucketName
      DatabaseInstallationFilesS3Prefix: !Ref CEVS3Prefix
      Description: CEV for RDS custom for Oracle
      Engine: custom-oracle-ee-cdb
      EngineVersion: 19.19-custom-oracle-home-mt
      KMSKeyId: !Ref KMSParameter
      Manifest: '{"mediaImportTemplateVersion":"2020-08-14","databaseInstallationFileNames":["V982063-01.zip"],"opatchFileNames":["p6880880_190000_Linux-x86-64.zip"],"psuRuPatchFileNames": ["p35042068_190000_Linux-x86-64.zip","p35050341_190000_Linux-x86-64.zip"],"otherPatchFileNames": ["p28730253_190000_Linux-x86-64.zip","p29213893_1919000DBRU_Generic.zip","p33125873_1919000DBRU_Linux-x86-64.zip","p35220732_190000_Linux-x86-64.zip","p35239280_190000_Generic.zip"],"installationParameters": { "unixGroupName": "dba","unixGroupId": 65789,"unixUname": "oracle","unixUid": 6758,"oracleHome": "/opt/oracle/product/19.19/dbhome_1","oracleBase": "/opt/oracle"}}'

  CustomOracleAprPatch:
    DependsOn: SSMExecutionRole
    Condition: ShouldAddCEV
    Type: AWS::RDS::CustomDBEngineVersion
    Properties:
      DatabaseInstallationFilesS3BucketName:  !Ref BucketName
      DatabaseInstallationFilesS3Prefix: !Ref CEVS3Prefix
      Description: CEV for RDS custom for Oracle
      Engine: custom-oracle-ee
      EngineVersion: 19.19-rdscustom-oracle-apr23psu-cev
      KMSKeyId: !Ref KMSParameter
      Manifest: '{"mediaImportTemplateVersion":"2020-08-14","databaseInstallationFileNames":["V982063-01.zip"],"opatchFileNames":["p6880880_190000_Linux-x86-64.zip"],"psuRuPatchFileNames": ["p35042068_190000_Linux-x86-64.zip","p35050341_190000_Linux-x86-64.zip"],"otherPatchFileNames": ["p28730253_190000_Linux-x86-64.zip","p29213893_1919000DBRU_Generic.zip","p33125873_1919000DBRU_Linux-x86-64.zip","p35220732_190000_Linux-x86-64.zip","p35239280_190000_Generic.zip"]}'

  RDSDBInstance:
    DependsOn: WaitCondition
    Type: AWS::RDS::DBInstance
    Properties:
      AllocatedStorage: 40
      DBInstanceIdentifier: !Sub '${DBInstanceIdentifier}-saz'
      AutoMinorVersionUpgrade: false
      CustomIAMInstanceProfile: AWSRDSCustomInstanceProfileForRdsCustomInstance
      DBInstanceClass: db.m5.large
      DBSubnetGroupName: rds-custom-private
      Engine: custom-oracle-ee
      EngineVersion: 19.18-rdscustom-oracle-jan23psu-cev
      KmsKeyId: !Ref KMSParameter
      MasterUsername: !Sub '{{resolve:secretsmanager:${RDSSecret}:SecretString:username}}'
      MasterUserPassword: !Sub '{{resolve:secretsmanager:${RDSSecret}:SecretString:password}}'
      StorageEncrypted: 'true'
      VPCSecurityGroups:
        - !Ref RDSSG
      DBName: ORCL
      PreferredBackupWindow: '03:00-04:00'
      PreferredMaintenanceWindow: 'Mon:05:00-Mon:06:00'
      PubliclyAccessible: false

  HARDSDBInstance:
    DependsOn: WaitCondition
    Type: AWS::RDS::DBInstance
    Properties:
      AllocatedStorage: 40
      DBInstanceIdentifier: !Sub '${DBInstanceIdentifier}-maz'
      AutoMinorVersionUpgrade: false
      CustomIAMInstanceProfile: AWSRDSCustomInstanceProfileForRdsCustomInstance
      DBInstanceClass: db.m5.large
      DBSubnetGroupName: rds-custom-private
      Engine: custom-oracle-ee
      EngineVersion: 19.19-rdscustom-oracle-apr23psu-cev
      KmsKeyId: !Ref KMSParameter
      MasterUsername: !Sub '{{resolve:secretsmanager:${RDSSecret}:SecretString:username}}'
      MasterUserPassword: !Sub '{{resolve:secretsmanager:${RDSSecret}:SecretString:password}}'
      StorageEncrypted: 'true'
      MultiAZ: 'true'
      VPCSecurityGroups:
        - !Ref RDSSG
      DBName: ORCL
      PreferredBackupWindow: '03:00-04:00'
      PreferredMaintenanceWindow: 'Mon:05:00-Mon:06:00'
      PubliclyAccessible: false

```

### Creating a Secrets Manager secret for a master password

The following example creates an AWS::RDS::DBInstance resource with a managed master user and password.
The secret used to authenticate the DB instance is shown in the `Secret` stack output.
You can remove this output if you don't need to pass the value to another CloudFormation stack.

For more information about managing passwords with Secrets Manager,
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
        "MyInstance": {
            "Properties": {
                "MasterUsername": "masteruser",
                "DBInstanceClass": "db.t3.micro",
                "Engine": "mysql",
                "AllocatedStorage": 5,
                "AutoMinorVersionUpgrade": false,
                "ManageMasterUserPassword": true,
                "MasterUserSecret": {
                    "KmsKeyId": {
                        "Ref": "KMSKey"
                    }
                }
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
                "Fn::GetAtt": "MyInstance.MasterUserSecret.SecretArn"
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
  MyInstance:
    Properties:
      MasterUsername: "masteruser"
      DBInstanceClass: "db.t3.micro"
      Engine: "mysql"
      AllocatedStorage: 5
      AutoMinorVersionUpgrade: false
      ManageMasterUserPassword: true
      MasterUserSecret:
        KmsKeyId: !Ref KMSKey
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
      Fn::GetAtt: MyInstance.MasterUserSecret.SecretArn

```

### Creating a DB instance with self-managed Active Directory domain

The following example creates a AWS CloudFormation stack with the AWS::RDS::DBInstance resource with Active Directory domain.

#### JSON

```json

{
    "Resources":
    {
        "RDSDBInstance":
        {
            "Type": "AWS::RDS::DBInstance",
            "Properties":
            {
                "AllocatedStorage": { "Ref" : "DBAllocatedStorage" },
                "DBInstanceClass": { "Ref" : "DBClass" },
                "DBSubnetGroupName": { "Ref" : "MyDBSubnetGroup" },
                "Engine": "sqlserver-ee",
                "LicenseModel": "license-included",
                "MasterUsername"    : { "Ref" : "DBUsername" },
                "MasterUserPassword": { "Ref" : "DBPassword" },
                "DomainAuthSecretArn": { "Ref" : "DomainSecretArn" },
                "DomainDnsIps":
                [
                    "10.0.0.1",
                    "10.0.0.2"
                ],
                "DomainFqdn": { "Ref" : "MyDomainFqdn" },
                "DomainOu": { "Ref" : "MyDomainOu" }
            },
            "UpdateReplacePolicy": "Delete",
            "DeletionPolicy": "Delete"
        }
    }
}
```

#### YAML

```yaml

Resources:
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      AllocatedStorage: !Ref DBAllocatedStorage
      DBInstanceClass: !Ref DBClass
      DBSubnetGroupName: !Ref MyDBSubnetGroup
      Engine: sqlserver-ee
      LicenseModel: license-included
      MasterUsername: !Ref DBUsername
      MasterUserPassword: !Ref DBPassword
      DomainAuthSecretArn: !Ref DomainSecretArn
      DomainDnsIps:
        - 10.0.0.1
        - 10.0.0.2
      DomainFqdn: !Ref MyDomainFqdn
      DomainOu: !Ref MyDomainOu
    UpdateReplacePolicy: Delete
    DeletionPolicy: Delete

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AdditionalStorageVolume

All content copied from https://docs.aws.amazon.com/.
