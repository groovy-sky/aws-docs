This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool Policies

A list of user pool policies. Contains the policy that sets password-complexity
requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PasswordPolicy" : PasswordPolicy,
  "SignInPolicy" : SignInPolicy
}

```

### YAML

```yaml

  PasswordPolicy:
    PasswordPolicy
  SignInPolicy:
    SignInPolicy

```

## Properties

`PasswordPolicy`

The password policy settings for a user pool, including complexity, history, and
length requirements.

_Required_: No

_Type_: [PasswordPolicy](aws-properties-cognito-userpool-passwordpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SignInPolicy`

The policy for allowed types of authentication in a user pool. To activate this
setting, your user pool must be in the [Essentials tier](../../../cognito/latest/developerguide/feature-plans-features-essentials.md) or higher.

_Required_: No

_Type_: [SignInPolicy](aws-properties-cognito-userpool-signinpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PasswordPolicy

PreTokenGenerationConfig

All content copied from https://docs.aws.amazon.com/.
