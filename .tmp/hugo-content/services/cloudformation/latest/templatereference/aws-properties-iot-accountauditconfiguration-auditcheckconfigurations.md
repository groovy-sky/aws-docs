This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::AccountAuditConfiguration AuditCheckConfigurations

The types of audit checks that can be performed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticatedCognitoRoleOverlyPermissiveCheck" : AuditCheckConfiguration,
  "CaCertificateExpiringCheck" : AuditCheckConfiguration,
  "CaCertificateKeyQualityCheck" : AuditCheckConfiguration,
  "ConflictingClientIdsCheck" : AuditCheckConfiguration,
  "DeviceCertificateAgeCheck" : DeviceCertAgeAuditCheckConfiguration,
  "DeviceCertificateExpiringCheck" : DeviceCertExpirationAuditCheckConfiguration,
  "DeviceCertificateKeyQualityCheck" : AuditCheckConfiguration,
  "DeviceCertificateSharedCheck" : AuditCheckConfiguration,
  "IntermediateCaRevokedForActiveDeviceCertificatesCheck" : AuditCheckConfiguration,
  "IotPolicyOverlyPermissiveCheck" : AuditCheckConfiguration,
  "IoTPolicyPotentialMisConfigurationCheck" : AuditCheckConfiguration,
  "IotRoleAliasAllowsAccessToUnusedServicesCheck" : AuditCheckConfiguration,
  "IotRoleAliasOverlyPermissiveCheck" : AuditCheckConfiguration,
  "LoggingDisabledCheck" : AuditCheckConfiguration,
  "RevokedCaCertificateStillActiveCheck" : AuditCheckConfiguration,
  "RevokedDeviceCertificateStillActiveCheck" : AuditCheckConfiguration,
  "UnauthenticatedCognitoRoleOverlyPermissiveCheck" : AuditCheckConfiguration
}

```

### YAML

```yaml

  AuthenticatedCognitoRoleOverlyPermissiveCheck:
    AuditCheckConfiguration
  CaCertificateExpiringCheck:
    AuditCheckConfiguration
  CaCertificateKeyQualityCheck:
    AuditCheckConfiguration
  ConflictingClientIdsCheck:
    AuditCheckConfiguration
  DeviceCertificateAgeCheck:
    DeviceCertAgeAuditCheckConfiguration
  DeviceCertificateExpiringCheck:
    DeviceCertExpirationAuditCheckConfiguration
  DeviceCertificateKeyQualityCheck:
    AuditCheckConfiguration
  DeviceCertificateSharedCheck:
    AuditCheckConfiguration
  IntermediateCaRevokedForActiveDeviceCertificatesCheck:
    AuditCheckConfiguration
  IotPolicyOverlyPermissiveCheck:
    AuditCheckConfiguration
  IoTPolicyPotentialMisConfigurationCheck:
    AuditCheckConfiguration
  IotRoleAliasAllowsAccessToUnusedServicesCheck:
    AuditCheckConfiguration
  IotRoleAliasOverlyPermissiveCheck:
    AuditCheckConfiguration
  LoggingDisabledCheck:
    AuditCheckConfiguration
  RevokedCaCertificateStillActiveCheck:
    AuditCheckConfiguration
  RevokedDeviceCertificateStillActiveCheck:
    AuditCheckConfiguration
  UnauthenticatedCognitoRoleOverlyPermissiveCheck:
    AuditCheckConfiguration

```

## Properties

`AuthenticatedCognitoRoleOverlyPermissiveCheck`

Checks the permissiveness of an authenticated Amazon Cognito identity pool role. For
this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been
used to connect to the AWS IoT message broker during the 31 days before the
audit is performed.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaCertificateExpiringCheck`

Checks if a CA certificate is expiring. This check applies to CA certificates expiring
within 30 days or that have expired.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaCertificateKeyQualityCheck`

Checks the quality of the CA certificate key. The quality checks if the key is in a
valid format, not expired, and if the key meets a minimum required size. This check applies
to CA certificates that are `ACTIVE` or `PENDING_TRANSFER`.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConflictingClientIdsCheck`

Checks if multiple devices connect using the same client ID.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceCertificateAgeCheck`

Checks when a device certificate has been active for a number of days greater than or equal to the number you specify.

_Required_: No

_Type_: [DeviceCertAgeAuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-devicecertageauditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceCertificateExpiringCheck`

Checks if a device certificate is expiring. By default, this check applies to device certificates expiring within 30 days or that have expired. You can modify this threshold by configuring the DeviceCertExpirationAuditCheckConfiguration.

_Required_: No

_Type_: [DeviceCertExpirationAuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-devicecertexpirationauditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceCertificateKeyQualityCheck`

Checks the quality of the device certificate key. The quality checks if the key is in a
valid format, not expired, signed by a registered certificate authority, and if the key
meets a minimum required size.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceCertificateSharedCheck`

Checks if multiple concurrent connections use the same X.509 certificate to authenticate
with AWS IoT.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntermediateCaRevokedForActiveDeviceCertificatesCheck`

Checks if device certificates are still active despite being revoked by an intermediate
CA.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotPolicyOverlyPermissiveCheck`

Checks the permissiveness of a policy attached to an authenticated Amazon Cognito
identity pool role.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IoTPolicyPotentialMisConfigurationCheck`

Checks if an AWS IoT policy is potentially misconfigured. Misconfigured
policies, including overly permissive policies, can cause security incidents like allowing
devices access to unintended resources. This check is a warning for you to make sure that
only intended actions are allowed before updating the policy.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotRoleAliasAllowsAccessToUnusedServicesCheck`

Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotRoleAliasOverlyPermissiveCheck`

Checks if the temporary credentials provided by AWS IoT role aliases are
overly permissive.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingDisabledCheck`

Checks if AWS IoT logs are disabled.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevokedCaCertificateStillActiveCheck`

Checks if a revoked CA certificate is still active.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevokedDeviceCertificateStillActiveCheck`

Checks if a revoked device certificate is still active.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnauthenticatedCognitoRoleOverlyPermissiveCheck`

Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too
permissive.

_Required_: No

_Type_: [AuditCheckConfiguration](aws-properties-iot-accountauditconfiguration-auditcheckconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuditCheckConfiguration

AuditNotificationTarget

All content copied from https://docs.aws.amazon.com/.
