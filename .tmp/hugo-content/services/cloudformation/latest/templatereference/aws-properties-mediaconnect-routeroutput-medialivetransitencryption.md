This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput MediaLiveTransitEncryption

The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive. This configuration determines whether encryption keys are automatically managed by the service or manually managed through AWS Secrets Manager.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionKeyConfiguration" : MediaLiveTransitEncryptionKeyConfiguration,
  "EncryptionKeyType" : String
}

```

### YAML

```yaml

  EncryptionKeyConfiguration:
    MediaLiveTransitEncryptionKeyConfiguration
  EncryptionKeyType: String

```

## Properties

`EncryptionKeyConfiguration`

The configuration details for the MediaLive encryption key.

_Required_: Yes

_Type_: [MediaLiveTransitEncryptionKeyConfiguration](aws-properties-mediaconnect-routeroutput-medialivetransitencryptionkeyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyType`

The type of encryption key to use for MediaLive transit encryption.

_Required_: No

_Type_: String

_Allowed values_: `SECRETS_MANAGER | AUTOMATIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaLiveInputRouterOutputConfiguration

MediaLiveTransitEncryptionKeyConfiguration

All content copied from https://docs.aws.amazon.com/.
