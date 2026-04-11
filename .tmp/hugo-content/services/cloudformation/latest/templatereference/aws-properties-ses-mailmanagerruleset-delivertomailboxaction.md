This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet DeliverToMailboxAction

This action to delivers an email to a mailbox.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "MailboxArn" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  MailboxArn: String
  RoleArn: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there
are configuration errors. For example, the mailbox ARN is no longer valid.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailboxArn`

The Amazon Resource Name (ARN) of a WorkMail organization to deliver the email to.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role to use to execute this action. The role must have access to
the workmail:DeliverToMailbox API.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchiveAction

DeliverToQBusinessAction

All content copied from https://docs.aws.amazon.com/.
