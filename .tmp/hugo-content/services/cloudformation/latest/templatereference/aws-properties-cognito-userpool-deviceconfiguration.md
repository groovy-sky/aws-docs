This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool DeviceConfiguration

The device-remembering configuration for a user pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChallengeRequiredOnNewDevice" : Boolean,
  "DeviceOnlyRememberedOnUserPrompt" : Boolean
}

```

### YAML

```yaml

  ChallengeRequiredOnNewDevice: Boolean
  DeviceOnlyRememberedOnUserPrompt: Boolean

```

## Properties

`ChallengeRequiredOnNewDevice`

When true, a remembered device can sign in with device authentication instead of SMS
and time-based one-time password (TOTP) factors for multi-factor authentication
(MFA).

###### Note

Whether or not `ChallengeRequiredOnNewDevice` is true, users who sign
in with devices that have not been confirmed or remembered must still provide a
second factor in a user pool that requires MFA.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceOnlyRememberedOnUserPrompt`

When true, Amazon Cognito doesn't automatically remember a user's device when your app sends a
`ConfirmDevice` API request. In your app, create a prompt for your user
to choose whether they want to remember their device. Return the user's choice in an
`UpdateDeviceStatus` API request.

When `DeviceOnlyRememberedOnUserPrompt` is `false`, Amazon
Cognito immediately remembers devices that you register in a `ConfirmDevice`
API request.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomSMSSender

EmailConfiguration

All content copied from https://docs.aws.amazon.com/.
