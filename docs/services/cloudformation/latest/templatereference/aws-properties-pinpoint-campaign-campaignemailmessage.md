This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CampaignEmailMessage

Specifies the content and "From" address for an email message that's sent to
recipients of a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Body" : String,
  "FromAddress" : String,
  "HtmlBody" : String,
  "Title" : String
}

```

### YAML

```yaml

  Body: String
  FromAddress: String
  HtmlBody: String
  Title: String

```

## Properties

`Body`

The body of the email for recipients whose email clients don't render HTML
content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromAddress`

The verified email address to send the email from. The default address is the
`FromAddress` specified for the email channel for the
application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HtmlBody`

The body of the email, in HTML format, for recipients whose email clients
render HTML content.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The subject line, or title, of the email.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CampaignCustomMessage

CampaignEventFilter

All content copied from https://docs.aws.amazon.com/.
