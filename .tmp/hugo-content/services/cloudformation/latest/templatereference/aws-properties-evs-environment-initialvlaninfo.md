This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment InitialVlanInfo

An object that represents an initial VLAN subnet for the Amazon EVS environment.
Amazon EVS creates initial VLAN subnets when you first create the environment.
Amazon EVS creates the following 10 VLAN subnets: host management VLAN, vMotion VLAN, vSAN VLAN, VTEP VLAN, Edge VTEP VLAN,
Management VM VLAN, HCX uplink VLAN, NSX uplink VLAN, expansion VLAN 1, expansion VLAN 2.

###### Note

For each Amazon EVS VLAN subnet, you must specify a non-overlapping CIDR block.
Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24.

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

The CIDR block that you provide to create an Amazon EVS VLAN subnet.
Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24.
Amazon EVS VLAN subnet CIDR blocks must not overlap with other subnets in the VPC.

_Required_: Yes

_Type_: String

_Pattern_: `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/(3[0-2]|[1-2][0-9]|[0-9])$`

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HostInfoForCreate

InitialVlans

All content copied from https://docs.aws.amazon.com/.
