This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective BurnRateConfiguration

This object defines the length of the look-back window used to calculate one burn rate metric
for this SLO. The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO. A burn rate of
exactly 1 indicates that the SLO goal will be met exactly.

For example, if you specify 60 as the number of minutes in the look-back window, the burn rate is calculated as the following:

_burn rate = error rate over the look-back window / (100% - attainment goal percentage)_

For more information about burn rates, see [Calculate burn rates](../../../amazoncloudwatch/latest/monitoring/cloudwatch-servicelevelobjectives.md#CloudWatch-ServiceLevelObjectives-burn).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LookBackWindowMinutes" : Integer
}

```

### YAML

```yaml

  LookBackWindowMinutes: Integer

```

## Properties

`LookBackWindowMinutes`

The number of minutes to use as the look-back window.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10080`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApplicationSignals::ServiceLevelObjective

CalendarInterval

All content copied from https://docs.aws.amazon.com/.
