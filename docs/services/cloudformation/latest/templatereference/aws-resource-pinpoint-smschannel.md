This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::SMSChannel

A _channel_ is a type of platform that you can deliver messages to.
To send an SMS text message, you send the message through the SMS channel. Before you
can use Amazon Pinpoint to send text messages, you have to enable the SMS channel for an
Amazon Pinpoint application.

The SMSChannel resource represents the status, sender ID, and other settings for the
SMS channel for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::SMSChannel",
  "Properties" : {
      "ApplicationId" : String,
      "Enabled" : Boolean,
      "SenderId" : String,
      "ShortCode" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::SMSChannel
Properties:
  ApplicationId: String
  Enabled: Boolean
  SenderId: String
  ShortCode: String

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the SMS channel applies
to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Enabled`

Specifies whether to enable the SMS channel for the application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SenderId`

The identity that you want to display on recipients' devices when they receive
messages from the SMS channel.

###### Note

SenderIDs are only supported in certain countries and regions. For more
information, see [Supported Countries\
and Regions](../../../pinpoint/latest/userguide/channels-sms-countries.md) in the _Amazon Pinpoint User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShortCode`

The registered short code that you want to use when you send messages through the SMS
channel.

###### Note

For information about obtaining a dedicated short code for sending SMS messages,
see [Requesting Dedicated Short Codes for SMS Messaging with Amazon Pinpoint](../../../pinpoint/latest/userguide/channels-sms-awssupport-short-code.md)
in the _Amazon Pinpoint User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application that the channel is associated with.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Id`

(Deprecated) An identifier for the SMS channel. This property is retained only
for backward compatibility.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceSegments

AWS::Pinpoint::SmsTemplate

All content copied from https://docs.aws.amazon.com/.
