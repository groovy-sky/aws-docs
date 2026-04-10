This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice SessionKeysAbpV10x

Session keys for ABP v1.0.x.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppSKey" : String,
  "NwkSKey" : String
}

```

### YAML

```yaml

  AppSKey: String
  NwkSKey: String

```

## Properties

`AppSKey`

The AppSKey value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NwkSKey`

The NwkKey value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OtaaV11

SessionKeysAbpV11

All content copied from https://docs.aws.amazon.com/.
