# BatchStopUpdateAction

Stop the service update. For more information on service updates and stopping them,
see [Stopping\
Service Updates](../../../../services/amazonelasticache/latest/dg/stopping-self-service-updates.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ServiceUpdateName**

The unique ID of the service update

Type: String

Required: Yes

**CacheClusterIds.member.N**

The cache cluster IDs

Type: Array of strings

Array Members: Maximum number of 20 items.

Required: No

**ReplicationGroupIds.member.N**

The replication group IDs

Type: Array of strings

Array Members: Maximum number of 20 items.

Required: No

## Response Elements

The following elements are returned by the service.

**ProcessedUpdateActions.ProcessedUpdateAction.N**

Update actions that have been processed successfully

Type: Array of [ProcessedUpdateAction](api-processedupdateaction.md) objects

**UnprocessedUpdateActions.UnprocessedUpdateAction.N**

Update actions that haven't been processed successfully

Type: Array of [UnprocessedUpdateAction](api-unprocessedupdateaction.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**ServiceUpdateNotFoundFault**

The service update doesn't exist

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/batchstopupdateaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/batchstopupdateaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchApplyUpdateAction

CompleteMigration

All content copied from https://docs.aws.amazon.com/.
