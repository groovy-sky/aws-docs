This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::Cluster

Specifies a cluster. All nodes in the cluster run the same
protocol-compliant engine software.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MemoryDB::Cluster",
  "Properties" : {
      "ACLName" : String,
      "AutoMinorVersionUpgrade" : Boolean,
      "ClusterEndpoint" : Endpoint,
      "ClusterName" : String,
      "DataTiering" : String,
      "Description" : String,
      "Engine" : String,
      "EngineVersion" : String,
      "FinalSnapshotName" : String,
      "IpDiscovery" : String,
      "KmsKeyId" : String,
      "MaintenanceWindow" : String,
      "MultiRegionClusterName" : String,
      "NetworkType" : String,
      "NodeType" : String,
      "NumReplicasPerShard" : Integer,
      "NumShards" : Integer,
      "ParameterGroupName" : String,
      "Port" : Integer,
      "SecurityGroupIds" : [ String, ... ],
      "SnapshotArns" : [ String, ... ],
      "SnapshotName" : String,
      "SnapshotRetentionLimit" : Integer,
      "SnapshotWindow" : String,
      "SnsTopicArn" : String,
      "SnsTopicStatus" : String,
      "SubnetGroupName" : String,
      "Tags" : [ Tag, ... ],
      "TLSEnabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::MemoryDB::Cluster
Properties:
  ACLName: String
  AutoMinorVersionUpgrade: Boolean
  ClusterEndpoint:
    Endpoint
  ClusterName: String
  DataTiering: String
  Description: String
  Engine: String
  EngineVersion: String
  FinalSnapshotName: String
  IpDiscovery: String
  KmsKeyId: String
  MaintenanceWindow: String
  MultiRegionClusterName: String
  NetworkType: String
  NodeType: String
  NumReplicasPerShard: Integer
  NumShards: Integer
  ParameterGroupName: String
  Port: Integer
  SecurityGroupIds:
    - String
  SnapshotArns:
    - String
  SnapshotName: String
  SnapshotRetentionLimit: Integer
  SnapshotWindow: String
  SnsTopicArn: String
  SnsTopicStatus: String
  SubnetGroupName: String
  Tags:
    - Tag
  TLSEnabled: Boolean

```

## Properties

`ACLName`

The name of the Access Control List to associate with the cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z][a-zA-Z0-9\-]*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoMinorVersionUpgrade`

When set to true, the cluster will automatically receive minor engine
version upgrades after launch.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterEndpoint`

The cluster's configuration endpoint.

_Required_: No

_Type_: [Endpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-memorydb-cluster-endpoint.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The name of the cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataTiering`

Enables data tiering. Data tiering is only supported for clusters using the r6gd node type.
This parameter must be set when using r6gd nodes. For more information, see [Data tiering](https://docs.aws.amazon.com/memorydb/latest/devguide/data-tiering.html).

_Required_: No

_Type_: String

_Allowed values_: `true | false`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The name of the engine used by the cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

The Redis engine version used by the cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FinalSnapshotName`

The user-supplied name of a final cluster snapshot. This is the unique
name that identifies the snapshot. MemoryDB creates the snapshot, and then
deletes the cluster immediately afterward.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpDiscovery`

The mechanism that the cluster uses to discover IP addresses. Returns 'ipv4' when DNS endpoints resolve to IPv4 addresses, or 'ipv6' when DNS endpoints resolve to IPv6 addresses.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the KMS key used to encrypt the cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceWindow`

Specifies the weekly time range during which maintenance on the cluster
is performed. It is specified as a range in the format `ddd:hh24:mi-ddd:hh24:mi`
(24H Clock UTC). The minimum maintenance window is a 60 minute period.

_Pattern_: `ddd:hh24:mi-ddd:hh24:mi`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiRegionClusterName`

The name of the multi-Region cluster that this cluster belongs to.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

The IP address type for the cluster. Returns 'ipv4' for IPv4 only, 'ipv6' for IPv6 only, or 'dual-stack' if the cluster supports both IPv4 and IPv6 addressing.

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dual_stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeType`

The cluster's node type.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumReplicasPerShard`

The number of replicas to apply to each shard.

_Default value_: `1`

_Maximum value_: `5`

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumShards`

The number of shards in the cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterGroupName`

The name of the parameter group used by the cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port used by the cluster.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

A list of security group names to associate with this cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotArns`

A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files
stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotName`

The name of a snapshot from which to restore data into the new cluster.
The snapshot status changes to restoring while the new cluster is being
created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotRetentionLimit`

The number of days for which MemoryDB retains automatic snapshots before
deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was
taken today is retained for 5 days before being deleted.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotWindow`

The daily time range (in UTC) during which MemoryDB begins taking a daily
snapshot of your shard. Example: 05:00-09:00 If you do not specify this parameter, MemoryDB automatically chooses an appropriate time range.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the SNS topic,
such as `arn:aws:memorydb:us-east-1:123456789012:mySNSTopic`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicStatus`

The SNS topic must be in Active status to receive notifications.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetGroupName`

The name of the subnet group used by the cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-memorydb-cluster-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TLSEnabled`

A flag to indicate if In-transit encryption is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

`ARN`

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the cluster,
such as `arn:aws:memorydb:us-east-1:123456789012:cluster/my-cluster`

`ClusterEndpoint.Address`

The address of the cluster's configuration endpoint.

`ClusterEndpoint.Port`

The port used by the cluster configuration endpoint.

`ParameterGroupStatus`

The status of the parameter group used by the cluster, for example
`active` or `applying`.

`Status`

The status of the cluster. For example, 'available', 'updating' or 'creating'.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Endpoint
