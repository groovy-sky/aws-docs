This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2

Creates a V2 automation rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::AutomationRuleV2",
  "Properties" : {
      "Actions" : [ AutomationRulesActionV2, ... ],
      "Criteria" : Criteria,
      "Description" : String,
      "RuleName" : String,
      "RuleOrder" : Number,
      "RuleStatus" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::AutomationRuleV2
Properties:
  Actions:
    - AutomationRulesActionV2
  Criteria:
    Criteria
  Description: String
  RuleName: String
  RuleOrder: Number
  RuleStatus: String
  Tags:
    Key: Value

```

## Properties

`Actions`

A list of actions to be performed when the rule criteria is met.

_Required_: Yes

_Type_: Array of [AutomationRulesActionV2](aws-properties-securityhub-automationrulev2-automationrulesactionv2.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Criteria`

The filtering type and configuration of the automation rule.

_Required_: Yes

_Type_: [Criteria](aws-properties-securityhub-automationrulev2-criteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the V2 automation rule.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the V2 automation rule.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleOrder`

The value for the rule priority.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleStatus`

The status of the V2 automation rule.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs associated with the V2 automation rule.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `AutomationRuleV2Arn` for the `AutomationRuleV2Arn` resource created: `arn:aws:securityhub:region:123456789012:automationrulev2/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the V2 automation rule was created.

`RuleArn`

The ARN of the V2 automation rule.

`RuleId`

The ID of the V2 automation rule.

`UpdatedAt`

The timestamp when the V2 automation rule was updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowUpdate

AutomationRulesActionV2

All content copied from https://docs.aws.amazon.com/.
