This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration

Use the `AWS::IoT::AccountAuditConfiguration` resource to configure or
reconfigure the Device Defender audit settings for your account. Settings include how audit
notifications are sent and which audit checks are enabled or disabled. For API reference,
see [UpdateAccountAuditConfiguration](https://docs.aws.amazon.com/iot/latest/apireference/API_UpdateAccountAuditConfiguration.html) and for detailed information on all available
audit checks, see [Audit\
checks](https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-audit-checks.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::AccountAuditConfiguration",
  "Properties" : {
      "AccountId" : String,
      "AuditCheckConfigurations" : AuditCheckConfigurations,
      "AuditNotificationTargetConfigurations" : AuditNotificationTargetConfigurations,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::AccountAuditConfiguration
Properties:
  AccountId: String
  AuditCheckConfigurations:
    AuditCheckConfigurations
  AuditNotificationTargetConfigurations:
    AuditNotificationTargetConfigurations
  RoleArn: String

```

## Properties

`AccountId`

The ID of the account. You can use the expression `!Sub "${AWS::AccountId}"`
to use your account ID.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuditCheckConfigurations`

Specifies which audit checks are enabled and disabled for this account.

Some data collection might start immediately when certain checks are enabled. When a
check is disabled, any data collected so far in relation to the check is deleted. To
disable a check, set the value of the `Enabled:` key to
`false`.

If an enabled check is removed from the template, it will also be disabled.

You can't disable a check if it's used by any scheduled audit. You must delete the check
from the scheduled audit or delete the scheduled audit itself to disable the check.

For more information on available audit checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html)

_Required_: Yes

_Type_: [AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuditNotificationTargetConfigurations`

Information about the targets to which audit notifications are sent.

_Required_: No

_Type_: [AuditNotificationTargetConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-accountauditconfiguration-auditnotificationtargetconfigurations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the role that grants permission to AWS IoT to access information about your devices, policies, certificates, and other items as
required when performing an audit.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the account ID.

## Examples

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Amazon Web Services IoT AccountAuditConfiguration Sample Template",
  "Resources": {
    "MyAccountAuditConfiguration": {
      "Type": "AWS::IoT::AccountAuditConfiguration",
      "Properties": {
        "AccountId": "${AWS::AccountId}",
        "AuditCheckConfigurations": {
          "AuthenticatedCognitoRoleOverlyPermissiveCheck": {
            "Enabled": true
          },
          "CaCertificateExpiringCheck": {
            "Enabled": true
          },
          "CaCertificateKeyQualityCheck": {
            "Enabled": true
          },
          "ConflictingClientIdsCheck": {
            "Enabled": true
          },
          "DeviceCertificateAgeCheck": {
            "Enabled": true,
            "Configuration": {
              "CertAgeThresholdInDays": 60
            }
          },
          "DeviceCertificateExpiringCheck": {
            "Enabled": true,
            "Configuration": {
              "CertExpirationThresholdInDays": 30
            }
          },
          "DeviceCertificateKeyQualityCheck": {
            "Enabled": true
          },
          "DeviceCertificateSharedCheck": {
            "Enabled": true
          },
          "IotPolicyOverlyPermissiveCheck": {
            "Enabled": true
          },
          "IotRoleAliasAllowsAccessToUnusedServicesCheck": {
            "Enabled": true
          },
          "IotRoleAliasOverlyPermissiveCheck": {
            "Enabled": true
          },
          "LoggingDisabledCheck": {
            "Enabled": true
          },
          "RevokedCaCertificateStillActiveCheck": {
            "Enabled": true
          },
          "RevokedDeviceCertificateStillActiveCheck": {
            "Enabled": true
          },
          "UnauthenticatedCognitoRoleOverlyPermissiveCheck": {
            "Enabled": true
          }
        },
        "AuditNotificationTargetConfigurations": {
          "Sns": {
            "TargetArn": "arn:aws:sns:us-east-1:123456789012:AuditNotifications",
            "RoleArn": "arn:aws:iam::123456789012:role/RoleForIoTAuditNotifications",
            "Enabled": true
          }
        },
        "RoleArn": "arn:aws:iam::123456789012:role/service-role/AWSIoTDeviceDefenderAudit"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Amazon Web Services IoT AccountAuditConfiguration Sample Template
Resources:
  MyAccountAuditConfiguration:
    Type: 'AWS::IoT::AccountAuditConfiguration'
    Properties:
      AccountId: !Sub '${AWS::AccountId}'
      AuditCheckConfigurations:
        AuthenticatedCognitoRoleOverlyPermissiveCheck:
          Enabled: true
        CaCertificateExpiringCheck:
          Enabled: true
        CaCertificateKeyQualityCheck:
          Enabled: true
        ConflictingClientIdsCheck:
          Enabled: true
        DeviceCertificateAgeCheck:
          Enabled: true
          Configuration:
            CertAgeThresholdInDays: 60
        DeviceCertificateExpiringCheck:
          Enabled: true
          Configuration:
            CertExpirationThresholdInDays: 30
        DeviceCertificateKeyQualityCheck:
          Enabled: true
        DeviceCertificateSharedCheck:
          Enabled: true
        IotPolicyOverlyPermissiveCheck:
          Enabled: true
        IotRoleAliasAllowsAccessToUnusedServicesCheck:
          Enabled: true
        IotRoleAliasOverlyPermissiveCheck:
          Enabled: true
        LoggingDisabledCheck:
          Enabled: true
        RevokedCaCertificateStillActiveCheck:
          Enabled: true
        RevokedDeviceCertificateStillActiveCheck:
          Enabled: true
        UnauthenticatedCognitoRoleOverlyPermissiveCheck:
          Enabled: true
      AuditNotificationTargetConfigurations:
        Sns:
          TargetArn: 'arn:aws:sns:us-east-1:123456789012:AuditNotifications'
          RoleArn: 'arn:aws:iam::123456789012:role/RoleForIoTAuditNotifications'
          Enabled: true
      RoleArn: 'arn:aws:iam::123456789012:role/service-role/AWSIoTDeviceDefenderAudit'
```

## See also

When you use CloudFormation to perform drift detection for
`AccountAuditConfiguration`, it won't compare values that aren't part of
the stack template. In `AccountAuditConfiguration`, specifying a
configuration for every check is optional, and skipped checks are interpreted as
disabled. To have accurate drift detection with CloudFormation, include configurations
(enabled or disabled) for all the 14 audit checks in your template. For more information
on the audit checks see [AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-accountauditconfiguration-auditcheckconfigurations.html).

For more information, see [Detecting\
unmanaged configuration changes to stacks and resources](../userguide/using-cfn-stack-drift.md) in the _user_
_guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS IoT

AuditCheckConfiguration
