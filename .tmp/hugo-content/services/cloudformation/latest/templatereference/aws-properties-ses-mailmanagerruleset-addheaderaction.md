This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet AddHeaderAction

The action to add a header to a message. When executed, this action will add the given
header to the message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderName" : String,
  "HeaderValue" : String
}

```

### YAML

```yaml

  HeaderName: String
  HeaderValue: String

```

## Properties

`HeaderName`

The name of the header to add to an email. The header must be prefixed with
"X-". Headers are added regardless of whether the header name pre-existed in
the email.

_Required_: Yes

_Type_: String

_Pattern_: `^[xX]\-[a-zA-Z0-9\-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderValue`

The value of the header to add to the email.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::MailManagerRuleSet

Analysis

All content copied from https://docs.aws.amazon.com/.
