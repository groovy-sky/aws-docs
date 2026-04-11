This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule CloudtrailParameters

Parameters specific to AWS CloudTrail telemetry configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvancedEventSelectors" : [ AdvancedEventSelector, ... ]
}

```

### YAML

```yaml

  AdvancedEventSelectors:
    - AdvancedEventSelector

```

## Properties

`AdvancedEventSelectors`

The advanced event selectors to use for filtering AWS CloudTrail events.

_Required_: Yes

_Type_: Array of [AdvancedEventSelector](aws-properties-observabilityadmin-telemetryrule-advancedeventselector.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedFieldSelector

Condition

All content copied from https://docs.aws.amazon.com/.
