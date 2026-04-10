This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule Rule

Receipt rules enable you to specify which actions Amazon SES should take when it receives
mail on behalf of one or more email addresses or domains that you own.

Each receipt rule defines a set of email addresses or domains that it applies to. If
the email addresses or domains match at least one recipient address of the message,
Amazon SES executes all of the receipt rule's actions on the message.

For information about setting up receipt rules, see the [Amazon SES\
Developer Guide](../../../ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ Action, ... ],
  "Enabled" : Boolean,
  "Name" : String,
  "Recipients" : [ String, ... ],
  "ScanEnabled" : Boolean,
  "TlsPolicy" : String
}

```

### YAML

```yaml

  Actions:
    - Action
  Enabled: Boolean
  Name: String
  Recipients:
    - String
  ScanEnabled: Boolean
  TlsPolicy: String

```

## Properties

`Actions`

An ordered list of actions to perform on messages that match at least one of the
recipient email addresses or domains specified in the receipt rule.

_Required_: No

_Type_: Array of [Action](aws-properties-ses-receiptrule-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

If `true`, the receipt rule is active. The default value is
`false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the receipt rule. The name must meet the following requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (\_), dashes
(-), or periods (.).

- Start and end with a letter or number.

- Contain 64 characters or fewer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Recipients`

The recipient domains and email addresses that the receipt rule applies to. If this
field is not specified, this rule matches all recipients on all verified domains.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScanEnabled`

If `true`, then messages that this receipt rule applies to are scanned for
spam and viruses. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlsPolicy`

Specifies whether Amazon SES should require that incoming email is delivered over a
connection encrypted with Transport Layer Security (TLS). If this parameter is set to
`Require`, Amazon SES bounces emails that are not received over TLS. The
default is `Optional`.

Valid Values: `Require | Optional`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaAction

S3Action

All content copied from https://docs.aws.amazon.com/.
