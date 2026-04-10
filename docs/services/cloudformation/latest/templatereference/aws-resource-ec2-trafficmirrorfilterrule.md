This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TrafficMirrorFilterRule

Creates a Traffic Mirror filter rule.

A Traffic Mirror rule defines the Traffic Mirror source traffic to mirror.

You need the Traffic Mirror filter ID when you create the rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TrafficMirrorFilterRule",
  "Properties" : {
      "Description" : String,
      "DestinationCidrBlock" : String,
      "DestinationPortRange" : TrafficMirrorPortRange,
      "Protocol" : Integer,
      "RuleAction" : String,
      "RuleNumber" : Integer,
      "SourceCidrBlock" : String,
      "SourcePortRange" : TrafficMirrorPortRange,
      "Tags" : [ Tag, ... ],
      "TrafficDirection" : String,
      "TrafficMirrorFilterId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TrafficMirrorFilterRule
Properties:
  Description: String
  DestinationCidrBlock: String
  DestinationPortRange:
    TrafficMirrorPortRange
  Protocol: Integer
  RuleAction: String
  RuleNumber: Integer
  SourceCidrBlock: String
  SourcePortRange:
    TrafficMirrorPortRange
  Tags:
    - Tag
  TrafficDirection: String
  TrafficMirrorFilterId: String

```

## Properties

`Description`

The description of the Traffic Mirror rule.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationCidrBlock`

The destination CIDR block to assign to the Traffic Mirror rule.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPortRange`

The destination port range.

_Required_: No

_Type_: [TrafficMirrorPortRange](aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol, for example UDP, to assign to the Traffic Mirror rule.

For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleAction`

The action to take on the filtered traffic.

_Required_: Yes

_Type_: String

_Allowed values_: `accept | reject`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleNumber`

The number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given
direction. The rules are processed in ascending order by rule number.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceCidrBlock`

The source CIDR block to assign to the Traffic Mirror rule.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePortRange`

The source port range.

_Required_: No

_Type_: [TrafficMirrorPortRange](aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags on Traffic Mirroring filter rules.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-trafficmirrorfilterrule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficDirection`

The type of traffic.

_Required_: Yes

_Type_: String

_Allowed values_: `ingress | egress`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficMirrorFilterId`

The ID of the filter that this rule is associated with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Traffic Mirror filter rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`TrafficMirrorFilterRuleId`

The ID of the Traffic Mirror rule.

## Examples

### Create a traffic mirror filter rule for inbound UDP traffic

This is a filter rule for UDP traffic.

#### JSON

```json

{
  "SampleTrafficMirrorFilterRule": {
    "Type": "AWS::EC2::TrafficMirrorFilterRule",
    "Properties": {
      "Description": "Example traffic mirror filter rule",
      "TrafficMirrorFilterId": "tmf-04812ff784EXAMPLE",
      "TrafficDirection": "ingress",
      "RuleNumber": 10,
      "DestinationCidrBlock": "10.0.0.0/16",
      "SourceCidrBlock": "10.0.0.0/16",
      "RuleAction": "accept",
      "Protocol": 17,
      "SourcePortRange": {
        "FromPort": 10,
        "ToPort": 50
      },
      "DestinationPortRange": {
        "FromPort": 50,
        "ToPort": 100
      }
    }
  }
}
```

#### YAML

```yaml

SampleTrafficMirrorFilterRule:
  Type: "AWS::EC2::TrafficMirrorFilterRule"
  Properties:
    Description: "Example traffic mirror filter rule"
    TrafficMirrorFilterId: "tmf-04812ff784EXAMPLE"
    TrafficDirection: "ingress"
    RuleNumber: 10
    DestinationCidrBlock: "10.0.0.0/16"
    SourceCidrBlock: "10.0.0.0/16"
    RuleAction: "accept"
    Protocol: 17
    SourcePortRange:
      FromPort: 10
      ToPort: 50
    DestinationPortRange:
      FromPort: 50
      ToPort: 100
```

## See also

- [Traffic mirror\
filters and filter rules](../../../vpc/latest/mirroring/traffic-mirroring-filters.md) in _Traffic_
_Mirroring_

- [CreateTrafficMirrorFilterRule](../../../../reference/awsec2/latest/apireference/api-createtrafficmirrorfilterrule.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
