This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup SummaryConfiguration

A complex type that specifies which Suricata rule metadata fields to use when displaying threat information. Contains:

- `RuleOptions` \- The Suricata rule options fields to extract and display

These settings affect how threat information appears in both the console and API responses. Summaries are available for rule groups you manage and for active threat defense AWS managed rule groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RuleOptions" : [ String, ... ]
}

```

### YAML

```yaml

  RuleOptions:
    - String

```

## Properties

`RuleOptions`

Specifies the selected rule options returned by `DescribeRuleGroupSummary`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatelessRulesAndCustomActions

Tag

All content copied from https://docs.aws.amazon.com/.
