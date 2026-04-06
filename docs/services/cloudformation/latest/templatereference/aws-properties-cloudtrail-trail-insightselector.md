This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Trail InsightSelector

A JSON string that contains a list of Insights types that are logged on a trail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventCategories" : [ String, ... ],
  "InsightType" : String
}

```

### YAML

```yaml

  EventCategories:
    - String
  InsightType: String

```

## Properties

`EventCategories`

Select the event category on which Insights should be enabled.

- If EventCategories is not provided, the specified Insights types are enabled on management API calls by default.

- If EventCategories is provided, the given event categories will overwrite the existing ones. For example,
if a trail already has Insights enabled on management events, and then a PutInsightSelectors request is made with only data events specified in EventCategories, Insights on management events will be disabled.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightType`

The type of Insights events to log on a trail. `ApiCallRateInsight` and
`ApiErrorRateInsight` are valid Insight types.

The `ApiCallRateInsight` Insights type analyzes write-only
management API calls that are aggregated per minute against a baseline API call volume.

The `ApiErrorRateInsight` Insights type analyzes management
API calls that result in error codes. The error is shown if the API call is
unsuccessful.

_Required_: No

_Type_: String

_Allowed values_: `ApiCallRateInsight | ApiErrorRateInsight`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventSelector

Tag
