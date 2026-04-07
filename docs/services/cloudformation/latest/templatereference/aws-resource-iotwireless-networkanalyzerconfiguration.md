This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::NetworkAnalyzerConfiguration

Network analyzer configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::NetworkAnalyzerConfiguration",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TraceContent" : TraceContent,
      "WirelessDevices" : [ String, ... ],
      "WirelessGateways" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::NetworkAnalyzerConfiguration
Properties:
  Description: String
  Name: String
  Tags:
    - Tag
  TraceContent:
    TraceContent
  WirelessDevices:
    - String
  WirelessGateways:
    - String

```

## Properties

`Description`

The description of the resource.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the network analyzer configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]+$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to attach to the specified resource. Tags are metadata that you can use to
manage a resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-networkanalyzerconfiguration-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TraceContent`

Trace content for your wireless gateway and wireless device resources.

_Required_: No

_Type_: [TraceContent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotwireless-networkanalyzerconfiguration-tracecontent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WirelessDevices`

Wireless device resources to add to the network analyzer configuration. Provide the
`WirelessDeviceId` of the resource to add in the input array.

_Required_: No

_Type_: Array of String

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WirelessGateways`

Wireless gateway resources to add to the network analyzer configuration. Provide the
`WirelessGatewayId` of the resource to add in the input array.

_Required_: No

_Type_: Array of String

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the network analyzer configuration.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
