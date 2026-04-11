This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel TtmlDestinationSettings

The setup of TTML captions in the output.

The parent of this entity is CaptionDestinationSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StyleControl" : String
}

```

### YAML

```yaml

  StyleControl: String

```

## Properties

`StyleControl`

When set to passthrough, passes through style and position information from a TTML-like input source (TTML,
SMPTE-TT, CFF-TT) to the CFF-TT output or TTML output.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimecodeConfig

UdpContainerSettings

All content copied from https://docs.aws.amazon.com/.
