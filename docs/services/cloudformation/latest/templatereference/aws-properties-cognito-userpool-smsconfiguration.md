This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool SmsConfiguration

User pool configuration for delivery of SMS messages with Amazon Simple Notification
Service. To send SMS messages with Amazon SNS in the AWS Region that you
want, the Amazon Cognito user pool uses an AWS Identity and Access Management (IAM) role in your
AWS account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExternalId" : String,
  "SnsCallerArn" : String,
  "SnsRegion" : String
}

```

### YAML

```yaml

  ExternalId: String
  SnsCallerArn: String
  SnsRegion: String

```

## Properties

`ExternalId`

The external ID provides additional security for your IAM role. You can use an
`ExternalId` with the IAM role that you use with Amazon SNS to send SMS
messages for your user pool. If you provide an `ExternalId`, your Amazon Cognito user
pool includes it in the request to assume your IAM role. You can configure the role
trust policy to require that Amazon Cognito, and any principal, provide the
`ExternalID`. If you use the Amazon Cognito Management Console to create a role
for SMS multi-factor authentication (MFA), Amazon Cognito creates a role with the required
permissions and a trust policy that demonstrates use of the
`ExternalId`.

For more information about the `ExternalId` of a role, see [How to use an\
external ID when granting access to your AWS resources to a third\
party](../../../iam/latest/userguide/id-roles-create-for-user-externalid.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsCallerArn`

The Amazon Resource Name (ARN) of the Amazon SNS caller. This is the ARN of the IAM role
in your AWS account that Amazon Cognito will use to send SMS messages. SMS
messages are subject to a [spending limit](../../../cognito/latest/developerguide/user-pool-settings-email-phone-verification.md).

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsRegion`

The AWS Region to use with Amazon SNS integration. You can choose the same Region as your
user pool, or a supported **Legacy Amazon SNS alternate**
**Region**.

Amazon Cognito resources in the Asia Pacific (Seoul) AWS Region must use your Amazon SNS
configuration in the Asia Pacific (Tokyo) Region. For more information, see [SMS message settings for Amazon Cognito user pools](../../../cognito/latest/developerguide/user-pool-sms-settings.md).

_Required_: No

_Type_: String

_Minimum_: `5`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignInPolicy

StringAttributeConstraints

All content copied from https://docs.aws.amazon.com/.
