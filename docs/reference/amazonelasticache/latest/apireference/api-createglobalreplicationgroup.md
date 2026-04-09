# CreateGlobalReplicationGroup

Global Datastore offers fully managed, fast, reliable and secure
cross-region replication. Using Global Datastore with Valkey or Redis OSS, you can create cross-region
read replica clusters for ElastiCache to enable low-latency reads and disaster
recovery across regions. For more information, see [Replication\
Across Regions Using Global Datastore](../../../../services/amazonelasticache/latest/dg/redis-global-datastore.md).

- The **GlobalReplicationGroupIdSuffix** is the
name of the Global datastore.

- The **PrimaryReplicationGroupId** represents the
name of the primary cluster that accepts writes and will replicate updates to
the secondary cluster.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**GlobalReplicationGroupIdSuffix**

The suffix name of a Global datastore. Amazon ElastiCache automatically applies a
prefix to the Global datastore ID when it is created. Each Amazon Region has its own
prefix. For instance, a Global datastore ID created in the US-West-1 region will begin
with "dsdfu" along with the suffix name you provide. The suffix, combined with the
auto-generated prefix, guarantees uniqueness of the Global datastore name across
multiple regions.

For a full list of Amazon Regions and their respective Global datastore iD prefixes,
see [Using the Amazon CLI with Global datastores](../../../../services/amazonelasticache/latest/dg/redis-global-datastores-cli.md).

Type: String

Required: Yes

**PrimaryReplicationGroupId**

The name of the primary cluster that accepts writes and will replicate updates to the
secondary cluster. This value is stored as a lowercase string.

Type: String

Required: Yes

**GlobalReplicationGroupDescription**

Provides details of the Global datastore

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**GlobalReplicationGroup**

Consists of a primary cluster that accepts writes and an associated secondary cluster
that resides in a different Amazon region. The secondary cluster accepts only reads. The
primary cluster automatically replicates updates to the secondary cluster.

- The **GlobalReplicationGroupIdSuffix** represents
the name of the Global datastore, which is what you use to associate a secondary
cluster.

Type: [GlobalReplicationGroup](api-globalreplicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**GlobalReplicationGroupAlreadyExistsFault**

The Global datastore name already exists.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidReplicationGroupState**

The requested replication group is not in the `available` state.

HTTP Status Code: 400

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/createglobalreplicationgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/createglobalreplicationgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateCacheSubnetGroup

CreateReplicationGroup

All content copied from https://docs.aws.amazon.com/.
