This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::VpcEndpointAssociation SubnetMapping

The ID for a subnet that's used in an association with a firewall. This is used in
`CreateFirewall`, `AssociateSubnets`, and `CreateVpcEndpointAssociation`. AWS Network Firewall
creates an instance of the associated firewall in each subnet that you specify, to filter
traffic in the subnet's Availability Zone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IPAddressType" : String,
  "SubnetId" : String
}

```

### YAML

```yaml

  IPAddressType: String
  SubnetId: String

```

## Properties

`IPAddressType`

The subnet's IP address type. You can't change the IP address type after you create the subnet.

_Required_: No

_Type_: String

_Allowed values_: `DUALSTACK | IPV4 | IPV6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The unique identifier for the subnet.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::NetworkFirewall::VpcEndpointAssociation

Tag
