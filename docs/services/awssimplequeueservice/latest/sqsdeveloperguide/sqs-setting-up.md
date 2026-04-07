# Setting up Amazon SQS

Before you can use Amazon SQS for the first time, you must complete the following steps:

## Step 1: Create an AWS account and IAM user

To access any AWS service, you first need to create an [AWS account](https://aws.amazon.com/), an Amazon.com account that can use AWS products. You can
use your AWS account to view your activity and usage reports and to manage
authentication and access.

To avoid using your AWS account root user for Amazon SQS actions, it is a best practice
to create an IAM user for each person who needs administrative access to Amazon SQS.

### Sign up for an AWS account

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

### Create a user with administrative access

After you sign up for an AWS account, secure your AWS account root user, enable AWS IAM Identity Center, and create an administrative user so that you
don't use the root user for everyday tasks.

###### Secure your AWS account root user

1. Sign in to the [AWS Management Console](https://console.aws.amazon.com/) as the account owner by choosing **Root user** and entering your AWS account email address. On the next page, enter your password.

For help signing in by using root user, see [Signing in as the root user](https://docs.aws.amazon.com/signin/latest/userguide/console-sign-in-tutorials.html#introduction-to-root-user-sign-in-tutorial) in the _AWS Sign-In User Guide_.

2. Turn on multi-factor authentication (MFA) for your root user.

For instructions, see [Enable a virtual MFA device for your AWS account root user (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/enable-virt-mfa-for-root.html) in the _IAM User Guide_.

###### Create a user with administrative access

1. Enable IAM Identity Center.

For instructions, see [Enabling\
    AWS IAM Identity Center](https://docs.aws.amazon.com/singlesignon/latest/userguide/get-set-up-for-idc.html) in the
    _AWS IAM Identity Center User Guide_.

2. In IAM Identity Center, grant administrative access to a user.

For a tutorial about using the IAM Identity Center directory as your identity source, see [Configure user access with the default IAM Identity Center directory](https://docs.aws.amazon.com/singlesignon/latest/userguide/quick-start-default-idc.html) in the
    _AWS IAM Identity Center User Guide_.

###### Sign in as the user with administrative access

- To sign in with your IAM Identity Center user, use the sign-in URL that was sent to your email address when you created the IAM Identity Center user.

For help signing in using an IAM Identity Center user, see [Signing in to the AWS access portal](https://docs.aws.amazon.com/signin/latest/userguide/iam-id-center-sign-in-tutorial.html) in the _AWS Sign-In User Guide_.

###### Assign access to additional users

1. In IAM Identity Center, create a permission set that follows the best practice of applying least-privilege permissions.

For instructions, see [Create a permission set](https://docs.aws.amazon.com/singlesignon/latest/userguide/get-started-create-a-permission-set.html) in the _AWS IAM Identity Center User Guide_.

2. Assign users to a group, and then assign single sign-on access to the group.

For instructions, see [Add groups](https://docs.aws.amazon.com/singlesignon/latest/userguide/addgroups.html) in the _AWS IAM Identity Center User Guide_.

## Step 2: Grant programmatic access

To use Amazon SQS actions (for example, using Java or through the AWS Command Line Interface), you need an
access key ID and a secret access key.

###### Note

The access key ID and secret access key are specific to AWS Identity and Access Management. Don't confuse
them with credentials for other AWS services, such as Amazon EC2 key pairs.

Users need programmatic access if they want to interact with AWS outside of the AWS Management Console. The way to grant programmatic access depends on the type of user that's accessing AWS.

To grant users programmatic access, choose one of the following options.

Which user needs programmatic access?ToByIAM(Recommended) Use console credentials as temporary credentials to sign programmatic requests to the AWS CLI, AWS SDKs, or AWS APIs.

Following the instructions for the interface that you want to use.

- For the AWS CLI, see [Login for AWS local development](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sign-in.html) in
the _AWS Command Line Interface User Guide_.

- For AWS SDKs, see [Login for AWS local development](https://docs.aws.amazon.com/sdkref/latest/guide/access-login.html) in the
_AWS SDKs and Tools Reference Guide_.

Workforce identity

(Users managed in IAM Identity Center)

Use temporary credentials to sign programmatic requests to the AWS CLI, AWS SDKs, or
AWS APIs.

Following the instructions for the interface that you want to use.

- For the AWS CLI, see [Configuring the AWS CLI to use AWS IAM Identity Center](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html) in the
_AWS Command Line Interface User Guide_.

- For AWS SDKs, tools, and AWS APIs, see [IAM Identity Center\
authentication](https://docs.aws.amazon.com/sdkref/latest/guide/access-sso.html) in the _AWS SDKs and Tools Reference Guide_.

IAMUse temporary credentials to sign programmatic requests to the AWS CLI, AWS SDKs, or
AWS APIs.Following the instructions in [Using temporary\
credentials with AWS resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html) in the _IAM User Guide_.IAM

(Not recommended)

Use long-term credentials to sign programmatic requests
to the AWS CLI, AWS SDKs, or AWS APIs.

Following the instructions for the interface that you want to use.

- For the AWS CLI, see [Authenticating using IAM user credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-authentication-user.html) in
the _AWS Command Line Interface User Guide_.

- For AWS SDKs and tools, see [Authenticate using long-term credentials](https://docs.aws.amazon.com/sdkref/latest/guide/access-iam-users.html) in the
_AWS SDKs and Tools Reference Guide_.

- For AWS APIs, see [Managing access keys for\
IAM users](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) in the _IAM User Guide_.

## Step 3: Get ready to use the example code

This guide includes examples that use the AWS SDK for Java. To run the example code,
follow the set-up instructions in [Getting Started with\
AWS SDK for Java 2.0](https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide).

You can develop AWS applications in other programming languages, such as Go,
JavaScript, Python and Ruby. For more information, see [Tools to Build on AWS](https://aws.amazon.com/developer/tools).

###### Note

You can explore Amazon SQS without writing code with tools such as the AWS Command Line Interface
(AWS CLI) or Windows PowerShell. You can find AWS CLI examples in the [Amazon SQS section](https://docs.aws.amazon.com/cli/latest/reference/sqs/index.html) of the
_AWS CLI Command Reference_. You can find Windows PowerShell examples
in the Amazon Simple Queue Service section of the _[AWS Tools for PowerShell Cmdlet Reference](https://docs.aws.amazon.com/powershell/latest/reference)_.

## Next steps

You are now ready for [Getting started](sqs-getting-started.md) with
managing Amazon SQS queues and messages using the AWS Management Console.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Getting started

Understanding the Amazon SQS console
