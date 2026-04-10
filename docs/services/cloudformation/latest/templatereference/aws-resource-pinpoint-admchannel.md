This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::ADMChannel

A _channel_ is a type of platform that you can deliver messages to.
You can use the ADM channel to send push notifications through the Amazon Device
Messaging (ADM) service to apps that run on Amazon devices, such as Kindle Fire tablets.
Before you can use Amazon Pinpoint to send messages to Amazon devices, you have to
enable the ADM channel for an Amazon Pinpoint application.

The ADMChannel resource represents the status and authentication settings for the ADM
channel for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::ADMChannel",
  "Properties" : {
      "ApplicationId" : String,
      "ClientId" : String,
      "ClientSecret" : String,
      "Enabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::ADMChannel
Properties:
  ApplicationId: String
  ClientId: String
  ClientSecret: String
  Enabled: Boolean

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the ADM channel applies
to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientId`

The Client ID that you received from Amazon to send messages by using ADM.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The Client Secret that you received from Amazon to send messages by using ADM.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether to enable the ADM channel for the application.

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

(Deprecated) An identifier for the ADM channel. This property is retained only
for backward compatibility.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Pinpoint

AWS::Pinpoint::APNSChannel

All content copied from https://docs.aws.amazon.com/.
