This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RelayAction

The action relays the email via SMTP to another specific SMTP server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "MailFrom" : String,
  "Relay" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  MailFrom: String
  Relay: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there
are configuration errors. For example, the specified relay has been deleted.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailFrom`

This action specifies whether to preserve or replace original mail from address while
relaying received emails to a destination server.

_Required_: No

_Type_: String

_Allowed values_: `REPLACE | PRESERVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Relay`

The identifier of the relay resource to be used when relaying an email.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeliverToQBusinessAction

ReplaceRecipientAction

All content copied from https://docs.aws.amazon.com/.
