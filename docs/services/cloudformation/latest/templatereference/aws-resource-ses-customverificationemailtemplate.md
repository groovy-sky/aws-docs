This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::CustomVerificationEmailTemplate

Represents a request to create a custom verification email template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::CustomVerificationEmailTemplate",
  "Properties" : {
      "FailureRedirectionURL" : String,
      "FromEmailAddress" : String,
      "SuccessRedirectionURL" : String,
      "Tags" : [ Tag, ... ],
      "TemplateContent" : String,
      "TemplateName" : String,
      "TemplateSubject" : String
    }
}

```

### YAML

```yaml

Type: AWS::SES::CustomVerificationEmailTemplate
Properties:
  FailureRedirectionURL: String
  FromEmailAddress: String
  SuccessRedirectionURL: String
  Tags:
    - Tag
  TemplateContent: String
  TemplateName: String
  TemplateSubject: String

```

## Properties

`FailureRedirectionURL`

The URL that the recipient of the verification email is sent to if his or her address
is not successfully verified.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromEmailAddress`

The email address that the custom verification email is sent from.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessRedirectionURL`

The URL that the recipient of the verification email is sent to if his or her address
is successfully verified.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of objects that define the tags (keys and values) to associate with the
custom verification email template.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-customverificationemailtemplate-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateContent`

The content of the custom verification email. The total size of the email must be less
than 10 MB. The message body may contain HTML, with some limitations. For more
information, see [Custom verification email frequently asked questions](../../../ses/latest/dg/creating-identities.md#send-email-verify-address-custom-faq) in the _Amazon SES_
_Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the custom verification email template.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateSubject`

The subject line of the custom verification email.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the template name. For example:

`{ "Ref": "MyCustomVerificationEmailTemplate" }`

For the Amazon SES custom verification email template
`MyCustomVerificationEmailTemplate`, `Ref` returns the name
of the custom verification email template.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Topic

Tag

All content copied from https://docs.aws.amazon.com/.
