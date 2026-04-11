This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy PolicyDetail

The configuration details for a lifecycle policy resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : Action,
  "ExclusionRules" : ExclusionRules,
  "Filter" : Filter
}

```

### YAML

```yaml

  Action:
    Action
  ExclusionRules:
    ExclusionRules
  Filter:
    Filter

```

## Properties

`Action`

Configuration details for the policy action.

_Required_: Yes

_Type_: [Action](aws-properties-imagebuilder-lifecyclepolicy-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionRules`

Additional rules to specify resources that should be exempt from policy actions.

_Required_: No

_Type_: [ExclusionRules](aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Specifies the resources that the lifecycle policy applies to.

_Required_: Yes

_Type_: [Filter](aws-properties-imagebuilder-lifecyclepolicy-filter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LastLaunched

RecipeSelection

All content copied from https://docs.aws.amazon.com/.
