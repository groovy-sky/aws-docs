This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12InboundEdiOptions

Contains options specific to processing inbound X12 EDI files.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcknowledgmentOptions" : X12AcknowledgmentOptions
}

```

### YAML

```yaml

  AcknowledgmentOptions:
    X12AcknowledgmentOptions

```

## Properties

`AcknowledgmentOptions`

Specifies acknowledgment options for inbound X12 EDI files. These options control how
functional and technical acknowledgments are handled.

_Required_: No

_Type_: [X12AcknowledgmentOptions](aws-properties-b2bi-partnership-x12acknowledgmentoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12FunctionalGroupHeaders

X12InterchangeControlHeaders

All content copied from https://docs.aws.amazon.com/.
