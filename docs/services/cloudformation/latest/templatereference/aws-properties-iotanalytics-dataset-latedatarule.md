This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset LateDataRule

A structure that contains the name and configuration information of a late data
rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RuleConfiguration" : LateDataRuleConfiguration,
  "RuleName" : String
}

```

### YAML

```yaml

  RuleConfiguration:
    LateDataRuleConfiguration
  RuleName: String

```

## Properties

`RuleConfiguration`

The information needed to configure the late data rule.

_Required_: Yes

_Type_: [LateDataRuleConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-latedataruleconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the late data rule.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IotEventsDestinationConfiguration

LateDataRuleConfiguration
