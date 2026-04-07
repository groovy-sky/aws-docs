# UpdateDeliveryConfiguration

Use this operation to update the configuration of a [delivery](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_Delivery.html) to change
either the S3 path pattern or the format of the delivered logs. You can't use this operation
to change the source or destination of the delivery.

## Request Syntax

```nohighlight

{
   "fieldDelimiter": "string",
   "id": "string",
   "recordFields": [ "string" ],
   "s3DeliveryConfiguration": {
      "enableHiveCompatiblePath": boolean,
      "suffixPath": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[fieldDelimiter](#API_UpdateDeliveryConfiguration_RequestSyntax)**

The field delimiter to use between record fields when the final output format of a
delivery is in `Plain`, `W3C`, or `Raw` format.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 5.

Required: No

**[id](#API_UpdateDeliveryConfiguration_RequestSyntax)**

The ID of the delivery to be updated by this request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `^[0-9A-Za-z]+$`

Required: Yes

**[recordFields](#API_UpdateDeliveryConfiguration_RequestSyntax)**

The list of record fields to be delivered to the destination, in order. If the delivery's
log source has mandatory fields, they must be included in this list.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 128 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[s3DeliveryConfiguration](#API_UpdateDeliveryConfiguration_RequestSyntax)**

This structure contains parameters that are valid only when the delivery's delivery
destination is an S3 bucket.

Type: [S3DeliveryConfiguration](api-s3deliveryconfiguration.md) object

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 400

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/logs-2014-03-28/UpdateDeliveryConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/UpdateDeliveryConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateAnomaly

UpdateLogAnomalyDetector
