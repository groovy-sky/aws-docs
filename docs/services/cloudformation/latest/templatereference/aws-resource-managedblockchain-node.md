This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Node

Creates a node on the specified blockchain network.

Applies to Hyperledger Fabric and Ethereum.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ManagedBlockchain::Node",
  "Properties" : {
      "MemberId" : String,
      "NetworkId" : String,
      "NodeConfiguration" : NodeConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::ManagedBlockchain::Node
Properties:
  MemberId: String
  NetworkId: String
  NodeConfiguration:
    NodeConfiguration

```

## Properties

`MemberId`

The unique identifier of the member to which the node belongs. Applies only to Hyperledger Fabric.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkId`

The unique identifier of the network for the node.

Ethereum public networks have the following `NetworkId` s:

- `n-ethereum-mainnet`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeConfiguration`

Configuration properties of a peer node.

_Required_: Yes

_Type_: [NodeConfiguration](aws-properties-managedblockchain-node-nodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the node ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the node.

`MemberId`

The unique identifier of the member in which the node is created. Applies only to Hyperledger Fabric.

`NetworkId`

The unique identifier of the network that the node is in.

`NodeId`

The unique identifier of the node.

## Examples

- [Create a node on an Ethereum network](#aws-resource-managedblockchain-node--examples--Create_a_node_on_an_Ethereum_network)

- [Create a node in a Hyperledger Fabric network member](#aws-resource-managedblockchain-node--examples--Create_a_node_in_a_Hyperledger_Fabric_network_member)

### Create a node on an Ethereum network

#### YAML

```yaml

Description: "Basic Ethereum node template"
Parameters:
  NetworkId:
    Type: String
  InstanceType:
    Type: String
  AvailabilityZone:
    Type: String
Resources:
  Node:
    Type: "AWS::ManagedBlockchain::Node"
    Properties:
      NetworkId: !Ref NetworkId
      NodeConfiguration:
        InstanceType: !Ref InstanceType
        AvailabilityZone: !Ref AvailabilityZone
```

#### JSON

```json

{
    "Description": "Basic Ethereum node template",
    "Parameters": {
        "NetworkId": {
            "Type": "String"
        },
        "InstanceType": {
            "Type": "String"
        },
        "AvailabilityZone": {
            "Type": "String"
        }
    },
    "Resources": {
        "Node": {
            "Type": "AWS::ManagedBlockchain::Node",
            "Properties": {
                "NetworkId": "NetworkId",
                "NodeConfiguration": {
                    "InstanceType": "InstanceType",
                    "AvailabilityZone": "AvailabilityZone"
                }
            }
        }
    }
}
```

### Create a node in a Hyperledger Fabric network member

#### YAML

```yaml

Description: "Basic Hyperledger Fabric node template"
Parameters:
  NetworkId:
    Type: String
  MemberId:
    Type: String
  InstanceType:
    Type: String
  AvailabilityZone:
    Type: String

Resources:
  InitialNode:
    Type: "AWS::ManagedBlockchain::Node"
    Properties:
      NetworkId: !Ref NetworkId
      MemberId: !Ref MemberId
      NodeConfiguration:
        InstanceType: !Ref InstanceType
        AvailabilityZone: !Ref AvailabilityZone

```

#### JSON

```json

{
  "Description": "Basic Hyperledger Fabric node template",
  "Parameters": {
    "NetworkId": {
      "Type": "String"
    },
    "MemberId": {
      "Type": "String"
    },
    "InstanceType": {
      "Type": "String"
    },
    "AvailabilityZone": {
      "Type": "String"
    }
  },
  "Resources": {
    "InitialNode": {
      "Type": "AWS::ManagedBlockchain::Node",
      "Properties": {
        "NetworkId": "NetworkId",
        "MemberId": "MemberId",
        "NodeConfiguration": {
          "InstanceType": "InstanceType",
          "AvailabilityZone": "AvailabilityZone"
        }
      }
    }
  }
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VotingPolicy

NodeConfiguration

All content copied from https://docs.aws.amazon.com/.
