This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12Envelope

A wrapper structure for an X12 definition object.

the X12 envelope ensures the integrity of the data and the efficiency of the information
exchange. The X12 message structure has hierarchical levels. From highest to the lowest,
they are:

- Interchange Envelope

- Functional Group

- Transaction Set

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Common" : X12OutboundEdiHeaders,
  "WrapOptions" : WrapOptions
}

```

### YAML

```yaml

  Common:
    X12OutboundEdiHeaders
  WrapOptions:
    WrapOptions

```

## Properties

`Common`

A container for the X12 outbound EDI headers.

_Required_: No

_Type_: [X12OutboundEdiHeaders](aws-properties-b2bi-partnership-x12outboundediheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WrapOptions`

Property description not available.

_Required_: No

_Type_: [WrapOptions](aws-properties-b2bi-partnership-wrapoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12Delimiters

X12FunctionalGroupHeaders

All content copied from https://docs.aws.amazon.com/.
