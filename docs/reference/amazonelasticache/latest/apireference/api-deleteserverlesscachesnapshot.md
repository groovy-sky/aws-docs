# DeleteServerlessCacheSnapshot

Deletes an existing serverless cache snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ServerlessCacheSnapshotName**

Idenfitier of the snapshot to be deleted. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**ServerlessCacheSnapshot**

The snapshot to be deleted. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: [ServerlessCacheSnapshot](api-serverlesscachesnapshot.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidServerlessCacheSnapshotStateFault**

The state of the serverless cache snapshot was not received. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 400

**ServerlessCacheSnapshotNotFoundFault**

This serverless cache snapshot could not be found or does not exist. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 404

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deleteserverlesscachesnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteServerlessCache

DeleteSnapshot

All content copied from https://docs.aws.amazon.com/.
