This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::InsightRule

Creates or updates a Contributor Insights rule. Rules evaluate log events in a CloudWatch Logs log group, enabling you to find contributor data
for the log events in that log group. For more information, see [Using Contributor Insights to Analyze High-Cardinality Data](../../../amazoncloudwatch/latest/monitoring/contributorinsights.md) in the _Amazon CloudWatch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::InsightRule",
  "Properties" : {
      "ApplyOnTransformedLogs" : Boolean,
      "RuleBody" : String,
      "RuleName" : String,
      "RuleState" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::InsightRule
Properties:
  ApplyOnTransformedLogs: Boolean
  RuleBody: String
  RuleName: String
  RuleState: String
  Tags:
    - Tag

```

## Properties

`ApplyOnTransformedLogs`

Determines whether the rules is evaluated on transformed versions of logs. Valid values are `TRUE` and `FALSE`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleBody`

The definition of the rule, as a JSON object.
For details about the syntax, see [Contributor Insights Rule Syntax](../../../amazoncloudwatch/latest/monitoring/contributorinsights-rulesyntax.md) in the _Amazon CloudWatch User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the rule.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RuleState`

The current state of the rule. Valid values are `ENABLED` and `DISABLED`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to associate with the Contributor Insights rule. You can
associate as many as 50 tags with a rule.

Tags can help you organize and categorize your resources. For more information,
see [Tagging Your Amazon CloudWatch Resources](../../../amazoncloudwatch/latest/monitoring/cloudwatch-tagging.md).

To be able to associate tags with a rule, you must have the `cloudwatch:TagResource` permission in addition to the `cloudwatch:PutInsightRule` permission.

_Required_: No

_Type_: Array of [Tag](aws-properties-cloudwatch-insightrule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the Contributor Insights rule, such as `arn:aws:cloudwatch:us-west-2:123456789012:insight-rule/MyInsightRuleName`.

`RuleName`

The name of the Contributor Insights rule.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudWatch::Dashboard

Tag

All content copied from https://docs.aws.amazon.com/.
