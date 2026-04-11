This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Member VotingPolicy

The voting rules for the network to decide if a proposal is accepted

Applies only to Hyperledger Fabric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApprovalThresholdPolicy" : ApprovalThresholdPolicy
}

```

### YAML

```yaml

  ApprovalThresholdPolicy:
    ApprovalThresholdPolicy

```

## Properties

`ApprovalThresholdPolicy`

Defines the rules for the network for voting on proposals, such as the percentage of `YES` votes required for the proposal to be approved and the duration of the proposal. The policy applies to all proposals and is specified when the network is created.

_Required_: No

_Type_: [ApprovalThresholdPolicy](aws-properties-managedblockchain-member-approvalthresholdpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkFrameworkConfiguration

AWS::ManagedBlockchain::Node

All content copied from https://docs.aws.amazon.com/.
