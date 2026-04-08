# AnomalyMonitor

This object continuously inspects your account's cost data for anomalies. It's based
on `MonitorType` and `MonitorSpecification`. The content consists
of detailed metadata and the current status of the monitor object.

## Contents

**MonitorName**

The name of the monitor.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: Yes

**MonitorType**

The type of the monitor.

Set this to `DIMENSIONAL` for an AWS managed monitor.
AWS managed monitors automatically track up to the top 5,000 values by
cost within a dimension of your choosing. Each dimension value is evaluated
independently. If you start incurring cost in a new value of your chosen dimension, it
will automatically be analyzed by an AWS managed monitor.

Set this to `CUSTOM` for a customer managed monitor. Customer managed
monitors let you select specific dimension values that get monitored in aggregate.

For more information about monitor types, see [Monitor\
types](../../../../services/cost-management/latest/userguide/getting-started-ad-monitor-type-def.md) in the _Billing and Cost Management User Guide_.

Type: String

Valid Values: `DIMENSIONAL | CUSTOM`

Required: Yes

**CreationDate**

The date when the monitor was created.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 40.

Pattern: `(\d{4}-\d{2}-\d{2})(T\d{2}:\d{2}:\d{2}Z)?`

Required: No

**DimensionalValueCount**

The value for evaluated dimensions.

Type: Integer

Valid Range: Minimum value of 0.

Required: No

**LastEvaluatedDate**

The date when the monitor last evaluated for anomalies.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 40.

Pattern: `(\d{4}-\d{2}-\d{2})(T\d{2}:\d{2}:\d{2}Z)?`

Required: No

**LastUpdatedDate**

The date when the monitor was last updated.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 40.

Pattern: `(\d{4}-\d{2}-\d{2})(T\d{2}:\d{2}:\d{2}Z)?`

Required: No

**MonitorArn**

The Amazon Resource Name (ARN) value.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Pattern: `[\S\s]*`

Required: No

**MonitorDimension**

For customer managed monitors, do not specify this field.

For AWS managed monitors, this field controls which cost dimension is
automatically analyzed by the monitor. For `TAG` and `COST_CATEGORY
            ` dimensions, you must also specify MonitorSpecification to configure the specific
tag or cost category key to analyze.

Type: String

Valid Values: `SERVICE | LINKED_ACCOUNT | TAG | COST_CATEGORY`

Required: No

**MonitorSpecification**

An [Expression](api-expression.md)
object used to control what costs the monitor analyzes for anomalies.

For AWS managed monitors:

- If MonitorDimension is `SERVICE` or `LINKED_ACCOUNT`, do
not specify this field

- If MonitorDimension is `TAG`, set this field to `{ "Tags": {
                          "Key": "your tag key" } }`

- If MonitorDimension is `COST_CATEGORY`, set this field to `{
                          "CostCategories": { "Key": "your cost category key" } }`

For customer managed monitors:

- To track linked accounts, set this field to `{ "Dimensions": { "Key":
                          "LINKED_ACCOUNT", "Values": [ "your list of up to 10 account IDs" ] } }
                          `

- To track cost allocation tags, set this field to `{ "Tags": { "Key":
                          "your tag key", "Values": [ "your list of up to 10 tag values" ] } }
                          `

- To track cost categories, set this field to `{ "CostCategories": { "Key":
                          "your cost category key", "Values": [ "your cost category value" ] } }
                          `

Type: [Expression](api-expression.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/anomalymonitor.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/anomalymonitor.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/anomalymonitor.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnomalyDateInterval

AnomalyScore
