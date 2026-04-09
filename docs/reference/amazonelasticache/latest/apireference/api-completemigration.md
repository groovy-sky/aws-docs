# CompleteMigration

Complete the migration of data.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ReplicationGroupId**

The ID of the replication group to which data is being migrated.

Type: String

Required: Yes

**Force**

Forces the migration to stop without ensuring that data is in sync. It is recommended
to use this option only to abort the migration and not recommended when application
wants to continue migration to ElastiCache.

Type: Boolean

Required: No

## Response Elements

The following element is returned by the service.

**ReplicationGroup**

Contains all of the attributes of a specific Valkey or Redis OSS replication group.

Type: [ReplicationGroup](api-replicationgroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidReplicationGroupState**

The requested replication group is not in the `available` state.

HTTP Status Code: 400

**ReplicationGroupNotFoundFault**

The specified replication group does not exist.

HTTP Status Code: 404

**ReplicationGroupNotUnderMigrationFault**

The designated replication group is not available for data migration.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/completemigration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/completemigration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchStopUpdateAction

CopyServerlessCacheSnapshot

All content copied from https://docs.aws.amazon.com/.
