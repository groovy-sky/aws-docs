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

_Type_: [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-lifecyclepolicy-action.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionRules`

Additional rules to specify resources that should be exempt from policy actions.

_Required_: No

_Type_: [ExclusionRules](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Specifies the resources that the lifecycle policy applies to.

_Required_: Yes

_Type_: [Filter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-lifecyclepolicy-filter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LastLaunched

RecipeSelection
