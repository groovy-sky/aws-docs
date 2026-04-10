This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput SrtEncryptionConfiguration

Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionKey" : SecretsManagerEncryptionKeyConfiguration
}

```

### YAML

```yaml

  EncryptionKey:
    SecretsManagerEncryptionKeyConfiguration

```

## Properties

`EncryptionKey`

Specifies the encryption key configuration used for encrypting SRT streams, including the key source and associated credentials.

_Required_: Yes

_Type_: [SecretsManagerEncryptionKeyConfiguration](aws-properties-mediaconnect-routeroutput-secretsmanagerencryptionkeyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SrtCallerRouterOutputConfiguration

SrtListenerRouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
