This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::GCMChannel

A _channel_ is a type of platform that you can deliver messages to.
You can use the GCM channel to send push notification messages to the Firebase Cloud
Messaging (FCM) service, which replaced the Google Cloud Messaging (GCM) service. Before
you use Amazon Pinpoint to send notifications to FCM, you have to enable the GCM channel
for an Amazon Pinpoint application.

The GCMChannel resource represents the status and authentication settings of the GCM
channel for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::GCMChannel",
  "Properties" : {
      "ApiKey" : String,
      "ApplicationId" : String,
      "DefaultAuthenticationMethod" : String,
      "Enabled" : Boolean,
      "ServiceJson" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::GCMChannel
Properties:
  ApiKey: String
  ApplicationId: String
  DefaultAuthenticationMethod: String
  Enabled: Boolean
  ServiceJson: String

```

## Properties

`ApiKey`

The Web API key, also called the _server key_, that you received
from Google to communicate with Google services.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the GCM channel applies
to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultAuthenticationMethod`

The default authentication method used for GCM. Values are either "TOKEN" or "KEY". Defaults to "KEY".

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether to enable the GCM channel for the Amazon Pinpoint
application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceJson`

The contents of the JSON file provided by Google during registration in order to generate an access token for authentication. For more information see [Migrate from legacy FCM APIs to HTTP v1](https://firebase.google.com/docs/cloud-messaging/migrate-v1).

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

(Deprecated) An identifier for the GCM channel. This property is retained only
for backward compatibility.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Pinpoint::EventStream

AWS::Pinpoint::InAppTemplate
