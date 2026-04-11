This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup NotificationConfiguration

A structure that specifies an Amazon SNS notification configuration for the
`NotificationConfigurations` property of the [AWS::AutoScaling::AutoScalingGroup](../userguide/aws-resource-autoscaling-autoscalinggroup.md) resource.

For an example template snippet, see [Configure Amazon EC2\
Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md).

For more information, see [Get Amazon SNS notifications\
when your Auto Scaling group scales](../../../autoscaling/ec2/userguide/asgettingnotifications.md) in the _Amazon EC2 Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NotificationTypes" : [ String, ... ],
  "TopicARN" : [ String, ... ]
}

```

### YAML

```yaml

  NotificationTypes:
    - String
  TopicARN:
    - String

```

## Properties

`NotificationTypes`

A list of event types that send a notification. Event types can include any of the
following types.

_Allowed values_:

- `autoscaling:EC2_INSTANCE_LAUNCH`

- `autoscaling:EC2_INSTANCE_LAUNCH_ERROR`

- `autoscaling:EC2_INSTANCE_TERMINATE`

- `autoscaling:EC2_INSTANCE_TERMINATE_ERROR`

- `autoscaling:TEST_NOTIFICATION`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicARN`

The Amazon Resource Name (ARN) of the Amazon SNS topic.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkInterfaceCountRequest

PerformanceFactorReferenceRequest

All content copied from https://docs.aws.amazon.com/.
