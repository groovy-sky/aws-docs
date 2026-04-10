This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterNetworkInterface

The `AWS::MediaConnect::RouterNetworkInterface` resource defines how the router communicates with the outside world.
Each router I/O needs a network interface, which determines how the router I/O connects to other resources and what security measures
protect the connection.

You can work with two types of router network interface:

- **Public network interfaces** \- allow communication over the public internet.

- **VPC network interfaces** \- allow communication within your Amazon Virtual Private Cloud (VPC).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::RouterNetworkInterface",
  "Properties" : {
      "Configuration" : RouterNetworkInterfaceConfiguration,
      "Name" : String,
      "RegionName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::RouterNetworkInterface
Properties:
  Configuration:
    RouterNetworkInterfaceConfiguration
  Name: String
  RegionName: String
  Tags:
    - Tag

```

## Properties

`Configuration`

The configuration settings for a router network interface.

_Required_: Yes

_Type_: [RouterNetworkInterfaceConfiguration](aws-properties-mediaconnect-routernetworkinterface-routernetworkinterfaceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the router network interface.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionName`

The AWS Region where the router network interface is located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs that can be used to tag and organize this router network interface.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediaconnect-routernetworkinterface-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the router network interface ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-west-2:111122223333:routerNetworkInterface:00e8efe67aa3"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the router network interface.

`AssociatedInputCount`

The number of router inputs associated with the network interface.

`AssociatedOutputCount`

The number of router outputs associated with the network interface.

`CreatedAt`

The timestamp when the router network interface was created.

`Id`

The unique identifier of the router network interface.

`NetworkInterfaceType`

The type of the router network interface.

`State`

The current state of the router network interface.

`UpdatedAt`

The timestamp when the router network interface was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

PublicRouterNetworkInterfaceConfiguration

All content copied from https://docs.aws.amazon.com/.
