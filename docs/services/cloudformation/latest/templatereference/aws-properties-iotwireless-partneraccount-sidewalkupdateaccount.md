This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::PartnerAccount SidewalkUpdateAccount

Sidewalk update.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppServerPrivateKey" : String
}

```

### YAML

```yaml

  AppServerPrivateKey: String

```

## Properties

`AppServerPrivateKey`

The new Sidewalk application server private key.

_Required_: No

_Type_: String

_Pattern_: `[a-fA-F0-9]{64}`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SidewalkAccountInfoWithFingerprint

Tag
