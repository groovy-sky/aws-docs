This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment CompromisedCredentialsActionsType

Settings for user pool actions when Amazon Cognito detects compromised credentials
with advanced security features in full-function `ENFORCED` mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventAction" : String
}

```

### YAML

```yaml

  EventAction: String

```

## Properties

`EventAction`

The action that Amazon Cognito takes when it detects compromised credentials.

_Required_: Yes

_Type_: String

_Allowed values_: `BLOCK | NO_ACTION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountTakeoverRiskConfigurationType

CompromisedCredentialsRiskConfigurationType

All content copied from https://docs.aws.amazon.com/.
