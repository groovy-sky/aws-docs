# RebootCacheCluster

Reboots some, or all, of the cache nodes within a provisioned cluster. This operation
applies any modified cache parameter groups to the cluster. The reboot operation takes
place as soon as possible, and results in a momentary outage to the cluster. During the
reboot, the cluster status is set to REBOOTING.

The reboot causes the contents of the cache (for each cache node being rebooted) to be
lost.

When the reboot is complete, a cluster event is created.

Rebooting a cluster is currently supported on Memcached, Valkey and Redis OSS (cluster mode
disabled) clusters. Rebooting is not supported on Valkey or Redis OSS (cluster mode enabled)
clusters.

If you make changes to parameters that require a Valkey or Redis OSS (cluster mode enabled) cluster
reboot for the changes to be applied, see [Rebooting a Cluster](../../../../services/amazonelasticache/latest/dg/nodes-rebooting.md) for an alternate process.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CacheClusterId**

The cluster identifier. This parameter is stored as a lowercase string.

Type: String

Required: Yes

**CacheNodeIdsToReboot.CacheNodeId.N**

A list of cache node IDs to reboot. A node ID is a numeric identifier (0001, 0002,
etc.). To reboot an entire cluster, specify all of the cache node IDs.

Type: Array of strings

Required: Yes

## Response Elements

The following element is returned by the service.

**CacheCluster**

Contains all of the attributes of a specific cluster.

Type: [CacheCluster](api-cachecluster.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CacheClusterNotFound**

The requested cluster ID does not refer to an existing cluster.

HTTP Status Code: 404

**InvalidCacheClusterState**

The requested cluster is not in the `available` state.

HTTP Status Code: 400

## Examples

### RebootCacheCluster

This example illustrates one usage of RebootCacheCluster.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=RebootCacheCluster
   &CacheClusterId=mycache
   &CacheNodeIdsToReboot.CacheNodeId.1=0001
   &CacheNodeIdsToReboot.CacheNodeId.2=0002
   &CacheNodeIdsToReboot.CacheNodeId.3=0003
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>
```

#### Sample Response

```

<RebootCacheClusterResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
   <RebootCacheClusterResult>
      <CacheCluster>
         <CacheClusterStatus>rebooting cluster nodes</CacheClusterStatus>
         <CacheParameterGroup>
            <CacheParameterGroupName>default.memcached1.4</CacheParameterGroupName>
            <ParameterApplyStatus>in-sync</ParameterApplyStatus>
            <CacheNodeIdsToReboot />
         </CacheParameterGroup>
         <CacheClusterId>mycache</CacheClusterId>
         <ConfigurationEndpoint>
            <Port>11211</Port>
            <Address>mycache.q68zge.cfg.use1devo.elmo-dev.amazonaws.com</Address>
         </ConfigurationEndpoint>
         <CacheNodeType>cache.m1.small</CacheNodeType>
         <Engine>memcached</Engine>
         <PendingModifiedValues />
         <PreferredAvailabilityZone>us-west-2b</PreferredAvailabilityZone>
         <CacheClusterCreateTime>2015-02-02T19:04:12.812Z</CacheClusterCreateTime>
         <EngineVersion>1.4.17</EngineVersion>
         <AutoMinorVersionUpgrade>true</AutoMinorVersionUpgrade>
         <PreferredMaintenanceWindow>wed:09:00-wed:10:00</PreferredMaintenanceWindow>
         <ClientDownloadLandingPage>https://console.aws.amazon.com/elasticache/home#client-download:</ClientDownloadLandingPage>
         <CacheSecurityGroups>
            <CacheSecurityGroup>
               <CacheSecurityGroupName>default</CacheSecurityGroupName>
               <Status>active</Status>
            </CacheSecurityGroup>
         </CacheSecurityGroups>
         <NumCacheNodes>3</NumCacheNodes>
      </CacheCluster>
   </RebootCacheClusterResult>
   <ResponseMetadata>
      <RequestId>cf7e6fc4-b9d1-11e3-8a16-7978bb24ffdf</RequestId>
   </ResponseMetadata>
</RebootCacheClusterResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/rebootcachecluster.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/rebootcachecluster.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RebalanceSlotsInGlobalReplicationGroup

RemoveTagsFromResource

All content copied from https://docs.aws.amazon.com/.
