This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership InboundEdiOptions

Contains options for processing inbound EDI files. These options allow for customizing
how incoming EDI documents are processed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "X12" : X12InboundEdiOptions
}

```

### YAML

```yaml

  X12:
    X12InboundEdiOptions

```

## Properties

`X12`

A structure that contains X12-specific options for processing inbound X12 EDI
files.

_Required_: No

_Type_: [X12InboundEdiOptions](aws-properties-b2bi-partnership-x12inboundedioptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapabilityOptions

OutboundEdiOptions

All content copied from https://docs.aws.amazon.com/.
