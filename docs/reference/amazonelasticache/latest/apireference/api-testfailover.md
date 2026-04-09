# TestFailover

Represents the input of a `TestFailover` operation which tests automatic
failover on a specified node group (called shard in the console) in a replication group
(called cluster in the console).

This API is designed for testing the behavior of your application in case of
ElastiCache failover. It is not designed to be an operational tool for initiating a
failover to overcome a problem you may have with the cluster. Moreover, in certain
conditions such as large-scale operational events, Amazon may block this API.

###### Note the following

- A customer can use this operation to test automatic failover on up to 15 shards
(called node groups in the ElastiCache API and Amazon CLI) in any rolling
24-hour period.

- If calling this operation on shards in different clusters (called replication
groups in the API and CLI), the calls can be made concurrently.

- If calling this operation multiple times on different shards in the same Valkey or Redis OSS (cluster mode enabled) replication group, the first node replacement must
complete before a subsequent call can be made.

- To determine whether the node replacement is complete you can check Events
using the Amazon ElastiCache console, the Amazon CLI, or the ElastiCache API.
Look for the following automatic failover related events, listed here in order
of occurrance:

1. Replication group message: `Test Failover API called for node
                                   group <node-group-id>`

2. Cache cluster message: `Failover from primary node
                                   <primary-node-id> to replica node <node-id>
                                   completed`

3. Replication group message: `Failover from primary node
                                   <primary-node-id> to replica node <node-id>
                                   completed`

4. Cache cluster message: `Recovering cache nodes
                                   <node-id>`

5. Cache cluster message: `Finished recovery for cache nodes
                                   <node-id>`

For more information see:

- [Viewing\
ElastiCache Events](../../../../services/amazonelasticache/latest/dg/ecevents-viewing.md) in the _ElastiCache User_
_Guide_

- [DescribeEvents](api-describeevents.md) in the ElastiCache API Reference

Also see, [Testing\
Multi-AZ](../../../../services/amazonelasticache/latest/dg/autofailover.md#auto-failover-test) in the _ElastiCache User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**NodeGroupId**

The name of the node group (called shard in the console) in this replication group on
which automatic failover is to be tested. You may test automatic failover on up to 15
node groups in any rolling 24-hour period.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4.

Pattern: `\d+`

Required: Yes

**ReplicationGroupId**

The name of the replication group (console: cluster) whose automatic failover is being
tested by this operation.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**ReplicationGroup**

Contains all of the attributes of a specific Valkey or Redis OSS replication group.

Type: [ReplicationGroup](api-replicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**APICallRateForCustomerExceeded**

The customer has exceeded the allowed rate of API calls.

HTTP Status Code: 400

**InvalidCacheClusterState**

The requested cluster is not in the `available` state.

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

**NodeGroupNotFoundFault**

The node group specified by the `NodeGroupId` parameter could not be found.
Please verify that the node group exists and that you spelled the
`NodeGroupId` value correctly.

HTTP Status Code: 404

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

**TestFailoverNotAvailableFault**

The `TestFailover` action is not available.

HTTP Status Code: 400

## Examples

### Example

The following example tests automatic failover on the Valkey or Redis OSS (cluster mode
disabled) replication group (console: cluster) `redis00`. Since there
is only one node group in Valkey or Redis OSS (cluster mode disabled) clusters, the
_NodeGroupId_ will always be
`<cluster-name>-0001`.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=TestFailover
   &NodeGroupId=redis00-0001
   &ReplicationGroupId=redis00
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20170418T192317Z
   &X-Amz-Credential=<credential>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/testfailover.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/testfailover.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartMigration

TestMigration

All content copied from https://docs.aws.amazon.com/.
