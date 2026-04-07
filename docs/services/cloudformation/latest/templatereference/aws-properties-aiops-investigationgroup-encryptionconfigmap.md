This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AIOps::InvestigationGroup EncryptionConfigMap

Use this structure if you want to use a customer managed AWS KMS key to
encrypt your investigation data. If you omit this parameter, CloudWatch investigations will use an AWS key to encrypt the data. For more information, see [Encryption of investigation data](../../../amazoncloudwatch/latest/monitoring/investigations-security.md#Investigations-KMS).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionConfigurationType" : String,
  "KmsKeyId" : String
}

```

### YAML

```yaml

  EncryptionConfigurationType: String
  KmsKeyId: String

```

## Properties

`EncryptionConfigurationType`

Displays whether investigation data is encrypted by a customer managed key or an AWS owned key.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

If the investigation group uses a customer managed key for encryption, this field
displays the ID of that key.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CrossAccountConfiguration

Tag
