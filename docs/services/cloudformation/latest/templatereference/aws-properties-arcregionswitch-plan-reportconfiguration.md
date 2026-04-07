This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan ReportConfiguration

Configuration for automatic report generation for plan executions. When configured, Region switch automatically generates a report after each plan execution that includes execution events, plan configuration, and CloudWatch alarm states.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReportOutput" : [ ReportOutputConfiguration, ... ]
}

```

### YAML

```yaml

  ReportOutput:
    - ReportOutputConfiguration

```

## Properties

`ReportOutput`

The output configuration for the report.

_Required_: No

_Type_: Array of [ReportOutputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-arcregionswitch-plan-reportoutputconfiguration.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RegionSwitchPlanConfiguration

ReportOutputConfiguration
