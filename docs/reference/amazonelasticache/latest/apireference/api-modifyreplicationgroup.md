# ModifyReplicationGroup

Modifies the settings for a replication group. This is limited to Valkey and Redis OSS 7 and above.

- [Scaling for Valkey or Redis OSS (cluster mode enabled)](../../../../services/amazonelasticache/latest/dg/scaling-redis-cluster-mode-enabled.md) in
the ElastiCache User Guide

- [ModifyReplicationGroupShardConfiguration](api-modifyreplicationgroupshardconfiguration.md) in the ElastiCache API
Reference

###### Note

This operation is valid for Valkey or Redis OSS only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ReplicationGroupId**

The identifier of the replication group to modify.

Type: String

Required: Yes

**ApplyImmediately**

If `true`, this parameter causes the modifications in this request and any
pending modifications to be applied, asynchronously and as soon as possible, regardless
of the `PreferredMaintenanceWindow` setting for the replication group.

If `false`, changes to the nodes in the replication group are applied on
the next maintenance reboot, or the next failure reboot, whichever occurs first.

Valid values: `true` \| `false`

Default: `false`

Type: Boolean

Required: No

**AuthToken**

Reserved parameter. The password used to access a password protected server. This
parameter must be specified with the `auth-token-update-strategy ` parameter.
Password constraints:

- Must be only printable ASCII characters

- Must be at least 16 characters and no more than 128 characters in
length

- Cannot contain any of the following characters: '/', '"', or '@', '%'

