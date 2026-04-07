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

_Type_: [EncryptionAtRest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-encryptionatrest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionInTransit`

The details for encryption in transit.

_Required_: No

_Type_: [EncryptionInTransit](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-encryptionintransit.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EncryptionAtRest

EncryptionInTransit
