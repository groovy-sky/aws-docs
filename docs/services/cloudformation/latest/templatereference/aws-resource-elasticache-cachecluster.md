This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::CacheCluster

The `AWS::ElastiCache::CacheCluster` type creates an Amazon ElastiCache cache cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::CacheCluster",
  "Properties" : {
      "AutoMinorVersionUpgrade" : Boolean,
      "AZMode" : String,
      "CacheNodeType" : String,
      "CacheParameterGroupName" : String,
      "CacheSecurityGroupNames" : [ String, ... ],
      "CacheSubnetGroupName" : String,
      "ClusterName" : String,
      "Engine" : String,
      "EngineVersion" : String,
      "IpDiscovery" : String,
      "LogDeliveryConfigurations" : [ LogDeliveryConfigurationRequest, ... ],
      "NetworkType" : String,
      "NotificationTopicArn" : String,
      "NumCacheNodes" : Integer,
      "Port" : Integer,
      "PreferredAvailabilityZone" : String,
      "PreferredAvailabilityZones" : [ String, ... ],
      "PreferredMaintenanceWindow" : String,
      "SnapshotArns" : [ String, ... ],
      "SnapshotName" : String,
      "SnapshotRetentionLimit" : Integer,
      "SnapshotWindow" : String,
      "Tags" : [ Tag, ... ],
      "TransitEncryptionEnabled" : Boolean,
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::CacheCluster
Properties:
  AutoMinorVersionUpgrade: Boolean
  AZMode: String
  CacheNodeType: String
  CacheParameterGroupName: String
  CacheSecurityGroupNames:
    - String
  CacheSubnetGroupName: String
  ClusterName: String
  Engine: String
  EngineVersion: String
  IpDiscovery: String
  LogDeliveryConfigurations:
    - LogDeliveryConfigurationRequest
  NetworkType: String
  NotificationTopicArn: String
  NumCacheNodes: Integer
  Port: Integer
  PreferredAvailabilityZone: String
  PreferredAvailabilityZones:
    - String
  PreferredMaintenanceWindow: String
  SnapshotArns:
    - String
  SnapshotName: String
  SnapshotRetentionLimit: Integer
  SnapshotWindow: String
  Tags:
    - Tag
  TransitEncryptionEnabled: Boolean
  VpcSecurityGroupIds:
    - String

```

## Properties

`AutoMinorVersionUpgrade`

If you are running Valkey 7.2 or later, or Redis OSS engine version 6.0 or later, set this parameter to yes if you want to opt-in to the
next minor version upgrade campaign. This parameter is disabled for previous versions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AZMode`

Specifies whether the nodes in this Memcached cluster are created in a single
Availability Zone or created across multiple Availability Zones in the cluster's
region.

This parameter is only supported for Memcached clusters.

If the `AZMode` and `PreferredAvailabilityZones` are not
specified, ElastiCache assumes `single-az` mode.

_Required_: No

_Type_: String

_Allowed values_: `single-az | cross-az`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheNodeType`

The compute and memory capacity of the nodes in the node group (shard).

The following node types are supported by ElastiCache. Generally speaking, the current generation types provide
more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.
Changing the CacheNodeType of a Memcached instance is currently not supported. If you need to scale using Memcached,
we recommend forcing a replacement update by changing the `LogicalResourceId` of the resource.

- General purpose:

- Current generation:

**M6g node types:** `cache.m6g.large`,
`cache.m6g.xlarge`, `cache.m6g.2xlarge`, `cache.m6g.4xlarge`,
`cache.m6g.8xlarge`, `cache.m6g.12xlarge`, `cache.m6g.16xlarge`,
`cache.m6g.24xlarge`

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
`cache.r6g.8xlarge`, `cache.r6g.12xlarge`, `cache.r6g.16xlarge`,
`cache.r6g.24xlarge`

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
Types by Region](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)

**Additional node type info**

- All current generation instance types are created in Amazon VPC by default.

- Valkey and Redis OSS append-only files (AOF) are not supported for T1 or T2 instances.

- Valkey and Redis OSS Multi-AZ with automatic failover is not supported on T1 instances.

- Redis OSS configuration variables `appendonly` and `appendfsync` are not supported on Redis OSS
version 2.8.22 and later.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheParameterGroupName`

The name of the parameter group to associate with this cluster. If this argument is omitted, the default
parameter group for the specified engine is used. You cannot use any parameter group which has
`cluster-enabled='yes'` when creating a cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheSecurityGroupNames`

A list of security group names to associate with this cluster.

Use this parameter only when you are creating a cluster outside of an Amazon Virtual Private Cloud (Amazon
VPC).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheSubnetGroupName`

The name of the subnet group to be used for the cluster.

Use this parameter only when you are creating a cluster in an Amazon Virtual Private Cloud (Amazon VPC).

###### Important

If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start
creating a cluster. For more information, see `AWS::ElastiCache::SubnetGroup.`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterName`

A name for the cache cluster. If you don't specify a name, CloudFormation generates a unique physical ID and
uses that ID for the cache cluster. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

The name must contain 1 to 50 alphanumeric characters or hyphens. The name must start with a letter and cannot
end with a hyphen or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Engine`

The name of the cache engine to be used for this cluster.

Valid values for this parameter are: `memcached` \| valkey \|
`redis`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The version number of the cache engine to be used for this cluster. To view the supported cache engine versions,
use the DescribeCacheEngineVersions operation.

**Important:** You can upgrade to a newer engine version (see [Selecting a Cache\
Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/SelectEngine.html#VersionManagement)), but you cannot downgrade to an earlier engine version. If you want to use an earlier
engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine
version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpDiscovery`

The network type you choose when modifying a cluster, either `ipv4` \|
`ipv6`. IPv6 is supported for workloads using Valkey 7.2 and above, Redis OSS engine version 6.2
to 7.1 and Memcached engine version 1.6.6 and above on all instances built on the [Nitro system](https://aws.amazon.com/ec2/nitro).

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogDeliveryConfigurations`

Specifies the destination, format and type of the logs.

_Required_: No

_Type_: Array of [LogDeliveryConfigurationRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-cachecluster-logdeliveryconfigurationrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

Must be either `ipv4` \| `ipv6` \| `dual_stack`. IPv6
is supported for workloads using Valkey 7.2 and above, Redis OSS engine version 6.2 to 7.1
and Memcached engine version 1.6.6 and above on all instances built on the [Nitro system](https://aws.amazon.com/ec2/nitro).

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dual_stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationTopicArn`

The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic
to which notifications are sent.

###### Note

The Amazon SNS topic owner must be the same as the cluster owner.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumCacheNodes`

The number of cache nodes that the cache cluster should have.

###### Note

However, if the `PreferredAvailabilityZone` and `PreferredAvailabilityZones ` properties
were not previously specified and you don't specify any new values, an update requires [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number on which each of the cache nodes accepts connections.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredAvailabilityZone`

The EC2 Availability Zone in which the cluster is created.

All nodes belonging to this cluster are placed in the preferred Availability Zone. If
you want to create your nodes across multiple Availability Zones, use
`PreferredAvailabilityZones`.

Default: System chosen Availability Zone.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredAvailabilityZones`

A list of the Availability Zones in which cache nodes are created. The order of the
zones in the list is not important.

This option is only supported on Memcached.

###### Note

If you are creating your cluster in an Amazon VPC (recommended) you can only
locate nodes in Availability Zones that are associated with the subnets in the
selected subnet group.

The number of Availability Zones listed must equal the value of
`NumCacheNodes`.

If you want all the nodes in the same Availability Zone, use
`PreferredAvailabilityZone` instead, or repeat the Availability Zone
multiple times in the list.

Default: System chosen Availability Zones.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

`SnapshotArns`

A single-element string list containing an Amazon Resource Name (ARN) that uniquely identifies a Valkey or Redis OSS RDB
snapshot file stored in Amazon S3. The snapshot file is used to populate the node group (shard). The Amazon S3 object
name in the ARN cannot contain any commas.

###### Note

This parameter is only valid if the `Engine` parameter is `redis`.

Example of an Amazon S3 ARN: `arn:aws:s3:::my_bucket/snapshot1.rdb`

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotName`

The name of a Valkey or Redis OSS snapshot from which to restore data into the new node group
(shard). The snapshot status changes to `restoring` while the new node group
(shard) is being created.

###### Note

This parameter is only valid if the `Engine` parameter is
`redis`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotRetentionLimit`

The number of days for which ElastiCache retains automatic snapshots before deleting
them. For example, if you set `SnapshotRetentionLimit` to 5, a snapshot taken
today is retained for 5 days before being deleted.

###### Note

This parameter is only valid if the `Engine` parameter is
`redis`.

Default: 0 (i.e., automatic backups are disabled for this cache cluster).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotWindow`

The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot
of your node group (shard).

Example: `05:00-09:00`

If you do not specify this parameter, ElastiCache automatically chooses an appropriate
time range.

###### Note

This parameter is only valid if the `Engine` parameter is
`redis`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to be added to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-cachecluster-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryptionEnabled`

A flag that enables in-transit encryption when set to true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

One or more VPC security groups associated with the cluster.

Use this parameter only when you are creating a cluster in an Amazon Virtual Private Cloud (Amazon VPC).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the resource name.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConfigurationEndpoint.Address`

The DNS hostname of the cache node.

###### Note

Valkey and Redis OSS (cluster mode disabled) replication groups don't have this attribute. Therefore, `Fn::GetAtt`
returns a value for this attribute only if the replication group is clustered. Otherwise, `Fn::GetAtt`
fails.

`ConfigurationEndpoint.Port`

The port number of the configuration endpoint for the Memcached cache cluster.

###### Note

Valkey and Redis OSS (cluster mode disabled) replication groups don't have this attribute. Therefore, `Fn::GetAtt`
returns a value for this attribute only if the replication group is clustered. Otherwise, `Fn::GetAtt`
fails.

`Id`

The resource name.

`RedisEndpoint.Address`

The DNS address of the configuration endpoint for the Valkey or Redis OSS cache cluster.

`RedisEndpoint.Port`

The port number of the configuration endpoint for the Valkey or Redis OSS cache cluster.

## Examples

- [Cluster in a Default VPC](#aws-resource-elasticache-cachecluster--examples--Cluster_in_a_Default_VPC)

- [Memcached Nodes in Multiple Availability Zones](#aws-resource-elasticache-cachecluster--examples--Memcached_Nodes_in_Multiple_Availability_Zones)

### Cluster in a Default VPC

The following snippet describes an Elasticache cluster in a security group that is in a
[default VPC](https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html). Usually, a security
group in a VPC requires the VPC ID to be specified. In this case, no VPC ID is needed
because the security group uses the default VPC. If you want to specify a VPC for the
security group, specify its `VpcId` property.

For the cache cluster, the `VpcSecurityGroupIds` property is used to
associate the cluster with the security group. Because the
`VpcSecurityGroupIds` property requires security group IDs (not security
group names), the template snippet uses the `Fn::GetAtt` function instead of
a `Ref` function on the `ElasticacheSecurityGroup` resource. The
`Ref` function will return the security group name. If you specify a VPC
ID for the security group, `Ref` returns the security group ID.

Note that `InstanceSecurityGroup` refers to the logical name of a security
group that is not actually defined in this example. To learn more about the
`SourceSecurityGroupName`
property, see [AWS::EC2::SecurityGroupIngress](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html).

#### JSON

```json

{
    "ElasticacheSecurityGroup": {
        "Type": "AWS::EC2::SecurityGroup",
        "Properties": {
            "GroupDescription": "Elasticache Security Group",
            "SecurityGroupIngress": [
                {
                    "IpProtocol": "tcp",
                    "FromPort": 11211,
                    "ToPort": 11211,
                    "SourceSecurityGroupName": {
                        "Ref": "InstanceSecurityGroup"
                    }
                }
            ]
        }
    },
    "ElasticacheCluster": {
        "Type": "AWS::ElastiCache::CacheCluster",
        "Properties": {
            "Engine": "memcached",
            "CacheNodeType": "cache.t2.micro",
            "NumCacheNodes": "1",
            "VpcSecurityGroupIds": [
                {
                    "Fn::GetAtt": [
                        "ElasticacheSecurityGroup",
                        "GroupId"
                    ]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

ElasticacheSecurityGroup:
  Type: 'AWS::EC2::SecurityGroup'
  Properties:
    GroupDescription: Elasticache Security Group
    SecurityGroupIngress:
      - IpProtocol: tcp
        FromPort: 11211
        ToPort: 11211
        SourceSecurityGroupName: !Ref InstanceSecurityGroup
ElasticacheCluster:
  Type: 'AWS::ElastiCache::CacheCluster'
  Properties:
    Engine: memcached
    CacheNodeType: cache.t2.micro
    NumCacheNodes: '1'
    VpcSecurityGroupIds:
      - !GetAtt
        - ElasticacheSecurityGroup
        - GroupId
```

### Memcached Nodes in Multiple Availability Zones

The following example launches a cache cluster with three nodes, where two nodes are created in us-west-2a and one is created in us-west-2b.

#### JSON

```json

{
    "myCacheCluster": {
        "Type": "AWS::ElastiCache::CacheCluster",
        "Properties": {
            "AZMode": "cross-az",
            "CacheNodeType": "cache.m3.medium",
            "Engine": "memcached",
            "NumCacheNodes": "3",
            "PreferredAvailabilityZones": [
                "us-west-2a",
                "us-west-2a",
                "us-west-2b"
            ]
        }
    }
}
```

#### YAML

```yaml

myCacheCluster:
  Type: 'AWS::ElastiCache::CacheCluster'
  Properties:
    AZMode: cross-az
    CacheNodeType: cache.m3.medium
    Engine: memcached
    NumCacheNodes: '3'
    PreferredAvailabilityZones:
      - us-west-2a
      - us-west-2a
      - us-west-2b
```

## See also

- [CreateCacheParameterGroup](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheParameterGroup.html) in the _Amazon ElastiCache API Reference Guide_

- [ModifyCacheParameterGroup](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html) in the _Amazon ElastiCache API Reference Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon ElastiCache

CloudWatchLogsDestinationDetails
