This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment ConnectivityInfo

The connectivity configuration for the environment.
Amazon EVS requires that you specify two route server peer IDs.
During environment creation, the route server endpoints peer with the NSX uplink VLAN for connectivity to the NSX overlay network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrivateRouteServerPeerings" : [ String, ... ]
}

```

### YAML

```yaml

  PrivateRouteServerPeerings:
    - String

```

## Properties

`PrivateRouteServerPeerings`

The unique IDs for private route server peers.

_Required_: Yes

_Type_: Array of String

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: Updates are not supported.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Check

HostInfoForCreate
