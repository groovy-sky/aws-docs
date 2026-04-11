This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis KPIConfiguration

The configuration of a KPI visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldWells" : KPIFieldWells,
  "Interactions" : VisualInteractionOptions,
  "KPIOptions" : KPIOptions,
  "SortConfiguration" : KPISortConfiguration
}

```

### YAML

```yaml

  FieldWells:
    KPIFieldWells
  Interactions:
    VisualInteractionOptions
  KPIOptions:
    KPIOptions
  SortConfiguration:
    KPISortConfiguration

```

## Properties

`FieldWells`

The field well configuration of a KPI visual.

_Required_: No

_Type_: [KPIFieldWells](aws-properties-quicksight-analysis-kpifieldwells.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KPIOptions`

The options that determine the presentation of a KPI visual.

_Required_: No

_Type_: [KPIOptions](aws-properties-quicksight-analysis-kpioptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SortConfiguration`

The sort configuration of a KPI visual.

_Required_: No

_Type_: [KPISortConfiguration](aws-properties-quicksight-analysis-kpisortconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KPIConditionalFormattingOption

KPIFieldWells

All content copied from https://docs.aws.amazon.com/.
