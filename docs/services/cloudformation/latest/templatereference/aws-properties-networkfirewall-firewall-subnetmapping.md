This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::Firewall SubnetMapping

The ID for a subnet that you want to associate with the firewall. AWS Network Firewall
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

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The unique identifier for the subnet.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AvailabilityZoneMapping

Tag

All content copied from https://docs.aws.amazon.com/.
