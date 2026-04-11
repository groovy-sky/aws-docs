AWS App Runner will no longer be open to new customers starting April 30, 2026. If you would like to use
App Runner, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS App Runner availability\
change](apprunner-availability-change.md).

# Troubleshooting App Runner identity and access

Use the following information to help you diagnose and fix common issues that you might encounter when working with AWS App Runner and IAM.

For other App Runner security topics, see [Security in App Runner](security.md).

###### Topics

- [I'm not authorized to perform an action in App Runner](#security_iam_troubleshoot-no-permissions)

- [I want to allow people outside of my AWS account to access my App Runner resources](#security_iam_troubleshoot-cross-account-access)

## I'm not authorized to perform an action in App Runner

If the AWS Management Console tells you that you're not authorized to perform an action, contact your administrator for assistance. Your administrator is the
person that provided you with your AWS sign-in credentials.

The following example error occurs when an IAM user named `marymajor` tries to use the console to view details about an App Runner service
but doesn't have `apprunner:DescribeService` permissions.

```nohighlight

User: arn:aws:iam::123456789012:user/marymajor is not authorized to perform: apprunner:DescribeService on resource: my-example-service
```

In this case, Mary asks her administrator to update her policies to allow her to access the
`my-example-service` resource using the `apprunner:DescribeService` action.

## I want to allow people outside of my AWS account to access my App Runner resources

You can create a role that users in other accounts or people outside of your organization can use to access your resources. You can specify who
is trusted to assume the role. For services that support resource-based policies or access control lists (ACLs), you can use those policies to grant
people access to your resources.

To learn more, consult the following:

- To learn whether App Runner supports these features, see [How App Runner works with IAM](security-iam-service-with-iam.md).

- To learn how to provide access to your resources across AWS accounts that you own, see [Providing access to an IAM user in another AWS account that you\
own](../../../iam/latest/userguide/id-roles-common-scenarios-aws-accounts.md) in the _IAM User Guide_.

- To learn how to provide access to your resources to third-party AWS accounts, see [Providing access to AWS accounts owned by third parties](../../../iam/latest/userguide/id-roles-common-scenarios-third-party.md) in the
_IAM User Guide_.

- To learn how to provide access through identity federation, see [Providing access to externally authenticated users (identity federation)](../../../iam/latest/userguide/id-roles-common-scenarios-federated-users.md) in the _IAM User Guide_.

- To learn the difference between using roles and resource-based policies for cross-account access, see [Cross account resource access in IAM](../../../iam/latest/userguide/access-policies-cross-account-resource-access.md) in the
_IAM User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS managed policies

Logging and monitoring

All content copied from https://docs.aws.amazon.com/.
