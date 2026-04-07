This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule

Retrieves the details of a specific organization centralization rule. This operation can
only be called by the organization's management account or a delegated administrator
account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ObservabilityAdmin::OrganizationTelemetryRule",
  "Properties" : {
      "Rule" : TelemetryRule,
      "RuleName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ObservabilityAdmin::OrganizationTelemetryRule
Properties:
  Rule:
    TelemetryRule
  RuleName: String
  Tags:
    - Tag

```

## Properties

`Rule`

The name of the organization telemetry rule.

_Required_: Yes

_Type_: [TelemetryRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationtelemetryrule-telemetryrule.html)

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

Lists all tags attached to the specified resource. Supports telemetry rule resources and
telemetry pipeline resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationtelemetryrule-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RuleArn`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ActionCondition
