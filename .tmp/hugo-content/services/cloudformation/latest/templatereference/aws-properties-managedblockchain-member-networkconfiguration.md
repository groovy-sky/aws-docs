This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Member NetworkConfiguration

Configuration properties of the network to which the member belongs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Framework" : String,
  "FrameworkVersion" : String,
  "Name" : String,
  "NetworkFrameworkConfiguration" : NetworkFrameworkConfiguration,
  "VotingPolicy" : VotingPolicy
}

```

### YAML

```yaml

  Description: String
  Framework: String
  FrameworkVersion: String
  Name: String
  NetworkFrameworkConfiguration:
    NetworkFrameworkConfiguration
  VotingPolicy:
    VotingPolicy

```

## Properties

`Description`

Attributes of the blockchain framework for the network.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Framework`

The blockchain framework that the network uses.

_Required_: Yes

_Type_: String

_Allowed values_: `HYPERLEDGER_FABRIC | ETHEREUM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameworkVersion`

The version of the blockchain framework that the network uses.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the network.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkFrameworkConfiguration`

Configuration properties relevant to the network for the blockchain framework that the network uses.

_Required_: No

_Type_: [NetworkFrameworkConfiguration](aws-properties-managedblockchain-member-networkframeworkconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VotingPolicy`

The voting rules that the network uses to decide if a proposal is accepted.

_Required_: Yes

_Type_: [VotingPolicy](aws-properties-managedblockchain-member-votingpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemberFrameworkConfiguration

NetworkFabricConfiguration

All content copied from https://docs.aws.amazon.com/.
