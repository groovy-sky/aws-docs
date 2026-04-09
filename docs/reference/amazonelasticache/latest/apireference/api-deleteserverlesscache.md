# DeleteServerlessCache

Deletes a specified existing serverless cache.

###### Note

`CreateServerlessCacheSnapshot` permission is required to create a final snapshot.
Without this permission, the API call will fail with an `Access Denied` exception.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ServerlessCacheName**

The identifier of the serverless cache to be deleted.

Type: String

Required: Yes

**FinalSnapshotName**

Name of the final snapshot to be taken before the serverless cache is deleted. Available for Valkey, Redis OSS and Serverless Memcached only.
Default: NULL, i.e. a final snapshot is not taken.

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**ServerlessCache**

Provides the details of the specified serverless cache that is about to be deleted.

Type: [ServerlessCache](api-serverlesscache.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidCredentials**

You must enter valid credentials.

HTTP Status Code: 408

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

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

## Examples

### DeleteServerlessCache

This example illustrates one usage of DeleteServerlessCache.

#### Sample Request

```

{
    "input": {
        "ServerlessCacheName": "my-serverless-cache"
    },
    "output": {
        "ServerlessCache": {
            "ServerlessCacheName": "my-serverless-cache",
            "Description": "A serverless cache.",
            "Status": "deleting",
            "Engine": "redis",
            "MajorEngineVersion": "7",
            "FullEngineVersion": "7.0",
            "SubnetIds": [
                "subnet-xxx8c982",
                "subnet-xxx382f3",
                "subnet-xxxb3e7c0"
            ],
            "SecurityGroupIds": [
                "sg-xxx0c9af"
            ],
            "Endpoint": {
                "Address": "my-serverless-cache-xxxxx.serverless.use1qa.cache.amazonaws.com",
                "Port": 6379
            },
            "ARN": "arn:aws:elasticache:us-east-1:222222222222:serverlesscache:my-serverless-cache",
            "SnapshotRetentionLimit": 10,
            "DailySnapshotTime": "11:00",
            "NetworkType": "ipv4"
        }
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/deleteserverlesscache.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/deleteserverlesscache.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteReplicationGroup

DeleteServerlessCacheSnapshot

All content copied from https://docs.aws.amazon.com/.
