This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction ScpActionDefinition

The service control policies (SCP) action definition details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PolicyId" : String,
  "TargetIds" : [ String, ... ]
}

```

### YAML

```yaml

  PolicyId: String
  TargetIds:
    - String

```

## Properties

`PolicyId`

The policy ID attached.

_Required_: Yes

_Type_: String

_Pattern_: `^p-[0-9a-zA-Z_]{8,128}$`

_Minimum_: `10`

_Maximum_: `130`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIds`

A list of target IDs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTag

SsmActionDefinition

All content copied from https://docs.aws.amazon.com/.
