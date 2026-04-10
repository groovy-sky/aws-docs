This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorSCEP::Connector IntuneConfiguration

Contains configuration details for use with Microsoft Intune. For information about using Connector for SCEP for Microsoft Intune, see [Using Connector for SCEP for Microsoft Intune](../../../privateca/latest/userguide/scep-connectorconnector-for-scep-intune.md).

When you use Connector for SCEP for Microsoft Intune, certain functionalities are enabled by accessing Microsoft Intune through the Microsoft API. Your use of the Connector for SCEP and accompanying AWS services doesn't remove your need to have a valid license for your use of the Microsoft Intune service. You should also review the [Microsoft Intune® App Protection Policies](https://learn.microsoft.com/en-us/mem/intune/apps/app-protection-policy).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AzureApplicationId" : String,
  "Domain" : String
}

```

### YAML

```yaml

  AzureApplicationId: String
  Domain: String

```

## Properties

`AzureApplicationId`

The directory (tenant) ID from your Microsoft Entra ID app registration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]{2,15}-[a-zA-Z0-9]{2,15}-[a-zA-Z0-9]{2,15}-[a-zA-Z0-9]{2,15}-[a-zA-Z0-9]{2,15}$`

_Minimum_: `15`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Domain`

The primary domain from your Microsoft Entra ID app registration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PCAConnectorSCEP::Connector

MobileDeviceManagement

All content copied from https://docs.aws.amazon.com/.
