# DescribeServerlessCaches

Returns information about a specific serverless cache.
If no identifier is specified, then the API returns information on all the serverless caches belonging to
this AWS account.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**MaxResults**

The maximum number of records in the response. If more records exist than the specified max-records value,
the next token is included in the response so that remaining results can be retrieved.
The default is 50.

Type: Integer

Required: No

**NextToken**

An optional marker returned from a prior request to support pagination of results from this operation.
If this parameter is specified, the response includes only records beyond the marker,
up to the value specified by MaxResults.

Type: String

Required: No

**ServerlessCacheName**

The identifier for the serverless cache. If this parameter is specified,
only information about that specific serverless cache is returned. Default: NULL

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**NextToken**

An optional marker returned from a prior request to support pagination of results from this operation.
If this parameter is specified, the response includes only records beyond the marker,
up to the value specified by MaxResults.

Type: String

**ServerlessCaches.member.N**

The serverless caches associated with a given description request.

Type: Array of [ServerlessCache](api-serverlesscache.md) objects

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

**ServerlessCacheNotFoundFault**

The serverless cache was not found or does not exist.

HTTP Status Code: 404

## Examples

### DescribeServerlessCaches

This example illustrates one usage of DescribeServerlessCaches.

#### Sample Request

```

{
    "input": {
    },
    "output": {
        "ServerlessCaches": [
            {
                "ServerlessCacheName": "my-serverless-cache",
                "Description": "A serverless cache.",
                "Status": "available",
                "Engine": "redis",
                "MajorEngineVersion": "7",
                "FullEngineVersion": "7.0",
                "SubnetIds": [
                    "subnet-xxx8c982",
                    "subnet-xxx382f3",
                    "subnet-xxxb3e7c0"
                ],
                "CacheUsageLimits": {
                    "DataStorage" : {
                        "Maximum" : 10,
                        "Unit" : "GB"
                    },
                    "ECPUPerSecond" : {
                        "Maximum" : 50000
                    }
                },
                "SecurityGroupIds": [
                    "sg-xxx0c9af"
                ],
                "Endpoint": {
                    "Address": "my-serverless-cache-xxxxxx.serverless.use1qa.cache.amazonaws.com",
                    "Port": 6379
                },
                "ARN": "arn:aws:elasticache:us-east-1:222222222222:serverlesscache:my-serverless-cache",
                "SnapshotRetentionLimit": 10,
                "DailySnapshotTime": "11:00",
                "NetworkType": "ipv4"
            }
        ]
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeserverlesscaches.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeserverlesscaches.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeReservedCacheNodesOfferings

DescribeServerlessCacheSnapshots

All content copied from https://docs.aws.amazon.com/.
