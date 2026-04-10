This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FraudDetector::Detector Rule

A rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "CreatedTime" : String,
  "Description" : String,
  "DetectorId" : String,
  "Expression" : String,
  "Language" : String,
  "LastUpdatedTime" : String,
  "Outcomes" : [ Outcome, ... ],
  "RuleId" : String,
  "RuleVersion" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  Arn: String
  CreatedTime: String
  Description: String
  DetectorId: String
  Expression: String
  Language: String
  LastUpdatedTime: String
  Outcomes:
    - Outcome
  RuleId: String
  RuleVersion: String
  Tags:
    - Tag

```

## Properties

`Arn`

The rule ARN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedTime`

Timestamp for when the rule was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The rule description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorId`

The detector for which the rule is associated.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The rule expression. A rule expression captures the business logic. For more information, see [Rule language reference](../../../frauddetector/latest/ug/rule-language-reference.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Language`

The rule language.

Valid Value: DETECTORPL

_Required_: No

_Type_: String

_Allowed values_: `DETECTORPL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUpdatedTime`

Timestamp for when the rule was last updated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outcomes`

The rule outcome.

_Required_: No

_Type_: Array of [Outcome](aws-properties-frauddetector-detector-outcome.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleId`

The rule ID.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleVersion`

The rule version.

_Required_: No

_Type_: String

_Pattern_: `^([1-9][0-9]*)$`

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-frauddetector-detector-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Outcome

Tag

All content copied from https://docs.aws.amazon.com/.
