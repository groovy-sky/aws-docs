This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::VoiceChannel

A _channel_ is a type of platform that you can deliver messages to.
To send a voice message, you send the message through the voice channel. Before you can
use Amazon Pinpoint to send voice messages, you have to enable the voice channel for an
Amazon Pinpoint application.

The VoiceChannel resource represents the status and other information about the voice
channel for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::VoiceChannel",
  "Properties" : {
      "ApplicationId" : String,
      "Enabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::VoiceChannel
Properties:
  ApplicationId: String
  Enabled: Boolean

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the voice channel
applies to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Enabled`

Specifies whether to enable the voice channel for the application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application that the channel is associated with.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Id`

(Deprecated) An identifier for the voice channel. This property is retained
only for backward compatibility.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Pinpoint::SmsTemplate

Next
