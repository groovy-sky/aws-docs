This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::PatchBaseline RuleGroup

The `RuleGroup` property type specifies a set of rules that define the
approval rules for an AWS Systems Manager patch baseline.

`RuleGroup` is the property type for the `ApprovalRules`
property of the [AWS::SSM::PatchBaseline](../userguide/aws-resource-ssm-patchbaseline.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PatchRules" : [ Rule, ... ]
}

```

### YAML

```yaml

  PatchRules:
    - Rule

```

## Properties

`PatchRules`

The rules that make up the rule group.

_Required_: No

_Type_: Array of [Rule](aws-properties-ssm-patchbaseline-rule.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

Tag

All content copied from https://docs.aws.amazon.com/.
