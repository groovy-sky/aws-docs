# DescribeDeliveryDestinations

Retrieves a list of the delivery destinations that have been created in the
account.

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

**[limit](#API_DescribeDeliveryDestinations_RequestSyntax)**

Optionally specify the maximum number of delivery destinations to return in the
response.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[nextToken](#API_DescribeDeliveryDestinations_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "deliveryDestinations": [
      {
         "arn": "string",
         "deliveryDestinationConfiguration": {
            "destinationResourceArn": "string"
         },
         "deliveryDestinationType": "string",
         "name": "string",
         "outputFormat": "string",
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

**[deliveryDestinations](#API_DescribeDeliveryDestinations_ResponseSyntax)**

An array of structures. Each structure contains information about one delivery destination
in the account.

Type: Array of [DeliveryDestination](api-deliverydestination.md) objects

**[nextToken](#API_DescribeDeliveryDestinations_ResponseSyntax)**

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describedeliverydestinations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describedeliverydestinations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDeliveries

DescribeDeliverySources

All content copied from https://docs.aws.amazon.com/.
