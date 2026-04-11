This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice AbpV11

ABP device object for create APIs for v1.1.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DevAddr" : String,
  "SessionKeys" : SessionKeysAbpV11
}

```

### YAML

```yaml

  DevAddr: String
  SessionKeys:
    SessionKeysAbpV11

```

## Properties

`DevAddr`

The DevAddr value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{8}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionKeys`

Session keys for ABP v1.1.

_Required_: Yes

_Type_: [SessionKeysAbpV11](aws-properties-iotwireless-wirelessdevice-sessionkeysabpv11.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AbpV10x

Application

All content copied from https://docs.aws.amazon.com/.
