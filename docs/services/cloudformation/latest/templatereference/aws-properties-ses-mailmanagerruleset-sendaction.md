This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet SendAction

Sends the email to the internet using the ses:SendRawEmail API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  RoleArn: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there
are configuration errors. For example, the caller does not have the permissions to call
the sendRawEmail API.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the role to use for this action. This role must have access to the
ses:SendRawEmail API.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Action

SnsAction

All content copied from https://docs.aws.amazon.com/.
