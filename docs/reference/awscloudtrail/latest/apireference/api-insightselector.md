# InsightSelector

A JSON string that contains a list of Insights types that are logged on a trail or event data store.

## Contents

**EventCategories**

Select the event category on which Insights should be enabled.

- If EventCategories is not provided, the specified Insights types are enabled on management API calls by default.

- If EventCategories is provided, the given event categories will overwrite the existing ones. For example,
if a trail already has Insights enabled on management events, and then a PutInsightSelectors request is made with only data events specified in EventCategories, Insights on management events will be disabled.

Type: Array of strings

Valid Values: `Management | Data`

Required: No

**InsightType**

The type of Insights events to log on a trail or event data store. `ApiCallRateInsight` and
`ApiErrorRateInsight` are valid Insight types.

The `ApiCallRateInsight` Insights type analyzes write-only
management API calls or read and write data API calls that are aggregated per minute against a baseline API call volume.

The `ApiErrorRateInsight` Insights type analyzes management and data
API calls that result in error codes. The error is shown if the API call is
unsuccessful.

Type: String

Valid Values: `ApiCallRateInsight | ApiErrorRateInsight`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/insightselector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/insightselector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/insightselector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngestionStatus

LookupAttribute

All content copied from https://docs.aws.amazon.com/.
