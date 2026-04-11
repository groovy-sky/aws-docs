# Required permissions for Spot placement score

By default, IAM identities (users, roles, or groups) don't have permission to use
[Spot placement score](spot-placement-score.md). To allow IAM identities to use Spot placement score, you
must create an IAM policy that grants permission to use the
`ec2:GetSpotPlacementScores` EC2 API action. You then attach the policy
to the IAM identities that require this permission.

The following is an example IAM policy that grants permission to use the
`ec2:GetSpotPlacementScores` EC2 API action.

JSON

```json

{
    "Version":"2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "ec2:GetSpotPlacementScores",
            "Resource": "*"
        }
    ]
}

```

For information about editing an IAM policy, see [Editing IAM policies](../../../iam/latest/userguide/access-policies-manage-edit.md)
in the _IAM User Guide_.

To provide access, add permissions to your users, groups, or roles:

- Users and groups in AWS IAM Identity Center:

Create a permission set. Follow the instructions in [Create a permission set](../../../singlesignon/latest/userguide/howtocreatepermissionset.md) in the _AWS IAM Identity Center User Guide_.

- Users managed in IAM through an identity provider:

Create a role for identity federation. Follow the instructions in [Create a role for a third-party identity provider (federation)](../../../iam/latest/userguide/id-roles-create-for-idp.md)
in the _IAM User Guide_.

- IAM users:

- Create a role that your user can assume. Follow the instructions in [Create a role for an IAM user](../../../iam/latest/userguide/id-roles-create-for-user.md) in the _IAM User Guide_.

- (Not recommended) Attach a policy directly to a user or add a user to a user group. Follow the instructions in [Adding permissions to a user (console)](../../../iam/latest/userguide/id-users-change-permissions.md#users_change_permissions-add-console) in the _IAM User Guide_.

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

How Spot placement score works

Calculate the Spot placement score
