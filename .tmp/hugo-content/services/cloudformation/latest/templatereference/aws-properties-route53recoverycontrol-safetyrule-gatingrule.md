This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::SafetyRule GatingRule

A gating rule verifies that a gating routing control or set of gating routing controls, evaluates as true, based on a rule
configuration that you specify, which allows a set of routing control state changes to complete.

For example, if you specify one gating routing control and you set the `Type` in
the rule configuration to `OR`, that indicates that you must set the gating routing control to `On`
for the rule to evaluate as true; that is, for the gating control switch to be On. When you do that, then
you can update the routing control states for the target routing controls that you specify in the gating
rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GatingControls" : [ String, ... ],
  "TargetControls" : [ String, ... ],
  "WaitPeriodMs" : Integer
}

```

### YAML

```yaml

  GatingControls:
    - String
  TargetControls:
    - String
  WaitPeriodMs: Integer

```

## Properties

`GatingControls`

An array of gating routing control Amazon Resource Names (ARNs). For a simple on-off switch, specify the ARN for
one routing control. The gating routing controls are evaluated by the rule configuration that you specify to
determine if the target routing control states can be changed.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`TargetControls`

An array of target routing control Amazon Resource Names (ARNs) for which the states can only be updated
if the rule configuration that you specify evaluates to true
for the gating routing control. As a simple example, if you have a single gating control, it acts as an overall
on-off switch for a set of target routing controls. You can use this to manually override automated failover,
for example.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WaitPeriodMs`

An evaluation period, in milliseconds (ms), during which any request against the target routing controls
will fail. This helps prevent flapping of state. The wait period is 5000 ms by default, but you can choose
a custom value.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssertionRule

RuleConfig

All content copied from https://docs.aws.amazon.com/.
