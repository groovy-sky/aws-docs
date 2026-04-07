This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectConnect::Lag

Creates a link aggregation group (LAG) in a specific Direct Connect location.
A LAG is a logical interface that uses the Link Aggregation Control Protocol
(LACP) to aggregate multiple interfaces, enabling you to treat them as a single
interface.

All connections in a LAG must use the same bandwidth (either 1Gbps, 10Gbps, 100Gbps,
or 400Gbps) and must terminate at the same Direct Connect endpoint.

You can have up to 10 dedicated connections per location. Regardless of this limit, if you
request more connections for the LAG than Direct Connect can allocate on a single endpoint, no LAG is
created.

If the AWS account used to create a LAG is a registered Direct Connect Partner, the LAG is
automatically enabled to host sub-connections. For a LAG owned by a partner, any associated virtual
interfaces cannot be directly configured.

###### Note

LAGs created using CloudFormation have no connections by default.

For more information, see [Direct Connect link aggregation groups (LAGs)](https://docs.aws.amazon.com/directconnect/latest/UserGuide/lags.html) in the _Direct Connect User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectConnect::Lag",
  "Properties" : {
      "ConnectionsBandwidth" : String,
      "LagName" : String,
      "Location" : String,
      "MinimumLinks" : Integer,
      "ProviderName" : String,
      "RequestMACSec" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DirectConnect::Lag
Properties:
  ConnectionsBandwidth: String
  LagName: String
  Location: String
  MinimumLinks: Integer
  ProviderName: String
  RequestMACSec: Boolean
  Tags:
    - Tag

```

## Properties

`ConnectionsBandwidth`

The individual bandwidth of the physical connections bundled by the LAG. The possible
values are 1Gbps, 10Gbps, 100Gbps, or 400 Gbps..

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*(M|G)bps$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LagName`

The name of the LAG.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w \-_,\/]{1,200}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

The location of the LAG.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumLinks`

The minimum number of physical dedicated connections that must be operational for the LAG itself to be operational.

###### Important

This property cannot be used when the LAG is being created for the first time using the template.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderName`

The name of the service provider associated with the LAG.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RequestMACSec`

Indicates whether the connection will support MAC Security (MACsec).

###### Note

All connections in the LAG must be capable of supporting MAC Security (MACsec). For information about MAC Security (MACsec) prerequisties, see [MACsec prerequisties](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) in the _Direct Connect User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags associated with the LAG.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directconnect-lag-tag.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the LAG.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LagArn`

The Amazon Resource Name (ARN) of the LAG.

`LagId`

The ID of the LAG.

`LagState`

The state of the LAG. The following are the possible values:

- `requested`: The initial state of a LAG. The LAG stays in the
requested state until the Letter of Authorization (LOA) is available.

- `pending`: The LAG has been approved and is being initialized.

- `available`: The network link is established and the LAG is ready for use.

- `down`: The network link is down.

- `deleting`: The LAG is being deleted.

- `deleted`: The LAG is deleted.

- `unknown`: The state of the LAG is not available.

## Examples

### Create LAG with one connection

This example shows a basic LAG which has one connection created on it. The bandwidth of the connections on the LAG is 10Gbps.

#### JSON

```json

{
  "Resources": {
    "myLag": {
      "Type": "AWS::DirectConnect::Lag",
      "Properties": {
        "LagName": "cfn-lag-example",
        "ConnectionsBandwidth": "10Gbps",
        "Location": "EqSY3",
        "Tags": [
          {
            "Key": "example-key",
            "Value": "example-value"
          }
        ]
      }
    },
    "myConnection": {
      "Type": "AWS::DirectConnect::Connection",
      "Properties": {
        "ConnectionName": "cfn-connection-example",
        "Bandwidth": "10Gbps",
        "Location": "EqSY3",
        "LagId": {
          "Ref": "myLag"
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  myLag:
    Type: AWS::DirectConnect::Lag
    Properties:
      LagName: cfn-lag-example
      ConnectionsBandwidth: 10Gbps
      Location: EqSY3
      Tags:
      - Key: example-key
        Value: example-value
  myConnection:
    Type: AWS::DirectConnect::Connection
    Properties:
      ConnectionName: cfn-connection-example
      Bandwidth: 10Gbps
      Location: EqSY3
      LagId: !Ref myLag
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DirectConnect::DirectConnectGatewayAssociation

Tag
