This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput RouterInputTransitEncryption

The transit encryption settings for a router input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionKeyConfiguration" : RouterInputTransitEncryptionKeyConfiguration,
  "EncryptionKeyType" : String
}

```

### YAML

```yaml

  EncryptionKeyConfiguration:
    RouterInputTransitEncryptionKeyConfiguration
  EncryptionKeyType: String

```

## Properties

`EncryptionKeyConfiguration`

Contains the configuration details for the encryption key used in transit encryption, including the key source and associated parameters.

_Required_: Yes

_Type_: [RouterInputTransitEncryptionKeyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-routerinput-routerinputtransitencryptionkeyconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyType`

Specifies the type of encryption key to use for transit encryption.

_Required_: No

_Type_: String

_Allowed values_: `SECRETS_MANAGER | AUTOMATIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RouterInputProtocolConfiguration

RouterInputTransitEncryptionKeyConfiguration
