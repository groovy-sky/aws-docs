This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow FlowTransitEncryptionKeyConfiguration

Configuration settings for flow transit encryption keys.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Automatic" : Json,
  "SecretsManager" : SecretsManagerEncryptionKeyConfiguration
}

```

### YAML

```yaml

  Automatic: Json
  SecretsManager:
    SecretsManagerEncryptionKeyConfiguration

```

## Properties

`Automatic`

Configuration settings for automatic encryption key management, where MediaConnect handles key creation and rotation.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretsManager`

The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.

_Required_: No

_Type_: [SecretsManagerEncryptionKeyConfiguration](aws-properties-mediaconnect-flow-secretsmanagerencryptionkeyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowTransitEncryption

Fmtp

All content copied from https://docs.aws.amazon.com/.
