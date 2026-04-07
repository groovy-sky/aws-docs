This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool LambdaConfig

A collection of user pool Lambda triggers. Amazon Cognito invokes triggers at several
possible stages of user pool operations. Triggers can modify the outcome of the
operations that invoked them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreateAuthChallenge" : String,
  "CustomEmailSender" : CustomEmailSender,
  "CustomMessage" : String,
  "CustomSMSSender" : CustomSMSSender,
  "DefineAuthChallenge" : String,
  "InboundFederation" : InboundFederation,
  "KMSKeyID" : String,
  "PostAuthentication" : String,
  "PostConfirmation" : String,
  "PreAuthentication" : String,
  "PreSignUp" : String,
  "PreTokenGeneration" : String,
  "PreTokenGenerationConfig" : PreTokenGenerationConfig,
  "UserMigration" : String,
  "VerifyAuthChallengeResponse" : String
}

```

### YAML

```yaml

  CreateAuthChallenge: String
  CustomEmailSender:
    CustomEmailSender
  CustomMessage: String
  CustomSMSSender:
    CustomSMSSender
  DefineAuthChallenge: String
  InboundFederation:
    InboundFederation
  KMSKeyID: String
  PostAuthentication: String
  PostConfirmation: String
  PreAuthentication: String
  PreSignUp: String
  PreTokenGeneration: String
  PreTokenGenerationConfig:
    PreTokenGenerationConfig
  UserMigration: String
  VerifyAuthChallengeResponse: String

```

## Properties

`CreateAuthChallenge`

The configuration of a create auth challenge Lambda trigger, one of three triggers in
the sequence of the [custom authentication challenge triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html).

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomEmailSender`

The configuration of a custom email sender Lambda trigger. This trigger routes all
email notifications from a user pool to a Lambda function that delivers the message using
custom logic.

_Required_: No

_Type_: [CustomEmailSender](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-customemailsender.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomMessage`

A custom message Lambda trigger. This trigger is an opportunity to customize all SMS
and email messages from your user pool. When a custom message trigger is active, your
user pool routes all messages to a Lambda function that returns a runtime-customized
message subject and body for your user pool to deliver to a user.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomSMSSender`

The configuration of a custom SMS sender Lambda trigger. This trigger routes all SMS
notifications from a user pool to a Lambda function that delivers the message using
custom logic.

_Required_: No

_Type_: [CustomSMSSender](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-customsmssender.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefineAuthChallenge`

The configuration of a define auth challenge Lambda trigger, one of three triggers in
the sequence of the [custom authentication challenge triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html).

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InboundFederation`

Property description not available.

_Required_: No

_Type_: [InboundFederation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-inboundfederation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KMSKeyID`

The ARN of an [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys). Amazon Cognito uses the key to encrypt codes and temporary passwords sent to
custom sender Lambda triggers.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostAuthentication`

The configuration of a [post authentication Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-authentication.html) in a user pool. This
trigger can take custom actions after a user signs in.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostConfirmation`

The configuration of a [post confirmation Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-post-confirmation.html) in a user pool. This
trigger can take custom actions after a user confirms their user account and their email
address or phone number.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreAuthentication`

The configuration of a [pre authentication trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-authentication.html) in a user pool. This trigger
can evaluate and modify user sign-in events.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreSignUp`

The configuration of a [pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html) in a user pool. This trigger
evaluates new users and can bypass confirmation, [link a federated user profile](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-identity-federation-consolidate-users.html), or block sign-up
requests.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreTokenGeneration`

The legacy configuration of a [pre token generation Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html) in a user
pool.

Set this parameter for legacy purposes. If you also set an ARN in
`PreTokenGenerationConfig`, its value must be identical to
`PreTokenGeneration`. For new instances of pre token generation triggers,
set the `LambdaArn` of `PreTokenGenerationConfig`.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreTokenGenerationConfig`

The detailed configuration of a [pre token generation Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html) in a user pool. If
you also set an ARN in `PreTokenGeneration`, its value must be identical to
`PreTokenGenerationConfig`.

_Required_: No

_Type_: [PreTokenGenerationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-pretokengenerationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserMigration`

The configuration of a [migrate user Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-migrate-user.html) in a user pool. This trigger
can create user profiles when users sign in or attempt to reset their password with
credentials that don't exist yet.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifyAuthChallengeResponse`

The configuration of a verify auth challenge Lambda trigger, one of three triggers in
the sequence of the [custom authentication challenge triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-challenge.html).

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InviteMessageTemplate

NumberAttributeConstraints
