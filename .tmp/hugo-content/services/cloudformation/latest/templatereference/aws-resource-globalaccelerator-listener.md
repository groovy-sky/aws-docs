This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GlobalAccelerator::Listener

The `AWS::GlobalAccelerator::Listener` resource is a Global Accelerator resource type that contains information about
how you create a listener to process inbound connections from clients to an accelerator. Connections arrive to assigned static
IP addresses on a port, port range, or list of port ranges that you specify.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GlobalAccelerator::Listener",
  "Properties" : {
      "AcceleratorArn" : String,
      "ClientAffinity" : String,
      "PortRanges" : [ PortRange, ... ],
      "Protocol" : String
    }
}

```

### YAML

```yaml

Type: AWS::GlobalAccelerator::Listener
Properties:
  AcceleratorArn: String
  ClientAffinity: String
  PortRanges:
    - PortRange
  Protocol: String

```

## Properties

`AcceleratorArn`

The Amazon Resource Name (ARN) of your accelerator.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientAffinity`

Client affinity lets you direct all requests from a user to the same endpoint, if you have stateful applications,
regardless of the port and protocol of the client request. Client affinity gives you control over whether to always
route each client to the same specific endpoint.

AWS Global Accelerator uses a consistent-flow hashing algorithm to choose the optimal endpoint for a connection. If client
affinity is `NONE`, Global Accelerator uses the "five-tuple" (5-tuple) properties—source IP address, source port,
destination IP address, destination port, and protocol—to select the hash value, and then chooses the best
endpoint. However, with this setting, if someone uses different ports to connect to Global Accelerator, their connections might not
be always routed to the same endpoint because the hash value changes.

If you want a given client to always be routed to the same endpoint, set client affinity to `SOURCE_IP`
instead. When you use the `SOURCE_IP` setting, Global Accelerator uses the "two-tuple" (2-tuple) properties—
source (client) IP address and destination IP address—to select the hash value.

The default value is `NONE`.

_Required_: No

_Type_: String

_Allowed values_: `NONE | SOURCE_IP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortRanges`

The list of port ranges for the connections from clients to the accelerator.

_Required_: Yes

_Type_: Array of [PortRange](aws-properties-globalaccelerator-listener-portrange.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocol for the connections from clients to the accelerator.

_Required_: Yes

_Type_: String

_Allowed values_: `TCP | UDP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the listener, such as
`arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh/listener/0123vxyz`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ListenerArn`

The ARN of the listener, such as
`arn:aws:globalaccelerator::012345678901:accelerator/1234abcd-abcd-1234-abcd-1234abcdefgh/listener/0123vxyz`.

## Examples

### Add a listener

These are examples to specify a listener.

#### JSON

```json

"Resources": {
    "Listener": {
        "Type": "AWS::GlobalAccelerator::Listener",
        "Properties": {
            "AcceleratorArn": {
                "Ref": "Accelerator"
            },
            "Protocol": "TCP",
            "PortRanges": [
                {
                    "FromPort": 80,
                    "ToPort": 80
                }
            ]
        }
    }
}
```

#### YAML

```yaml

Listener:
  Type: AWS::GlobalAccelerator::Listener
  Properties:
    AcceleratorArn:
      Ref: Accelerator
    Protocol: TCP
    PortRanges:
    - FromPort: 80
      ToPort: 80
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortOverride

PortRange

All content copied from https://docs.aws.amazon.com/.
