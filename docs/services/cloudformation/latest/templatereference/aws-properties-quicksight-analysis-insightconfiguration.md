This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis InsightConfiguration

The configuration of an insight visual.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Computations" : [ Computation, ... ],
  "CustomNarrative" : CustomNarrativeOptions,
  "Interactions" : VisualInteractionOptions
}

```

### YAML

```yaml

  Computations:
    - Computation
  CustomNarrative:
    CustomNarrativeOptions
  Interactions:
    VisualInteractionOptions

```

## Properties

`Computations`

The computations configurations of the insight visual

_Required_: No

_Type_: Array of [Computation](aws-properties-quicksight-analysis-computation.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomNarrative`

The custom narrative of the insight visual.

_Required_: No

_Type_: [CustomNarrativeOptions](aws-properties-quicksight-analysis-customnarrativeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InnerFilter

InsightVisual

All content copied from https://docs.aws.amazon.com/.
