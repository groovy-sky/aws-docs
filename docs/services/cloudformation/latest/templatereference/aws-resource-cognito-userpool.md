This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool

The `AWS::Cognito::UserPool` resource creates an Amazon Cognito user pool.
For more information on working with Amazon Cognito user pools, see [Amazon Cognito\
User Pools](../../../cognito/latest/developerguide/cognito-user-identity-pools.md) and [CreateUserPool](../../../../reference/cognito-user-identity-pools/latest/apireference/api-createuserpool.md).

###### Note

If you don't specify a value for a parameter, Amazon Cognito sets it to a default
value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPool",
  "Properties" : {
      "AccountRecoverySetting" : AccountRecoverySetting,
      "AdminCreateUserConfig" : AdminCreateUserConfig,
      "AliasAttributes" : [ String, ... ],
      "AutoVerifiedAttributes" : [ String, ... ],
      "DeletionProtection" : String,
      "DeviceConfiguration" : DeviceConfiguration,
      "EmailAuthenticationMessage" : String,
      "EmailAuthenticationSubject" : String,
      "EmailConfiguration" : EmailConfiguration,
      "EmailVerificationMessage" : String,
      "EmailVerificationSubject" : String,
      "EnabledMfas" : [ String, ... ],
      "LambdaConfig" : LambdaConfig,
      "MfaConfiguration" : String,
      "Policies" : Policies,
      "Schema" : [ SchemaAttribute, ... ],
      "SmsAuthenticationMessage" : String,
      "SmsConfiguration" : SmsConfiguration,
      "SmsVerificationMessage" : String,
      "UserAttributeUpdateSettings" : UserAttributeUpdateSettings,
      "UsernameAttributes" : [ String, ... ],
      "UsernameConfiguration" : UsernameConfiguration,
      "UserPoolAddOns" : UserPoolAddOns,
      "UserPoolName" : String,
      "UserPoolTags" : {Key: Value, ...},
      "UserPoolTier" : String,
      "VerificationMessageTemplate" : VerificationMessageTemplate,
      "WebAuthnRelyingPartyID" : String,
      "WebAuthnUserVerification" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPool
Properties:
  AccountRecoverySetting:
    AccountRecoverySetting
  AdminCreateUserConfig:
    AdminCreateUserConfig
  AliasAttributes:
    - String
  AutoVerifiedAttributes:
    - String
  DeletionProtection: String
  DeviceConfiguration:
    DeviceConfiguration
  EmailAuthenticationMessage: String
  EmailAuthenticationSubject: String
  EmailConfiguration:
    EmailConfiguration
  EmailVerificationMessage: String
  EmailVerificationSubject: String
  EnabledMfas:
    - String
  LambdaConfig:
    LambdaConfig
  MfaConfiguration: String
  Policies:
    Policies
  Schema:
    - SchemaAttribute
  SmsAuthenticationMessage: String
  SmsConfiguration:
    SmsConfiguration
  SmsVerificationMessage: String
  UserAttributeUpdateSettings:
    UserAttributeUpdateSettings
  UsernameAttributes:
    - String
  UsernameConfiguration:
    UsernameConfiguration
  UserPoolAddOns:
    UserPoolAddOns
  UserPoolName: String
  UserPoolTags:
    Key: Value
  UserPoolTier: String
  VerificationMessageTemplate:
    VerificationMessageTemplate
  WebAuthnRelyingPartyID: String
  WebAuthnUserVerification: String

```

## Properties

`AccountRecoverySetting`

The available verified method a user can use to recover their password when they call
`ForgotPassword`. You can use this setting to define a preferred method
when a user has more than one method available. With this setting, SMS doesn't qualify
for a valid password recovery mechanism if the user also has SMS multi-factor
authentication (MFA) activated. In the absence of this setting, Amazon Cognito uses the legacy
behavior to determine the recovery method where SMS is preferred through email.

_Required_: No

_Type_: [AccountRecoverySetting](aws-properties-cognito-userpool-accountrecoverysetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdminCreateUserConfig`

The settings for administrator creation of users in a user pool. Contains settings for
allowing user sign-up, customizing invitation messages to new users, and the amount of
time before temporary passwords expire.

_Required_: No

_Type_: [AdminCreateUserConfig](aws-properties-cognito-userpool-admincreateuserconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AliasAttributes`

Attributes supported as an alias for this user pool. For more information about alias
attributes, see [Customizing sign-in attributes](../../../cognito/latest/developerguide/user-pool-settings-attributes.md#user-pool-settings-aliases).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoVerifiedAttributes`

The attributes that you want your user pool to automatically verify. For more
information, see [Verifying contact information at sign-up](../../../cognito/latest/developerguide/signing-up-users-in-your-app.md#allowing-users-to-sign-up-and-confirm-themselves).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtection`

When active, `DeletionProtection` prevents accidental deletion of your user
pool. Before you can delete a user pool that you have protected against deletion, you
must deactivate this feature.

When you try to delete a protected user pool in a `DeleteUserPool` API request,
Amazon Cognito returns an `InvalidParameterException` error. To delete a protected user pool,
send a new `DeleteUserPool` request after you deactivate deletion protection in an
`UpdateUserPool` API request.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceConfiguration`

The device-remembering configuration for a user pool. Device remembering or device
tracking is a "Remember me on this device" option for user pools that perform
authentication with the device key of a trusted device in the back end, instead of a
user-provided MFA code. For more information about device authentication, see [Working with user devices in your user pool](../../../cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.md). A null value indicates that
you have deactivated device remembering in your user pool.

###### Note

When you provide a value for any `DeviceConfiguration` field, you
activate the Amazon Cognito device-remembering feature. For more information, see [Working with devices](../../../cognito/latest/developerguide/amazon-cognito-user-pools-device-tracking.md).

_Required_: No

_Type_: [DeviceConfiguration](aws-properties-cognito-userpool-deviceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailAuthenticationMessage`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailAuthenticationSubject`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailConfiguration`

The email configuration of your user pool. The email configuration type sets your
preferred sending method, AWS Region, and sender for messages from your user
pool.

_Required_: No

_Type_: [EmailConfiguration](aws-properties-cognito-userpool-emailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailVerificationMessage`

This parameter is no longer used. See [VerificationMessageTemplateType](../userguide/aws-properties-cognito-userpool-verificationmessagetemplate.md).

_Required_: No

_Type_: String

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailVerificationSubject`

This parameter is no longer used. See [VerificationMessageTemplateType](../userguide/aws-properties-cognito-userpool-verificationmessagetemplate.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnabledMfas`

Set enabled MFA options on a specified user pool. To disable all MFAs after it has
been enabled, set `MfaConfiguration` to `OFF` and remove
EnabledMfas. MFAs can only be all disabled if `MfaConfiguration` is
`OFF`. After you enable `SMS_MFA`, you can only disable it by
setting `MfaConfiguration` to `OFF`. Can be one of the following
values:

- `SMS_MFA` \- Enables MFA with SMS for the user pool. To select this
option, you must also provide values for `SmsConfiguration`.

- `SOFTWARE_TOKEN_MFA` \- Enables software token MFA for the user
pool.

- `EMAIL_OTP` \- Enables MFA with email for the user pool. To select
this option, you must provide values for `EmailConfiguration` and
within those, set `EmailSendingAccount` to
`DEVELOPER`.

Allowed values: `SMS_MFA` \| `SOFTWARE_TOKEN_MFA` \|
`EMAIL_OTP`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaConfig`

A collection of user pool Lambda triggers. Amazon Cognito invokes triggers at several possible
stages of authentication operations. Triggers can modify the outcome of the operations
that invoked them.

_Required_: No

_Type_: [LambdaConfig](aws-properties-cognito-userpool-lambdaconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MfaConfiguration`

Displays the state of multi-factor authentication (MFA) as on, off, or optional. When
`ON`, all users must set up MFA before they can sign in. When
`OPTIONAL`, your application must make a client-side determination of
whether a user wants to register an MFA device. For user pools with adaptive
authentication with threat protection, choose `OPTIONAL`.

When `MfaConfiguration` is `OPTIONAL`, managed login
doesn't automatically prompt users to set up MFA. Amazon Cognito generates MFA prompts in
API responses and in managed login for users who have chosen and configured a preferred
MFA factor.

_Required_: No

_Type_: String

_Allowed values_: `OFF | ON | OPTIONAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policies`

A list of user pool policies. Contains the policy that sets password-complexity
requirements.

_Required_: No

_Type_: [Policies](aws-properties-cognito-userpool-policies.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

An array of attributes for the new user pool. You can add custom attributes and modify
the properties of default attributes. The specifications in this parameter set the
required attributes in your user pool. For more information, see [Working with user attributes](../../../cognito/latest/developerguide/user-pool-settings-attributes.md).

_Required_: No

_Type_: Array of [SchemaAttribute](aws-properties-cognito-userpool-schemaattribute.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmsAuthenticationMessage`

The contents of the SMS authentication message.

_Required_: No

_Type_: String

_Minimum_: `6`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmsConfiguration`

The settings for your Amazon Cognito user pool to send SMS messages with Amazon Simple Notification Service. To send SMS
messages with Amazon SNS in the AWS Region that you want, the Amazon Cognito user pool uses an
AWS Identity and Access Management (IAM) role in your AWS account. For more information see
[SMS message settings](../../../cognito/latest/developerguide/user-pool-sms-settings.md).

_Required_: No

_Type_: [SmsConfiguration](aws-properties-cognito-userpool-smsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmsVerificationMessage`

This parameter is no longer used. See [VerificationMessageTemplateType](../userguide/aws-properties-cognito-userpool-verificationmessagetemplate.md).

_Required_: No

_Type_: String

_Minimum_: `6`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAttributeUpdateSettings`

The settings for updates to user attributes. These settings include the property `AttributesRequireVerificationBeforeUpdate`,
a user-pool setting that tells Amazon Cognito how to handle changes to the value of your users' email address and phone number attributes. For
more information, see [Verifying updates to email addresses and phone numbers](../../../cognito/latest/developerguide/user-pool-settings-email-phone-verification.md#user-pool-settings-verifications-verify-attribute-updates).

_Required_: No

_Type_: [UserAttributeUpdateSettings](aws-properties-cognito-userpool-userattributeupdatesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsernameAttributes`

Specifies whether a user can use an email address or phone number as a username when
they sign up.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsernameConfiguration`

Sets the case sensitivity option for sign-in usernames. When
`CaseSensitive` is `false` (case insensitive), users can sign
in with any combination of capital and lowercase letters. For example,
`username`, `USERNAME`, or `UserName`, or for
email, `email@example.com` or `EMaiL@eXamplE.Com`. For most use
cases, set case sensitivity to `false` as a best practice. When usernames and
email addresses are case insensitive, Amazon Cognito treats any variation in case as the same
user, and prevents a case variation from being assigned to the same attribute for a
different user.

When `CaseSensitive` is `true` (case sensitive), Amazon Cognito
interprets `USERNAME` and `UserName` as distinct users.

This configuration is immutable after you set it.

_Required_: No

_Type_: [UsernameConfiguration](aws-properties-cognito-userpool-usernameconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolAddOns`

Contains settings for activation of threat protection, including the operating
mode and additional authentication types. To log user security information but take
no action, set to `AUDIT`. To configure automatic security responses to
potentially unwanted traffic to your user pool, set to `ENFORCED`.

For more information, see [Adding advanced security to a user pool](../../../cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.md). To activate this setting, your user pool must be on the [Plus tier](../../../cognito/latest/developerguide/feature-plans-features-plus.md).

_Required_: No

_Type_: [UserPoolAddOns](aws-properties-cognito-userpool-userpooladdons.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolName`

A friendly name for your user pool.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolTags`

The tag keys and values to assign to the user pool. A tag is a label that you can use
to categorize and manage user pools in different ways, such as by purpose, owner,
environment, or other criteria.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolTier`

The user pool [feature plan](../../../cognito/latest/developerguide/cognito-sign-in-feature-plans.md), or tier. This parameter determines the
eligibility of the user pool for features like managed login, access-token
customization, and threat protection. Defaults to `ESSENTIALS`.

_Required_: No

_Type_: String

_Allowed values_: `LITE | ESSENTIALS | PLUS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerificationMessageTemplate`

The template for the verification message that your user pool delivers to users who
set an email address or phone number attribute.

Set the email message type that corresponds to your `DefaultEmailOption`
selection. For `CONFIRM_WITH_LINK`, specify an
`EmailMessageByLink` and leave `EmailMessage` blank. For
`CONFIRM_WITH_CODE`, specify an `EmailMessage` and leave
`EmailMessageByLink` blank. When you supply both parameters with either
choice, Amazon Cognito returns an error.

_Required_: No

_Type_: [VerificationMessageTemplate](aws-properties-cognito-userpool-verificationmessagetemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebAuthnRelyingPartyID`

Sets or displays the authentication domain, typically your user pool domain, that
passkey providers must use as a relying party (RP) in their configuration.

Under the following conditions, the passkey relying party ID must be the
fully-qualified domain name of your custom domain:

- The user pool is configured for passkey authentication.

- The user pool has a custom domain, whether or not it also has a prefix
domain.

- Your application performs authentication with managed login or the classic
hosted UI.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebAuthnUserVerification`

When `required`, users can only register and sign in users with passkeys
that are capable of [user\
verification](https://www.w3.org/TR/webauthn-2). When `preferred`, your user pool doesn't
require the use of authenticators with user verification but encourages it.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `9`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, such as
`us-east-2_zgaEXAMPLE`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the user pool, such as
`arn:aws:cognito-idp:us-east-1:123412341234:userpool/us-east-1_123412341`.

`ProviderName`

A friendly name for the IdP.

`ProviderURL`

The URL of the provider of the Amazon Cognito user pool, specified as a
`String`.

`UserPoolId`

The ID of the user pool.

## Examples

### Creating a new user pool

The following example creates a user pool with values for all possible
parameters.

#### JSON

```json

{
    "UserPool": {
        "Properties": {
            "AccountRecoverySetting": {
                "RecoveryMechanisms": [
                    {
                        "Name": "verified_email",
                        "Priority": 1
                    },
                    {
                        "Name": "verified_phone_number",
                        "Priority": 2
                    }
                ]
            },
            "AdminCreateUserConfig": {
                "AllowAdminCreateUserOnly": false,
                "InviteMessageTemplate": {
                    "EmailMessage": "Your username is {username} and your temporary password is {####}",
                    "EmailSubject": "Welcome to MyApp",
                    "SMSMessage": "Your username for MyApp is {username} and your temporary password is {####}."
                }
            },
            "AliasAttributes": [
                "email",
                "phone_number",
                "preferred_username"
            ],
            "AutoVerifiedAttributes": [
                "email",
                "phone_number"
            ],
            "DeletionProtection": "ACTIVE",
            "DeviceConfiguration": {
                "ChallengeRequiredOnNewDevice": true,
                "DeviceOnlyRememberedOnUserPrompt": true
            },
            "EmailAuthenticationMessage": "This is your sign-in code: \"{####}\"",
            "EmailAuthenticationSubject": "Your sign-in code",
            "EmailConfiguration": {
                "ConfigurationSet": "my-test-configuration-set",
                "EmailSendingAccount": "DEVELOPER",
                "From": "admin@example.com",
                "ReplyToEmailAddress": "no-reply@example.com",
                "SourceArn": "arn:aws:ses:us-west-2:123456789012:identity/admin@example.com"
            },
            "EmailVerificationMessage": "This is your verification code: \"{####}\"",
            "EmailVerificationSubject": "Your verification code",
            "EnabledMfas": [
                "EMAIL_OTP",
                "SMS_MFA",
                "SOFTWARE_TOKEN_MFA"
            ],
            "LambdaConfig": {
                "CreateAuthChallenge": "arn:aws:lambda:us-west-2:123456789012:function:my-CreateAuthChallenge-function",
                "CustomEmailSender": {
                    "LambdaArn": "arn:aws:lambda:us-west-2:123456789012:function:my-CustomEmailSender-function",
                    "LambdaVersion": "V1_0"
                },
                "CustomMessage": "arn:aws:lambda:us-west-2:123456789012:function:my-CustomMessage-function",
                "CustomSMSSender": {
                    "LambdaArn": "arn:aws:lambda:us-west-2:123456789012:function:my-CustomSMSSender-function",
                    "LambdaVersion": "V1_0"
                },
                "DefineAuthChallenge": "arn:aws:lambda:us-west-2:123456789012:function:my-DefineAuthChallenge-function",
                "KMSKeyID": "arn:aws:kms:us-west-2:123456789012:key/4d43904c-8edf-4bb4-9fca-fb1a80e41cbe",
                "PostAuthentication": "arn:aws:lambda:us-west-2:123456789012:function:my-PostAuthentication-function",
                "PostConfirmation": "arn:aws:lambda:us-west-2:123456789012:function:my-PostConfirmation-function",
                "PreAuthentication": "arn:aws:lambda:us-west-2:123456789012:function:my-PreAuthentication-function",
                "PreSignUp": "arn:aws:lambda:us-west-2:123456789012:function:my-PreSignUp-function",
                "PreTokenGenerationConfig": {
                    "LambdaArn": "arn:aws:lambda:us-west-2:123456789012:function:my-PreTokenGenerationConfig-function",
                    "LambdaVersion": "V2_0"
                },
                "UserMigration": "arn:aws:lambda:us-west-2:123456789012:function:my-UserMigration-function",
                "VerifyAuthChallengeResponse": "arn:aws:lambda:us-west-2:123456789012:function:my-VerifyAuthChallengeResponse-function"
            },
            "MfaConfiguration": "OPTIONAL",
            "Policies": {
                "PasswordPolicy": {
                    "MinimumLength": 12,
                    "PasswordHistorySize": 22,
                    "RequireLowercase": true,
                    "RequireNumbers": true,
                    "RequireSymbols": true,
                    "RequireUppercase": true,
                    "TemporaryPasswordValidityDays": 7
                },
                "SignInPolicy": {
                    "AllowedFirstAuthFactors": [
                        "EMAIL_OTP",
                        "SMS_OTP",
                        "WEB_AUTHN",
                        "PASSWORD"
                    ]
                }
            },
            "Schema": [
                {
                    "AttributeDataType": "String",
                    "DeveloperOnlyAttribute": false,
                    "Mutable": true,
                    "Name": "ResidenceType",
                    "Required": false,
                    "StringAttributeConstraints": {
                        "MaxLength": "999",
                        "MinLength": "1"
                    }
                },
                {
                    "AttributeDataType": "Number",
                    "DeveloperOnlyAttribute": true,
                    "Mutable": true,
                    "Name": "NumberOfResidents",
                    "NumberAttributeConstraints": {
                        "MaxValue": "9999999",
                        "MinValue": "1"
                    },
                    "Required": false
                },
                {
                    "AttributeDataType": "String",
                    "DeveloperOnlyAttribute": false,
                    "Mutable": true,
                    "Name": "email",
                    "Required": true,
                    "StringAttributeConstraints": {
                        "MaxLength": "99",
                        "MinLength": "1"
                    }
                }
            ],
            "SmsAuthenticationMessage": "This is your sign-in code: \"{####}\"",
            "SmsConfiguration": {
                "ExternalId": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
                "SnsCallerArn": "arn:aws:iam::123456789012:role/service-role/my-test-SMS-Role",
                "SnsRegion": "us-west-2"
            },
            "SmsVerificationMessage": "This is your verification code: \"{####}\"",
            "UserAttributeUpdateSettings": {
                "AttributesRequireVerificationBeforeUpdate": [
                    "email",
                    "phone_number"
                ]
            },
            "UsernameConfiguration": {
                "CaseSensitive": true
            },
            "UserPoolAddOns": {
                "AdvancedSecurityAdditionalFlows": {
                    "CustomAuthMode": "AUDIT"
                },
                "AdvancedSecurityMode": "AUDIT"
            },
            "UserPoolName": "Example_CloudFormation_UserPool",
            "UserPoolTags": {
                "administrator": "Jie",
                "tenant": "ExampleCorp"
            },
            "UserPoolTier": "PLUS",
            "VerificationMessageTemplate": {
                "DefaultEmailOption": "CONFIRM_WITH_CODE",
                "EmailMessage": "This is your verification code: \"{####}\"",
                "EmailMessageByLink": "Choose this link to {##verify your email##}",
                "EmailSubject": "Your confirmation code",
                "EmailSubjectByLink": "Your confirmation link",
                "SmsMessage": "This is your verification code: \"{####}\""
            },
            "WebAuthnRelyingPartyID": "auth.example.com",
            "WebAuthnUserVerification": "preferred"
        },
        "Type": "AWS::Cognito::UserPool"
    }
}
```

#### YAML

```yaml

ExampleUserPool:
    Type: AWS::Cognito::UserPool
    Properties:
      AccountRecoverySetting:
        RecoveryMechanisms:
          - Name: verified_email
            Priority: 1
          - Name: verified_phone_number
            Priority: 2
      AdminCreateUserConfig:
        AllowAdminCreateUserOnly: false
        InviteMessageTemplate:
          EmailMessage: Your username is {username} and your temporary password is {####}
          EmailSubject: Welcome to MyApp
          SMSMessage: Your username for MyApp is {username} and your temporary password is {####}.
      AliasAttributes:
        - email
        - phone_number
        - preferred_username
      AutoVerifiedAttributes:
        - email
        - phone_number
      DeletionProtection: ACTIVE
      DeviceConfiguration:
        ChallengeRequiredOnNewDevice: true
        DeviceOnlyRememberedOnUserPrompt: true
      EmailAuthenticationMessage: 'This is your sign-in code: "{####}"'
      EmailAuthenticationSubject: Your sign-in code
      EmailConfiguration:
        ConfigurationSet: my-test-configuration-set
        EmailSendingAccount: DEVELOPER
        From: admin@example.com
        ReplyToEmailAddress: no-reply@example.com
        SourceArn: arn:aws:ses:us-west-2:123456789012:identity/admin@example.com
      EmailVerificationMessage: 'This is your verification code: "{####}"'
      EmailVerificationSubject: Your verification code
      EnabledMfas:
        - EMAIL_OTP
        - SMS_MFA
        - SOFTWARE_TOKEN_MFA
      LambdaConfig:
        CreateAuthChallenge: arn:aws:lambda:us-west-2:123456789012:function:my-CreateAuthChallenge-function
        CustomEmailSender:
          LambdaArn: arn:aws:lambda:us-west-2:123456789012:function:my-CustomEmailSender-function
          LambdaVersion: V1_0
        CustomMessage: arn:aws:lambda:us-west-2:123456789012:function:my-CustomMessage-function
        CustomSMSSender:
          LambdaArn: arn:aws:lambda:us-west-2:123456789012:function:my-CustomSMSSender-function
          LambdaVersion: V1_0
        DefineAuthChallenge: arn:aws:lambda:us-west-2:123456789012:function:my-DefineAuthChallenge-function
        KMSKeyID: arn:aws:kms:us-west-2:123456789012:key/a1b2c3d4-5678-90ab-cdef-EXAMPLE22222
        PostAuthentication: arn:aws:lambda:us-west-2:123456789012:function:my-PostAuthentication-function
        PostConfirmation: arn:aws:lambda:us-west-2:123456789012:function:my-PostConfirmation-function
        PreAuthentication: arn:aws:lambda:us-west-2:123456789012:function:my-PreAuthentication-function
        PreSignUp: arn:aws:lambda:us-west-2:123456789012:function:my-PreSignUp-function
        PreTokenGenerationConfig:
          LambdaArn: arn:aws:lambda:us-west-2:123456789012:function:my-PreTokenGeneration-function
          LambdaVersion: V2_0
        UserMigration: arn:aws:lambda:us-west-2:123456789012:function:my-UserMigration-function
        VerifyAuthChallengeResponse: arn:aws:lambda:us-west-2:123456789012:function:my-VerifyAuthChallengeResponse-function
      MfaConfiguration: OPTIONAL
      Policies:
        PasswordPolicy:
          MinimumLength: 12
          PasswordHistorySize: 22
          RequireLowercase: true
          RequireNumbers: true
          RequireSymbols: true
          RequireUppercase: true
          TemporaryPasswordValidityDays: 7
        SignInPolicy:
          AllowedFirstAuthFactors:
            - EMAIL_OTP
            - SMS_OTP
            - WEB_AUTHN
            - PASSWORD
      Schema:
        - AttributeDataType: String
          DeveloperOnlyAttribute: false
          Mutable: true
          Name: ResidenceType
          Required: false
          StringAttributeConstraints:
            MaxLength: "999"
            MinLength: "1"
        - AttributeDataType: Number
          DeveloperOnlyAttribute: true
          Mutable: true
          Name: NumberOfResidents
          NumberAttributeConstraints:
            MaxValue: "9999999"
            MinValue: "1"
          Required: false
        - AttributeDataType: String
          DeveloperOnlyAttribute: false
          Mutable: true
          Name: email
          Required: true
          StringAttributeConstraints:
            MaxLength: "99"
            MinLength: "1"
      SmsAuthenticationMessage: 'This is your sign-in code: "{####}"'
      SmsConfiguration:
        ExternalId: a1b2c3d4-5678-90ab-cdef-EXAMPLE11111
        SnsCallerArn: arn:aws:iam::123456789012:role/service-role/my-test-SMS-Role
        SnsRegion: us-west-2
      SmsVerificationMessage: 'This is your verification code: "{####}"'
      UserAttributeUpdateSettings:
        AttributesRequireVerificationBeforeUpdate:
          - email
          - phone_number
      UsernameConfiguration:
        CaseSensitive: true
      UserPoolAddOns:
        AdvancedSecurityAdditionalFlows:
          CustomAuthMode: AUDIT
        AdvancedSecurityMode: AUDIT
      UserPoolName: Example_CloudFormation_UserPool
      UserPoolTags:
        administrator: Jie
        tenant: ExampleCorp
      UserPoolTier: PLUS
      VerificationMessageTemplate:
        DefaultEmailOption: CONFIRM_WITH_CODE
        EmailMessage: 'This is your verification code: "{####}"'
        EmailMessageByLink: Choose this link to {##verify your email##}
        EmailSubject: Your confirmation code
        EmailSubjectByLink: Your confirmation link
        SmsMessage: 'This is your verification code: "{####}"'
      WebAuthnRelyingPartyID: auth.example.com
      WebAuthnUserVerification: preferred
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::Terms

AccountRecoverySetting

All content copied from https://docs.aws.amazon.com/.
