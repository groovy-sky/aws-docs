# GetDeliveryDestinationPolicy

Retrieves the delivery destination policy assigned to the delivery destination that you
specify. For more information about delivery destinations and their policies, see [PutDeliveryDestinationPolicy](api-putdeliverydestinationpolicy.md).

## Request Syntax

```nohighlight

{
   "deliveryDestinationName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[deliveryDestinationName](#API_GetDeliveryDestinationPolicy_RequestSyntax)**

The name of the delivery destination that you want to retrieve the policy of.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: Yes

## Response Syntax

```nohighlight

{
   "policy": {
      "deliveryDestinationPolicy": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[policy](#API_GetDeliveryDestinationPolicy_ResponseSyntax)**

The IAM policy for this delivery destination.

Type: [Policy](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Policy.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/GetDeliveryDestinationPolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/GetDeliveryDestinationPolicy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetDeliveryDestination

GetDeliverySource
