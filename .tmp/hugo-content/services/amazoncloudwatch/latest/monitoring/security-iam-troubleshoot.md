# Troubleshooting Amazon CloudWatch identity and access

Use the following information to help you diagnose and fix common issues that you might
encounter when working with CloudWatch and IAM.

###### Topics

- [I am not authorized to perform an action in CloudWatch](#security_iam_troubleshoot-no-permissions)

- [I am not authorized to perform iam:PassRole](#security_iam_troubleshoot-passrole)

- [I want to allow people outside of my AWS account to access my CloudWatch resources](#security_iam_troubleshoot-cross-account-access)

## I am not authorized to perform an action in CloudWatch

If you receive an error that you're not authorized to perform an action, your
policies must be updated to allow you to perform the action.

The following example error occurs when the `mateojackson` IAM user
tries to use the console to view details about a fictional
`my-example-widget` resource but doesn't
have the fictional `cloudwatch:GetWidget` permissions.

```replaceable
User: arn:aws:iam::123456789012:user/mateojackson is not authorized to perform: cloudwatch:GetWidget on resource: my-example-widget
```

In this case, the policy for the `mateojackson` user must be updated to allow access to the
`my-example-widget` resource by using the
`cloudwatch:GetWidget` action.

If you need help, contact your AWS administrator. Your administrator is the person who provided you with your sign-in credentials.

## I am not authorized to perform iam:PassRole

If you receive an error that you're not authorized to perform the `iam:PassRole` action, your policies must be updated to allow you to pass a role to CloudWatch.

Some AWS services allow you to pass an existing role to that service instead of creating a new service role or service-linked role. To do
this, you must have permissions to pass the role to the service.

The following example error occurs when an IAM user named `marymajor` tries to use the console to perform an action in
CloudWatch. However, the action requires the service to have permissions that are granted by a service role. Mary does not have permissions to pass the
role to the service.

```nohighlight

User: arn:aws:iam::123456789012:user/marymajor is not authorized to perform: iam:PassRole
```

In this case, Mary's policies must be updated to allow her to perform the `iam:PassRole` action.

If you need help, contact your AWS administrator. Your administrator is the person who provided you with your sign-in credentials.

## I want to allow people outside of my AWS account to access my CloudWatch resources

You can create a role that users in other accounts or people outside of your organization can use to access your resources. You can specify who
is trusted to assume the role. For services that support resource-based policies or access control lists (ACLs), you can use those policies to grant
people access to your resources.

To learn more, consult the following:

- To learn whether CloudWatch supports these features, see [How Amazon CloudWatch works with IAM](security-iam-service-with-iam.md).

- To learn how to provide access to your resources across AWS accounts that you own, see [Providing access to an IAM user in another AWS account that you\
own](../../../iam/latest/userguide/id-roles-common-scenarios-aws-accounts.md) in the _IAM User Guide_.

- To learn how to provide access to your resources to third-party AWS accounts, see [Providing access to AWS accounts owned by third parties](../../../iam/latest/userguide/id-roles-common-scenarios-third-party.md) in the
_IAM User Guide_.

- To learn how to provide access through identity federation, see [Providing access to externally authenticated users (identity federation)](../../../iam/latest/userguide/id-roles-common-scenarios-federated-users.md) in the _IAM User Guide_.

- To learn the difference between using roles and resource-based policies for cross-account access, see [Cross account resource access in IAM](../../../iam/latest/userguide/access-policies-cross-account-resource-access.md) in the
_IAM User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Identity-based policy examples

CloudWatch dashboard permissions update

All content copied from https://docs.aws.amazon.com/.
