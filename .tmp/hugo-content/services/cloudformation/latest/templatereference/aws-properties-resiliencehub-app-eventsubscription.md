This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::App EventSubscription

Indicates an event you would like to subscribe and get notification for. Currently,
AWS Resilience Hub supports notifications only for **Drift**
**detected** and **Scheduled assessment failure**
events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventType" : String,
  "Name" : String,
  "SnsTopicArn" : String
}

```

### YAML

```yaml

  EventType: String
  Name: String
  SnsTopicArn: String

```

## Properties

`EventType`

The type of event you would like to subscribe and get notification for. Currently, AWS Resilience Hub supports notifications only for **Drift**
**detected** ( `DriftDetected`) and **Scheduled**
**assessment failure** ( `ScheduledAssessmentFailure`) events.

_Required_: Yes

_Type_: String

_Allowed values_: `ScheduledAssessmentFailure | DriftDetected`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Unique name to identify an event subscription.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicArn`

Amazon Resource Name (ARN) of the Amazon Simple Notification Service topic. The format for this ARN
is: `arn:partition:sns:region:account:topic-name`. For more information about ARNs,
see [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the
_AWS General Reference_ guide.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-iso|aws-iso-[a-z]{1}|aws-us-gov):[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:([a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-[0-9]):[0-9]{12}:[A-Za-z0-9/][A-Za-z0-9:_/+.-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ResilienceHub::App

PermissionModel

All content copied from https://docs.aws.amazon.com/.
