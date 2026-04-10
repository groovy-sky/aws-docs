This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment AccountTakeoverActionsType

A list of account-takeover actions for each level of risk that Amazon Cognito might
assess with advanced security features.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HighAction" : AccountTakeoverActionType,
  "LowAction" : AccountTakeoverActionType,
  "MediumAction" : AccountTakeoverActionType
}

```

### YAML

```yaml

  HighAction:
    AccountTakeoverActionType
  LowAction:
    AccountTakeoverActionType
  MediumAction:
    AccountTakeoverActionType

```

## Properties

`HighAction`

The action that you assign to a high-risk assessment by threat protection.

_Required_: No

_Type_: [AccountTakeoverActionType](aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LowAction`

The action that you assign to a low-risk assessment by threat protection.

_Required_: No

_Type_: [AccountTakeoverActionType](aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediumAction`

The action that you assign to a medium-risk assessment by threat protection.

_Required_: No

_Type_: [AccountTakeoverActionType](aws-properties-cognito-userpoolriskconfigurationattachment-accounttakeoveractiontype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::UserPoolRiskConfigurationAttachment

AccountTakeoverActionType

All content copied from https://docs.aws.amazon.com/.
