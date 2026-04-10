This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool AdvancedSecurityAdditionalFlows

Threat protection configuration options for additional authentication types in your
user pool, including custom
authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomAuthMode" : String
}

```

### YAML

```yaml

  CustomAuthMode: String

```

## Properties

`CustomAuthMode`

The operating mode of threat protection in custom authentication with [Custom\
authentication challenge Lambda triggers](../../../cognito/latest/developerguide/user-pool-lambda-challenge.md).

_Required_: No

_Type_: String

_Allowed values_: `AUDIT | ENFORCED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdminCreateUserConfig

CustomEmailSender

All content copied from https://docs.aws.amazon.com/.
