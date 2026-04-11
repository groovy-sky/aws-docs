This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ManagedBlockchain::Member ApprovalThresholdPolicy

A policy type that defines the voting rules for the network. The rules decide if a proposal is approved. Approval may be based on criteria such as the percentage of `YES` votes and the duration of the proposal. The policy applies to all proposals and is specified when the network is created.

Applies only to Hyperledger Fabric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProposalDurationInHours" : Integer,
  "ThresholdComparator" : String,
  "ThresholdPercentage" : Integer
}

```

### YAML

```yaml

  ProposalDurationInHours: Integer
  ThresholdComparator: String
  ThresholdPercentage: Integer

```

## Properties

`ProposalDurationInHours`

The duration from the time that a proposal is created until it expires. If members cast neither the required number of `YES` votes to approve the proposal nor the number of `NO` votes required to reject it before the duration expires, the proposal is `EXPIRED` and `ProposalActions` aren't carried out.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `168`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThresholdComparator`

Determines whether the vote percentage must be greater than the
`ThresholdPercentage` or must be greater than or equal to the
`ThresholdPercentage` to be approved.

_Required_: No

_Type_: String

_Allowed values_: `GREATER_THAN | GREATER_THAN_OR_EQUAL_TO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThresholdPercentage`

The percentage of votes among all members that must be `YES` for a proposal to be approved. For example, a `ThresholdPercentage` value of `50` indicates 50%. The `ThresholdComparator` determines the precise comparison. If a `ThresholdPercentage` value of `50` is specified on a network with 10 members, along with a `ThresholdComparator` value of `GREATER_THAN`, this indicates that 6 `YES` votes are required for the proposal to be approved.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ManagedBlockchain::Member

MemberConfiguration

All content copied from https://docs.aws.amazon.com/.
