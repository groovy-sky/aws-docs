# CreateDelivery

Creates a _delivery_. A delivery is a connection between a logical
_delivery source_ and a logical _delivery destination_
that you have already created.

Only some AWS services support being configured as a delivery source using
this operation. These services are listed as **Supported \[V2**
**Permissions\]** in the table at [Enabling logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

A delivery destination can represent a log group in CloudWatch Logs, an Amazon S3 bucket, a delivery stream in Firehose, or X-Ray.

To configure logs delivery between a supported AWS service and a
destination, you must do the following:

- Create a delivery source, which is a logical object that represents the resource that
is actually sending the logs. For more information, see [PutDeliverySource](api-putdeliverysource.md).

- Create a _delivery destination_, which is a logical object that
represents the actual delivery destination. For more information, see [PutDeliveryDestination](api-putdeliverydestination.md).

- If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy](api-putdeliverydestinationpolicy.md) in the destination account to assign an IAM policy to the destination. This policy allows delivery to that destination.

- Use `CreateDelivery` to create a _delivery_ by pairing
exactly one delivery source and one delivery destination.

You can configure a single delivery source to send logs to multiple destinations by
creating multiple deliveries. You can also create multiple deliveries to configure multiple
delivery sources to send logs to the same delivery destination.

To update an existing delivery configuration, use [UpdateDeliveryConfiguration](api-updatedeliveryconfiguration.md).

## Request Syntax

```nohighlight

{
   "deliveryDestinationArn": "string",
   "deliverySourceName": "string",
   "fieldDelimiter": "string",
   "recordFields": [ "string" ],
   "s3DeliveryConfiguration": {
      "enableHiveCompatiblePath": boolean,
      "suffixPath": "string"
   },
   "tags": {
      "string" : "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[deliveryDestinationArn](#API_CreateDelivery_RequestSyntax)**

The ARN of the delivery destination to use for this delivery.

Type: String

Required: Yes

**[deliverySourceName](#API_CreateDelivery_RequestSyntax)**

The name of the delivery source to use for this delivery.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: Yes

**[fieldDelimiter](#API_CreateDelivery_RequestSyntax)**

The field delimiter to use between record fields when the final output format of a
delivery is in `Plain`, `W3C`, or `Raw` format.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 5.

Required: No

**[recordFields](#API_CreateDelivery_RequestSyntax)**

The list of record fields to be delivered to the destination, in order. If the delivery's
log source has mandatory fields, they must be included in this list.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 128 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**[s3DeliveryConfiguration](#API_CreateDelivery_RequestSyntax)**

This structure contains parameters that are valid only when the delivery's delivery
destination is an S3 bucket.

Type: [S3DeliveryConfiguration](api-s3deliveryconfiguration.md) object

Required: No

**[tags](#API_CreateDelivery_RequestSyntax)**

An optional list of key-value pairs to associate with the resource.

For more information about tagging, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md)

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

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

**[delivery](#API_CreateDelivery_ResponseSyntax)**

A structure that contains information about the delivery that you just created.

Type: [Delivery](api-delivery.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/createdelivery.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/createdelivery.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/createdelivery.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/createdelivery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/createdelivery.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/createdelivery.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/createdelivery.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/createdelivery.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/createdelivery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/createdelivery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CancelImportTask

CreateExportTask

All content copied from https://docs.aws.amazon.com/.
