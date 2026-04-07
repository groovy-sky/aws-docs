This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction Definition

The definition is where you specify all of the type-specific parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IamActionDefinition" : IamActionDefinition,
  "ScpActionDefinition" : ScpActionDefinition,
  "SsmActionDefinition" : SsmActionDefinition
}

```

### YAML

```yaml

  IamActionDefinition:
    IamActionDefinition
  ScpActionDefinition:
    ScpActionDefinition
  SsmActionDefinition:
    SsmActionDefinition

```

## Properties

`IamActionDefinition`

The AWS Identity and Access Management (IAM) action definition details.

_Required_: No

_Type_: [IamActionDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-iamactiondefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScpActionDefinition`

The service control policies (SCP) action definition details.

_Required_: No

_Type_: [ScpActionDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-scpactiondefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SsmActionDefinition`

The Amazon EC2 Systems Manager (SSM) action definition details.

_Required_: No

_Type_: [SsmActionDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-budgets-budgetsaction-ssmactiondefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ActionThreshold

IamActionDefinition
