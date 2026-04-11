# ConnectionNotification

Describes a connection notification for a VPC endpoint or VPC endpoint
service.

## Contents

**ConnectionEvents.N**

The events for the notification. Valid values are `Accept`,
`Connect`, `Delete`, and `Reject`.

Type: Array of strings

Required: No

**connectionNotificationArn**

The ARN of the SNS topic for the notification.

Type: String

Required: No

**connectionNotificationId**

The ID of the notification.

Type: String

Required: No

**connectionNotificationState**

The state of the notification.

Type: String

Valid Values: `Enabled | Disabled`

Required: No

**connectionNotificationType**

The type of notification.

Type: String

Valid Values: `Topic`

Required: No

**serviceId**

The ID of the endpoint service.

Type: String

Required: No

**serviceRegion**

The Region for the endpoint service.

Type: String

Required: No

**vpcEndpointId**

The ID of the VPC endpoint.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/connectionnotification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/connectionnotification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/connectionnotification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ConnectionLogResponseOptions

ConnectionTrackingConfiguration
