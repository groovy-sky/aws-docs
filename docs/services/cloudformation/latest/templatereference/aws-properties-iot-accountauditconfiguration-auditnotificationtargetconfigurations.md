This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration AuditNotificationTargetConfigurations

The configuration of the audit notification target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Sns" : AuditNotificationTarget
}

```

### YAML

```yaml

  Sns:
    AuditNotificationTarget

```

## Properties

`Sns`

The `Sns` notification target.

_Required_: No

_Type_: [AuditNotificationTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-accountauditconfiguration-auditnotificationtarget.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AuditNotificationTarget

CertAgeCheckCustomConfiguration
