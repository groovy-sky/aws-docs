This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDB::DBCluster

The `AWS::DocDB::DBCluster` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBCluster.
Amazon DocumentDB is a fully managed, MongoDB-compatible document database engine. For more information, see
[DBCluster](../../../documentdb/latest/developerguide/api-dbcluster.md) in the _Amazon DocumentDB Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DocDB::DBCluster",
  "Properties" : {
      "AvailabilityZones" : [ String, ... ],
      "BackupRetentionPeriod" : Integer,
      "CopyTagsToSnapshot" : Boolean,
      "DBClusterIdentifier" : String,
      "DBClusterParameterGroupName" : String,
      "DBSubnetGroupName" : String,
      "DeletionProtection" : Boolean,
      "EnableCloudwatchLogsExports" : [ String, ... ],
      "EngineVersion" : String,
      "GlobalClusterIdentifier" : String,
      "KmsKeyId" : String,
      "ManageMasterUserPassword" : Boolean,
      "MasterUsername" : String,
      "MasterUserPassword" : String,
      "MasterUserSecretKmsKeyId" : String,
      "NetworkType" : String,
      "Port" : Integer,
      "PreferredBackupWindow" : String,
      "PreferredMaintenanceWindow" : String,
      "RestoreToTime" : String,
      "RestoreType" : String,
      "RotateMasterUserPassword" : Boolean,
      "ServerlessV2ScalingConfiguration" : ServerlessV2ScalingConfiguration,
      "SnapshotIdentifier" : String,
      "SourceDBClusterIdentifier" : String,
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

Type: AWS::DocDB::DBCluster
Properties:
  AvailabilityZones:
    - String
  BackupRetentionPeriod: Integer
  CopyTagsToSnapshot: Boolean
  DBClusterIdentifier: String
  DBClusterParameterGroupName: String
  DBSubnetGroupName: String
  DeletionProtection: Boolean
  EnableCloudwatchLogsExports:
    - String
  EngineVersion: String
  GlobalClusterIdentifier: String
  KmsKeyId: String
  ManageMasterUserPassword: Boolean
  MasterUsername: String
  MasterUserPassword: String
  MasterUserSecretKmsKeyId: String
  NetworkType: String
  Port: Integer
  PreferredBackupWindow: String
  PreferredMaintenanceWindow: String
  RestoreToTime: String
  RestoreType: String
  RotateMasterUserPassword: Boolean
  ServerlessV2ScalingConfiguration:
    ServerlessV2ScalingConfiguration
  SnapshotIdentifier: String
  SourceDBClusterIdentifier: String
  StorageEncrypted: Boolean
  StorageType: String
  Tags:
    - Tag
  UseLatestRestorableTime: Boolean
  VpcSecurityGroupIds:
    - String

```

## Properties

`AvailabilityZones`

A list of Amazon EC2 Availability Zones that instances in the
cluster can be created in.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupRetentionPeriod`

The number of days for which automated backups are retained. You
must specify a minimum value of 1.

Default: 1

Constraints:

- Must be a value from 1 to 35.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyTagsToSnapshot`

Set to `true` to copy all tags from the source cluster
snapshot to the target cluster snapshot, and otherwise
`false`. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterIdentifier`

The cluster identifier. This parameter is stored as a lowercase
string.

Constraints:

- Must contain from 1 to 63 letters, numbers, or hyphens.

- The first character must be a letter.

- Cannot end with a hyphen or contain two consecutive hyphens.

Example: `my-cluster`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBClusterParameterGroupName`

The name of the cluster parameter group to associate with this
cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSubnetGroupName`

A subnet group to associate with this cluster.

Constraints: Must match the name of an existing
`DBSubnetGroup`. Must not be default.

Example: `mySubnetgroup`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeletionProtection`

Protects clusters from being accidentally deleted. If enabled, the
cluster cannot be deleted unless it is modified and
`DeletionProtection` is disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableCloudwatchLogsExports`

The list of log types that need to be enabled for exporting to Amazon CloudWatch
Logs. You can enable audit logs or profiler logs. For more information, see
[Auditing Amazon DocumentDB Events](../../../documentdb/latest/developerguide/event-auditing.md)
and [Profiling Amazon DocumentDB Operations](../../../documentdb/latest/developerguide/profiling.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

The version number of the database engine to use. The `--engine-version` will default to the latest major engine version.
For production workloads, we recommend explicitly declaring this parameter with the intended major engine version.

If you intend to trigger an in-place upgrade, please refer to [Amazon DocumentDB in-place major version upgrade](../../../documentdb/latest/developerguide/docdb-mvu.md).
Note that for an in-place engine version upgrade, you need to remove other cluster properties changes (e.g. SecurityGroupId) from the CFN template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalClusterIdentifier`

The cluster identifier of the new global cluster.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z][0-9A-Za-z-:._]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The AWS KMS key identifier for an encrypted cluster.

The AWS KMS key identifier is the Amazon Resource Name (ARN) for the AWS KMS encryption key. If you are creating a cluster using the same AWS account that owns the AWS KMS encryption key that is used to encrypt the new cluster, you can use the AWS KMS key alias instead of the ARN for the AWS KMS encryption key.

If an encryption key is not specified in `KmsKeyId`:

- If the `StorageEncrypted` parameter is
`true`, Amazon DocumentDB uses your default encryption key.

AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS Regions.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManageMasterUserPassword`

Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.

Constraint: You can't manage the master user password with Amazon Web Services Secrets Manager if `MasterUserPassword` is specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUsername`

The name of the master user for the cluster.

Constraints:

- Must be from 1 to 63 letters or numbers.

- The first character must be a letter.

- Cannot be a reserved word for the chosen database engine.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MasterUserPassword`

The password for the master database user. This password can
contain any printable ASCII character except forward slash (/),
double quote ("), or the "at" symbol (@).

Constraints: Must contain from 8 to 100 characters.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserSecretKmsKeyId`

The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.
This setting is valid only if the master user password is managed by Amazon DocumentDB in Amazon Web Services Secrets Manager for the DB cluster.

The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.

If you don't specify `MasterUserSecretKmsKeyId`, then the `aws/secretsmanager` KMS key is used to encrypt the secret.
If the secret is in a different Amazon Web Services account, then you can't use the `aws/secretsmanager` KMS key to encrypt the secret, and you must use a customer managed KMS key.

There is a default KMS key for your Amazon Web Services account.
Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

The network type of the cluster.

The network type is determined by the `DBSubnetGroup` specified for the cluster.
A `DBSubnetGroup` can support only the IPv4 protocol or the IPv4 and the IPv6 protocols ( `DUAL`).

For more information, see [DocumentDB clusters in a VPC](../../../documentdb/latest/developerguide/vpc-clusters.md) in the Amazon DocumentDB Developer Guide.

Valid Values: `IPV4` \| `DUAL`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

Specifies the port that the database engine is listening on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredBackupWindow`

The daily time range during which automated backups are created if
automated backups are enabled using the `BackupRetentionPeriod` parameter.

The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region.

Constraints:

- Must be in the format `hh24:mi-hh24:mi`.

- Must be in Universal Coordinated Time (UTC).

- Must not conflict with the preferred maintenance window.

- Must be at least 30 minutes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur,
in Universal Coordinated Time (UTC).

Format: `ddd:hh24:mi-ddd:hh24:mi`

The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week.

Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun

Constraints: Minimum 30-minute window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreToTime`

The date and time to restore the cluster to.

Valid values: A time in Universal Coordinated Time (UTC) format.

Constraints:

- Must be before the latest restorable time for the instance.

- Must be specified if the `UseLatestRestorableTime` parameter is not provided.

- Cannot be specified if the `UseLatestRestorableTime` parameter is `true`.

- Cannot be specified if the `RestoreType` parameter is `copy-on-write`.

Example: `2015-03-07T23:45:00Z`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreType`

The type of restore to be performed. You can specify one of the following values:

- `full-copy` \- The new DB cluster is restored as a full copy of the
source DB cluster.

- `copy-on-write` \- The new DB cluster is restored as a clone of the
source DB cluster.

Constraints: You can't specify `copy-on-write` if the engine version of the source DB cluster is earlier than 1.11.

If you don't specify a `RestoreType` value, then the new DB cluster is
restored as a full copy of the source DB cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotateMasterUserPassword`

Specifies whether to rotate the secret managed by Amazon Web Services Secrets Manager for the master user password.

This setting is valid only if the master user password is managed by Amazon DocumentDB in Amazon Web Services Secrets Manager for the cluster.
The secret value contains the updated password.

Constraint: You must apply the change immediately when rotating the master user password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerlessV2ScalingConfiguration`

Contains the scaling configuration of an Amazon DocumentDB Serverless cluster.

_Required_: No

_Type_: [ServerlessV2ScalingConfiguration](aws-properties-docdb-dbcluster-serverlessv2scalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotIdentifier`

The identifier for the snapshot or cluster snapshot to restore from.

You can use either the name or the Amazon Resource Name (ARN) to specify a cluster
snapshot. However, you can use only the ARN to specify a snapshot.

Constraints:

- Must match the identifier of an existing snapshot.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceDBClusterIdentifier`

The identifier of the source cluster from which to restore.

Constraints:

- Must match the identifier of an existing `DBCluster`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageEncrypted`

Specifies whether the cluster is encrypted.

If you specify `SourceDBClusterIdentifier` or `SnapshotIdentifier` and don’t specify `StorageEncrypted`, the encryption property is inherited from the source cluster or snapshot (unless `KMSKeyId` is specified, in which case the restored cluster will be encrypted with that KMS key).
If the source is encrypted and `StorageEncrypted` is specified to be true, the restored cluster will be encrypted (if you want to use a different KMS key, specify the `KMSKeyId` property as well).
If the source is unencrypted and `StorageEncrypted` is specified to be true, then the `KMSKeyId` property must be specified.
If the source is encrypted, don’t specify `StorageEncrypted` to be false as opting out of encryption is not allowed.

_Required_: Conditional

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageType`

The storage type to associate with the DB cluster.

For information on storage types for Amazon DocumentDB clusters, see
Cluster storage configurations in the _Amazon DocumentDB Developer Guide_.

Valid values for storage type - `standard | iopt1`

Default value is `standard `

###### Note

When you create an Amazon DocumentDB cluster with the storage type set to `iopt1`, the storage type is returned
in the response. The storage type isn't returned when you set it to `standard`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be assigned to the cluster.

_Required_: No

_Type_: Array of [Tag](aws-properties-docdb-dbcluster-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLatestRestorableTime`

A value that is set to `true` to restore the cluster to the latest
restorable backup time, and `false` otherwise.

Default: `false`

Constraints: Cannot be specified if the `RestoreToTime` parameter is
provided.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

A list of EC2 VPC security groups to associate with this cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DBClusterIdentifier, such as `mycluster`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClusterResourceId`

The resource id for the cluster; for example: `cluster-ABCD1234EFGH5678IJKL90MNOP`.
The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.

`Endpoint`

The connection endpoint for the cluster, such as
`sample-cluster.cluster-cozrlsfrcjoc.us-east-1.docdb.amazonaws.com`.

`Port`

The port number on which the cluster accepts connections. For example: `27017`.

`ReadEndpoint`

The reader endpoint for the cluster. For example: `sample-cluster.cluster-ro-cozrlsfrcjoc.us-east-1.docdb.amazonaws.com`.

## Examples

#### JSON

```json

{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myDBInstance" : {
         "Type" : "AWS::DocDB::DBCluster",
         "Properties" : {
            "BackupRetentionPeriod" : 8,
            "DBClusterIdentifier" : "sample-cluster",
            "DBClusterParameterGroupName" : "default.docdb3.6",
            "DBSubnetGroupName" : "default",
            "KmsKeyId" : "your-kms-key-id",
            "MasterUsername" : "your-master-username",
            "MasterUserPassword" : "your-master-user-password",
            "Port" : "27017",
            "PreferredBackupWindow" : "07:34-08:04",
            "PreferredMaintenanceWindow" : "sat:04:51-sat:05:21",
            "SnapshotIdentifier" : "sample-cluster-snapshot-id",
            "StorageEncrypted" : true,
            "Tags" : [ {"Key" : "String", "Value" : "String"} ]
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Resources:
   myDBInstance:
      Type: "AWS::DocDB::DBCluster"
      Properties:
         BackupRetentionPeriod : 8
         DBClusterIdentifier : "sample-cluster"
         DBClusterParameterGroupName : "default.docdb3.6"
         DBSubnetGroupName : "default"
         KmsKeyId : "your-kms-key-id"
         MasterUsername : "your-master-username"
         MasterUserPassword : "your-master-user-password"
         Port : "27017"
         PreferredBackupWindow : "07:34-08:04"
         PreferredMaintenanceWindow : "sat:04:51-sat:05:21"
         SnapshotIdentifier : "sample-cluster-snapshot-id"
         StorageEncrypted : true
         Tags:
            -
               Key: "String"
               Value: "String"
```

## See also

- [DBCluster](../../../documentdb/latest/developerguide/api-dbcluster.md)

- [CreateDBCluster](../../../documentdb/latest/developerguide/api-createdbcluster.md)

- [DeleteDBCluster](../../../documentdb/latest/developerguide/api-deletedbcluster.md)

- [DescribeDBClusters](../../../documentdb/latest/developerguide/api-describedbclusters.md)

- [ModifyDBCluster](../../../documentdb/latest/developerguide/api-modifydbcluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon DocumentDB (with MongoDB compatibility)

ServerlessV2ScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
