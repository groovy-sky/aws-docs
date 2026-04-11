This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration CertAgeCheckCustomConfiguration

Configuration structure containing settings for the device certificate age check.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertAgeThresholdInDays" : String
}

```

### YAML

```yaml

  CertAgeThresholdInDays: String

```

## Properties

`CertAgeThresholdInDays`

The number of days that defines when a device certificate is considered to have aged. The check will report a finding if a certificate has been active for a number of days greater than or equal to this threshold value.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuditNotificationTargetConfigurations

CertExpirationCheckCustomConfiguration

All content copied from https://docs.aws.amazon.com/.
