This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer InputConversion

Contains the input formatting options for an inbound transformer (takes an X12-formatted
EDI document as input and converts it to JSON or XML.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvancedOptions" : AdvancedOptions,
  "FormatOptions" : FormatOptions,
  "FromFormat" : String
}

```

### YAML

```yaml

  AdvancedOptions:
    AdvancedOptions
  FormatOptions:
    FormatOptions
  FromFormat: String

```

## Properties

`AdvancedOptions`

Specifies advanced options for the input conversion process. These options provide
additional control over how EDI files are processed during transformation.

_Required_: No

_Type_: [AdvancedOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-advancedoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatOptions`

A structure that contains the formatting options for an inbound transformer.

_Required_: No

_Type_: [FormatOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-formatoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FromFormat`

The format for the transformer input: currently on `X12` is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `X12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FormatOptions

Mapping
