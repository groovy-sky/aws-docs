This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::EventDataStore InsightSelector

A JSON string that contains a list of Insights types that are logged on an event data store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InsightType" : String
}

```

### YAML

```yaml

  InsightType: String

```

## Properties

`InsightType`

The type of Insights events to log on an event data store. `ApiCallRateInsight` and
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

ContextKeySelector

Tag
