# DeliverySource

This structure contains information about one _delivery source_ in your
account. A delivery source is an AWS resource that sends logs to an AWS destination. The destination can be CloudWatch Logs, Amazon S3, or
Firehose.

Only some AWS services support being configured as a delivery source. These
services are listed as **Supported \[V2 Permissions\]** in the
table at [Enabling logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

To configure logs delivery between a supported AWS service and a
destination, you must do the following:

- Create a delivery source, which is a logical object that represents the resource that
is actually sending the logs. For more information, see [PutDeliverySource](api-putdeliverysource.md).

- Create a _delivery destination_, which is a logical object that
represents the actual delivery destination. For more information, see [PutDeliveryDestination](api-putdeliverydestination.md).

- If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy](api-putdeliverydestinationpolicy.md) in the destination account to assign an IAM policy to the destination. This policy allows delivery to that destination.

- Create a _delivery_ by pairing exactly one delivery source and one
delivery destination. For more information, see [CreateDelivery](api-createdelivery.md).

You can configure a single delivery source to send logs to multiple destinations by
creating multiple deliveries. You can also create multiple deliveries to configure multiple
delivery sources to send logs to the same delivery destination.

## Contents

**arn**

The Amazon Resource Name (ARN) that uniquely identifies this delivery source.

Type: String

Required: No

**logType**

The type of log that the source is sending. For valid values for this parameter, see the
documentation for the source service.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w]*`

Required: No

**name**

The unique name of the delivery source.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: No

**resourceArns**

This array contains the ARN of the AWS resource that sends logs and is
represented by this delivery source. Currently, only one ARN can be in the array.

Type: Array of strings

Required: No

**service**

The AWS service that is sending logs.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w_-]*`

Required: No

**tags**

The tags that have been assigned to this delivery source.

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/DeliverySource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/DeliverySource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/DeliverySource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeliveryDestinationConfiguration

Destination
