# CopyServerlessCacheSnapshot

Creates a copy of an existing serverless cache’s snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceServerlessCacheSnapshotName**

The identifier of the existing serverless cache’s snapshot to be copied. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: Yes

**TargetServerlessCacheSnapshotName**

The identifier for the snapshot to be created. Available for Valkey, Redis OSS and Serverless Memcached only. This value is stored as a lowercase string.

Type: String

Required: Yes

**KmsKeyId**

The identifier of the KMS key used to encrypt the target snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

**Tags.Tag.N**

A list of tags to be added to the target snapshot resource. A tag is a key-value pair. Available for Valkey, Redis OSS and Serverless Memcached only. Default: NULL

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**ServerlessCacheSnapshot**

The response for the attempt to copy the serverless cache snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: [ServerlessCacheSnapshot](api-serverlesscachesnapshot.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**InvalidServerlessCacheSnapshotStateFault**

The state of the serverless cache snapshot was not received. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 400

**ServerlessCacheSnapshotAlreadyExistsFault**

A serverless cache snapshot with this name already exists. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 400

**ServerlessCacheSnapshotNotFoundFault**

This serverless cache snapshot could not be found or does not exist. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 404

**ServerlessCacheSnapshotQuotaExceededFault**

The number of serverless cache snapshots exceeds the customer snapshot quota. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 400

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

## Examples

### CopyServerlessCacheSnapshot

This example illustrates one usage of CopyServerlessCacheSnapshot.

#### Sample Request

```

{
    "input": {
        "SourceServerlessCacheSnapshotName": "my-serverless-cache-snapshot",
        "TargetServerlessCacheSnapshotName": "my-serverless-cache-snapshot-copy"
    },
    "output": {
        "ServerlessCacheSnapshot": {
            "ServerlessCacheSnapshotName": "my-serverless-cache-snapshot-copy",
            "ARN": "arn:aws:elasticache:us-east-1:222222222222:serverlesscachesnapshot:my-serverless-cache-snapshot-copy",
            "SnapshotType": "manual",
            "Status": "creating",
            "ServerlessCacheConfiguration": {
                "ServerlessCacheName": "my-serverless-cache",
                "Engine": "redis",
                "MajorEngineVersion": "7"
            }
        }
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/copyserverlesscachesnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompleteMigration

CopySnapshot

All content copied from https://docs.aws.amazon.com/.
