# DeliveryDestination

This structure contains information about one _delivery destination_ in
your account. A delivery destination is an AWS resource that represents an
AWS service that logs can be sent to. CloudWatch Logs, Amazon S3,
Firehose, and X-Ray are supported as delivery destinations.

To configure logs delivery between a supported AWS service and a
destination, you must do the following:

- Create a delivery source, which is a logical object that represents the resource that
is actually sending the logs. For more information, see [PutDeliverySource](api-putdeliverysource.md).

- Create a _delivery destination_, which is a logical object that
represents the actual delivery destination.

- If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy](api-putdeliverydestinationpolicy.md) in the destination account to assign an IAM policy to the destination. This policy allows delivery to that destination.

- Create a _delivery_ by pairing exactly one delivery source and one
delivery destination. For more information, see [CreateDelivery](api-createdelivery.md).

You can configure a single delivery source to send logs to multiple destinations by
creating multiple deliveries. You can also create multiple deliveries to configure multiple
delivery sources to send logs to the same delivery destination.

## Contents

**arn**

The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.

Type: String

Required: No

**deliveryDestinationConfiguration**

A structure that contains the ARN of the AWS resource that will receive the
logs.

Type: [DeliveryDestinationConfiguration](api-deliverydestinationconfiguration.md) object

Required: No

**deliveryDestinationType**

Displays whether this delivery destination is CloudWatch Logs, Amazon S3,
Firehose, or X-Ray.

Type: String

Valid Values: `S3 | CWL | FH | XRAY`

Required: No

**name**

The name of this delivery destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: No

**outputFormat**

The format of the logs that are sent to this delivery destination.

Type: String

Valid Values: `json | plain | w3c | raw | parquet`

Required: No

**tags**

The tags that have been assigned to this delivery destination.

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deliverydestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deliverydestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deliverydestination.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Delivery

DeliveryDestinationConfiguration
