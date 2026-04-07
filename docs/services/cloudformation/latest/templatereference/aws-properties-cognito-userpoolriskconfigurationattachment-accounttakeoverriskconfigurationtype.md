This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment AccountTakeoverRiskConfigurationType

The settings for automated responses and notification templates for adaptive
authentication with advanced security features.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : AccountTakeoverActionsType,
  "NotifyConfiguration" : NotifyConfigurationType
}

```

### YAML

```yaml

  Actions:
    AccountTakeoverActionsType
  NotifyConfiguration:
    NotifyConfigurationType

```

## Properties

`Actions`

A list of account-takeover actions for each level of risk that Amazon Cognito might assess with
threat protection.

_Required_: Yes

_Type_: [AccountTakeoverActionsType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractionstype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotifyConfiguration`

The settings for composing and sending an email message when threat protection
assesses a risk level with adaptive authentication. When you choose to notify users in
`AccountTakeoverRiskConfiguration`, Amazon Cognito sends an email message using
the method and template that you set with this data type.

_Required_: No

_Type_: [NotifyConfigurationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AccountTakeoverActionType

CompromisedCredentialsActionsType
