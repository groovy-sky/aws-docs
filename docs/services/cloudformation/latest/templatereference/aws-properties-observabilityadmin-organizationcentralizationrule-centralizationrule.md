This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationCentralizationRule CentralizationRule

Defines how telemetry data should be centralized across an AWS Organization, including
source and destination configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : CentralizationRuleDestination,
  "Source" : CentralizationRuleSource
}

```

### YAML

```yaml

  Destination:
    CentralizationRuleDestination
  Source:
    CentralizationRuleSource

```

## Properties

`Destination`

Configuration determining where the telemetry data should be centralized, backed up, as
well as encryption configuration for the primary and backup destinations.

_Required_: Yes

_Type_: [CentralizationRuleDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Configuration determining the source of the telemetry data to be centralized.

_Required_: Yes

_Type_: [CentralizationRuleSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrulesource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ObservabilityAdmin::OrganizationCentralizationRule

CentralizationRuleDestination
