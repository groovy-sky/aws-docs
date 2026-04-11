This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration DeviceCertAgeAuditCheckConfiguration

Configuration for the device certificate age audit check.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : CertAgeCheckCustomConfiguration,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Configuration:
    CertAgeCheckCustomConfiguration
  Enabled: Boolean

```

## Properties

`Configuration`

Configuration settings for the device certificate age check, including the threshold in days for certificate age. This configuration is of type `CertAgeCheckCustomConfiguration`.

_Required_: No

_Type_: [CertAgeCheckCustomConfiguration](aws-properties-iot-accountauditconfiguration-certagecheckcustomconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

True if this audit check is enabled for this account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertExpirationCheckCustomConfiguration

DeviceCertExpirationAuditCheckConfiguration

All content copied from https://docs.aws.amazon.com/.
