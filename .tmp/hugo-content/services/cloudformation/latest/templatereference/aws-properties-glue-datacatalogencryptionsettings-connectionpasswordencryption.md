This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DataCatalogEncryptionSettings ConnectionPasswordEncryption

The data structure used by the Data Catalog to encrypt the password as part of
`CreateConnection` or `UpdateConnection` and store it in the
`ENCRYPTED_PASSWORD` field in the connection properties. You can enable catalog
encryption or only password encryption.

When a `CreationConnection` request arrives containing a password, the Data
Catalog first encrypts the password using your AWS KMS key. It then encrypts the whole
connection object again if catalog encryption is also enabled.

This encryption requires that you set AWS KMS key permissions to enable or restrict access
on the password key according to your security requirements. For example, you might want only
administrators to have decrypt permission on the password key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "ReturnConnectionPasswordEncrypted" : Boolean
}

```

### YAML

```yaml

  KmsKeyId: String
  ReturnConnectionPasswordEncrypted: Boolean

```

## Properties

`KmsKeyId`

An AWS KMS key that is used to encrypt the connection password.

If connection password protection is enabled, the caller of
`CreateConnection` and `UpdateConnection` needs at least
`kms:Encrypt` permission on the specified AWS KMS key, to encrypt
passwords before storing them in the Data Catalog. You can set the decrypt permission to
enable or restrict access on the password key according to your security
requirements.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReturnConnectionPasswordEncrypted`

When the `ReturnConnectionPasswordEncrypted` flag is set to "true", passwords remain encrypted in the responses of `GetConnection` and `GetConnections`. This encryption takes effect independently from catalog encryption.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::DataCatalogEncryptionSettings

DataCatalogEncryptionSettings

All content copied from https://docs.aws.amazon.com/.
