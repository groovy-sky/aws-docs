---
title: "AWSIQFullAccess"
---

# AWSIQFullAccess

**Description**: Provides full access to AWS IQ

`AWSIQFullAccess` is an [AWS managed policy](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies).

## Using this policy

You can attach `AWSIQFullAccess` to your users, groups, and roles.

## Policy details

- **Type**: AWS managed policy

- **Creation time**: April 04, 2019, 23:13 UTC

- **Edited time:** September 25, 2019, 20:22 UTC

- **ARN**:
`arn:aws:iam::aws:policy/AWSIQFullAccess`

## Policy version

**Policy version:** v2 (default)

The policy's default version is the version that defines the permissions for the policy. When a user or role with the policy makes a
request to access an AWS resource, AWS checks the default version of the policy to determine whether to allow the request.

## JSON policy document

```json

{
  "Version" : "2012-10-17",
  "Statement" : [
    {
      "Action" : [
        "iq:*",
        "iq-permission:*"
      ],
      "Effect" : "Allow",
      "Resource" : "*"
    },
    {
      "Effect" : "Allow",
      "Action" : "iam:CreateServiceLinkedRole",
      "Resource" : "*",
      "Condition" : {
        "StringEquals" : {
          "iam:AWSServiceName" : [
            "permission.iq.amazonaws.com",
            "contract.iq.amazonaws.com"
          ]
        }
      }
    }
  ]
}
```

## Learn more

- [Create a permission set using AWS managed policies in IAM Identity Center](../../../singlesignon/latest/userguide/howtocreatepermissionset.md)

- [Adding and removing IAM identity permissions](../../../iam/latest/userguide/access-policies-manage-attach-detach.md)

- [Understand versioning for IAM policies](../../../iam/latest/userguide/access-policies-managed-versioning.md)

- [Get started with AWS managed policies and move toward least-privilege permissions](../../../iam/latest/userguide/best-practices.md#bp-use-aws-defined-policies)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSIQContractServiceRolePolicy

AWSIQPermissionServiceRolePolicy

All content copied from https://docs.aws.amazon.com/.
