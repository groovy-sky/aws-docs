This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::TrustAnchor NotificationSetting

Customizable notification settings that will be applied to notification events.
IAM Roles Anywhere consumes these settings while notifying across multiple channels - CloudWatch metrics, EventBridge, and Health Dashboard.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Channel" : String,
  "Enabled" : Boolean,
  "Event" : String,
  "Threshold" : Number
}

```

### YAML

```yaml

  Channel: String
  Enabled: Boolean
  Event: String
  Threshold: Number

```

## Properties

`Channel`

The specified channel of notification.
IAM Roles Anywhere uses CloudWatch metrics, EventBridge, and Health Dashboard to notify for an event.

###### Note

In the absence of a specific channel, IAM Roles Anywhere applies this setting to 'ALL' channels.

_Required_: No

_Type_: String

_Allowed values_: `ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether the notification setting is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Event`

The event to which this notification setting is applied.

_Required_: Yes

_Type_: String

_Allowed values_: `CA_CERTIFICATE_EXPIRY | END_ENTITY_CERTIFICATE_EXPIRY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The number of days before a notification event. This value is required for a notification setting that is enabled.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `360`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::RolesAnywhere::TrustAnchor

Source
