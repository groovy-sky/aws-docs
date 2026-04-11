This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction SsmActionDefinition

The Amazon EC2 Systems Manager (SSM) action definition details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceIds" : [ String, ... ],
  "Region" : String,
  "Subtype" : String
}

```

### YAML

```yaml

  InstanceIds:
    - String
  Region: String
  Subtype: String

```

## Properties

`InstanceIds`

The EC2 and RDS instance IDs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The Region to run the (SSM) document.

_Required_: Yes

_Type_: String

_Pattern_: `^\w{2,4}-\w+(-\w+)?-\d$`

_Minimum_: `9`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtype`

The action subType.

_Required_: Yes

_Type_: String

_Allowed values_: `STOP_EC2_INSTANCES | STOP_RDS_INSTANCES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScpActionDefinition

Subscriber

All content copied from https://docs.aws.amazon.com/.
