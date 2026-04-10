This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection Phase1DHGroupNumbersRequestListValue

Specifies a Diffie-Hellman group number for the VPN tunnel for phase 1 IKE
negotiations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : Integer
}

```

### YAML

```yaml

  Value: Integer

```

## Properties

`Value`

The Diffie-Hellmann group number.

_Required_: No

_Type_: Integer

_Allowed values_: `2 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IKEVersionsRequestListValue

Phase1EncryptionAlgorithmsRequestListValue

All content copied from https://docs.aws.amazon.com/.
