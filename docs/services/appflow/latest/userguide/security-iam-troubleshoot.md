# Troubleshooting Amazon AppFlow identity and access

Use the following information to help you diagnose and fix common issues that you might
encounter when working with Amazon AppFlow and IAM.

###### Topics

- [I am not authorized to perform an action in Amazon AppFlow](#security_iam_troubleshoot-no-permissions)

- [I am not authorized to perform iam:PassRole](#security_iam_troubleshoot-passrole)

- [I'm an administrator and want to allow others to access Amazon AppFlow](#security_iam_troubleshoot-admin-delegate)

- [I want to allow people outside of my AWS account to access my Amazon AppFlow resources](#security_iam_troubleshoot-cross-account-access)

## I am not authorized to perform an action in Amazon AppFlow

If you receive an error that you're not authorized to perform an action, your
policies must be updated to allow you to perform the action.

The following example error occurs when the `mateojackson` IAM user
tries to use the console to view details about a fictional
`my-example-widget` resource but doesn't
have the fictional `appflow:GetWidget` permissions.

```replaceable
User: arn:aws:iam::123456789012:user/mateojackson is not authorized to perform: appflow:GetWidget on resource: my-example-widget
```

In this case, the policy for the `mateojackson` user must be updated to allow access to the
`my-example-widget` resource by using the
`appflow:GetWidget` action.

If you need help, contact your AWS administrator. Your administrator is the person who provided you with your sign-in credentials.

## I am not authorized to perform iam:PassRole

If you receive an error that you're not authorized to perform the `iam:PassRole` action, your policies must be updated to allow you to pass a role to Amazon AppFlow.

Some AWS services allow you to pass an existing role to that service instead of creating a new service role or service-linked role. To do
this, you must have permissions to pass the role to the service.

The following example error occurs when an IAM user named `marymajor` tries to use the console to perform an action in
Amazon AppFlow. However, the action requires the service to have permissions that are granted by a service role. Mary does not have permissions to pass the
role to the service.

```nohighlight

User: arn:aws:iam::123456789012:user/marymajor is not authorized to perform: iam:PassRole
```

In this case, Mary's policies must be updated to allow her to perform the `iam:PassRole` action.

If you need help, contact your AWS administrator. Your administrator is the person who provided you with your sign-in credentials.

## I'm an administrator and want to allow others to access Amazon AppFlow

To allow others to access Amazon AppFlow, you must grant permission to the people or applications that need access. If you are using AWS IAM Identity Center
to manage people and applications, you assign permission sets to users or groups to define their level of access. Permission sets automatically create
and assign IAM policies to IAM roles that are associated with the person or application. For more information, see [Permission sets](../../../singlesignon/latest/userguide/permissionsetsconcept.md) in the _AWS IAM Identity Center User Guide_.

If you are not using IAM Identity Center, you must create IAM entities (users or roles) for the people or applications that need access. You must then attach
a policy to the entity that grants them the correct permissions in Amazon AppFlow. After the permissions are granted, provide the credentials to the user
or application developer. They will use those credentials to access AWS. To learn more about creating IAM users, groups, policies, and permissions,
see [IAM Identities](../../../iam/latest/userguide/id.md) and [Policies and permissions in IAM](../../../iam/latest/userguide/access-policies.md) in the _IAM User Guide_.

## I want to allow people outside of my AWS account to access my Amazon AppFlow resources

You can create a role that users in other accounts or people outside of your organization can use to access your resources. You can specify who
is trusted to assume the role. For services that support resource-based policies or access control lists (ACLs), you can use those policies to grant
people access to your resources.

To learn more, consult the following:

- To learn whether Amazon AppFlow supports these features, see [How Amazon AppFlow works with IAM](security-iam-service-with-iam.md).

- To learn how to provide access to your resources across AWS accounts that you own, see [Providing access to an IAM user in another AWS account that you\
own](../../../iam/latest/userguide/id-roles-common-scenarios-aws-accounts.md) in the _IAM User Guide_.

- To learn how to provide access to your resources to third-party AWS accounts, see [Providing access to AWS accounts owned by third parties](../../../iam/latest/userguide/id-roles-common-scenarios-third-party.md) in the
_IAM User Guide_.

- To learn how to provide access through identity federation, see [Providing access to externally authenticated users (identity federation)](../../../iam/latest/userguide/id-roles-common-scenarios-federated-users.md) in the _IAM User Guide_.

- To learn the difference between using roles and resource-based policies for cross-account access, see [Cross account resource access in IAM](../../../iam/latest/userguide/access-policies-cross-account-resource-access.md) in the
_IAM User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS managed policies

Compliance validation

All content copied from https://docs.aws.amazon.com/.
