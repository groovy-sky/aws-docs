This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Gateway

Creates a gateway, which is a virtual or edge device that delivers industrial data streams
from local servers to AWS IoT SiteWise. For more information, see [Ingesting data using a gateway](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html) in the
_AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::Gateway",
  "Properties" : {
      "GatewayCapabilitySummaries" : [ GatewayCapabilitySummary, ... ],
      "GatewayName" : String,
      "GatewayPlatform" : GatewayPlatform,
      "GatewayVersion" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::Gateway
Properties:
  GatewayCapabilitySummaries:
    - GatewayCapabilitySummary
  GatewayName: String
  GatewayPlatform:
    GatewayPlatform
  GatewayVersion: String
  Tags:
    - Tag

```

## Properties

`GatewayCapabilitySummaries`

A list of gateway capability summaries that each contain a namespace and status. Each
gateway capability defines data sources for the gateway. To retrieve a capability
configuration's definition, use [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html).

_Required_: No

_Type_: Array of [GatewayCapabilitySummary](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-gateway-gatewaycapabilitysummary.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GatewayName`

A unique name for the gateway.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GatewayPlatform`

The gateway's platform. You can only specify one platform in a gateway.

_Required_: Yes

_Type_: [GatewayPlatform](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-gateway-gatewayplatform.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GatewayVersion`

The version of the gateway. A value of `3` indicates an MQTT-enabled, V3
gateway, while `2` indicates a Classic streams, V2 gateway.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of key-value pairs that contain metadata for the gateway. For more information, see
[Tagging your AWS IoT SiteWise\
resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-gateway-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `GatewayId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GatewayId`

The ID for the gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

GatewayCapabilitySummary
