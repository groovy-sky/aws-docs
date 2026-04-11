# Setting up DynamoDB (web service)

To use the Amazon DynamoDB web service:

1. [Sign up for\
    AWS.](#SettingUp.DynamoWebService.SignUpForAWS)

2. [Get an AWS access\
    key](#SettingUp.DynamoWebService.GetCredentials) (used to access DynamoDB programmatically).

###### Note

If you plan to interact with DynamoDB only through the AWS Management Console, you don't
need an AWS access key, and you can skip ahead to [Using the console](accessingdynamodb.md#ConsoleDynamoDB).

3. [Configure your\
    credentials](#SettingUp.DynamoWebService.ConfigureCredentials) (used to access DynamoDB programmatically).

## Signing up for AWS

To use the DynamoDB service, you must have an AWS account. If you don't already
have an account, you are prompted to create one when you sign up. You're not charged
for any AWS services that you sign up for unless you use them.

###### To sign up for AWS

1. Open [https://portal.aws.amazon.com/billing/signup](https://portal.aws.amazon.com/billing/signup).

2. Follow the online instructions.

Part of the sign-up procedure involves receiving a phone call or text message and entering
    a verification code on the phone keypad.

When you sign up for an AWS account, an _AWS account root user_ is created. The root user has access to all AWS services
    and resources in the account. As a security best practice, assign administrative access to a user, and use only the root user to perform [tasks that require root user access](../../../iam/latest/userguide/id-root-user.md#root-user-tasks).

## Granting programmatic access

Before you can access DynamoDB programmatically or through the AWS Command Line Interface (AWS CLI),
you must have programmatic access. You don't need programmatic access if you plan to
use the DynamoDB console only.

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

## Configuring your credentials

Before you can access DynamoDB programmatically or through the AWS CLI, you must
configure your credentials to enable authorization for your applications.

There are several ways to do this. For example, you can manually create the
credentials file to store your access key ID and secret access key. You also can use
the AWS CLI command `aws configure` to automatically create the file.
Alternatively, you can use environment variables. For more information about
configuring your credentials, see the programming-specific AWS SDK developer
guide.

To install and configure the AWS CLI, see [Using the AWS CLI](accessingdynamodb.md#Tools.CLI).

## Integrating with other DynamoDB services

You can integrate DynamoDB with many other AWS services. For more information, see
the following:

- [Using DynamoDB with other AWS services](otherservices.md)

- [CloudFormation for DynamoDB](../../../cloudformation/latest/userguide/aws-resource-dynamodb-table.md)

- [Using AWS Backup with DynamoDB](backuprestore-howitworksaws.md)

- [AWS Identity and Access Management (IAM) and DynamoDB](identity-and-access-mgmt.md)

- [Using\
AWS Lambda with Amazon DynamoDB](../../../lambda/latest/dg/with-ddb.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting up DynamoDB

Setting up DynamoDB local (downloadable version)

All content copied from https://docs.aws.amazon.com/.
