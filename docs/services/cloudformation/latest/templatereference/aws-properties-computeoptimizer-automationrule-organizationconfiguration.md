This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule OrganizationConfiguration

Configuration settings for organization-wide automation rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountIds" : [ String, ... ],
  "RuleApplyOrder" : String
}

```

### YAML

```yaml

  AccountIds:
    - String
  RuleApplyOrder: String

```

## Properties

`AccountIds`

List of specific AWS account IDs where the organization rule should be applied.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleApplyOrder`

Specifies when organization rules should be applied relative to account rules.

_Required_: No

_Type_: String

_Allowed values_: `BeforeAccountRules | AfterAccountRules`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IntegerCriteriaCondition

ResourceTagsCriteriaCondition
