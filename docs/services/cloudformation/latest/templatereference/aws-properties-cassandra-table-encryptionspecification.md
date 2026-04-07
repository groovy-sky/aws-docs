This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table EncryptionSpecification

Specifies the encryption at rest option selected for the table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionType" : String,
  "KmsKeyIdentifier" : String
}

```

### YAML

```yaml

  EncryptionType: String
  KmsKeyIdentifier: String

```

## Properties

`EncryptionType`

The encryption at rest options for the table.

- **AWS owned key** (default) - `AWS_OWNED_KMS_KEY`

- **Customer managed key** \- `CUSTOMER_MANAGED_KMS_KEY`

###### Important

If you choose `CUSTOMER_MANAGED_KMS_KEY`, a `kms_key_identifier` in the format of a
key ARN is required.

Valid values: `CUSTOMER_MANAGED_KMS_KEY` \| `AWS_OWNED_KMS_KEY`.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_OWNED_KMS_KEY | CUSTOMER_MANAGED_KMS_KEY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyIdentifier`

Requires a `kms_key_identifier` in the format of a
key ARN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Column

ProvisionedThroughput
