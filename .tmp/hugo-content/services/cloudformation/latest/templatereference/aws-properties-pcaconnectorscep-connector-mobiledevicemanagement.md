This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorSCEP::Connector MobileDeviceManagement

If you don't supply a value, by default Connector for SCEP creates a connector for general-purpose use. A general-purpose connector is designed to work with clients or endpoints that support the SCEP protocol, except Connector for SCEP for Microsoft Intune. For information about considerations and limitations with using Connector for SCEP, see [Considerations and Limitations](../../../privateca/latest/userguide/scep-connectorc4scep-considerations-limitations.md).

If you provide an `IntuneConfiguration`, Connector for SCEP creates a connector for use with Microsoft Intune, and you manage the challenge passwords using Microsoft Intune. For more information, see [Using Connector for SCEP for Microsoft Intune](../../../privateca/latest/userguide/scep-connectorconnector-for-scep-intune.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Intune" : IntuneConfiguration
}

```

### YAML

```yaml

  Intune:
    IntuneConfiguration

```

## Properties

`Intune`

Configuration settings for use with Microsoft Intune. For information about using Connector for SCEP for Microsoft Intune, see [Using Connector for SCEP for Microsoft Intune](../../../privateca/latest/userguide/scep-connectorconnector-for-scep-intune.md).

_Required_: Yes

_Type_: [IntuneConfiguration](aws-properties-pcaconnectorscep-connector-intuneconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntuneConfiguration

OpenIdConfiguration

All content copied from https://docs.aws.amazon.com/.
