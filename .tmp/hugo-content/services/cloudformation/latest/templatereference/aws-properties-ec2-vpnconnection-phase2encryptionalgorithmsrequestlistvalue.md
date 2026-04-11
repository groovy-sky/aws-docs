This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection Phase2EncryptionAlgorithmsRequestListValue

Specifies the encryption algorithm for the VPN tunnel for phase 2 IKE
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

The encryption algorithm.

_Required_: No

_Type_: String

_Allowed values_: `AES128 | AES256 | AES128-GCM-16 | AES256-GCM-16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Phase2DHGroupNumbersRequestListValue

Phase2IntegrityAlgorithmsRequestListValue

All content copied from https://docs.aws.amazon.com/.
