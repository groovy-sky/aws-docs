This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTManagedIntegrations::ManagedThing CapabilityReportEndpoint

The endpoints used in the capability report.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Capabilities" : [ CapabilityReportCapability, ... ],
  "DeviceTypes" : [ String, ... ],
  "Id" : String
}

```

### YAML

```yaml

  Capabilities:
    - CapabilityReportCapability
  DeviceTypes:
    - String
  Id: String

```

## Properties

`Capabilities`

The capabilities used in the capability report.

_Required_: Yes

_Type_: Array of [CapabilityReportCapability](aws-properties-iotmanagedintegrations-managedthing-capabilityreportcapability.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceTypes`

The type of device.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The id of the endpoint used in the capability report.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapabilityReportCapability

AWS::IoTManagedIntegrations::ProvisioningProfile

All content copied from https://docs.aws.amazon.com/.
