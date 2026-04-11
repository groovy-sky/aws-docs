This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ReplicationGroup

The `AWS::ElastiCache::ReplicationGroup` resource creates an Amazon ElastiCache (Valkey or Redis OSS) replication
group. A Valkey or Redis OSS (cluster mode disabled) replication group is a collection of cache clusters, where one of the clusters
is a primary read-write cluster and the others are read-only replicas.

A Valkey or Redis OSS (cluster mode enabled) cluster is comprised of from 1 to 90 shards (API/CLI: node groups). Each shard
has a primary node and up to 5 read-only replica nodes. The configuration can range from 90 shards and 0 replicas to
15 shards and 5 replicas, which is the maximum number or replicas allowed.

The node or shard limit can be increased to a maximum of 500 per cluster if the engine version is Valkey 7.2 or higher, or Redis OSS 5.0.6 or
higher. For example, you can choose to configure a 500 node cluster that ranges between 83 shards (one primary and 5
replicas per shard) and 500 shards (single primary and no replicas). Make sure there are enough available IP
addresses to accommodate the increase. Common pitfalls include the subnets in the subnet group have too small a CIDR
range or the subnets are shared and heavily used by other clusters. For more information, see [Creating a Subnet\
Group](../../../amazonelasticache/latest/dg/subnetgroups-creating.md). For versions below 5.0.6, the limit is 250 per cluster.

