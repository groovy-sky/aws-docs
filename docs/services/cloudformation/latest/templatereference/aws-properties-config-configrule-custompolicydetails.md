This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigRule CustomPolicyDetails

Provides the CustomPolicyDetails, the rule owner ( `
                    AWS
                ` for managed rules, `CUSTOM_POLICY` for Custom Policy rules, and `CUSTOM_LAMBDA` for Custom Lambda rules), the rule
identifier, and the events that cause the evaluation of your AWS
resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableDebugLogDelivery" : Boolean,
  "PolicyRuntime" : String,
  "PolicyText" : String
}

```

### YAML

```yaml

  EnableDebugLogDelivery: Boolean
  PolicyRuntime: String
  PolicyText: String

```

## Properties

`EnableDebugLogDelivery`

The boolean expression for enabling debug logging for your AWS Config Custom Policy rule. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyRuntime`

The runtime system for your AWS Config Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by AWS Config Custom Policy rules. For more information about Guard, see the [Guard GitHub\
Repository](https://github.com/aws-cloudformation/cloudformation-guard).

_Required_: No

_Type_: String

_Pattern_: `guard\-2\.x\.x`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyText`

The policy definition containing the logic for your AWS Config Custom Policy rule.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Compliance

EvaluationModeConfiguration

All content copied from https://docs.aws.amazon.com/.
