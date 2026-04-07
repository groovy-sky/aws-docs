This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DataCatalogEncryptionSettings DataCatalogEncryptionSettings

Contains configuration information for maintaining Data Catalog security.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionPasswordEncryption" : ConnectionPasswordEncryption,
  "EncryptionAtRest" : EncryptionAtRest
}

```

### YAML

```yaml

  ConnectionPasswordEncryption:
    ConnectionPasswordEncryption
  EncryptionAtRest:
    EncryptionAtRest

```

## Properties

`ConnectionPasswordEncryption`

When connection password protection is enabled, the Data Catalog uses a customer-provided
key to encrypt the password as part of `CreateConnection` or
`UpdateConnection` and store it in the `ENCRYPTED_PASSWORD` field in
the connection properties. You can enable catalog encryption or only password
encryption.

_Required_: No

_Type_: [ConnectionPasswordEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionAtRest`

Specifies the encryption-at-rest configuration for the Data Catalog.

_Required_: No

_Type_: [EncryptionAtRest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-datacatalogencryptionsettings-encryptionatrest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectionPasswordEncryption

EncryptionAtRest
