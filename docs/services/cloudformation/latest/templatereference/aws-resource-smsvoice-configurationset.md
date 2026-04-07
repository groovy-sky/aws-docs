This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ConfigurationSet

Creates a new configuration set. After you create the configuration set, you can add
one or more event destinations to it.

A configuration set is a set of rules that you apply to the SMS and voice messages
that you send.

When you send a message, you can optionally specify a single configuration set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::ConfigurationSet",
  "Properties" : {
      "ConfigurationSetName" : String,
      "DefaultSenderId" : String,
      "EventDestinations" : [ EventDestination, ... ],
      "MessageFeedbackEnabled" : Boolean,
      "ProtectConfigurationId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::ConfigurationSet
Properties:
  ConfigurationSetName: String
  DefaultSenderId: String
  EventDestinations:
    - EventDestination
  MessageFeedbackEnabled: Boolean
  ProtectConfigurationId: String
  Tags:
    - Tag

```

## Properties

`ConfigurationSetName`

The name of the ConfigurationSet.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultSenderId`

The default sender ID used by the ConfigurationSet.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventDestinations`

An array of EventDestination objects that describe any events to log and where to log
them.

_Required_: No

_Type_: Array of [EventDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-configurationset-eventdestination.html)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageFeedbackEnabled`

Set to true to enable feedback for the message.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectConfigurationId`

The unique identifier for the protect configuration.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_:/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key and value pair tags that's associated with the new configuration set.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-smsvoice-configurationset-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `ConfigurationSetName`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the ConfigurationSet.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS End User Messaging SMS

CloudWatchLogsDestination
