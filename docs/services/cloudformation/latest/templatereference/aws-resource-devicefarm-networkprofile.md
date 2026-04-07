This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::NetworkProfile

Creates a network profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DeviceFarm::NetworkProfile",
  "Properties" : {
      "Description" : String,
      "DownlinkBandwidthBits" : Integer,
      "DownlinkDelayMs" : Integer,
      "DownlinkJitterMs" : Integer,
      "DownlinkLossPercent" : Integer,
      "Name" : String,
      "ProjectArn" : String,
      "Tags" : [ Tag, ... ],
      "UplinkBandwidthBits" : Integer,
      "UplinkDelayMs" : Integer,
      "UplinkJitterMs" : Integer,
      "UplinkLossPercent" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::DeviceFarm::NetworkProfile
Properties:
  Description: String
  DownlinkBandwidthBits: Integer
  DownlinkDelayMs: Integer
  DownlinkJitterMs: Integer
  DownlinkLossPercent: Integer
  Name: String
  ProjectArn: String
  Tags:
    - Tag
  UplinkBandwidthBits: Integer
  UplinkDelayMs: Integer
  UplinkJitterMs: Integer
  UplinkLossPercent: Integer

```

## Properties

`Description`

The description of the network profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DownlinkBandwidthBits`

The data throughput rate in bits per second, as an integer from 0 to
104857600.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `104857600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DownlinkDelayMs`

Delay time for all packets to destination in milliseconds as an integer from 0 to
2000.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DownlinkJitterMs`

Time variation in the delay of received packets in milliseconds as an integer from
0 to 2000.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DownlinkLossPercent`

Proportion of received packets that fail to arrive from 0 to 100 percent.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the network profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectArn`

The Amazon Resource Name (ARN) of the specified project.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.+`

_Minimum_: `32`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the
_AWS CloudFormation guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-devicefarm-networkprofile-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UplinkBandwidthBits`

The data throughput rate in bits per second, as an integer from 0 to
104857600.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `104857600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UplinkDelayMs`

Delay time for all packets to destination in milliseconds as an integer from 0 to
2000.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UplinkJitterMs`

Time variation in the delay of received packets in milliseconds as an integer from
0 to 2000.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UplinkLossPercent`

Proportion of transmitted packets that fail to arrive from 0 to 100
percent.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

Not supported for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the network profile. See [Amazon resource names](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the
_General Reference guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
