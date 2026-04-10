This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterNetworkInterface PublicRouterNetworkInterfaceConfiguration

The configuration settings for a public router network interface, including the list of allowed CIDR blocks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowRules" : [ PublicRouterNetworkInterfaceRule, ... ]
}

```

### YAML

```yaml

  AllowRules:
    - PublicRouterNetworkInterfaceRule

```

## Properties

`AllowRules`

The list of allowed CIDR blocks for the public router network interface.

_Required_: Yes

_Type_: Array of [PublicRouterNetworkInterfaceRule](aws-properties-mediaconnect-routernetworkinterface-publicrouternetworkinterfacerule.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaConnect::RouterNetworkInterface

PublicRouterNetworkInterfaceRule

All content copied from https://docs.aws.amazon.com/.
