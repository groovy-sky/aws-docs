This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPool ProvisionedCidr

The CIDR provisioned to the IPAM pool. A CIDR is a representation of an IP address and its associated network mask (or netmask)
and refers to a range of IP addresses. An IPv4 CIDR example is `10.24.34.0/23`. An IPv6 CIDR example is `2001:DB8::/32`.

###### Note

This resource type does not allow you to provision a CIDR using the netmask length. To provision a CIDR using netmask length, use [AWS::EC2::IPAMPoolCidr](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampoolcidr.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String
}

```

### YAML

```yaml

  Cidr: String

```

## Properties

`Cidr`

The CIDR provisioned to the IPAM pool. A CIDR is a representation of an IP address and its associated network mask (or netmask)
and refers to a range of IP addresses. An IPv4 CIDR example is `10.24.34.0/23`. An IPv6 CIDR example is `2001:DB8::/32`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::IPAMPool

SourceResource
