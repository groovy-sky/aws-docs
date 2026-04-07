This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CertificateManager::Account ExpiryEventsConfiguration

Object containing expiration events options associated with an AWS account. For more information, see [ExpiryEventsConfiguration](../../../../reference/acm/latest/apireference/api-expiryeventsconfiguration.md) in the API reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DaysBeforeExpiry" : Integer
}

```

### YAML

```yaml

  DaysBeforeExpiry: Integer

```

## Properties

`DaysBeforeExpiry`

This option specifies the number of days prior to certificate expiration when ACM
starts generating `EventBridge` events. ACM sends one event per day per
certificate until the certificate expires. By default, accounts receive events starting
45 days before certificate expiration.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `45`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CertificateManager::Account

AWS::CertificateManager::Certificate
