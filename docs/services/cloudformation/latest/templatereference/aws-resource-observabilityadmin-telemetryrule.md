This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule

Creates a telemetry rule that defines how telemetry should be configured for AWS
resources in your account. The rule specifies which resources should have telemetry enabled
and how that telemetry data should be collected based on resource type, telemetry type, and
selection criteria.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ObservabilityAdmin::TelemetryRule",
  "Properties" : {
      "Rule" : TelemetryRule,
      "RuleName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ObservabilityAdmin::TelemetryRule
Properties:
  Rule:
    TelemetryRule
  RuleName: String
  Tags:
    - Tag

```

## Properties

`Rule`

Retrieves the details of a specific telemetry rule in your account.

_Required_: Yes

_Type_: [TelemetryRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-telemetryrule-telemetryrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the telemetry rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z-]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Lists all tags attached to the specified resource. Supports telemetry rule resources and
telemetry pipeline resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-telemetryrule-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RuleArn`

The Amazon Resource Name (ARN) of the telemetry rule.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TelemetryPipelineStatusReason

ActionCondition
