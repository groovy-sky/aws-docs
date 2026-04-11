This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MPA::ApprovalTeam ApprovalStrategy

Strategy for how an approval team grants approval.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MofN" : MofNApprovalStrategy
}

```

### YAML

```yaml

  MofN:
    MofNApprovalStrategy

```

## Properties

`MofN`

Minimum number of approvals (M) required for a total number of approvers (N).

_Required_: Yes

_Type_: [MofNApprovalStrategy](aws-properties-mpa-approvalteam-mofnapprovalstrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MPA::ApprovalTeam

Approver

All content copied from https://docs.aws.amazon.com/.
