This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool UserPoolAddOns

User pool add-ons. Contains settings for activation of threat protection. To log user
security information but take no action, set to `AUDIT`. To configure
automatic security responses to risky traffic to your user pool, set to
`ENFORCED`.

For more information, see [Adding advanced security to a user pool](../../../cognito/latest/developerguide/cognito-user-pool-settings-advanced-security.md). To activate this setting, your
user pool must be on the [Plus\
tier](../../../cognito/latest/developerguide/feature-plans-features-plus.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvancedSecurityAdditionalFlows" : AdvancedSecurityAdditionalFlows,
  "AdvancedSecurityMode" : String
}

```

### YAML

```yaml

  AdvancedSecurityAdditionalFlows:
    AdvancedSecurityAdditionalFlows
  AdvancedSecurityMode: String

```

## Properties

`AdvancedSecurityAdditionalFlows`

Threat protection configuration options for additional authentication types in your
user pool, including custom
authentication.

_Required_: No

_Type_: [AdvancedSecurityAdditionalFlows](aws-properties-cognito-userpool-advancedsecurityadditionalflows.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedSecurityMode`

The operating mode of threat protection for standard authentication types in your user
pool, including username-password and secure remote password (SRP) authentication.

_Required_: No

_Type_: String

_Allowed values_: `OFF | AUDIT | ENFORCED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UsernameConfiguration

VerificationMessageTemplate

All content copied from https://docs.aws.amazon.com/.
