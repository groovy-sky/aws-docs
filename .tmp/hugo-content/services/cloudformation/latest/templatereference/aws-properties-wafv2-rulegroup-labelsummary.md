This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup LabelSummary

List of labels used by one or more of the rules of a [AWS::WAFv2::RuleGroup](aws-resource-wafv2-rulegroup.md). This
summary object is used for the following rule group lists:

- `AvailableLabels` \- Labels that rules add to matching requests. These
labels are defined in the `RuleLabels` for a rule.

- `ConsumedLabels` \- Labels that rules match against. These labels are
defined in a `LabelMatchStatement` specification, in the [Statement](../userguide/aws-properties-wafv2-webacl-notstatement.md#cfn-wafv2-webacl-notstatement-statement) definition of a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

An individual label specification.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_:-]{1,1024}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelMatchStatement

NotStatement

All content copied from https://docs.aws.amazon.com/.
