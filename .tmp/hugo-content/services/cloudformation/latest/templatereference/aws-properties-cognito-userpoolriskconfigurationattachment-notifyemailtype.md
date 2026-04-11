This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment NotifyEmailType

The template for email messages that advanced security features sends to a user when
your threat protection automated response has a _Notify_
action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HtmlBody" : String,
  "Subject" : String,
  "TextBody" : String
}

```

### YAML

```yaml

  HtmlBody: String
  Subject: String
  TextBody: String

```

## Properties

`HtmlBody`

The body of an email notification formatted in HTML. Choose an `HtmlBody`
or a `TextBody` to send an HTML-formatted or plaintext message,
respectively.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]+`

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subject`

The subject of the threat protection email notification.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextBody`

The body of an email notification formatted in plaintext. Choose an
`HtmlBody` or a `TextBody` to send an HTML-formatted or
plaintext message, respectively.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]+`

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotifyConfigurationType

RiskExceptionConfigurationType

All content copied from https://docs.aws.amazon.com/.
