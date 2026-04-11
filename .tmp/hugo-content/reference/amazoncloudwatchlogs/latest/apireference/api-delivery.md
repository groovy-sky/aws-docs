# Delivery

This structure contains information about one _delivery_ in your
account.

A delivery is a connection between a logical _delivery source_ and a
logical _delivery destination_.

For more information, see [CreateDelivery](api-createdelivery.md).

To update an existing delivery configuration, use [UpdateDeliveryConfiguration](api-updatedeliveryconfiguration.md).

## Contents

**arn**

The Amazon Resource Name (ARN) that uniquely identifies this delivery.

Type: String

Required: No

**deliveryDestinationArn**

The ARN of the delivery destination that is associated with this delivery.

Type: String

Required: No

**deliveryDestinationType**

Displays whether the delivery destination associated with this delivery is CloudWatch Logs, Amazon S3, Firehose, or X-Ray.

Type: String

Valid Values: `S3 | CWL | FH | XRAY`

Required: No

**deliverySourceName**

The name of the delivery source that is associated with this delivery.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: No

**fieldDelimiter**

The field delimiter that is used between record fields when the final output format of a
delivery is in `Plain`, `W3C`, or `Raw` format.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 5.

Required: No

**id**

The unique ID that identifies this delivery in your account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `^[0-9A-Za-z]+$`

Required: No

**recordFields**

The record fields used in this delivery.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 128 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**s3DeliveryConfiguration**

This structure contains delivery configurations that apply only when the delivery
destination resource is an S3 bucket.

Type: [S3DeliveryConfiguration](api-s3deliveryconfiguration.md) object

Required: No

**tags**

The tags that have been assigned to this delivery.

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/delivery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/delivery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/delivery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteKeys

DeliveryDestination

All content copied from https://docs.aws.amazon.com/.
