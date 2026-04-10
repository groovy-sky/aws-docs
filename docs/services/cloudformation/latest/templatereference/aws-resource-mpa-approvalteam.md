This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MPA::ApprovalTeam

Creates a new approval team. For more information, see [Approval team](../../../mpa/latest/userguide/mpa-concepts.md) in the _Multi-party approval User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MPA::ApprovalTeam",
  "Properties" : {
      "ApprovalStrategy" : ApprovalStrategy,
      "Approvers" : [ Approver, ... ],
      "Description" : String,
      "Name" : String,
      "Policies" : [ Policy, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MPA::ApprovalTeam
Properties:
  ApprovalStrategy:
    ApprovalStrategy
  Approvers:
    - Approver
  Description: String
  Name: String
  Policies:
    - Policy
  Tags:
    - Tag

```

## Properties

`ApprovalStrategy`

Contains details for how an approval team grants approval.

_Required_: Yes

_Type_: [ApprovalStrategy](aws-properties-mpa-approvalteam-approvalstrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Approvers`

Contains details for an approver.

_Required_: Yes

_Type_: Array of [Approver](aws-properties-mpa-approvalteam-approver.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Description for the team.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the team.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policies`

Contains details for a policy. Policies define what operations a team that define the permissions for team resources.

_Required_: Yes

_Type_: Array of [Policy](aws-properties-mpa-approvalteam-policy.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags that you have added to the specified resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-mpa-approvalteam-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Amazon Resource Name (ARN) for the team.

`CreationTime`

Timestamp when the team was created.

`LastUpdateTime`

Timestamp when the team was last updated.

`NumberOfApprovers`

Total number of approvers in the team.

`Status`

Status for the team. For more information, see [Team health](../../../mpa/latest/userguide/mpa-health.md) in the _Multi-party approval User Guide_.

`StatusCode`

Status code for the team. For more information, see [Team health](../../../mpa/latest/userguide/mpa-health.md) in the _Multi-party approval User Guide_.

`StatusMessage`

Message describing the status for the team.

`UpdateSessionArn`

Timestamp when the team was last updated.

`VersionId`

Version ID for the team.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Multi-party approval

ApprovalStrategy

All content copied from https://docs.aws.amazon.com/.
