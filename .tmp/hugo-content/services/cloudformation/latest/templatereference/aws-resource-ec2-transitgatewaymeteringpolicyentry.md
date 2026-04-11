This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayMeteringPolicyEntry

Creates an entry in a transit gateway metering policy to define traffic measurement rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TransitGatewayMeteringPolicyEntry",
  "Properties" : {
      "DestinationCidrBlock" : String,
      "DestinationPortRange" : String,
      "DestinationTransitGatewayAttachmentId" : String,
      "DestinationTransitGatewayAttachmentType" : String,
      "MeteredAccount" : String,
      "PolicyRuleNumber" : Integer,
      "Protocol" : String,
      "SourceCidrBlock" : String,
      "SourcePortRange" : String,
      "SourceTransitGatewayAttachmentId" : String,
      "SourceTransitGatewayAttachmentType" : String,
      "TransitGatewayMeteringPolicyId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TransitGatewayMeteringPolicyEntry
Properties:
  DestinationCidrBlock: String
  DestinationPortRange: String
  DestinationTransitGatewayAttachmentId: String
  DestinationTransitGatewayAttachmentType: String
  MeteredAccount: String
  PolicyRuleNumber: Integer
  Protocol: String
  SourceCidrBlock: String
  SourcePortRange: String
  SourceTransitGatewayAttachmentId: String
  SourceTransitGatewayAttachmentType: String
  TransitGatewayMeteringPolicyId: String

```

## Properties

`DestinationCidrBlock`

Describes an IPv4 CIDR block.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPortRange`

Describes a range of ports.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationTransitGatewayAttachmentId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationTransitGatewayAttachmentType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `vpc | vpn | direct-connect-gateway | peering | network-function | vpn-concentrator | client-vpn`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeteredAccount`

The AWS account ID to which the metered traffic is attributed.

_Required_: Yes

_Type_: String

_Allowed values_: `source-attachment-owner | destination-attachment-owner | transit-gateway-owner`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyRuleNumber`

The rule number of the metering policy entry.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocol`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceCidrBlock`

Describes an IPv4 CIDR block.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePortRange`

Describes a range of ports.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceTransitGatewayAttachmentId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceTransitGatewayAttachmentType`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `vpc | vpn | direct-connect-gateway | peering | network-function | vpn-concentrator | client-vpn`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitGatewayMeteringPolicyId`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`State`

The state of the metering policy entry.

`UpdateEffectiveAt`

The date and time when the metering policy entry update becomes effective.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::TransitGatewayMulticastDomain

All content copied from https://docs.aws.amazon.com/.
