This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment CompromisedCredentialsRiskConfigurationType

Settings for compromised-credentials actions and authentication-event sources with
advanced security features in full-function `ENFORCED` mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : CompromisedCredentialsActionsType,
  "EventFilter" : [ String, ... ]
}

```

### YAML

```yaml

  Actions:
    CompromisedCredentialsActionsType
  EventFilter:
    - String

```

## Properties

`Actions`

Settings for the actions that you want your user pool to take when Amazon Cognito detects
compromised credentials.

_Required_: Yes

_Type_: [CompromisedCredentialsActionsType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsactionstype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventFilter`

Settings for the sign-in activity where you want to configure compromised-credentials
actions. Defaults to all events.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CompromisedCredentialsActionsType

NotifyConfigurationType
