# StartMigration

Start the migration of data.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CustomerNodeEndpointList.member.N**

List of endpoints from which data should be migrated. For Valkey or Redis OSS (cluster mode
disabled), the list should have only one element.

Type: Array of [CustomerNodeEndpoint](api-customernodeendpoint.md) objects

Required: Yes

**ReplicationGroupId**

The ID of the replication group to which data should be migrated.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**ReplicationGroup**

Contains all of the attributes of a specific Valkey or Redis OSS replication group.

Type: [ReplicationGroup](api-replicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidReplicationGroupState**

The requested replication group is not in the `available` state.

HTTP Status Code: 400

**ReplicationGroupAlreadyUnderMigrationFault**

The targeted replication group is not available.

HTTP Status Code: 400

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/startmigration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/startmigration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RevokeCacheSecurityGroupIngress

TestFailover

All content copied from https://docs.aws.amazon.com/.
