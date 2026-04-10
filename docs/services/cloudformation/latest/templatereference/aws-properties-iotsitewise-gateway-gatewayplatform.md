This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Gateway GatewayPlatform

The gateway's platform configuration. You can only specify one platform type in a
gateway.

(Legacy only) For Greengrass V1 gateways, specify the `greengrass` parameter
with a valid Greengrass group ARN.

For Greengrass V2 gateways, specify the `greengrassV2` parameter with a valid
core device thing name. If creating a V3 gateway ( `gatewayVersion=3`), you must
also specify the `coreDeviceOperatingSystem`.

For Siemens Industrial Edge gateways, specify the `siemensIE` parameter with a
valid IoT Core thing name.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GreengrassV2" : GreengrassV2,
  "SiemensIE" : SiemensIE
}

```

### YAML

```yaml

  GreengrassV2:
    GreengrassV2
  SiemensIE:
    SiemensIE

```

## Properties

`GreengrassV2`

A gateway that runs on AWS IoT Greengrass V2.

_Required_: No

_Type_: [GreengrassV2](aws-properties-iotsitewise-gateway-greengrassv2.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SiemensIE`

An AWS IoT SiteWise Edge gateway that runs on a Siemens Industrial Edge
Device.

_Required_: No

_Type_: [SiemensIE](aws-properties-iotsitewise-gateway-siemensie.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GatewayCapabilitySummary

GreengrassV2

All content copied from https://docs.aws.amazon.com/.
