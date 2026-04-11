This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration EncryptionConfiguration

Configurations related to encryption for the security configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AtRestEncryptionConfiguration" : AtRestEncryptionConfiguration,
  "InTransitEncryptionConfiguration" : InTransitEncryptionConfiguration
}

```

### YAML

```yaml

  AtRestEncryptionConfiguration:
    AtRestEncryptionConfiguration
  InTransitEncryptionConfiguration:
    InTransitEncryptionConfiguration

```

## Properties

`AtRestEncryptionConfiguration`

Property description not available.

_Required_: No

_Type_: [AtRestEncryptionConfiguration](aws-properties-emrcontainers-securityconfiguration-atrestencryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InTransitEncryptionConfiguration`

In-transit encryption-related input for the security configuration.

_Required_: No

_Type_: [InTransitEncryptionConfiguration](aws-properties-emrcontainers-securityconfiguration-intransitencryptionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksInfo

IAMConfiguration

All content copied from https://docs.aws.amazon.com/.
