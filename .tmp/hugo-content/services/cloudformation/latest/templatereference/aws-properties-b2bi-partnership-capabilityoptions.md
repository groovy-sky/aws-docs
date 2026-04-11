This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership CapabilityOptions

Contains the details for an Outbound EDI capability.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InboundEdi" : InboundEdiOptions,
  "OutboundEdi" : OutboundEdiOptions
}

```

### YAML

```yaml

  InboundEdi:
    InboundEdiOptions
  OutboundEdi:
    OutboundEdiOptions

```

## Properties

`InboundEdi`

A structure that contains the inbound EDI options for the capability.

_Required_: No

_Type_: [InboundEdiOptions](aws-properties-b2bi-partnership-inboundedioptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundEdi`

A structure that contains the outbound EDI options.

_Required_: No

_Type_: [OutboundEdiOptions](aws-properties-b2bi-partnership-outboundedioptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::B2BI::Partnership

InboundEdiOptions

All content copied from https://docs.aws.amazon.com/.
