This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment AccountTakeoverActionType

The automated response to a risk level for adaptive authentication in full-function,
or `ENFORCED`, mode. You can assign an action to each risk level that
advanced security features evaluates.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventAction" : String,
  "Notify" : Boolean
}

```

### YAML

```yaml

  EventAction: String
  Notify: Boolean

```

## Properties

`EventAction`

The action to take for the attempted account takeover action for the associated risk
level. Valid values are as follows:

- `BLOCK`: Block the request.

- `MFA_IF_CONFIGURED`: Present an MFA challenge if possible. MFA is
possible if the user pool has active MFA methods that the user can set up. For
example, if the user pool only supports SMS message MFA but the user
doesn't have a phone number attribute, MFA setup isn't possible. If MFA
setup isn't possible, allow the request.

- `MFA_REQUIRED`: Present an MFA challenge if possible. Block the
request if a user hasn't set up MFA. To sign in with required MFA, users must
have an email address or phone number attribute, or a registered TOTP
factor.

- `NO_ACTION`: Take no action. Permit sign-in.

_Required_: Yes

_Type_: String

_Allowed values_: `BLOCK | MFA_IF_CONFIGURED | MFA_REQUIRED | NO_ACTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Notify`

Determines whether Amazon Cognito sends a user a notification message when your user pools
assesses a user's session at the associated risk level.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountTakeoverActionsType

AccountTakeoverRiskConfigurationType

All content copied from https://docs.aws.amazon.com/.
