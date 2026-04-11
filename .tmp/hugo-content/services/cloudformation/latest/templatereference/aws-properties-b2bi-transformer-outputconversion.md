This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer OutputConversion

Contains the formatting options for an outbound transformer (takes JSON or XML as input
and converts it to an EDI document (currently only X12 format is supported).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvancedOptions" : AdvancedOptions,
  "FormatOptions" : FormatOptions,
  "ToFormat" : String
}

```

### YAML

```yaml

  AdvancedOptions:
    AdvancedOptions
  FormatOptions:
    FormatOptions
  ToFormat: String

```

## Properties

`AdvancedOptions`

Property description not available.

_Required_: No

_Type_: [AdvancedOptions](aws-properties-b2bi-transformer-advancedoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatOptions`

A structure that contains the X12 transaction set and version for the transformer
output.

_Required_: No

_Type_: [FormatOptions](aws-properties-b2bi-transformer-formatoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToFormat`

The format for the output from an outbound transformer: only X12 is currently
supported.

_Required_: Yes

_Type_: String

_Allowed values_: `X12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Mapping

SampleDocumentKeys

All content copied from https://docs.aws.amazon.com/.
