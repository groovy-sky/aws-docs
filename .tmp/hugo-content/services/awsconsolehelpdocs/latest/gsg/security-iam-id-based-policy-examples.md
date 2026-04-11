# Identity-based policy examples for AWS User Experience Customization

By default, users and roles don't have permission to get or modify UXC resources. To
grant users permission to perform actions on resources, an IAM administrator can create
IAM policies. To learn how to create an IAM identity-based policy by using these example
JSON policy documents, see [Create IAM policies\
(console)](../../../iam/latest/userguide/access-policies-create-console.md) in the _IAM User Guide_.

###### Topics

- [Policy best practices](#security_iam_id-based-policy-examples-best-practices)

- [Read-only access to UXC Account Customizations](#security_iam_id-based-policy-examples-readonly)

- [Full access to UXC Account Customizations](#security_iam_id-based-policy-examples-fullaccess)

## Policy best practices

Identity-based policies determine whether someone can create, access, or delete User Experience Customization resources in your
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

## Read-only access to UXC Account Customizations

The following example shows how to create a policy that allows read-only access to
UXC Account Customizations:

```

{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "uxc:GetAccountCustomizations",
                "uxc:ListServices"
            ],
            "Resource": "*"
        }
    ]
}
```

## Full access to UXC Account Customizations

The following example shows how to create a policy that allows full access to UXC
Account Customizations:

```

{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "uxc:*"
            ],
            "Resource": "*"
        }
    ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

How User Experience Customization works with IAM

AWS managed policies

All content copied from https://docs.aws.amazon.com/.
