This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::PartnerAccount SidewalkAccountInfoWithFingerprint

Information about a Sidewalk account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmazonId" : String,
  "Arn" : String,
  "Fingerprint" : String
}

```

### YAML

```yaml

  AmazonId: String
  Arn: String
  Fingerprint: String

```

## Properties

`AmazonId`

The Sidewalk Amazon ID.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Arn`

The Amazon Resource Name (ARN) of the resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fingerprint`

The fingerprint of the Sidewalk application server private key.

_Required_: No

_Type_: String

_Pattern_: `[a-fA-F0-9]{64}`

_Minimum_: `64`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SidewalkAccountInfo

SidewalkUpdateAccount
