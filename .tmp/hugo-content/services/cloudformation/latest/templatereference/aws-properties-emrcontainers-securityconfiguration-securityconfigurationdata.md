This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration SecurityConfigurationData

Configurations related to the security configuration for the request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationConfiguration" : AuthenticationConfiguration,
  "AuthorizationConfiguration" : AuthorizationConfiguration,
  "EncryptionConfiguration" : EncryptionConfiguration
}

```

### YAML

```yaml

  AuthenticationConfiguration:
    AuthenticationConfiguration
  AuthorizationConfiguration:
    AuthorizationConfiguration
  EncryptionConfiguration:
    EncryptionConfiguration

```

## Properties

`AuthenticationConfiguration`

Property description not available.

_Required_: No

_Type_: [AuthenticationConfiguration](aws-properties-emrcontainers-securityconfiguration-authenticationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizationConfiguration`

Authorization-related configuration input for the security configuration.

_Required_: No

_Type_: [AuthorizationConfiguration](aws-properties-emrcontainers-securityconfiguration-authorizationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionConfiguration`

Property description not available.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-emrcontainers-securityconfiguration-encryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecureNamespaceInfo

Tag

All content copied from https://docs.aws.amazon.com/.
