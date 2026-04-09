# GetDelivery

Returns complete information about one logical _delivery_. A delivery
is a connection between a [_delivery_\
_source_](api-putdeliverysource.md) and a [_delivery destination_](api-putdeliverydestination.md).

A delivery source represents an AWS resource that sends logs to an logs
delivery destination. The destination can be CloudWatch Logs, Amazon S3, or Firehose. Only some AWS services support being configured as a delivery
source. These services are listed in [Enable logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

You need to specify the delivery `id` in this operation. You can find the IDs
of the deliveries in your account with the [DescribeDeliveries](api-describedeliveries.md) operation.

## Request Syntax

```nohighlight

{
   "id": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[id](#API_GetDelivery_RequestSyntax)**

The ID of the delivery that you want to retrieve.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `^[0-9A-Za-z]+$`

Required: Yes

## Response Syntax

```nohighlight

{
   "delivery": {
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
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[delivery](#API_GetDelivery_ResponseSyntax)**

A structure that contains information about the delivery.

Type: [Delivery](api-delivery.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/getdelivery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/getdelivery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/getdelivery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/getdelivery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/getdelivery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/getdelivery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/getdelivery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/getdelivery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/getdelivery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/getdelivery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetDataProtectionPolicy

GetDeliveryDestination

All content copied from https://docs.aws.amazon.com/.
