This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterNetworkInterface PublicRouterNetworkInterfaceRule

A rule that allows a specific CIDR block to access the public router network interface.

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

The CIDR block that is allowed to access the public router network interface.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicRouterNetworkInterfaceConfiguration

RouterNetworkInterfaceConfiguration

All content copied from https://docs.aws.amazon.com/.
