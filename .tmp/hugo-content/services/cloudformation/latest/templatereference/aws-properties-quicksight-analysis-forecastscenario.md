This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ForecastScenario

The forecast scenario of a forecast in the line chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WhatIfPointScenario" : WhatIfPointScenario,
  "WhatIfRangeScenario" : WhatIfRangeScenario
}

```

### YAML

```yaml

  WhatIfPointScenario:
    WhatIfPointScenario
  WhatIfRangeScenario:
    WhatIfRangeScenario

```

## Properties

`WhatIfPointScenario`

The what-if analysis forecast setup with the target date.

_Required_: No

_Type_: [WhatIfPointScenario](aws-properties-quicksight-analysis-whatifpointscenario.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhatIfRangeScenario`

The what-if analysis forecast setup with the date range.

_Required_: No

_Type_: [WhatIfRangeScenario](aws-properties-quicksight-analysis-whatifrangescenario.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ForecastConfiguration

FormatConfiguration

All content copied from https://docs.aws.amazon.com/.
