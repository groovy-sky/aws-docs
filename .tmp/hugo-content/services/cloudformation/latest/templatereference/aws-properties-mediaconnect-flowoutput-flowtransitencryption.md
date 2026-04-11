This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowOutput FlowTransitEncryption

The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionKeyConfiguration" : FlowTransitEncryptionKeyConfiguration,
  "EncryptionKeyType" : String
}

```

### YAML

```yaml

  EncryptionKeyConfiguration:
    FlowTransitEncryptionKeyConfiguration
  EncryptionKeyType: String

```

## Properties

`EncryptionKeyConfiguration`

The configuration details for the encryption key.

_Required_: Yes

_Type_: [FlowTransitEncryptionKeyConfiguration](aws-properties-mediaconnect-flowoutput-flowtransitencryptionkeyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyType`

The type of encryption key to use for flow transit encryption.

_Required_: No

_Type_: String

_Allowed values_: `SECRETS_MANAGER | AUTOMATIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Encryption

FlowTransitEncryptionKeyConfiguration

All content copied from https://docs.aws.amazon.com/.
