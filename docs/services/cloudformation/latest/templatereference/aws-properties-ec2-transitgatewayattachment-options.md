This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TransitGatewayAttachment Options

Describes the VPC attachment options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplianceModeSupport" : String,
  "DnsSupport" : String,
  "Ipv6Support" : String,
  "SecurityGroupReferencingSupport" : String
}

```

### YAML

```yaml

  ApplianceModeSupport: String
  DnsSupport: String
  Ipv6Support: String
  SecurityGroupReferencingSupport: String

```

## Properties

`ApplianceModeSupport`

Enable or disable appliance mode support. The default is `disable`.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsSupport`

Enable or disable DNS support. The default is `disable`.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Support`

Enable or disable IPv6 support. The default is `disable`.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupReferencingSupport`

Enables you to reference a security group across VPCs attached to a transit gateway (TGW). Use this option to simplify security group management and control of instance-to-instance traffic across VPCs that are connected by transit gateway. You can also use this option to migrate from VPC peering (which was the only option that supported security group referencing) to transit gateways (which now also support security group referencing). This option is disabled by default and there are no additional costs to use this feature.

For important information about this feature, see [Create a transit gateway](../../../vpc/latest/tgw/tgw-transit-gateways.md#create-tgw) in the _AWS Transit Gateway Guide_.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::TransitGatewayAttachment

Tag

All content copied from https://docs.aws.amazon.com/.
