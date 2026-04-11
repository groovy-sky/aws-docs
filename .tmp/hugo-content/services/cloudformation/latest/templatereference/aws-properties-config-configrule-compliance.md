This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigRule Compliance

Indicates whether an AWS resource or AWS Config rule is
compliant and provides the number of contributors that affect the
compliance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

Indicates whether an AWS resource or AWS Config rule is
compliant.

A resource is compliant if it complies with all of the AWS Config rules that evaluate it. A resource is noncompliant if it does
not comply with one or more of these rules.

A rule is compliant if all of the resources that the rule
evaluates comply with it. A rule is noncompliant if any of these
resources do not comply.

AWS Config returns the `INSUFFICIENT_DATA` value
when no evaluation results are available for the AWS resource or AWS Config rule.

For the `Compliance` data type, AWS Config supports
only `COMPLIANT`, `NON_COMPLIANT`, and
`INSUFFICIENT_DATA` values. AWS Config does not
support the `NOT_APPLICABLE` value for the
`Compliance` data type.

_Required_: No

_Type_: String

_Allowed values_: `COMPLIANT | NON_COMPLIANT | NOT_APPLICABLE | INSUFFICIENT_DATA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Config::ConfigRule

CustomPolicyDetails

All content copied from https://docs.aws.amazon.com/.
