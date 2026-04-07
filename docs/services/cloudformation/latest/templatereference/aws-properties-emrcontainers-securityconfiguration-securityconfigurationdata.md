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

_Type_: [AuthenticationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrcontainers-securityconfiguration-authenticationconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizationConfiguration`

Authorization-related configuration input for the security configuration.

_Required_: No

_Type_: [AuthorizationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrcontainers-securityconfiguration-authorizationconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionConfiguration`

Property description not available.

_Required_: No

_Type_: [EncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrcontainers-securityconfiguration-encryptionconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SecureNamespaceInfo

Tag
