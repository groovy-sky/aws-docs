This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::Template Template

An object that defines the email template to use for an email message, and the values
to use for any message variables in that template. An _email_
_template_ is a type of message template that contains content that you
want to reuse in email messages that you send. You can specifiy the email template by providing
the name or ARN of an _email template_
previously saved in your Amazon SES account or by providing the full template content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HtmlPart" : String,
  "SubjectPart" : String,
  "TemplateName" : String,
  "TextPart" : String
}

```

### YAML

```yaml

  HtmlPart: String
  SubjectPart: String
  TemplateName: String
  TextPart: String

```

## Properties

`HtmlPart`

The HTML body of the email.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubjectPart`

The subject line of the email.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the template. You will refer to this name when you send email using the
`SendEmail` or `SendBulkEmail` operations.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,64}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TextPart`

The email body that is visible to recipients whose email clients do not display HTML
content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::SES::Tenant
