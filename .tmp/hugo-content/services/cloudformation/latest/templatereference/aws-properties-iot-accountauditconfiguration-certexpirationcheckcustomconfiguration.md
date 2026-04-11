This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration CertExpirationCheckCustomConfiguration

Configuration structure containing settings for the device certificate expiration check.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertExpirationThresholdInDays" : String
}

```

### YAML

```yaml

  CertExpirationThresholdInDays: String

```

## Properties

`CertExpirationThresholdInDays`

The number of days before expiration that defines when a device certificate is considered to be approaching expiration. The check will report a finding if a certificate will expire within this number of days.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertAgeCheckCustomConfiguration

DeviceCertAgeAuditCheckConfiguration

All content copied from https://docs.aws.amazon.com/.
