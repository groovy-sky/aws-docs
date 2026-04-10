This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleIpExpression

An IP address expression matching certain IP addresses within a given range of IP
addresses.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Evaluate" : RuleIpToEvaluate,
  "Operator" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Evaluate:
    RuleIpToEvaluate
  Operator: String
  Values:
    - String

```

## Properties

`Evaluate`

The IP address to evaluate in this condition.

_Required_: Yes

_Type_: [RuleIpToEvaluate](aws-properties-ses-mailmanagerruleset-ruleiptoevaluate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The operator to evaluate the IP address.

_Required_: Yes

_Type_: String

_Allowed values_: `CIDR_MATCHES | NOT_CIDR_MATCHES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The IP CIDR blocks in format "x.y.z.w/n" (eg 10.0.0.0/8) to match with the
email's IP address. For the operator CIDR\_MATCHES, if multiple values are given, they
are evaluated as an OR. That is, if the IP address is contained within any of the given
CIDR ranges, the condition is deemed to match. For NOT\_CIDR\_MATCHES, if multiple CIDR
ranges are given, the condition is deemed to match if the IP address is not contained in
any of the given CIDR ranges.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `43 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleDmarcExpression

RuleIpToEvaluate

All content copied from https://docs.aws.amazon.com/.
