This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign Template

Specifies the name and version of the message template to use for the
message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Version" : String
}

```

### YAML

```yaml

  Name: String
  Version: String

```

## Properties

`Name`

The name of the message template to use for the message. If specified, this
value must match the name of an existing message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The unique identifier for the version of the message template to use for the
message. If specified, this value must match the identifier for an existing
template version. To retrieve a list of versions and version identifiers for a
template, use the [Template Versions](../../../../reference/pinpoint/latest/apireference/templates-template-name-template-type-versions.md) resource.

If you don't specify a value for this property, Amazon Pinpoint uses the
_active version_ of the template. The _active_
_version_ is typically the version of a template that's been most
recently reviewed and approved for use, depending on your workflow. It isn't
necessarily the latest version of a template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SetDimension

TemplateConfiguration

All content copied from https://docs.aws.amazon.com/.
