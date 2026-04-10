This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment HostInfoForCreate

An object that represents a host.

###### Note

You cannot use `dedicatedHostId` and `placementGroupId` together in the same `HostInfoForCreate` object.
This results in a `ValidationException` response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DedicatedHostId" : String,
  "HostName" : String,
  "InstanceType" : String,
  "KeyName" : String,
  "PlacementGroupId" : String
}

```

### YAML

```yaml

  DedicatedHostId: String
  HostName: String
  InstanceType: String
  KeyName: String
  PlacementGroupId: String

```

## Properties

`DedicatedHostId`

The unique ID of the Amazon EC2 Dedicated Host.

_Required_: No

_Type_: String

_Pattern_: `^h-[a-f0-9]{8}([a-f0-9]{9})?$`

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: Updates are not supported.

`HostName`

The DNS hostname of the host.
DNS hostnames for hosts must be unique across Amazon EVS environments and within VCF.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`InstanceType`

The EC2 instance type that represents the host.

###### Note

Currently, Amazon EVS supports only the `i4i.metal` instance type.

_Required_: Yes

_Type_: String

_Allowed values_: `i4i.metal`

_Update requires_: Updates are not supported.

`KeyName`

The name of the SSH key that is used to access the host.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: Updates are not supported.

`PlacementGroupId`

The unique ID of the placement group where the host is placed.

_Required_: No

_Type_: String

_Pattern_: `^pg-[a-f0-9]{8}([a-f0-9]{9})?$`

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectivityInfo

InitialVlanInfo

All content copied from https://docs.aws.amazon.com/.
