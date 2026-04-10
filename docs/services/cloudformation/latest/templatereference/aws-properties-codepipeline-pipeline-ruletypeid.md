This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline RuleTypeId

The ID for the rule type, which is made up of the combined values for category, owner,
provider, and version. For more information about conditions, see [Stage\
conditions](../../../codepipeline/latest/userguide/stage-conditions.md). For more information about rules, see the [AWS CodePipeline rule reference](../../../codepipeline/latest/userguide/rule-reference.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : String,
  "Owner" : String,
  "Provider" : String,
  "Version" : String
}

```

### YAML

```yaml

  Category: String
  Owner: String
  Provider: String
  Version: String

```

## Properties

`Category`

A category defines what kind of rule can be run in the stage, and constrains the
provider type for the rule. The valid category is `Rule`.

_Required_: No

_Type_: String

_Allowed values_: `Rule`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Owner`

The creator of the rule being called. The valid value for the `Owner` field
in the rule category is `AWS`.

_Required_: No

_Type_: String

_Allowed values_: `AWS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Provider`

The rule provider, such as the `DeploymentWindow` rule. For a list of rule
provider names, see the rules listed in the [AWS CodePipeline rule\
reference](../../../codepipeline/latest/userguide/rule-reference.md).

_Required_: No

_Type_: String

_Pattern_: `[0-9A-Za-z_-]+`

_Minimum_: `1`

_Maximum_: `35`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

A string that describes the rule version.

_Required_: No

_Type_: String

_Pattern_: `[0-9A-Za-z_-]+`

_Minimum_: `1`

_Maximum_: `9`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleDeclaration

StageDeclaration

All content copied from https://docs.aws.amazon.com/.
