This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment InitialVlans

The initial VLAN subnets for the environment.
Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24.
Amazon EVS VLAN subnet CIDR blocks must not overlap with other subnets in the VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EdgeVTep" : InitialVlanInfo,
  "ExpansionVlan1" : InitialVlanInfo,
  "ExpansionVlan2" : InitialVlanInfo,
  "Hcx" : InitialVlanInfo,
  "HcxNetworkAclId" : String,
  "IsHcxPublic" : Boolean,
  "NsxUpLink" : InitialVlanInfo,
  "VmkManagement" : InitialVlanInfo,
  "VmManagement" : InitialVlanInfo,
  "VMotion" : InitialVlanInfo,
  "VSan" : InitialVlanInfo,
  "VTep" : InitialVlanInfo
}

```

### YAML

```yaml

  EdgeVTep:
    InitialVlanInfo
  ExpansionVlan1:
    InitialVlanInfo
  ExpansionVlan2:
    InitialVlanInfo
  Hcx:
    InitialVlanInfo
  HcxNetworkAclId: String
  IsHcxPublic: Boolean
  NsxUpLink:
    InitialVlanInfo
  VmkManagement:
    InitialVlanInfo
  VmManagement:
    InitialVlanInfo
  VMotion:
    InitialVlanInfo
  VSan:
    InitialVlanInfo
  VTep:
    InitialVlanInfo

```

## Properties

`EdgeVTep`

The edge VTEP VLAN subnet.
This VLAN subnet manages traffic flowing between the internal network and external networks, including internet access and other site connections.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`ExpansionVlan1`

An additional VLAN subnet that can be used to extend VCF capabilities once configured.
For example, you can configure an expansion VLAN subnet to use NSX Federation for centralized management and synchronization
of multiple NSX deployments across different locations.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`ExpansionVlan2`

An additional VLAN subnet that can be used to extend VCF capabilities once configured.
For example, you can configure an expansion VLAN subnet to use NSX Federation for centralized management and synchronization
of multiple NSX deployments across different locations.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`Hcx`

The HCX VLAN subnet.
This VLAN subnet allows the HCX Interconnnect (IX) and HCX Network Extension (NE) to reach their peers and enable HCX Service Mesh creation.

If you plan to use a public HCX VLAN subnet, the following requirements must be met:

- Must have a /28 netmask and be allocated from the IPAM public pool. Required for HCX internet access configuration.

- The HCX public VLAN CIDR block must be added to the VPC as a secondary CIDR block.

- Must have at least two Elastic IP addresses to be allocated from the public IPAM pool for HCX components.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`HcxNetworkAclId`

A unique ID for a network access control list that the HCX VLAN uses.
Required when `isHcxPublic` is set to `true`.

_Required_: No

_Type_: String

_Pattern_: `^acl-[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsHcxPublic`

Determines if the HCX VLAN that Amazon EVS provisions is public or private.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NsxUpLink`

The NSX uplink VLAN subnet.
This VLAN subnet allows connectivity to the NSX overlay network.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`VmkManagement`

The host VMkernel management VLAN subnet.
This VLAN subnet carries traffic for managing ESX hosts and communicating with VMware vCenter Server.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`VmManagement`

The VM management VLAN subnet.
This VLAN subnet carries traffic for vSphere virtual machines.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`VMotion`

The vMotion VLAN subnet.
This VLAN subnet carries traffic for vSphere vMotion.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`VSan`

The vSAN VLAN subnet.
This VLAN subnet carries the communication between ESX hosts to implement a vSAN shared storage pool.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

`VTep`

The VTEP VLAN subnet.
This VLAN subnet handles internal network traffic between virtual machines within a VCF instance.

_Required_: Yes

_Type_: [InitialVlanInfo](aws-properties-evs-environment-initialvlaninfo.md)

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InitialVlanInfo

LicenseInfo

All content copied from https://docs.aws.amazon.com/.
