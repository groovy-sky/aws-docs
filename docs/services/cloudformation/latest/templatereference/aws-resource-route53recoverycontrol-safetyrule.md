This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::SafetyRule

Creates a safety rule in a control panel in Amazon Route 53 Application Recovery Controller. Safety rules in Amazon Route 53
Application Recovery Controller let you add safeguards
around changing routing control states, and enabling and disabling routing controls,
to help prevent unwanted outcomes. Note that the name of a safety rule must be unique within a control panel.

There are two types of safety rules in Route 53 ARC: assertion rules and gating rules.

Assertion rule: An assertion rule enforces that, when you change a routing control state,
certain criteria are met. For example, the criteria might be that at least one routing
control state is `On` after the transaction completes so that traffic continues to be directed to at
least one cell for the application. This prevents a fail-open scenario.

Gating rule: A gating rule lets you configure a gating routing control as an overall on-off switch for a
group of routing controls. Or, you can configure more complex gating scenarios, for example, by
configuring multiple gating routing controls.

For more information, see [Safety rules](../../../r53recovery/latest/dg/routing-control-safety-rules.md)
in the Amazon Route 53 Application Recovery Controller Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryControl::SafetyRule",
  "Properties" : {
      "AssertionRule" : AssertionRule,
      "ControlPanelArn" : String,
      "GatingRule" : GatingRule,
      "Name" : String,
      "RuleConfig" : RuleConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryControl::SafetyRule
Properties:
  AssertionRule:
    AssertionRule
  ControlPanelArn: String
  GatingRule:
    GatingRule
  Name: String
  RuleConfig:
    RuleConfig
  Tags:
    - Tag

```

## Properties

`AssertionRule`

An assertion rule enforces that, when you change a routing control state, that the criteria that you set in the rule configuration is met.
Otherwise, the change to the routing control is not accepted. For example, the criteria might be that at least one routing
control state is `On` after the transaction so that traffic continues to flow to at least one cell for the application.
This ensures that you avoid a fail-open scenario.

_Required_: No

_Type_: [AssertionRule](aws-properties-route53recoverycontrol-safetyrule-assertionrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlPanelArn`

The Amazon Resource Name (ARN) of the control panel.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`GatingRule`

A gating rule verifies that a gating routing control or set of gating routing controls, evaluates as true, based on a rule
configuration that you specify, which allows a set of routing control state changes to complete.

For example, if you specify one gating routing control and you set the `Type` in
the rule configuration to `OR`, that indicates that you must set the gating routing control to `On`
for the rule to evaluate as true; that is, for the gating control switch to be On. When you do that, then
you can update the routing control states for the target routing controls that you specify in the gating
rule.

_Required_: No

_Type_: [GatingRule](aws-properties-route53recoverycontrol-safetyrule-gatingrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the assertion rule. The name must be unique within a control panel.
You can use any non-white space character in the name except the following: & > < ' (single quote) " (double quote) ; (semicolon)

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleConfig`

The criteria that you set for specific assertion controls (routing controls) that designate
how many control states must be `ON` as the result of a transaction. For example, if you have three
assertion controls, you might specify `ATLEAST 2` for your rule configuration. This means
that at least two assertion controls must be `ON`, so that at least two AWS Regions
have traffic flowing to them.

_Required_: No

_Type_: [RuleConfig](aws-properties-route53recoverycontrol-safetyrule-ruleconfig.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

The tags associated with the safety rule.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53recoverycontrol-safetyrule-tag.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `SafetyRuleArn` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SafetyRuleArn`

The Amazon Resource Name (ARN) of the safety rule.

`Status`

The deployment status of the safety rule. Status can be one of the following: PENDING, DEPLOYED, PENDING\_DELETION.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Route53RecoveryControl::RoutingControl

AssertionRule

All content copied from https://docs.aws.amazon.com/.
