# PutDeliveryDestination

Creates or updates a logical _delivery destination_. A delivery
destination is an AWS resource that represents an AWS service
that logs can be sent to. CloudWatch Logs, Amazon S3, and Firehose are
supported as logs delivery destinations and X-Ray as the trace delivery
destination.

To configure logs delivery between a supported AWS service and a
destination, you must do the following:

- Create a delivery source, which is a logical object that represents the resource that
is actually sending the logs. For more information, see [PutDeliverySource](api-putdeliverysource.md).

- Use `PutDeliveryDestination` to create a _delivery_
_destination_ in the same account of the actual delivery destination. The
delivery destination that you create is a logical object that represents the actual
delivery destination.

- If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy](api-putdeliverydestinationpolicy.md) in the destination account to assign an IAM policy to the destination. This policy allows delivery to that destination.

- Use `CreateDelivery` to create a _delivery_ by pairing
exactly one delivery source and one delivery destination. For more information, see [CreateDelivery](api-createdelivery.md).

You can configure a single delivery source to send logs to multiple destinations by
creating multiple deliveries. You can also create multiple deliveries to configure multiple
delivery sources to send logs to the same delivery destination.

Only some AWS services support being configured as a delivery source. These
services are listed as **Supported \[V2 Permissions\]** in the
table at [Enabling logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

If you use this operation to update an existing delivery destination, all the current
delivery destination parameters are overwritten with the new parameter values that you
specify.

## Request Syntax

```nohighlight

{
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
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[deliveryDestinationConfiguration](#API_PutDeliveryDestination_RequestSyntax)**

A structure that contains the ARN of the AWS resource that will receive the
logs.

###### Note

`deliveryDestinationConfiguration` is required for CloudWatch Logs,
Amazon S3, Firehose log delivery destinations and not required for
X-Ray trace delivery destinations. `deliveryDestinationType` is
needed for X-Ray trace delivery destinations but not required for other logs
delivery destinations.

Type: [DeliveryDestinationConfiguration](api-deliverydestinationconfiguration.md) object

Required: No

**[deliveryDestinationType](#API_PutDeliveryDestination_RequestSyntax)**

The type of delivery destination. This parameter specifies the target service where log
data will be delivered. Valid values include:

- `S3` \- Amazon S3 for long-term storage and analytics

- `CWL` \- CloudWatch Logs for centralized log management

- `FH` \- Amazon Kinesis Data Firehose for real-time data streaming

- `XRAY` \- AWS
X-Ray for distributed tracing and application monitoring

The delivery destination type determines the format and configuration options available
for log delivery.

Type: String

Valid Values: `S3 | CWL | FH | XRAY`

Required: No

**[name](#API_PutDeliveryDestination_RequestSyntax)**

A name for this delivery destination. This name must be unique for all delivery
destinations in your account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: Yes

**[outputFormat](#API_PutDeliveryDestination_RequestSyntax)**

The format for the logs that this delivery destination will receive.

Type: String

Valid Values: `json | plain | w3c | raw | parquet`

Required: No

**[tags](#API_PutDeliveryDestination_RequestSyntax)**

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
   "deliveryDestination": {
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
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[deliveryDestination](#API_PutDeliveryDestination_ResponseSyntax)**

A structure containing information about the delivery destination that you just created or
updated.

Type: [DeliveryDestination](api-deliverydestination.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putdeliverydestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putdeliverydestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutDataProtectionPolicy

PutDeliveryDestinationPolicy

All content copied from https://docs.aws.amazon.com/.
