This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBCluster

The `AWS::Neptune::DBCluster` resource creates an Amazon Neptune DB cluster.
Neptune is a fully managed graph database.

###### Note

Currently, you can create this resource only in AWS Regions in which Amazon Neptune is
supported.

If no `DeletionPolicy` is set for `AWS::Neptune::DBCluster`
resources, the default deletion behavior is that the entire volume will be deleted without a snapshot.
To retain a backup of the volume, the `DeletionPolicy` should be set to `Snapshot`.
For more information about how CloudFormation deletes resources,
see [DeletionPolicy Attribute](../userguide/aws-attribute-deletionpolicy.md).

You can use `AWS::Neptune::DBCluster.DeletionProtection` to help guard against
unintended deletion of your DB cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Neptune::DBCluster",
  "Properties" : {
      "AssociatedRoles" : [ DBClusterRole, ... ],
      "AvailabilityZones" : [ String, ... ],
      "BackupRetentionPeriod" : Integer,
      "CopyTagsToSnapshot" : Boolean,
      "DBClusterIdentifier" : String,
      "DBClusterParameterGroupName" : String,
      "DBInstanceParameterGroupName" : String,
      "DBPort" : Integer,
      "DBSubnetGroupName" : String,
      "DeletionProtection" : Boolean,
      "EnableCloudwatchLogsExports" : [ String, ... ],
      "EngineVersion" : String,
      "IamAuthEnabled" : Boolean,
      "KmsKeyId" : String,
      "Port" : String,
      "PreferredBackupWindow" : String,
      "PreferredMaintenanceWindow" : String,
      "RestoreToTime" : String,
      "RestoreType" : String,
      "ServerlessScalingConfiguration" : ServerlessScalingConfiguration,
      "SnapshotIdentifier" : String,
      "SourceDBClusterIdentifier" : String,
      "StorageEncrypted" : Boolean,
      "Tags" : [ Tag, ... ],
      "UseLatestRestorableTime" : Boolean,
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Neptune::DBCluster
Properties:
  AssociatedRoles:
    - DBClusterRole
  AvailabilityZones:
    - String
  BackupRetentionPeriod: Integer
  CopyTagsToSnapshot: Boolean
  DBClusterIdentifier: String
  DBClusterParameterGroupName: String
  DBInstanceParameterGroupName: String
  DBPort: Integer
  DBSubnetGroupName: String
  DeletionProtection: Boolean
  EnableCloudwatchLogsExports:
    - String
  EngineVersion: String
  IamAuthEnabled: Boolean
  KmsKeyId: String
  Port: String
  PreferredBackupWindow: String
  PreferredMaintenanceWindow: String
  RestoreToTime: String
  RestoreType: String
  ServerlessScalingConfiguration:
    ServerlessScalingConfiguration
  SnapshotIdentifier: String
  SourceDBClusterIdentifier: String
  StorageEncrypted: Boolean
  Tags:
    - Tag
  UseLatestRestorableTime: Boolean
  VpcSecurityGroupIds:
    - String

```

## Properties

`AssociatedRoles`

Provides a list of the Amazon Identity and Access Management (IAM) roles that are associated
with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the
DB cluster to access other Amazon services on your behalf.

_Required_: No

_Type_: Array of [DBClusterRole](aws-properties-neptune-dbcluster-dbclusterrole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZones`

Provides the list of EC2 Availability Zones that instances in the DB cluster can be
created in.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupRetentionPeriod`

Specifies the number of days for which automatic DB snapshots are retained.

An update may require some interruption. See [ModifyDBInstance](../../../neptune/latest/userguide/api-instances.md#ModifyDBInstance) in the Amazon Neptune User Guide for more information.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTagsToSnapshot`

_If set to `true`, tags are copied to any snapshot of_
_the DB cluster that is created._

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterIdentifier`

Contains a user-supplied DB cluster identifier. This identifier is the unique key that
identifies a DB cluster.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBClusterParameterGroupName`

Provides the name of the DB cluster parameter group.

An update may require some interruption. See [ModifyDBInstance](../../../neptune/latest/userguide/api-instances.md#ModifyDBInstance)
in the Amazon Neptune User Guide for more information.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBInstanceParameterGroupName`

The name of the DB parameter group to apply to all instances of the DB cluster.
Used only in case of a major engine version upgrade request

Note that when you apply a parameter group using `DBInstanceParameterGroupName`,
parameter changes are applied immediately, not during the next maintenance window.

###### Constraints

- The DB parameter group must be in the same DB parameter group family
as the target DB cluster version.

- The `DBInstanceParameterGroupName` parameter is only
valid for major engine version upgrades.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBPort`

The port number on which the DB instances in the DB cluster accept connections.

If not specified, the default port used is `8182`.

###### Note

The `Port` property will soon be deprecated. Please update existing templates to use
the new `DBPort` property that has the same functionality.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSubnetGroupName`

Specifies information on the subnet group associated with the DB cluster, including the
name, description, and subnets in the subnet group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeletionProtection`

Indicates whether or not the DB cluster has deletion protection
enabled. The database can't be deleted when deletion protection is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableCloudwatchLogsExports`

Specifies a list of log types that are enabled for export to CloudWatch Logs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

Indicates the database engine version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamAuthEnabled`

True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts
is enabled, and otherwise false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The Amazon Resource Name (ARN) of the KMS key that is used to encrypt the database instances in the DB cluster,
such as `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef`. If you enable the
`StorageEncrypted` property but don't specify this property, the default KMS key is used. If you specify
this property, you must set the `StorageEncrypted` property to `true`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The port number on which the DB instances in the DB cluster accept connections.

If not specified, the default port used is `8182`.

###### Note

This property will soon be deprecated. Please update existing templates to use
the new `DBPort` property that has the same functionality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredBackupWindow`

Specifies the daily time range during which automated backups are created if automated
backups are enabled, as determined by the `BackupRetentionPeriod`.

An update may require some interruption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

Specifies the weekly time range during which system maintenance can occur, in Universal
Coordinated Time (UTC).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreToTime`

Creates a new DB cluster from a DB snapshot or DB cluster snapshot.

If a DB snapshot is specified, the target DB cluster is created from the source DB
snapshot with a default configuration and default security group.

If a DB cluster snapshot is specified, the target DB cluster is created from the source DB
cluster restore point with the same configuration as the original source DB cluster, except
that the new DB cluster is created with the default security group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestoreType`

Creates a new DB cluster from a DB snapshot or DB cluster snapshot.

If a DB snapshot is specified, the target DB cluster is created from the source DB
snapshot with a default configuration and default security group.

If a DB cluster snapshot is specified, the target DB cluster is created from the source DB
cluster restore point with the same configuration as the original source DB cluster, except
that the new DB cluster is created with the default security group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerlessScalingConfiguration`

Property description not available.

_Required_: No

_Type_: [ServerlessScalingConfiguration](aws-properties-neptune-dbcluster-serverlessscalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotIdentifier`

Specifies the identifier for a DB cluster snapshot. Must match the identifier
of an existing snapshot.

After you restore a DB cluster using a `SnapshotIdentifier`,
you must specify the same `SnapshotIdentifier` for any future
updates to the DB cluster. When you specify this property for an update, the DB
cluster is not restored from the snapshot again, and the data in the database is not
changed.

However, if you don't specify the `SnapshotIdentifier`, an empty
DB cluster is created, and the original DB cluster is deleted. If you specify a
property that is different from the previous snapshot restore property, the DB
cluster is restored from the snapshot specified by the `SnapshotIdentifier`,
and the original DB cluster is deleted.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceDBClusterIdentifier`

Creates a new DB cluster from a DB snapshot or DB cluster snapshot.

If a DB snapshot is specified, the target DB cluster is created from the source DB
snapshot with a default configuration and default security group.

If a DB cluster snapshot is specified, the target DB cluster is created from the source DB
cluster restore point with the same configuration as the original source DB cluster, except
that the new DB cluster is created with the default security group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageEncrypted`

Indicates whether the DB cluster is encrypted.

If you specify the `KmsKeyId` property, then you must enable encryption and set this property to
`true`.

If you enable the `StorageEncrypted` property but don't specify the `KmsKeyId` property, then the
default KMS key is used. If you specify the `KmsKeyId` property, then that KMS key is used to encrypt the
database instances in the DB cluster.

If you specify the `SourceDBClusterIdentifier` property, and don't specify this property or disable it, the
value is inherited from the source DB cluster. If the source DB cluster is encrypted, the `KmsKeyId` property from
the source cluster is used.

If you specify the `DBSnapshotIdentifier` and don't specify this property or disable it, the value is
inherited from the snapshot and the specified `KmsKeyId` property from the snapshot is used.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to this cluster.

_Required_: No

_Type_: Array of [Tag](aws-properties-neptune-dbcluster-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLatestRestorableTime`

Creates a new DB cluster from a DB snapshot or DB cluster snapshot.

If a DB snapshot is specified, the target DB cluster is created from the source DB
snapshot with a default configuration and default security group.

If a DB cluster snapshot is specified, the target DB cluster is created from the source DB
cluster restore point with the same configuration as the original source DB cluster, except
that the new DB cluster is created with the default security group.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSecurityGroupIds`

Provides a list of VPC security groups that the DB cluster belongs to.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClusterResourceId`

The resource id for the DB cluster. For example: `cluster-ABCD1234EFGH5678IJKL90MNOP`.
The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.

`Endpoint`

The connection endpoint for the DB cluster. For example:
`mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`

`Port`

The port number on which the DB instances in the DB cluster accept connections.

`ReadEndpoint`

The reader endpoint for the DB cluster. For example:
`mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Neptune

DBClusterRole

All content copied from https://docs.aws.amazon.com/.
