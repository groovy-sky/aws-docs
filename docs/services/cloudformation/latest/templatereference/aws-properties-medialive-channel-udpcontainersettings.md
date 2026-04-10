This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel UdpContainerSettings

The configuration of a UDP output.

The parent of this entity is UdpOutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "M2tsSettings" : M2tsSettings
}

```

### YAML

```yaml

  M2tsSettings:
    M2tsSettings

```

## Properties

`M2tsSettings`

The M2TS configuration for this UDP output.

_Required_: No

_Type_: [M2tsSettings](aws-properties-medialive-channel-m2tssettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TtmlDestinationSettings

UdpGroupSettings

All content copied from https://docs.aws.amazon.com/.