To request a limit increase, see [Amazon Service Limits](../../../../general/latest/gr/aws-service-limits.md) and choose the limit type **Nodes per cluster per instance**
**type**.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::ReplicationGroup",
  "Properties" : {
      "AtRestEncryptionEnabled" : Boolean,
      "AuthToken" : String,
      "AutomaticFailoverEnabled" : Boolean,
      "AutoMinorVersionUpgrade" : Boolean,
      "CacheNodeType" : String,
      "CacheParameterGroupName" : String,
      "CacheSecurityGroupNames" : [ String, ... ],
      "CacheSubnetGroupName" : String,
      "ClusterMode" : String,
      "DataTieringEnabled" : Boolean,
      "Engine" : String,
      "EngineVersion" : String,
      "GlobalReplicationGroupId" : String,
      "IpDiscovery" : String,
      "KmsKeyId" : String,
      "LogDeliveryConfigurations" : [ LogDeliveryConfigurationRequest, ... ],
      "MultiAZEnabled" : Boolean,
      "NetworkType" : String,
      "NodeGroupConfiguration" : [ NodeGroupConfiguration, ... ],
      "NotificationTopicArn" : String,
      "NumCacheClusters" : Integer,
      "NumNodeGroups" : Integer,
      "Port" : Integer,
      "PreferredCacheClusterAZs" : [ String, ... ],
      "PreferredMaintenanceWindow" : String,
      "PrimaryClusterId" : String,
      "ReplicasPerNodeGroup" : Integer,
      "ReplicationGroupDescription" : String,
      "ReplicationGroupId" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SnapshotArns" : [ String, ... ],
      "SnapshotName" : String,
      "SnapshotRetentionLimit" : Integer,
      "SnapshottingClusterId" : String,
      "SnapshotWindow" : String,
      "Tags" : [ Tag, ... ],
      "TransitEncryptionEnabled" : Boolean,
      "TransitEncryptionMode" : String,
      "UserGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::ReplicationGroup
Properties:
  AtRestEncryptionEnabled: Boolean
  AuthToken: String
  AutomaticFailoverEnabled: Boolean
  AutoMinorVersionUpgrade: Boolean
  CacheNodeType: String
  CacheParameterGroupName: String
  CacheSecurityGroupNames:
    - String
  CacheSubnetGroupName: String
  ClusterMode: String
  DataTieringEnabled: Boolean
  Engine: String
  EngineVersion: String
  GlobalReplicationGroupId: String
  IpDiscovery: String
  KmsKeyId: String
  LogDeliveryConfigurations:
    - LogDeliveryConfigurationRequest
  MultiAZEnabled: Boolean
  NetworkType: String
  NodeGroupConfiguration:
    - NodeGroupConfiguration
  NotificationTopicArn: String
  NumCacheClusters: Integer
  NumNodeGroups: Integer
  Port: Integer
  PreferredCacheClusterAZs:
    - String
  PreferredMaintenanceWindow: String
  PrimaryClusterId: String
  ReplicasPerNodeGroup: Integer
  ReplicationGroupDescription: String
  ReplicationGroupId: String
  SecurityGroupIds:
    - String
  SnapshotArns:
    - String
  SnapshotName: String
  SnapshotRetentionLimit: Integer
  SnapshottingClusterId: String
  SnapshotWindow: String
  Tags:
    - Tag
  TransitEncryptionEnabled: Boolean
  TransitEncryptionMode: String
  UserGroupIds:
    - String

```

## Properties

`AtRestEncryptionEnabled`

A flag that enables encryption at rest when set to `true`.

**Required:** Only available when creating a replication group in an Amazon VPC
using Redis OSS version `3.2.6` or `4.x` onward.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthToken`

**Reserved parameter.** The password used to access a password protected
server.

`AuthToken` can be specified only on replication groups where `TransitEncryptionEnabled`
is `true`. For more information, see [Authenticating Valkey or Redis OSS users with the AUTH Command](../../../amazonelasticache/latest/dg/auth.md).

###### Important

For HIPAA compliance, you must specify `TransitEncryptionEnabled` as `true`, an
`AuthToken`, and a `CacheSubnetGroup`.

Password constraints:

- Must be only printable ASCII characters.

- Must be at least 16 characters and no more than 128 characters in length.

- Nonalphanumeric characters are restricted to (!, &, #, $, ^, <, >, -, ).

For more information, see [AUTH password](http://redis.io/commands/AUTH) at
http://redis.io/commands/AUTH.

###### Note

If ADDING the AuthToken, update requires [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticFailoverEnabled`

Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary
fails.

`AutomaticFailoverEnabled` must be enabled for Valkey or Redis OSS (cluster mode enabled) replication
groups.

Default: false

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoMinorVersionUpgrade`

If you are running Valkey 7.2 or later, or Redis OSS 6.0 or later, set this parameter to yes if you want to opt-in to the
next minor version upgrade campaign. This parameter is disabled for previous versions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheNodeType`

The compute and memory capacity of the nodes in the node group (shard).

The following node types are supported by ElastiCache. Generally speaking, the current generation types provide
more memory and computational power at lower cost when compared to their equivalent previous generation
counterparts.

- General purpose:

- Current generation:

**M6g node types:** `cache.m6g.large`,
`cache.m6g.xlarge`, `cache.m6g.2xlarge`, `cache.m6g.4xlarge`,
`cache.m6g.12xlarge`, `cache.m6g.24xlarge`

**M5 node types:** `cache.m5.large`, `cache.m5.xlarge`,
`cache.m5.2xlarge`, `cache.m5.4xlarge`, `cache.m5.12xlarge`,
`cache.m5.24xlarge`

**M4 node types:** `cache.m4.large`, `cache.m4.xlarge`,
`cache.m4.2xlarge`, `cache.m4.4xlarge`, `cache.m4.10xlarge`

**T4g node types:** `cache.t4g.micro`, `cache.t4g.small`,
`cache.t4g.medium`

**T3 node types:** `cache.t3.micro`, `cache.t3.small`,
`cache.t3.medium`

**T2 node types:** `cache.t2.micro`, `cache.t2.small`,
`cache.t2.medium`

- Previous generation: (not recommended)

**T1 node types:** `cache.t1.micro`

**M1 node types:** `cache.m1.small`, `cache.m1.medium`,
`cache.m1.large`, `cache.m1.xlarge`

**M3 node types:** `cache.m3.medium`, `cache.m3.large`,
`cache.m3.xlarge`, `cache.m3.2xlarge`

- Compute optimized:

- Previous generation: (not recommended)

**C1 node types:** `cache.c1.xlarge`

- Memory optimized:

- Current generation:

**R6gd node types:** `cache.r6gd.xlarge`,
`cache.r6gd.2xlarge`, `cache.r6gd.4xlarge`, `cache.r6gd.8xlarge`,
`cache.r6gd.12xlarge`, `cache.r6gd.16xlarge`

###### Note

The `r6gd` family is available in the following regions: `us-east-2`,
`us-east-1`, `us-west-2`, `us-west-1`, `eu-west-1`,
`eu-central-1`, `ap-northeast-1`, `ap-southeast-1`,
`ap-southeast-2`.

**R6g node types:** `cache.r6g.large`,
`cache.r6g.xlarge`, `cache.r6g.2xlarge`, `cache.r6g.4xlarge`,
`cache.r6g.12xlarge`, `cache.r6g.24xlarge`

**R5 node types:** `cache.r5.large`, `cache.r5.xlarge`,
`cache.r5.2xlarge`, `cache.r5.4xlarge`, `cache.r5.12xlarge`,
`cache.r5.24xlarge`

**R4 node types:** `cache.r4.large`, `cache.r4.xlarge`,
`cache.r4.2xlarge`, `cache.r4.4xlarge`, `cache.r4.8xlarge`,
`cache.r4.16xlarge`

- Previous generation: (not recommended)

**M2 node types:** `cache.m2.xlarge`, `cache.m2.2xlarge`,
`cache.m2.4xlarge`

**R3 node types:** `cache.r3.large`, `cache.r3.xlarge`,
`cache.r3.2xlarge`, `cache.r3.4xlarge`, `cache.r3.8xlarge`

For region availability, see [Supported Node\
Types by Amazon Region](../../../amazonelasticache/latest/dg/cachenodes-supportedtypes.md#CacheNodes.SupportedTypesByRegion)

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheParameterGroupName`

The name of the parameter group to associate with this replication group. If this
argument is omitted, the default cache parameter group for the specified engine is
used.

If you are running Valkey or Redis OSS version 3.2.4 or later, only one node group (shard), and want
to use a default parameter group, we recommend that you specify the parameter group by
name.

- To create a Valkey or Redis OSS (cluster mode disabled) replication group, use
`CacheParameterGroupName=default.redis3.2`.

- To create a Valkey or Redis OSS (cluster mode enabled) replication group, use
`CacheParameterGroupName=default.redis3.2.cluster.on`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheSecurityGroupNames`

A list of cache security group names to associate with this replication group.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheSubnetGroupName`

The name of the cache subnet group to be used for the replication group.

###### Important

If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start
creating a cluster. For more information, see [AWS::ElastiCache::SubnetGroup](../userguide/aws-properties-elasticache-subnetgroup.md).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterMode`

The mode can be enabled or disabled. To change the cluster mode from disabled to enabled, you must first set the cluster mode to
compatible. The compatible mode allows your Valkey or Redis OSS clients to connect using both cluster mode enabled and cluster mode
disabled. After you migrate all Valkey or Redis OSS clients to use cluster mode enabled, you can then complete cluster mode
configuration and set the cluster mode to enabled. For more information, see [Modify cluster mode](../../../amazonelasticache/latest/dg/modify-cluster-mode.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTieringEnabled`

Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This
parameter must be set to true when using r6gd nodes. For more information, see [Data tiering](../../../amazonelasticache/latest/dg/data-tiering.md).

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Engine`

The name of the cache engine to be used for the clusters in this replication group.
The value must be set to `valkey` or `redis`.

###### Note

Upgrading an existing engine from redis to valkey is done through in-place migration, and requires a parameter group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

The version number of the cache engine to be used for the clusters in this replication
group. To view the supported cache engine versions, use the
`DescribeCacheEngineVersions` operation.

**Important:** You can upgrade to a newer engine version
(see [Selecting\
a Cache Engine and Version](../../../amazonelasticache/latest/dg/selectengine.md#VersionManagement)) in the _ElastiCache User_
_Guide_, but you cannot downgrade to an earlier engine version. If you want
to use an earlier engine version, you must delete the existing cluster or replication
group and create it anew with the earlier engine version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalReplicationGroupId`

The name of the Global datastore

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpDiscovery`

The network type you choose when creating a replication group, either
`ipv4` \| `ipv6`. IPv6 is supported for workloads using Valkey 7.2 and above, Redis OSS engine version 6.2
to 7.1 or Memcached engine version 1.6.6 and above on all instances built on
the [Nitro system](https://aws.amazon.com/ec2/nitro).

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the KMS key used to encrypt the disk on the cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogDeliveryConfigurations`

Specifies the destination, format and type of the logs.

_Required_: No

_Type_: Array of [LogDeliveryConfigurationRequest](aws-properties-elasticache-replicationgroup-logdeliveryconfigurationrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiAZEnabled`

A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see [Minimizing Downtime:\
Multi-AZ](../../../amazonelasticache/latest/dg/autofailover.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

Must be either `ipv4` \| `ipv6` \| `dual_stack`. IPv6
is supported for workloads using Valkey 7.2 and above, Redis OSS engine version 6.2
to 7.1 and Memcached engine version 1.6.6 and above on all instances built on the [Nitro system](https://aws.amazon.com/ec2/nitro).

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dual_stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeGroupConfiguration`

`NodeGroupConfiguration ` is a property of the `AWS::ElastiCache::ReplicationGroup`
resource that configures an Amazon ElastiCache (ElastiCache) Valkey or Redis OSS cluster node group.

If you set [UseOnlineResharding](../userguide/aws-attribute-updatepolicy.md#cfn-attributes-updatepolicy-useonlineresharding) to `true`, you can update `NodeGroupConfiguration` without
interruption. When `UseOnlineResharding` is set to `false`, or is not specified, updating
`NodeGroupConfiguration` results in [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: No

_Type_: [Array](aws-properties-elasticache-replicationgroup-nodegroupconfiguration.md) of [NodeGroupConfiguration](aws-properties-elasticache-replicationgroup-nodegroupconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationTopicArn`

The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic
to which notifications are sent.

###### Note

The Amazon SNS topic owner must be the same as the cluster owner.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumCacheClusters`

The number of clusters this replication group initially has.

This parameter is not used if there is more than one node group (shard). You should
use `ReplicasPerNodeGroup` instead.

If `AutomaticFailoverEnabled` is `true`, the value of this
parameter must be at least 2. If `AutomaticFailoverEnabled` is
`false` you can omit this parameter (it will default to 1), or you can
explicitly set it to a value between 2 and 6.

The maximum permitted value for `NumCacheClusters` is 6 (1 primary plus 5
replicas).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumNodeGroups`

An optional parameter that specifies the number of node groups (shards) for this Valkey or Redis OSS (cluster mode enabled)
replication group. For Valkey or Redis OSS (cluster mode disabled) either omit this parameter or set it to 1.

If you set [UseOnlineResharding](../userguide/aws-attribute-updatepolicy.md#cfn-attributes-updatepolicy-useonlineresharding) to `true`, you can update `NumNodeGroups` without interruption.
When `UseOnlineResharding` is set to `false`, or is not specified, updating
`NumNodeGroups` results in [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

Default: 1

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number on which each member of the replication group accepts
connections.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredCacheClusterAZs`

A list of EC2 Availability Zones in which the replication group's clusters are
created. The order of the Availability Zones in the list is the order in which clusters
are allocated. The primary cluster is created in the first AZ in the list.

This parameter is not used if there is more than one node group (shard). You should
use `NodeGroupConfiguration` instead.

###### Note

If you are creating your replication group in an Amazon VPC (recommended), you can
only locate clusters in Availability Zones associated with the subnets in the
selected subnet group.

The number of Availability Zones listed must equal the value of
`NumCacheClusters`.

Default: system chosen Availability Zones.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredMaintenanceWindow`

Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range
in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.

Valid values for `ddd` are:

- `sun`

- `mon`

- `tue`

- `wed`

- `thu`

- `fri`

- `sat`

Example: `sun:23:00-mon:01:30`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryClusterId`

The identifier of the cluster that serves as the primary for this replication group.
This cluster must already exist and have a status of `available`.

This parameter is not required if `NumCacheClusters`,
`NumNodeGroups`, or `ReplicasPerNodeGroup` is
specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicasPerNodeGroup`

An optional parameter that specifies the number of replica nodes in each node group
(shard). Valid values are 0 to 5.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplicationGroupDescription`

A user-created description for the replication group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationGroupId`

The replication group identifier. This parameter is stored as a lowercase
string.

Constraints:

- A name must contain from 1 to 40 alphanumeric characters or hyphens.

- The first character must be a letter.

- A name cannot end with a hyphen or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

One or more Amazon VPC security groups associated with this replication group.

Use this parameter only when you are creating a replication group in an Amazon Virtual
Private Cloud (Amazon VPC).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotArns`

A list of Amazon Resource Names (ARN) that uniquely identify the Valkey or Redis OSS RDB snapshot
files stored in Amazon S3. The snapshot files are used to populate the new replication
group. The Amazon S3 object name in the ARN cannot contain any commas. The new
replication group will have the number of node groups (console: shards) specified by the
parameter _NumNodeGroups_ or the number of node groups configured by
_NodeGroupConfiguration_ regardless of the number of ARNs
specified here.

Example of an Amazon S3 ARN: `arn:aws:s3:::my_bucket/snapshot1.rdb`

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotName`

The name of a snapshot from which to restore data into the new replication group. The
snapshot status changes to `restoring` while the new replication group is
being created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotRetentionLimit`

The number of days for which ElastiCache retains automatic snapshots before deleting
them. For example, if you set `SnapshotRetentionLimit` to 5, a snapshot that
was taken today is retained for 5 days before being deleted.

Default: 0 (i.e., automatic backups are disabled for this cluster).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshottingClusterId`

The cluster ID that is used as the daily snapshot source for the replication group.
This parameter cannot be set for Valkey or Redis OSS (cluster mode enabled) replication groups.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotWindow`

The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot
of your node group (shard).

Example: `05:00-09:00`

If you do not specify this parameter, ElastiCache automatically chooses an appropriate
time range.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to be added to this resource. Tags are comma-separated key,value pairs
(e.g. Key= `myKey`, Value= `myKeyValue`. You can include multiple
tags as shown following: Key= `myKey`, Value= `myKeyValue`
Key= `mySecondKey`, Value= `mySecondKeyValue`. Tags on
replication groups will be replicated to all nodes.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticache-replicationgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryptionEnabled`

A flag that enables in-transit encryption when set to `true`.

This parameter is only available when creating a replication group in an Amazon VPC
using Valkey version `7.2` and above, Redis OSS version `3.2.6`, or Redis OSS version `4.x` and above, and the cluster is being created in an Amazon VPC.

If you enable in-transit encryption, you must also specify a value for `CacheSubnetGroup`.

###### Note

TransitEncryptionEnabled is required when creating a new valkey replication group.

Default: `false`

###### Important

For HIPAA compliance, you must specify `TransitEncryptionEnabled` as `true`, an
`AuthToken`, and a `CacheSubnetGroup`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryptionMode`

A setting that allows you to migrate your clients to use in-transit encryption, with
no downtime.

When setting `TransitEncryptionEnabled` to `true`, you can set
your `TransitEncryptionMode` to `preferred` in the same request,
to allow both encrypted and unencrypted connections at the same time. Once you migrate
all your Valkey or Redis OSS clients to use encrypted connections you can modify the value to
`required` to allow encrypted connections only.

Setting `TransitEncryptionMode` to `required` is a two-step
process that requires you to first set the `TransitEncryptionMode` to
`preferred`, after that you can set `TransitEncryptionMode` to
`required`.

This process will not trigger the replacement of the replication group.

_Required_: No

_Type_: String

_Allowed values_: `preferred | required`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserGroupIds`

The ID of user group to associate with the replication group.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref returns the resource
name.

For more information about using the Ref function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConfigurationEndPoint.Address`

The DNS hostname of the cache node.

###### Note

Valkey or Redis OSS (cluster mode disabled) replication groups don't have this attribute. Therefore, `Fn::GetAtt`
returns a value for this attribute only if the replication group is clustered. Otherwise, `Fn::GetAtt`
fails. For Valkey or Redis OSS (cluster mode disabled) replication groups, use the `PrimaryEndpoint` or
`ReadEndpoint` attributes.

`ConfigurationEndPoint.Port`

The port number that the cache engine is listening on.

`PrimaryEndPoint.Address`

The DNS address of the primary read-write cache node.

`PrimaryEndPoint.Port`

The number of the port that the primary read-write cache engine is listening on.

`ReadEndPoint.Addresses`

A string with a list of endpoints for the primary and read-only replicas. The order of the addresses maps to the
order of the ports from the `ReadEndPoint.Ports` attribute.

`ReadEndPoint.Addresses.List`

A string with a list of endpoints for the read-only replicas. The order of the addresses maps to the order of
the ports from the `ReadEndPoint.Ports` attribute.

`ReadEndPoint.Ports`

A string with a list of ports for the read-only replicas. The order of the ports maps to the order of the
addresses from the `ReadEndPoint.Addresses` attribute.

`ReadEndPoint.Ports.List`

A string with a list of ports for the read-only replicas. The order of the ports maps to the order of the
addresses from the ReadEndPoint.Addresses attribute.

`ReaderEndPoint.Address`

The address of the reader endpoint.

`ReaderEndPoint.Port`

The port used by the reader endpoint.

`ReplicationGroupId`

The ID of the replication group to which data is to be migrated.

## Examples

- [Declare a Replication Group with Two Nodes](#aws-resource-elasticache-replicationgroup--examples--Declare_a_Replication_Group_with_Two_Nodes)

- [Declare a Replication Group with Two Node Groups](#aws-resource-elasticache-replicationgroup--examples--Declare_a_Replication_Group_with_Two_Node_Groups)

### Declare a Replication Group with Two Nodes

The following example declares a replication group with two nodes and automatic failover enabled.

#### JSON

```json

{
    "myReplicationGroup": {
        "Type": "AWS::ElastiCache::ReplicationGroup",
        "Properties": {
            "ReplicationGroupDescription": "my description",
            "NumCacheClusters": "2",
            "Engine": "redis",
            "CacheNodeType": "cache.m3.medium",
            "AutomaticFailoverEnabled": "true",
            "CacheSubnetGroupName": "subnetgroup",
            "EngineVersion": "2.8.6",
            "PreferredMaintenanceWindow": "wed:09:25-wed:22:30",
            "SnapshotRetentionLimit": "4",
            "SnapshotWindow": "03:30-05:30"
        }
    }
}
```

#### YAML

```yaml

myReplicationGroup:
  Type: 'AWS::ElastiCache::ReplicationGroup'
  Properties:
    ReplicationGroupDescription: my description
    NumCacheClusters: '2'
    Engine: redis
    CacheNodeType: cache.m3.medium
    AutomaticFailoverEnabled: 'true'
    CacheSubnetGroupName: subnetgroup
    EngineVersion: 2.8.6
    PreferredMaintenanceWindow: 'wed:09:25-wed:22:30'
    SnapshotRetentionLimit: '4'
    SnapshotWindow: '03:30-05:30'
```

### Declare a Replication Group with Two Node Groups

The following example declares a replication group with two nodes groups (shards) with three replicas in each group.

#### JSON

```json

{
    "BasicReplicationGroup": {
        "Type": "AWS::ElastiCache::ReplicationGroup",
        "Properties": {
            "AutomaticFailoverEnabled": true,
            "CacheNodeType": "cache.r3.large",
            "CacheSubnetGroupName": {
                "Ref": "CacheSubnetGroup"
            },
            "Engine": "redis",
            "EngineVersion": "3.2",
            "NumNodeGroups": "2",
            "ReplicasPerNodeGroup": "3",
            "Port": 6379,
            "PreferredMaintenanceWindow": "sun:05:00-sun:09:00",
            "ReplicationGroupDescription": "A sample replication group",
            "SecurityGroupIds": [
                {
                    "Ref": "ReplicationGroupSG"
                }
            ],
            "SnapshotRetentionLimit": 5,
            "SnapshotWindow": "10:00-12:00"
        }
    }
}
```

#### YAML

```yaml

BasicReplicationGroup:
  Type: 'AWS::ElastiCache::ReplicationGroup'
  Properties:
    AutomaticFailoverEnabled: true
    CacheNodeType: cache.r3.large
    CacheSubnetGroupName: !Ref CacheSubnetGroup
    Engine: redis
    EngineVersion: '3.2'
    NumNodeGroups: '2'
    ReplicasPerNodeGroup: '3'
    Port: 6379
    PreferredMaintenanceWindow: 'sun:05:00-sun:09:00'
    ReplicationGroupDescription: A sample replication group
    SecurityGroupIds:
      - !Ref ReplicationGroupSG
    SnapshotRetentionLimit: 5
    SnapshotWindow: '10:00-12:00'
```

## See also

[CreateReplicationGroup](../../../../reference/amazonelasticache/latest/apireference/api-createreplicationgroup.md) in the _Amazon ElastiCache API Reference Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CloudWatchLogsDestinationDetails

All content copied from https://docs.aws.amazon.com/.
