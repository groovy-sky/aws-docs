This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::Connection

Creates a connection between a customer network and a specific Direct Connect location.

A connection links your internal network to an Direct Connect location over a standard Ethernet fiber-optic
cable. One end of the cable is connected to your router, the other to an Direct Connect router.

To find the locations for your Region, use [DescribeLocations](../../../../reference/directconnect/latest/apireference/api-describelocations.md).

You can automatically add the new connection to a link aggregation group (LAG) by
specifying a LAG ID in the request. This ensures that the new connection is allocated on the
same Direct Connect endpoint that hosts the specified LAG. If there are no available ports on the endpoint,
the request fails and no connection is created.

###### Important

Upon creation, a connection will enter the `requested` state. The connection cannot be used with other resources until it is approved and advances to `available` or `down`.

In order to create a virtual interface using a connection defined in the same CloudFormation template, it is recommended to create a LAG and create the virtual interface using the LAG. Member connections can be added to the LAG in the CloudFormation template.

For more information, see [Dedicated Direct Connect connections](../../../directconnect/latest/userguide/dedicated-connection.md) in the _Direct Connect User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectConnect::Connection",
  "Properties" : {
      "Bandwidth" : String,
      "ConnectionName" : String,
      "LagId" : String,
      "Location" : String,
      "ProviderName" : String,
      "RequestMACSec" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DirectConnect::Connection
Properties:
  Bandwidth: String
  ConnectionName: String
  LagId: String
  Location: String
  ProviderName: String
  RequestMACSec: Boolean
  Tags:
    - Tag

```

## Properties

`Bandwidth`

The bandwidth of the connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*(M|G)bps$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectionName`

The name of the connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w \-_,\/]{1,200}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LagId`

The ID or ARN of the LAG to associate the connection with.

Connectivity will be interrupted when associating or disassociating a connection with a LAG.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[a-z-]*:directconnect:[a-z0-9-]+:[0-9]{12}:dxlag/)?dxlag-[a-zA-Z0-9]{8,21}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Location`

The location of the connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProviderName`

The name of the service provider associated with the connection.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RequestMACSec`

Indicates whether you want the connection to support MAC Security (MACsec).

MAC Security (MACsec) is unavailable on hosted connections. For information about MAC Security (MACsec) prerequisites, see [MAC Security in Direct Connect](../../../directconnect/latest/userguide/macsec.md) in the _Direct Connect User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the connection.

_Required_: No

_Type_: Array of [Tag](aws-properties-directconnect-connection-tag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the connection.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectionArn`

The Amazon Resource Name (ARN) of the connection.

`ConnectionId`

The ID of the connection.

`ConnectionState`

The state of the connection. The following are the possible values:

- `ordering`: The initial state of a hosted connection provisioned on an interconnect. The connection stays in the ordering state until the owner of the hosted connection confirms or declines the connection order.

- `requested`: The initial state of a standard connection. The connection stays in the requested state until the Letter of Authorization (LOA) is sent to the customer.

- `pending`: The connection has been approved and is being initialized.

- `available`: The network link is up and the connection is ready for use.

- `down`: The network link is down.

- `deleting`: The connection is being deleted.

- `deleted`: The connection has been deleted.

- `rejected`: A hosted connection in the `ordering` state enters the `rejected` state if it is deleted by the customer.

- `unknown`: The state of the connection is not available.

## Examples

### Basic connection

This example shows a basic connection. The bandwidth of the connection is 10Gbps.

#### JSON

```json

{
  "Resources": {
    "myConnection": {
      "Type": "AWS::DirectConnect::Connection",
      "Properties": {
        "ConnectionName": "cfn-connection-example",
        "Bandwidth": "10Gbps",
        "Location": "EqSY3",
        "Tags": [
          {
            "Key": "example-key",
            "Value": "example-value"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  myConnection:
    Type: AWS::DirectConnect::Connection
    Properties:
      ConnectionName: cfn-connection-example
      Bandwidth: 10Gbps
      Location: EqSY3
      Tags:
      - Key: example-key
        Value: example-value
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Direct Connect

Tag

All content copied from https://docs.aws.amazon.com/.
