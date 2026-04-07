This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::APNSSandboxChannel

A _channel_ is a type of platform that you can deliver messages to.
You can use the APNs sandbox channel to send push notification messages to the sandbox
environment of the Apple Push Notification service (APNs). Before you can use Amazon
Pinpoint to send notifications to the APNs sandbox environment, you have to enable the
APNs sandbox channel for an Amazon Pinpoint application.

The APNSSandboxChannel resource represents the status and authentication settings of
the APNs sandbox channel for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::APNSSandboxChannel",
  "Properties" : {
      "ApplicationId" : String,
      "BundleId" : String,
      "Certificate" : String,
      "DefaultAuthenticationMethod" : String,
      "Enabled" : Boolean,
      "PrivateKey" : String,
      "TeamId" : String,
      "TokenKey" : String,
      "TokenKeyId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::APNSSandboxChannel
Properties:
  ApplicationId: String
  BundleId: String
  Certificate: String
  DefaultAuthenticationMethod: String
  Enabled: Boolean
  PrivateKey: String
  TeamId: String
  TokenKey: String
  TokenKeyId: String

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the APNs sandbox
channel applies to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BundleId`

The bundle identifier that's assigned to your iOS app. This identifier is used for
APNs tokens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Certificate`

The APNs client certificate that you received from Apple. Specify this value if you
want Amazon Pinpoint to communicate with APNs by using an APNs certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultAuthenticationMethod`

The default authentication method that you want Amazon Pinpoint to use when
authenticating with APNs. Valid options are `key` or
`certificate`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether to enable the APNs Sandbox channel for the Amazon Pinpoint
application.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateKey`

The private key for the APNs client certificate that you want Amazon Pinpoint to use
to communicate with APNs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeamId`

The identifier that's assigned to your Apple Developer Account team. This identifier
is used for APNs tokens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenKey`

The authentication key to use for APNs tokens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenKeyId`

The key identifier that's assigned to your APNs signing key. Specify this value if you
want Amazon Pinpoint to communicate with APNs by using APNs tokens.

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

(Deprecated) An identifier for the APNs sandbox channel. This property is
retained only for backward compatibility.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Pinpoint::APNSChannel

AWS::Pinpoint::APNSVoipChannel
