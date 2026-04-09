# ModifyGlobalReplicationGroup

Modifies the settings for a Global datastore.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ApplyImmediately**

This parameter causes the modifications in this request and any pending modifications
to be applied, asynchronously and as soon as possible. Modifications to Global
Replication Groups cannot be requested to be applied in PreferredMaintenceWindow.

Type: Boolean

Required: Yes

**GlobalReplicationGroupId**

The name of the Global datastore

Type: String

Required: Yes

**AutomaticFailoverEnabled**

Determines whether a read replica is automatically promoted to read/write primary if
the existing primary encounters a failure.

Type: Boolean

Required: No

**CacheNodeType**

A valid cache node type that you want to scale this Global datastore to.

Type: String

Required: No

**CacheParameterGroupName**

The name of the cache parameter group to use with the Global datastore. It must be
compatible with the major engine version used by the Global datastore.

Type: String

Required: No

**Engine**

Modifies the engine listed in a global replication group message. The options are valkey, memcached or redis.

Type: String

Required: No

**EngineVersion**

The upgraded version of the cache engine to be run on the clusters in the Global
datastore.

Type: String

Required: No

**GlobalReplicationGroupDescription**

A description of the Global datastore

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

**GlobalReplicationGroupNotFoundFault**

The Global datastore does not exist

HTTP Status Code: 404

**InvalidGlobalReplicationGroupState**

The Global datastore is not available or in primary-only state.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/modifyglobalreplicationgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyCacheSubnetGroup

ModifyReplicationGroup

All content copied from https://docs.aws.amazon.com/.
