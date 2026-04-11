# Amazon ECR public identity-based policy examples

By default, users and roles don't have permission to create or modify Amazon ECR Public
resources. To grant users permission to perform actions on the
resources that they need, an IAM administrator can create IAM policies.

To learn how to create an IAM identity-based policy by using these example JSON policy
documents, see [Create IAM policies (console)](../../../iam/latest/userguide/access-policies-create-console.md) in the
_IAM User Guide_.

For details about actions and resource types defined by Amazon ECR, including the format of the ARNs for each of the resource types, see [Actions, resources, and condition keys for Amazon Elastic Container Registry](../../../service-authorization/latest/reference/ecr.md) in the _Service Authorization Reference_.

###### Topics

- [Policy best practices](#security_iam_service-with-iam-policy-best-practices)

- [Using the Amazon ECR console](#security_iam_id-based-policy-examples-console)

- [Allow users to view their own permissions](#security_iam_id-based-policy-examples-view-own-permissions)

- [Accessing an Amazon ECR public repository](#security_iam_id-based-policy-examples-access-one-bucket)

## Policy best practices

Identity-based policies determine whether someone can create, access, or delete Amazon ECR Public resources in your
account. These actions can incur costs for your AWS account. When you create or edit identity-based policies, follow these guidelines and
recommendations:

- **Get started with AWS managed policies and move toward least-privilege permissions**
– To get started granting permissions to your users and workloads, use the _AWS_
_managed policies_ that grant permissions for many common use cases. They are
available in your AWS account. We recommend that you reduce permissions further by
defining AWS customer managed policies that are specific to your use cases. For more information, see
[AWS managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies) or [AWS managed policies for job functions](../../../iam/latest/userguide/access-policies-job-functions.md) in the _IAM User Guide_.

- **Apply least-privilege permissions** –
When you set permissions with IAM policies, grant only the permissions required to
perform a task. You do this by defining the actions that can be taken on specific resources
under specific conditions, also known as _least-privilege permissions_.
For more information about using IAM to apply permissions, see [Policies and permissions in IAM](../../../iam/latest/userguide/access-policies.md) in the _IAM User Guide_.

- **Use conditions in IAM policies to further restrict access**
– You can add a condition to your policies to limit access to actions and resources. For example, you can write a policy condition to specify that all requests must
be sent using SSL. You can also use conditions to grant access to service actions
if they are used through a specific AWS service, such as CloudFormation. For more information, see
[IAM JSON policy elements: Condition](../../../iam/latest/userguide/reference-policies-elements-condition.md) in the _IAM User Guide_.

- **Use IAM Access Analyzer to validate your IAM policies to ensure secure and functional permissions**
– IAM Access Analyzer validates new and existing policies so that the policies adhere to the IAM policy language (JSON) and IAM best practices.
IAM Access Analyzer provides more than 100 policy checks and actionable recommendations to help
you author secure and functional policies. For more information, see [Validate policies with IAM Access Analyzer](../../../iam/latest/userguide/access-analyzer-policy-validation.md) in the _IAM User Guide_.

- **Require multi-factor authentication (MFA)** –
If you have a scenario that requires IAM users or a root user in your AWS account, turn on MFA for additional security. To require
MFA when API operations are called, add MFA conditions to your policies. For
more information, see [Secure API access with MFA](../../../iam/latest/userguide/id-credentials-mfa-configure-api-require.md) in the _IAM User Guide_.

For more information about best practices in IAM, see [Security best practices in IAM](../../../iam/latest/userguide/best-practices.md) in the _IAM User Guide_.

## Using the Amazon ECR console

###### Note

These permissions are only for Amazon ECR Public resources. For permissions related
to your private Amazon ECR resources, see [Amazon ECR identity-based policy examples](../userguide/security-iam-id-based-policy-examples.md) in the
_Amazon Elastic Container Registry User Guide_.

To access the Amazon ECR console, you must have a minimum set of permissions. These
permissions must allow you to list and view details about the Amazon ECR Public resources
in your AWS account. If you create an identity-based policy that is more
restrictive than the minimum required permissions, the console won't function as
intended for entities (users or roles) with that policy.

To ensure that those entities can still use the Amazon ECR console, add the
`AmazonElasticContainerRegistryPublicReadOnly` AWS managed policy
to the entities. For more information, see [Adding Permissions to a User](../../../iam/latest/userguide/id-users-change-permissions.md#users_change_permissions-add-console) in the
_IAM User Guide_:

To view the permissions for this policy, see [AmazonElasticContainerRegistryPublicReadOnly](../../../aws-managed-policy/latest/reference/amazonelasticcontainerregistrypublicreadonly.md) in the _AWS Managed Policy_
_Reference_.

You don't need to allow minimum console permissions for users that are making
calls only to the AWS CLI or the AWS API. Instead, allow access to only the actions
that match the API operation that you're trying to perform.

## Allow users to view their own permissions

This example shows how you might create a policy that allows IAM users to view the inline and managed policies that are attached to their user
identity. This policy includes permissions to complete this action on the console or programmatically using the AWS CLI or AWS API.

```json

{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "ViewOwnUserInfo",
            "Effect": "Allow",
            "Action": [
                "iam:GetUserPolicy",
                "iam:ListGroupsForUser",
                "iam:ListAttachedUserPolicies",
                "iam:ListUserPolicies",
                "iam:GetUser"
            ],
            "Resource": ["arn:aws:iam::*:user/${aws:username}"]
        },
        {
            "Sid": "NavigateInConsole",
            "Effect": "Allow",
            "Action": [
                "iam:GetGroupPolicy",
                "iam:GetPolicyVersion",
                "iam:GetPolicy",
                "iam:ListAttachedGroupPolicies",
                "iam:ListGroupPolicies",
                "iam:ListPolicyVersions",
                "iam:ListPolicies",
                "iam:ListUsers"
            ],
            "Resource": "*"
        }
    ]
}
```

## Accessing an Amazon ECR public repository

In this example, you want to grant an IAM user in your AWS account access to
one of your Amazon ECR public repositories, `my-repo`.

JSON

```json

{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Sid":"GetAuthorizationToken",
         "Effect":"Allow",
         "Action":[
            "ecr-public:GetAuthorizationToken"
         ],
         "Resource":"*"
      },
      {
         "Sid":"ManageRepositoryContents",
         "Effect":"Allow",
         "Action":[
                "ecr-public:BatchCheckLayerAvailability",
                "ecr-public:GetRepositoryPolicy",
                "ecr-public:DescribeRepositories",
                "ecr-public:DescribeImages",
                "ecr-public:InitiateLayerUpload",
                "ecr-public:UploadLayerPart",
                "ecr-public:CompleteLayerUpload",
                "ecr-public:PutImage"
         ],
         "Resource":"arn:aws:ecr-public::123456789012:repository/my-repo"
      }
   ]
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS managed policies for Amazon ECR Public

Using tag-based access control in Amazon ECR public

All content copied from https://docs.aws.amazon.com/.
