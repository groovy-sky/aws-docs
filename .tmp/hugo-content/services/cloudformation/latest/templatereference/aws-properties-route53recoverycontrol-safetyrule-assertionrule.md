This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryControl::SafetyRule AssertionRule

An assertion rule enforces that, when you change a routing control state, that the criteria that you set in the rule configuration is met.
Otherwise, the change to the routing control is not accepted. For example, the criteria might be that at least one routing
control state is `On` after the transaction so that traffic continues to flow to at least one cell for the application.
This ensures that you avoid a fail-open scenario.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssertedControls" : [ String, ... ],
  "WaitPeriodMs" : Integer
}

```

### YAML

```yaml

  AssertedControls:
    - String
  WaitPeriodMs: Integer

```

## Properties

`AssertedControls`

The routing controls that are part of transactions that are evaluated to determine if a request
to change a routing control state is allowed. For example, you might include three routing controls,
one for each of three AWS Regions.

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

AWS::Route53RecoveryControl::SafetyRule

GatingRule

All content copied from https://docs.aws.amazon.com/.