For more information, see AUTH password at [AUTH](http://redis.io/commands/AUTH).

Type: String

Required: No

**AuthTokenUpdateStrategy**

Specifies the strategy to use to update the AUTH token. This parameter must be
specified with the `auth-token` parameter. Possible values:

- ROTATE - default, if no update strategy is provided

- SET - allowed only after ROTATE

- DELETE - allowed only when transitioning to RBAC

For more information, see [Authenticating Users with AUTH](../../../../services/amazonelasticache/latest/dg/auth.md)

Type: String

Valid Values: `SET | ROTATE | DELETE`

Required: No

**AutomaticFailoverEnabled**

Determines whether a read replica is automatically promoted to read/write primary if
the existing primary encounters a failure.

Valid values: `true` \| `false`

Type: Boolean

Required: No

**AutoMinorVersionUpgrade**

If you are running Valkey or Redis OSS engine version 6.0 or later, set this parameter to yes if
you want to opt-in to the next auto minor version upgrade campaign. This parameter is
disabled for previous versions.

Type: Boolean

Required: No

**CacheNodeType**

A valid cache node type that you want to scale this replication group to.

Type: String

Required: No

**CacheParameterGroupName**

The name of the cache parameter group to apply to all of the clusters in this
replication group. This change is asynchronously applied as soon as possible for
parameters when the `ApplyImmediately` parameter is specified as
`true` for this request.

Type: String

Required: No

**CacheSecurityGroupNames.CacheSecurityGroupName.N**

A list of cache security group names to authorize for the clusters in this replication
group. This change is asynchronously applied as soon as possible.

This parameter can be used only with replication group containing clusters running
outside of an Amazon Virtual Private Cloud (Amazon VPC).

Constraints: Must contain no more than 255 alphanumeric characters. Must not be
`Default`.

Type: Array of strings

Required: No

**ClusterMode**

Enabled or Disabled. To modify cluster mode from Disabled to Enabled, you must first
set the cluster mode to Compatible. Compatible mode allows your Valkey or Redis OSS clients to connect
using both cluster mode enabled and cluster mode disabled. After you migrate all Valkey or Redis OSS
clients to use cluster mode enabled, you can then complete cluster mode configuration
and set the cluster mode to Enabled.

Type: String

Valid Values: `enabled | disabled | compatible`

Required: No

**Engine**

Modifies the engine listed in a replication group message. The options are valkey, memcached or redis.

Type: String

Required: No

**EngineVersion**

The upgraded version of the cache engine to be run on the clusters in the replication
group.

**Important:** You can upgrade to a newer engine version
(see [Selecting\
a Cache Engine and Version](../../../../services/amazonelasticache/latest/dg/selectengine.md#VersionManagement)), but you cannot downgrade to an earlier engine
version. If you want to use an earlier engine version, you must delete the existing
replication group and create it anew with the earlier engine version.

Type: String

Required: No

**IpDiscovery**

The network type you choose when modifying a cluster, either `ipv4` \|
`ipv6`. IPv6 is supported for workloads using Valkey 7.2 and above, Redis OSS engine version 6.2
to 7.1 and Memcached engine version 1.6.6 and above on all instances built on the [Nitro system](http://aws.amazon.com/ec2/nitro).

Type: String

Valid Values: `ipv4 | ipv6`

Required: No

**LogDeliveryConfigurations.LogDeliveryConfigurationRequest.N**

Specifies the destination, format and type of the logs.

Type: Array of [LogDeliveryConfigurationRequest](api-logdeliveryconfigurationrequest.md) objects

Required: No

**MultiAZEnabled**

A flag to indicate MultiAZ is enabled.

Type: Boolean

Required: No

**NodeGroupId**

Deprecated. This parameter is not used.

Type: String

Required: No

**NotificationTopicArn**

The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications are
sent.

###### Note

The Amazon SNS topic owner must be same as the replication group owner.

Type: String

Required: No

**NotificationTopicStatus**

The status of the Amazon SNS notification topic for the replication group.
Notifications are sent only if the status is `active`.

Valid values: `active` \| `inactive`

Type: String

Required: No

**PreferredMaintenanceWindow**

Specifies the weekly time range during which maintenance on the cluster is performed.
It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The
minimum maintenance window is a 60 minute period.

Valid values for `ddd` are:

- `sun`

- `mon`

- `tue`

- `wed`

- `thu`

- `fri`

- `sat`

Example: `sun:23:00-mon:01:30`

Type: String

Required: No

**PrimaryClusterId**

For replication groups with a single primary, if this parameter is specified,
ElastiCache promotes the specified cluster in the specified replication group to the
primary role. The nodes of all other clusters in the replication group are read
replicas.

Type: String

Required: No

**RemoveUserGroups**

Removes the user group associated with this replication group.

Type: Boolean

Required: No

**ReplicationGroupDescription**

A description for the replication group. Maximum length is 255 characters.

Type: String

Required: No

**SecurityGroupIds.SecurityGroupId.N**

Specifies the VPC Security Groups associated with the clusters in the replication
group.

This parameter can be used only with replication group containing clusters running in
an Amazon Virtual Private Cloud (Amazon VPC).

Type: Array of strings

Required: No

**SnapshotRetentionLimit**

The number of days for which ElastiCache retains automatic node group (shard)
snapshots before deleting them. For example, if you set
`SnapshotRetentionLimit` to 5, a snapshot that was taken today is
retained for 5 days before being deleted.

**Important** If the value of SnapshotRetentionLimit is
set to zero (0), backups are turned off.

Type: Integer

Required: No

**SnapshottingClusterId**

The cluster ID that is used as the daily snapshot source for the replication group.
This parameter cannot be set for Valkey or Redis OSS (cluster mode enabled) replication groups.

Type: String

Required: No

**SnapshotWindow**

The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot
of the node group (shard) specified by `SnapshottingClusterId`.

Example: `05:00-09:00`

If you do not specify this parameter, ElastiCache automatically chooses an appropriate
time range.

Type: String

Required: No

**TransitEncryptionEnabled**

A flag that enables in-transit encryption when set to true. If you are enabling
in-transit encryption for an existing cluster, you must also set
`TransitEncryptionMode` to `preferred`.

Type: Boolean

Required: No

**TransitEncryptionMode**

A setting that allows you to migrate your clients to use in-transit encryption, with
no downtime.

You must set `TransitEncryptionEnabled` to `true`, for your
existing cluster, and set `TransitEncryptionMode` to `preferred`
in the same request to allow both encrypted and unencrypted connections at the same
time. Once you migrate all your Valkey or Redis OSS clients to use encrypted connections you can set
the value to `required` to allow encrypted connections only.

Setting `TransitEncryptionMode` to `required` is a two-step
process that requires you to first set the `TransitEncryptionMode` to
`preferred`, after that you can set `TransitEncryptionMode` to
`required`.

Type: String

Valid Values: `preferred | required`

Required: No

**UserGroupIdsToAdd.member.N**

The ID of the user group you are associating with the replication group.

Type: Array of strings

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: No

**UserGroupIdsToRemove.member.N**

The ID of the user group to disassociate from the replication group, meaning the users
in the group no longer can access the replication group.

Type: Array of strings

Length Constraints: Minimum length of 1.

Pattern: `[a-zA-Z][a-zA-Z0-9\-]*`

Required: No

## Response Elements

The following element is returned by the service.

**ReplicationGroup**

Contains all of the attributes of a specific Valkey or Redis OSS replication group.

Type: [ReplicationGroup](api-replicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheClusterNotFound**

The requested cluster ID does not refer to an existing cluster.

HTTP Status Code: 404

**CacheParameterGroupNotFound**

The requested cache parameter group name does not refer to an existing cache parameter
group.

HTTP Status Code: 404

**CacheSecurityGroupNotFound**

The requested cache security group name does not refer to an existing cache security
group.

HTTP Status Code: 404

**InsufficientCacheClusterCapacity**

The requested cache node type is not available in the specified Availability Zone. For
more information, see [InsufficientCacheClusterCapacity](../../../../services/amazonelasticache/latest/dg/errormessages.md#ErrorMessages.INSUFFICIENT_CACHE_CLUSTER_CAPACITY) in the ElastiCache User Guide.

HTTP Status Code: 400

**InvalidCacheClusterState**

The requested cluster is not in the `available` state.

HTTP Status Code: 400

**InvalidCacheSecurityGroupState**

The current state of the cache security group does not allow deletion.

HTTP Status Code: 400

**InvalidKMSKeyFault**

The KMS key supplied is not valid.

HTTP Status Code: 400

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidReplicationGroupState**

The requested replication group is not in the `available` state.

HTTP Status Code: 400

**InvalidUserGroupState**

The user group is not in an active state.

HTTP Status Code: 400

**InvalidVPCNetworkStateFault**

The VPC network is in an invalid state.

HTTP Status Code: 400

**NodeQuotaForClusterExceeded**

The request cannot be processed because it would exceed the allowed number of cache
nodes in a single cluster.

HTTP Status Code: 400

**NodeQuotaForCustomerExceeded**

The request cannot be processed because it would exceed the allowed number of cache
nodes per customer.

HTTP Status Code: 400

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

**UserGroupNotFound**

The user group was not found or does not exist

HTTP Status Code: 404

## Examples

### ModifyReplicationGroup

This example illustrates one usage of ModifyReplicationGroup.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=ModifyReplicationGroup
   &ApplyImmediately=false
   &ReplicationGroupId=my-repgroup
   &PrimaryClusterId=my-replica-1
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<ModifyReplicationGroupResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <ModifyReplicationGroupResult>
      <ReplicationGroup>
         <SnapshottingClusterId>my-redis-primary</SnapshottingClusterId>
         <MemberClusters>
            <ClusterId>my-redis-primary</ClusterId>
            <ClusterId>my-replica-1</ClusterId>
         </MemberClusters>
         <NodeGroups>
            <NodeGroup>
               <NodeGroupId>0001</NodeGroupId>
               <PrimaryEndpoint>
                  <Port>6379</Port>
                  <Address>my-repgroup.q68zge.ng.0001.use1devo.elmo-dev.amazonaws.com</Address>
               </PrimaryEndpoint>
               <Status>available</Status>
               <NodeGroupMembers>
                  <NodeGroupMember>
                     <CacheClusterId>my-redis-primary</CacheClusterId>
                     <ReadEndpoint>
                        <Port>6379</Port>
                        <Address>my-redis-primary.q68zge.0001.use1devo.elmo-dev.amazonaws.com</Address>
                     </ReadEndpoint>
                     <PreferredAvailabilityZone>us-west-2c</PreferredAvailabilityZone>
                     <CacheNodeId>0001</CacheNodeId>
                     <CurrentRole>primary</CurrentRole>
                  </NodeGroupMember>
                  <NodeGroupMember>
                     <CacheClusterId>my-replica-1</CacheClusterId>
                     <ReadEndpoint>
                        <Port>6379</Port>
                        <Address>my-replica-1.q68zge.0001.use1devo.elmo-dev.amazonaws.com</Address>
                     </ReadEndpoint>
                     <PreferredAvailabilityZone>us-west-2b</PreferredAvailabilityZone>
                     <CacheNodeId>0001</CacheNodeId>
                     <CurrentRole>replica</CurrentRole>
                  </NodeGroupMember>
               </NodeGroupMembers>
            </NodeGroup>
         </NodeGroups>
         <ReplicationGroupId>my-repgroup</ReplicationGroupId>
         <Status>available</Status>
         <PendingModifiedValues>
            <PrimaryClusterId>my-replica-1</PrimaryClusterId>
         </PendingModifiedValues>
         <Description>My replication group</Description>
      </ReplicationGroup>
   </ModifyReplicationGroupResult>
   <ResponseMetadata>
      <RequestId>6fd0aad6-b9d7-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</ModifyReplicationGroupResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/modifyreplicationgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/modifyreplicationgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyGlobalReplicationGroup

ModifyReplicationGroupShardConfiguration

All content copied from https://docs.aws.amazon.com/.
