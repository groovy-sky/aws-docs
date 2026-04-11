This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::EventRule

Creates an [`EventRule`](../../../notifications/latest/userguide/glossary.md) that
is associated with a specified `NotificationConfiguration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::EventRule",
  "Properties" : {
      "EventPattern" : String,
      "EventType" : String,
      "NotificationConfigurationArn" : String,
      "Regions" : [ String, ... ],
      "Source" : String
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::EventRule
Properties:
  EventPattern: String
  EventType: String
  NotificationConfigurationArn: String
  Regions:
    - String
  Source: String

```

## Properties

`EventPattern`

An additional event pattern used to further filter the events this `EventRule` receives.

For more information, see [Amazon EventBridge event patterns](../../../eventbridge/latest/userguide/eb-event-patterns.md) in the _Amazon EventBridge User Guide._

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventType`

The event type this rule should match with the EventBridge events. It must match with atleast one of the valid EventBridge event types. For example, Amazon EC2 Instance State change Notification and Amazon CloudWatch State Change. For more information, see [Event delivery from AWS services](../../../eventbridge/latest/userguide/eb-service-event.md#eb-service-event-delivery-level) in the _Amazon EventBridge User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9 \-\(\)])+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationConfigurationArn`

The ARN for the `NotificationConfiguration` associated with this `EventRule`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:notifications::[0-9]{12}:configuration/[a-z0-9]{27}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Regions`

A list of AWS Regions that send events to this `EventRule`.

_Required_: Yes

_Type_: Array of String

_Maximum_: `25`

_Minimum_: `2 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The event source this rule should match with the EventBridge event sources. It must match with atleast one of the valid EventBridge event sources. Only AWS service sourced events are supported.
For example, `aws.ec2` and `aws.cloudwatch`. For more information, see [Event delivery from AWS services](../../../eventbridge/latest/userguide/eb-service-event.md#eb-service-event-delivery-level) in the _Amazon EventBridge User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^aws.([a-z0-9\-])+$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the `EventRule`. CloudFormation stack generates this ARN and then uses this ARN associated with the `NotificationConfiguration`.

`CreationTime`

The creation time of the `EventRule`.

`ManagedRules`

A list of Amazon EventBridge Managed Rule ARNs associated with this `EventRule`.

###### Note

These are created by AWS User Notifications within your account so your `EventRules` can function.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Notifications::ChannelAssociation

EventRuleStatusSummary

All content copied from https://docs.aws.amazon.com/.
