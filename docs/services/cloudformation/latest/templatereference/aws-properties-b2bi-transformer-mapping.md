This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer Mapping

Specifies the mapping template for the transformer. This template is used to map the
parsed EDI file using JSONata or XSLT.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Template" : String,
  "TemplateLanguage" : String
}

```

### YAML

```yaml

  Template: String
  TemplateLanguage: String

```

## Properties

`Template`

A string that represents the mapping template, in the transformation language specified
in `templateLanguage`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `350000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateLanguage`

The transformation language for the template, either XSLT or JSONATA.

_Required_: Yes

_Type_: String

_Allowed values_: `XSLT | JSONATA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputConversion

OutputConversion

All content copied from https://docs.aws.amazon.com/.
