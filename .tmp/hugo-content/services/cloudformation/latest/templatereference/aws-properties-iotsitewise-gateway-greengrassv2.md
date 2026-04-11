This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Gateway GreengrassV2

Contains details for a gateway that runs on AWS IoT Greengrass V2. To create a gateway that runs on AWS IoT Greengrass V2,
you must deploy the IoT SiteWise Edge component to your gateway device. Your [Greengrass\
device role](../../../greengrass/v2/developerguide/device-service-role.md) must use the `AWSIoTSiteWiseEdgeAccess` policy. For more
information, see [Using AWS IoT SiteWise at the edge](../../../iot-sitewise/latest/userguide/sw-gateways.md) in the
_AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CoreDeviceOperatingSystem" : String,
  "CoreDeviceThingName" : String
}

```

### YAML

```yaml

  CoreDeviceOperatingSystem: String
  CoreDeviceThingName: String

```

## Properties

`CoreDeviceOperatingSystem`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `LINUX_AARCH64 | LINUX_AMD64 | WINDOWS_AMD64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CoreDeviceThingName`

The name of the AWS IoT thing for your AWS IoT Greengrass V2 core device.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayPlatform

SiemensIE

All content copied from https://docs.aws.amazon.com/.
