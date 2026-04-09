AWS App Runner will no longer be open to new customers starting April 30, 2026. If you would like to use
App Runner, sign up prior to that date. Existing customers can continue to use the service as normal. For more information, see
[AWS App Runner availability\
change](apprunner-availability-change.md).

# Setting up for App Runner

If you're a new AWS customer, complete the setup prerequisites that are listed on this page before you start using AWS App Runner.

For these setup procedures, you use the AWS Identity and Access Management (IAM) service. For complete information about IAM, see the following reference materials:

- [AWS Identity and Access Management (IAM)](https://aws.amazon.com/iam)

- [IAM User Guide](../../../iam/latest/userguide.md)

## Sign up for an AWS account

If you do not have an AWS account, complete the following steps to create one.

###### To sign up for an AWS account

1. Open [https://portal.aws.amazon.com/billing/signup](https://portal.aws.amazon.com/billing/signup).

2. Follow the online instructions.

Part of the sign-up procedure involves receiving a phone call or text message and entering
    a verification code on the phone keypad.

When you sign up for an AWS account, an _AWS account root user_ is created. The root user has access to all AWS services
    and resources in the account. As a security best practice, assign administrative access to a user, and use only the root user to perform [tasks that require root user access](../../../iam/latest/userguide/id-root-user.md#root-user-tasks).

AWS sends you a confirmation email after the sign-up process is
complete. At any time, you can view your current account activity and manage your account by
going to [https://aws.amazon.com/](https://aws.amazon.com/) and choosing **My**
**Account**.

## Create a user with administrative access

After you sign up for an AWS account, secure your AWS account root user, enable AWS IAM Identity Center, and create an administrative user so that you
don't use the root user for everyday tasks.

###### Secure your AWS account root user

1. Sign in to the [AWS Management Console](https://console.aws.amazon.com/) as the account owner by choosing **Root user** and entering your AWS account email address. On the next page, enter your password.

For help signing in by using root user, see [Signing in as the root user](../../../signin/latest/userguide/console-sign-in-tutorials.md#introduction-to-root-user-sign-in-tutorial) in the _AWS Sign-In User Guide_.

2. Turn on multi-factor authentication (MFA) for your root user.

For instructions, see [Enable a virtual MFA device for your AWS account root user (console)](../../../iam/latest/userguide/enable-virt-mfa-for-root.md) in the _IAM User Guide_.

###### Create a user with administrative access

1. Enable IAM Identity Center.

For instructions, see [Enabling\
    AWS IAM Identity Center](../../../singlesignon/latest/userguide/get-set-up-for-idc.md) in the
    _AWS IAM Identity Center User Guide_.

2. In IAM Identity Center, grant administrative access to a user.

For a tutorial about using the IAM Identity Center directory as your identity source, see [Configure user access with the default IAM Identity Center directory](../../../singlesignon/latest/userguide/quick-start-default-idc.md) in the
    _AWS IAM Identity Center User Guide_.

###### Sign in as the user with administrative access

- To sign in with your IAM Identity Center user, use the sign-in URL that was sent to your email address when you created the IAM Identity Center user.

For help signing in using an IAM Identity Center user, see [Signing in to the AWS access portal](../../../signin/latest/userguide/iam-id-center-sign-in-tutorial.md) in the _AWS Sign-In User Guide_.

###### Assign access to additional users

1. In IAM Identity Center, create a permission set that follows the best practice of applying least-privilege permissions.

For instructions, see [Create a permission set](../../../singlesignon/latest/userguide/get-started-create-a-permission-set.md) in the _AWS IAM Identity Center User Guide_.

2. Assign users to a group, and then assign single sign-on access to the group.

For instructions, see [Add groups](../../../singlesignon/latest/userguide/addgroups.md) in the _AWS IAM Identity Center User Guide_.

## Grant programmatic access

Users need programmatic access if they want to interact with AWS outside of the AWS Management Console. The way to grant programmatic access depends on the type of user that's accessing AWS.

To grant users programmatic access, choose one of the following options.

Which user needs programmatic access?ToByIAM(Recommended) Use console credentials as temporary credentials to sign programmatic requests to the AWS CLI, AWS SDKs, or AWS APIs.

Following the instructions for the interface that you want to use.

- For the AWS CLI, see [Login for AWS local development](../../../cli/latest/userguide/cli-configure-sign-in.md) in
the _AWS Command Line Interface User Guide_.

- For AWS SDKs, see [Login for AWS local development](../../../../reference/sdkref/latest/guide/access-login.md) in the
_AWS SDKs and Tools Reference Guide_.

Workforce identity

(Users managed in IAM Identity Center)

Use temporary credentials to sign programmatic requests to the AWS CLI, AWS SDKs, or
AWS APIs.

Following the instructions for the interface that you want to use.

- For the AWS CLI, see [Configuring the AWS CLI to use AWS IAM Identity Center](../../../cli/latest/userguide/cli-configure-sso.md) in the
_AWS Command Line Interface User Guide_.

- For AWS SDKs, tools, and AWS APIs, see [IAM Identity Center\
authentication](../../../../reference/sdkref/latest/guide/access-sso.md) in the _AWS SDKs and Tools Reference Guide_.

IAMUse temporary credentials to sign programmatic requests to the AWS CLI, AWS SDKs, or
AWS APIs.Following the instructions in [Using temporary\
credentials with AWS resources](../../../iam/latest/userguide/id-credentials-temp-use-resources.md) in the _IAM User Guide_.IAM

(Not recommended)

Use long-term credentials to sign programmatic requests
to the AWS CLI, AWS SDKs, or AWS APIs.

Following the instructions for the interface that you want to use.

- For the AWS CLI, see [Authenticating using IAM user credentials](../../../cli/latest/userguide/cli-authentication-user.md) in
the _AWS Command Line Interface User Guide_.

- For AWS SDKs and tools, see [Authenticate using long-term credentials](../../../../reference/sdkref/latest/guide/access-iam-users.md) in the
_AWS SDKs and Tools Reference Guide_.

- For AWS APIs, see [Managing access keys for\
IAM users](../../../iam/latest/userguide/id-credentials-access-keys.md) in the _IAM User Guide_.

## What's next

You completed the prerequisite steps. To deploy your first application to App Runner, see [Getting started with App Runner](getting-started.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Availability change

Getting started

All content copied from https://docs.aws.amazon.com/.
