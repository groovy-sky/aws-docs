This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::SubscriptionDefinitionVersion Subscription

Subscriptions define how MQTT messages can be exchanged between devices, functions, and
connectors in the group, and with AWS IoT or the local shadow service. A
subscription defines a message source, message target, and a topic (or subject) that's used
to route messages from the source to the target. A subscription defines the message flow in
one direction, from the source to the target. For two-way communication, you must set up
two subscriptions, one for each direction.

In an CloudFormation template, the `Subscriptions` property of the [`AWS::Greengrass::SubscriptionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-subscriptiondefinitionversion.html) resource contains a list of `Subscription` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Source" : String,
  "Subject" : String,
  "Target" : String
}

```

### YAML

```yaml

  Id: String
  Source: String
  Subject: String
  Target: String

```

## Properties

`Id`

A descriptive or arbitrary ID for the subscription. This value must be unique within the
subscription definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Source`

The originator of the message. The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN,
`cloud` (which represents the AWS IoT cloud), or
`GGShadowService`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subject`

The MQTT topic used to route the message.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Target`

The destination of the message. The value can be a thing ARN, the ARN of a Lambda function alias (recommended) or version, a connector ARN,
`cloud` (which represents the AWS IoT cloud), or
`GGShadowService`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Subscription](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-subscription.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Greengrass::SubscriptionDefinitionVersion

Next
