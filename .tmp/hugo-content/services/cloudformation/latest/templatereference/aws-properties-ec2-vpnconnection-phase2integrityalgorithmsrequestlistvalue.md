This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection Phase2IntegrityAlgorithmsRequestListValue

Specifies the integrity algorithm for the VPN tunnel for phase 2 IKE
negotiations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : String
}

```

### YAML

```yaml

  Value: String

```

## Properties

`Value`

The integrity algorithm.

_Required_: No

_Type_: String

_Allowed values_: `SHA1 | SHA2-256 | SHA2-384 | SHA2-512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Phase2EncryptionAlgorithmsRequestListValue

Tag

All content copied from https://docs.aws.amazon.com/.
