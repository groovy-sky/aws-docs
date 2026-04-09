# MarkerRecordedEventAttributes

Provides the details of the `MarkerRecorded` event.

## Contents

**decisionTaskCompletedEventId**

The ID of the `DecisionTaskCompleted` event corresponding to the decision task that resulted in the
`RecordMarker` decision that requested this marker. This information can be useful for diagnosing problems by tracing back the chain of
events leading up to this event.

Type: Long

Required: Yes

**markerName**

The name of the marker.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**details**

The details of the marker.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/markerrecordedeventattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/markerrecordedeventattributes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/markerrecordedeventattributes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionTimedOutEventAttributes

RecordMarkerDecisionAttributes

All content copied from https://docs.aws.amazon.com/.
