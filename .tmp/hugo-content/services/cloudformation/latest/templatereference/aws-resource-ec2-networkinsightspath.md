This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsPath

Specifies a path to analyze for reachability.

VPC Reachability Analyzer enables you to analyze and debug network reachability between
two resources in your virtual private cloud (VPC). For more information, see the [Reachability Analyzer User Guide](../../../vpc/latest/reachability/what-is-reachability-analyzer.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInsightsPath",
  "Properties" : {
      "Destination" : String,
      "DestinationIp" : String,
      "DestinationPort" : Integer,
      "FilterAtDestination" : PathFilter,
      "FilterAtSource" : PathFilter,
      "Protocol" : String,
      "Source" : String,
      "SourceIp" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInsightsPath
Properties:
  Destination: String
  DestinationIp: String
  DestinationPort: Integer
  FilterAtDestination:
    PathFilter
  FilterAtSource:
    PathFilter
  Protocol: String
  Source: String
  SourceIp: String
  Tags:
    - Tag

```

## Properties

`Destination`

The ID or ARN of the destination. If the resource is in another account, you must specify an ARN.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationIp`

The IP address of the destination.

_Required_: No

_Type_: String

_Pattern_: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPort`

The destination port.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterAtDestination`

Scopes the analysis to network paths that match specific filters at the destination. If you specify
this parameter, you can't specify the parameter for the destination IP address.

_Required_: No

_Type_: [PathFilter](aws-properties-ec2-networkinsightspath-pathfilter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterAtSource`

Scopes the analysis to network paths that match specific filters at the source. If you specify
this parameter, you can't specify the parameters for the source IP address or the destination port.

_Required_: No

_Type_: [PathFilter](aws-properties-ec2-networkinsightspath-pathfilter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

The protocol.

_Required_: Yes

_Type_: String

_Allowed values_: `tcp | udp`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Source`

The ID or ARN of the source. If the resource is in another account, you must specify an ARN.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceIp`

The IP address of the source.

_Required_: No

_Type_: String

_Pattern_: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to add to the path.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-networkinsightspath-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the path.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedDate`

The time stamp when the path was created.

`DestinationArn`

The Amazon Resource Name (ARN) of the destination.

`NetworkInsightsPathArn`

The Amazon Resource Name (ARN) of the path.

`NetworkInsightsPathId`

The ID of the path.

`SourceArn`

The Amazon Resource Name (ARN) of the source.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransitGatewayRouteTableRoute

FilterPortRange

All content copied from https://docs.aws.amazon.com/.
