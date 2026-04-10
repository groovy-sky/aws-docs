This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterNetworkInterface RouterNetworkInterfaceConfiguration

The configuration settings for a router network interface.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Public" : PublicRouterNetworkInterfaceConfiguration,
  "Vpc" : VpcRouterNetworkInterfaceConfiguration
}

```

### YAML

```yaml

  Public:
    PublicRouterNetworkInterfaceConfiguration
  Vpc:
    VpcRouterNetworkInterfaceConfiguration

```

## Properties

`Public`

The configuration settings for a public router network interface, including the list of allowed CIDR blocks.

_Required_: No

_Type_: [PublicRouterNetworkInterfaceConfiguration](aws-properties-mediaconnect-routernetworkinterface-publicrouternetworkinterfaceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vpc`

The configuration settings for a router network interface within a VPC, including the security group IDs and subnet ID.

_Required_: No

_Type_: [VpcRouterNetworkInterfaceConfiguration](aws-properties-mediaconnect-routernetworkinterface-vpcrouternetworkinterfaceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicRouterNetworkInterfaceRule

Tag

All content copied from https://docs.aws.amazon.com/.
