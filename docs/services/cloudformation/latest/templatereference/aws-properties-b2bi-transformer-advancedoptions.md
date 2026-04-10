This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer AdvancedOptions

A structure that contains advanced options for EDI processing. Currently, only X12
advanced options are supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "X12" : X12AdvancedOptions
}

```

### YAML

```yaml

  X12:
    X12AdvancedOptions

```

## Properties

`X12`

A structure that contains X12-specific advanced options, such as split options for
processing X12 EDI files.

_Required_: No

_Type_: [X12AdvancedOptions](aws-properties-b2bi-transformer-x12advancedoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::B2BI::Transformer

FormatOptions

All content copied from https://docs.aws.amazon.com/.
