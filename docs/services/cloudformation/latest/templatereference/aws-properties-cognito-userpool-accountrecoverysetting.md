This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool AccountRecoverySetting

The available verified method a user can use to recover their password when they call
`ForgotPassword`. You can use this setting to define a preferred method
when a user has more than one method available. With this setting, SMS doesn't qualify
for a valid password recovery mechanism if the user also has SMS multi-factor
authentication (MFA) activated. In the absence of this setting, Amazon Cognito uses the legacy
behavior to determine the recovery method where SMS is preferred through email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecoveryMechanisms" : [ RecoveryOption, ... ]
}

```

### YAML

```yaml

  RecoveryMechanisms:
    - RecoveryOption

```

## Properties

`RecoveryMechanisms`

The list of options and priorities for user message delivery in forgot-password
operations. Sets or displays user pool preferences for email or SMS message priority,
whether users should fall back to a second delivery method, and whether passwords should
only be reset by administrators.

_Required_: No

_Type_: Array of [RecoveryOption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-recoveryoption.html)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Cognito::UserPool

AdminCreateUserConfig
