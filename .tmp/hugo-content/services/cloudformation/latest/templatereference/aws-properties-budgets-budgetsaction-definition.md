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

_Type_: [IamActionDefinition](aws-properties-budgets-budgetsaction-iamactiondefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScpActionDefinition`

The service control policies (SCP) action definition details.

_Required_: No

_Type_: [ScpActionDefinition](aws-properties-budgets-budgetsaction-scpactiondefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SsmActionDefinition`

The Amazon EC2 Systems Manager (SSM) action definition details.

_Required_: No

_Type_: [SsmActionDefinition](aws-properties-budgets-budgetsaction-ssmactiondefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionThreshold

IamActionDefinition

All content copied from https://docs.aws.amazon.com/.
