This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster EncryptionInfo

Includes encryption-related information, such as the Amazon KMS key used for encrypting data at rest and whether you want MSK to encrypt your data in transit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionAtRest" : EncryptionAtRest,
  "EncryptionInTransit" : EncryptionInTransit
}

```

### YAML

```yaml

  EncryptionAtRest:
    EncryptionAtRest
  EncryptionInTransit:
    EncryptionInTransit

```

## Properties

`EncryptionAtRest`

The data-volume encryption details.

_Required_: No

_Type_: [EncryptionAtRest](aws-properties-msk-cluster-encryptionatrest.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionInTransit`

The details for encryption in transit.

_Required_: No

_Type_: [EncryptionInTransit](aws-properties-msk-cluster-encryptionintransit.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionAtRest

EncryptionInTransit

All content copied from https://docs.aws.amazon.com/.
