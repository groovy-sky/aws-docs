# DecisionTask

A structure that represents a decision task. Decision tasks are sent to deciders in order for them to make decisions.

## Contents

**events**

A paginated list of history events of the workflow execution. The decider uses this during the processing of the decision task.

Type: Array of [HistoryEvent](api-historyevent.md) objects

Required: Yes

**startedEventId**

The ID of the `DecisionTaskStarted` event recorded in the history.

Type: Long

Required: Yes

**taskToken**

The opaque string used as a handle on the task. This token is used by workers to communicate progress and response information back to the system about the task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**workflowExecution**

The workflow execution for which this decision task was created.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

**workflowType**

The type of the workflow execution for which this decision task was created.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**nextPageToken**

If a `NextPageToken` was returned by a previous call, there are more
results available. To retrieve the next page of results, make the call again using the returned token in
`nextPageToken`. Keep all other arguments unchanged.

The configured `maximumPageSize` determines how many results can be returned in a single call.

Type: String

Length Constraints: Maximum length of 2048.

Required: No

**previousStartedEventId**

The ID of the DecisionTaskStarted event of the previous decision task of this workflow execution that was processed by the decider. This can be used to determine the events in the history new since the last decision task received by the decider.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/decisiontask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/decisiontask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/decisiontask.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Decision

DecisionTaskCompletedEventAttributes

All content copied from https://docs.aws.amazon.com/.
