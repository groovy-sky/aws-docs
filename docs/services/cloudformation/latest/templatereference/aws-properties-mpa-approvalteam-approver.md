This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MPA::ApprovalTeam Approver

Contains details for an approver.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApproverId" : String,
  "PrimaryIdentityId" : String,
  "PrimaryIdentitySourceArn" : String,
  "PrimaryIdentityStatus" : String,
  "ResponseTime" : String
}

```

### YAML

```yaml

  ApproverId: String
  PrimaryIdentityId: String
  PrimaryIdentitySourceArn: String
  PrimaryIdentityStatus: String
  ResponseTime: String

```

## Properties

`ApproverId`

ID for the approver.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryIdentityId`

ID for the user.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryIdentitySourceArn`

Amazon Resource Name (ARN) for the identity source. The identity source manages the user authentication for approvers.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryIdentityStatus`

Status for the identity source. For example, if an approver has accepted a team invitation with a user authentication method managed by the identity source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseTime`

Timestamp when the approver responded to an approval team invitation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApprovalStrategy

MofNApprovalStrategy
