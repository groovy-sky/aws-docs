This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::EmailTemplate

Creates a message template that you can use in messages that are sent through the
email channel. A _message template_ is a set of content and settings
that you can define, save, and reuse in messages for any of your Amazon Pinpoint
applications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::EmailTemplate",
  "Properties" : {
      "DefaultSubstitutions" : String,
      "HtmlPart" : String,
      "Subject" : String,
      "Tags" : [ Tag, ... ],
      "TemplateDescription" : String,
      "TemplateName" : String,
      "TextPart" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::EmailTemplate
Properties:
  DefaultSubstitutions: String
  HtmlPart: String
  Subject: String
  Tags:
    - Tag
  TemplateDescription: String
  TemplateName: String
  TextPart: String

```

## Properties

`DefaultSubstitutions`

A JSON object that specifies the default values to use for message variables
in the message template. This object is a set of key-value pairs. Each key
defines a message variable in the template. The corresponding value defines the
default value for that variable. When you create a message that's based on the
template, you can override these defaults with message-specific and
address-specific variables and values.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HtmlPart`

The message body, in HTML format, to use in email messages that are based on
the message template. We recommend using HTML format for email clients that
render HTML content. You can include links, formatted text, and more in an HTML
message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subject`

The subject line, or title, to use in email messages that are based on the
message template.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateDescription`

A custom description of the message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the message template.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TextPart`

The message body, in plain text format, to use in email messages that are
based on the message template. We recommend using plain text format for email
clients that don't render HTML content and clients that are connected to
high-latency networks, such as mobile devices.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the name of the message template
( `TemplateName`).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the message template.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Pinpoint::EmailChannel

AWS::Pinpoint::EventStream

All content copied from https://docs.aws.amazon.com/.
