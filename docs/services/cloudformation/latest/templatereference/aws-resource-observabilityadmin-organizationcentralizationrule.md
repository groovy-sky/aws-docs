This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule

Defines how telemetry data should be centralized across an AWS Organization, including
source and destination configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ObservabilityAdmin::OrganizationCentralizationRule",
  "Properties" : {
      "Rule" : CentralizationRule,
      "RuleName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ObservabilityAdmin::OrganizationCentralizationRule
Properties:
  Rule:
    CentralizationRule
  RuleName: String
  Tags:
    - Tag

```

## Properties

`Rule`

Property description not available.

_Required_: Yes

_Type_: [CentralizationRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the organization centralization rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z-]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A key-value pair to filter resources based on tags associated with the resource. For
more information about tags, see [What are tags?](https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/what-are-tags.html)

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationcentralizationrule-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RuleArn`

The Amazon Resource Name (ARN) of the organization centralization rule.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatch Observability Admin

CentralizationRule
