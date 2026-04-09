# CreateServerlessCacheSnapshot

This API creates a copy of an entire ServerlessCache at a specific moment in time. Available for Valkey, Redis OSS and Serverless Memcached only.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ServerlessCacheName**

The name of an existing serverless cache. The snapshot is created from this cache. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: Yes

**ServerlessCacheSnapshotName**

The name for the snapshot being created. Must be unique for the customer account. Available for Valkey, Redis OSS and Serverless Memcached only.
Must be between 1 and 255 characters. This value is stored as a lowercase string.

Type: String

Required: Yes

**KmsKeyId**

The ID of the KMS key used to encrypt the snapshot. Available for Valkey, Redis OSS and Serverless Memcached only. Default: NULL

Type: String

Required: No

**Tags.Tag.N**

A list of tags to be added to the snapshot resource. A tag is a key-value pair. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: Array of [Tag](api-tag.md) objects

Required: No

## Response Elements

The following element is returned by the service.

**ServerlessCacheSnapshot**

The state of a serverless cache snapshot at a specific point in time, to the millisecond. Available for Valkey, Redis OSS and Serverless Memcached only.

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

**InvalidServerlessCacheStateFault**

The account for these credentials is not currently active.

HTTP Status Code: 400

**ServerlessCacheNotFoundFault**

The serverless cache was not found or does not exist.

HTTP Status Code: 404

**ServerlessCacheSnapshotAlreadyExistsFault**

A serverless cache snapshot with this name already exists. Available for Valkey, Redis OSS and Serverless Memcached only.

HTTP Status Code: 400

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

### CreateServerlessCacheSnapshot

This example illustrates one usage of CreateServerlessCacheSnapshot.

#### Sample Request

```

{
    "input": {
        "ServerlessCacheSnapshotName": "my-serverless-cache-snapshot",
        "ServerlessCacheName": "my-serverless-cache"
    },
    "output": {
        "ServerlessCacheSnapshot": {
            "ServerlessCacheSnapshotName": "my-serverless-cache-snapshot",
            "ARN": "arn:aws:elasticache:us-east-1:222222222222:serverlesscachesnapshot:my-serverless-cache-snapshot",
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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/createserverlesscachesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/createserverlesscachesnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateServerlessCache

CreateSnapshot

All content copied from https://docs.aws.amazon.com/.
