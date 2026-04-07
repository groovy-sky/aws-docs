# DescribeDeliveries

Retrieves a list of the deliveries that have been created in the account.

A _delivery_ is a connection between a [_delivery_\
_source_](api-putdeliverysource.md) and a [_delivery destination_](api-putdeliverydestination.md).

A delivery source represents an AWS resource that sends logs to an logs
delivery destination. The destination can be CloudWatch Logs, Amazon S3, Firehose or X-Ray. Only some AWS services support being
configured as a delivery source. These services are listed in [Enable logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

## Request Syntax

```nohighlight

{
   "limit": number,
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[limit](#API_DescribeDeliveries_RequestSyntax)**

Optionally specify the maximum number of deliveries to return in the response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[nextToken](#API_DescribeDeliveries_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "deliveries": [
      {
         "arn": "string",
         "deliveryDestinationArn": "string",
         "deliveryDestinationType": "string",
         "deliverySourceName": "string",
         "fieldDelimiter": "string",
         "id": "string",
         "recordFields": [ "string" ],
         "s3DeliveryConfiguration": {
            "enableHiveCompatiblePath": boolean,
            "suffixPath": "string"
         },
         "tags": {
            "string" : "string"
         }
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[deliveries](#API_DescribeDeliveries_ResponseSyntax)**

An array of structures. Each structure contains information about one delivery in the
account.

Type: Array of [Delivery](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Delivery.html) objects

**[nextToken](#API_DescribeDeliveries_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ServiceQuotaExceededException**

This request exceeds a service quota.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 400

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/DescribeDeliveries)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/DescribeDeliveries)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeConfigurationTemplates

DescribeDeliveryDestinations
