This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration DeviceCertExpirationAuditCheckConfiguration

Configuration for the device certificate expiration audit check.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : CertExpirationCheckCustomConfiguration,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Configuration:
    CertExpirationCheckCustomConfiguration
  Enabled: Boolean

```

## Properties

`Configuration`

Configuration settings for the device certificate expiration check, including the threshold in days before expiration. This configuration is of type `CertExpirationCheckCustomConfiguration`

_Required_: No

_Type_: [CertExpirationCheckCustomConfiguration](aws-properties-iot-accountauditconfiguration-certexpirationcheckcustomconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

True if this audit check is enabled for this account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeviceCertAgeAuditCheckConfiguration

AWS::IoT::Authorizer

All content copied from https://docs.aws.amazon.com/.
