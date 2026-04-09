# HistoryEvent

Event within a workflow execution. A history event can be one of these types:

- `ActivityTaskCancelRequested` – A `RequestCancelActivityTask` decision was received by the
system.

- `ActivityTaskCanceled` – The activity task was successfully canceled.

- `ActivityTaskCompleted` – An activity worker successfully completed an activity task by calling
[RespondActivityTaskCompleted](api-respondactivitytaskcompleted.md).

- `ActivityTaskFailed` – An activity worker failed an activity task by calling
[RespondActivityTaskFailed](api-respondactivitytaskfailed.md).

- `ActivityTaskScheduled` – An activity task was scheduled for execution.

- `ActivityTaskStarted` – The scheduled activity task was dispatched to a worker.

- `ActivityTaskTimedOut` – The activity task timed out.

- `CancelTimerFailed` – Failed to process CancelTimer decision. This happens when the decision isn't
configured properly, for example no timer exists with the specified timer Id.

- `CancelWorkflowExecutionFailed` – A request to cancel a workflow execution failed.

- `ChildWorkflowExecutionCanceled` – A child workflow execution, started by this workflow execution, was
canceled and closed.

- `ChildWorkflowExecutionCompleted` – A child workflow execution, started by this workflow execution,
completed successfully and was closed.

- `ChildWorkflowExecutionFailed` – A child workflow execution,
started by this workflow execution, failed to complete successfully and was closed.

- `ChildWorkflowExecutionStarted` – A child workflow execution was successfully started.

- `ChildWorkflowExecutionTerminated` – A child workflow execution, started by this workflow execution, was
terminated.

- `ChildWorkflowExecutionTimedOut` – A child workflow execution, started by this workflow execution, timed
out and was closed.

- `CompleteWorkflowExecutionFailed` – The workflow execution failed to complete.

- `ContinueAsNewWorkflowExecutionFailed` – The workflow execution failed to complete after being continued
as a new workflow execution.

- `DecisionTaskCompleted` – The decider successfully completed a decision task by calling
[RespondDecisionTaskCompleted](api-responddecisiontaskcompleted.md).

- `DecisionTaskScheduled` – A decision task was scheduled for the workflow execution.

- `DecisionTaskStarted` – The decision task was dispatched to a decider.

- `DecisionTaskTimedOut` – The decision task timed out.

- `ExternalWorkflowExecutionCancelRequested` – Request to cancel an external workflow execution was
successfully delivered to the target execution.

- `ExternalWorkflowExecutionSignaled` – A signal, requested by this workflow execution, was successfully
delivered to the target external workflow execution.

- `FailWorkflowExecutionFailed` – A request to mark a workflow execution as failed, itself failed.

- `MarkerRecorded` – A marker was recorded in the workflow history as the result of a
`RecordMarker` decision.

- `RecordMarkerFailed` – A `RecordMarker` decision was returned as failed.

- `RequestCancelActivityTaskFailed` – Failed to process RequestCancelActivityTask decision. This happens
when the decision isn't configured properly.

- `RequestCancelExternalWorkflowExecutionFailed` – Request to cancel an external workflow execution
failed.

- `RequestCancelExternalWorkflowExecutionInitiated` – A request was made to request the cancellation of an
external workflow execution.

- `ScheduleActivityTaskFailed` – Failed to process ScheduleActivityTask decision. This happens when the
decision isn't configured properly, for example the activity type specified isn't registered.

- `SignalExternalWorkflowExecutionFailed` – The request to signal an external workflow execution
failed.

- `SignalExternalWorkflowExecutionInitiated` – A request to signal an external workflow was made.

- `StartActivityTaskFailed` – A scheduled activity task failed to start.

- `StartChildWorkflowExecutionFailed` – Failed to process StartChildWorkflowExecution decision. This happens
when the decision isn't configured properly, for example the workflow type specified isn't registered.

- `StartChildWorkflowExecutionInitiated` – A request was made to start a child workflow execution.

- `StartTimerFailed` – Failed to process StartTimer decision. This happens when the decision isn't
configured properly, for example a timer already exists with the specified timer Id.

- `TimerCanceled` – A timer, previously started for this workflow execution, was successfully canceled.

- `TimerFired` – A timer, previously started for this workflow execution, fired.

- `TimerStarted` – A timer was started for the workflow execution due to a `StartTimer`
decision.

- `WorkflowExecutionCancelRequested` – A request to cancel this workflow execution was made.

- `WorkflowExecutionCanceled` – The workflow execution was successfully canceled and closed.

- `WorkflowExecutionCompleted` – The workflow execution was closed due to successful completion.

- `WorkflowExecutionContinuedAsNew` – The workflow execution was closed and a new execution of the same type
was created with the same workflowId.

- `WorkflowExecutionFailed` – The workflow execution closed due to a failure.

- `WorkflowExecutionSignaled` – An external signal was received for the workflow execution.

- `WorkflowExecutionStarted` – The workflow execution was started.

- `WorkflowExecutionTerminated` – The workflow execution was terminated.

- `WorkflowExecutionTimedOut` – The workflow execution was closed because a time out was exceeded.

## Contents

**eventId**

The system generated ID of the event. This ID uniquely identifies the event with in the workflow execution history.

Type: Long

Required: Yes

**eventTimestamp**

The date and time when the event occurred.

Type: Timestamp

Required: Yes

**eventType**

The type of the history event.

Type: String

Valid Values: `WorkflowExecutionStarted | WorkflowExecutionCancelRequested | WorkflowExecutionCompleted | CompleteWorkflowExecutionFailed | WorkflowExecutionFailed | FailWorkflowExecutionFailed | WorkflowExecutionTimedOut | WorkflowExecutionCanceled | CancelWorkflowExecutionFailed | WorkflowExecutionContinuedAsNew | ContinueAsNewWorkflowExecutionFailed | WorkflowExecutionTerminated | DecisionTaskScheduled | DecisionTaskStarted | DecisionTaskCompleted | DecisionTaskTimedOut | ActivityTaskScheduled | ScheduleActivityTaskFailed | ActivityTaskStarted | ActivityTaskCompleted | ActivityTaskFailed | ActivityTaskTimedOut | ActivityTaskCanceled | ActivityTaskCancelRequested | RequestCancelActivityTaskFailed | WorkflowExecutionSignaled | MarkerRecorded | RecordMarkerFailed | TimerStarted | StartTimerFailed | TimerFired | TimerCanceled | CancelTimerFailed | StartChildWorkflowExecutionInitiated | StartChildWorkflowExecutionFailed | ChildWorkflowExecutionStarted | ChildWorkflowExecutionCompleted | ChildWorkflowExecutionFailed | ChildWorkflowExecutionTimedOut | ChildWorkflowExecutionCanceled | ChildWorkflowExecutionTerminated | SignalExternalWorkflowExecutionInitiated | SignalExternalWorkflowExecutionFailed | ExternalWorkflowExecutionSignaled | RequestCancelExternalWorkflowExecutionInitiated | RequestCancelExternalWorkflowExecutionFailed | ExternalWorkflowExecutionCancelRequested | LambdaFunctionScheduled | LambdaFunctionStarted | LambdaFunctionCompleted | LambdaFunctionFailed | LambdaFunctionTimedOut | ScheduleLambdaFunctionFailed | StartLambdaFunctionFailed`

Required: Yes

**activityTaskCanceledEventAttributes**

If the event is of type `ActivityTaskCanceled` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskCanceledEventAttributes](api-activitytaskcanceledeventattributes.md) object

Required: No

**activityTaskCancelRequestedEventAttributes**

If the event is of type `ActivityTaskcancelRequested` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskCancelRequestedEventAttributes](api-activitytaskcancelrequestedeventattributes.md) object

Required: No

**activityTaskCompletedEventAttributes**

If the event is of type `ActivityTaskCompleted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskCompletedEventAttributes](api-activitytaskcompletedeventattributes.md) object

Required: No

**activityTaskFailedEventAttributes**

If the event is of type `ActivityTaskFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskFailedEventAttributes](api-activitytaskfailedeventattributes.md) object

Required: No

**activityTaskScheduledEventAttributes**

If the event is of type `ActivityTaskScheduled` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskScheduledEventAttributes](api-activitytaskscheduledeventattributes.md) object

Required: No

**activityTaskStartedEventAttributes**

If the event is of type `ActivityTaskStarted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskStartedEventAttributes](api-activitytaskstartedeventattributes.md) object

Required: No

**activityTaskTimedOutEventAttributes**

If the event is of type `ActivityTaskTimedOut` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ActivityTaskTimedOutEventAttributes](api-activitytasktimedouteventattributes.md) object

Required: No

**cancelTimerFailedEventAttributes**

If the event is of type `CancelTimerFailed` then this member is set and provides detailed information
about the event. It isn't set for other event types.

Type: [CancelTimerFailedEventAttributes](api-canceltimerfailedeventattributes.md) object

Required: No

**cancelWorkflowExecutionFailedEventAttributes**

If the event is of type `CancelWorkflowExecutionFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [CancelWorkflowExecutionFailedEventAttributes](api-cancelworkflowexecutionfailedeventattributes.md) object

Required: No

**childWorkflowExecutionCanceledEventAttributes**

If the event is of type `ChildWorkflowExecutionCanceled` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [ChildWorkflowExecutionCanceledEventAttributes](api-childworkflowexecutioncanceledeventattributes.md) object

Required: No

**childWorkflowExecutionCompletedEventAttributes**

If the event is of type `ChildWorkflowExecutionCompleted` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [ChildWorkflowExecutionCompletedEventAttributes](api-childworkflowexecutioncompletedeventattributes.md) object

Required: No

**childWorkflowExecutionFailedEventAttributes**

If the event is of type `ChildWorkflowExecutionFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ChildWorkflowExecutionFailedEventAttributes](api-childworkflowexecutionfailedeventattributes.md) object

Required: No

**childWorkflowExecutionStartedEventAttributes**

If the event is of type `ChildWorkflowExecutionStarted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ChildWorkflowExecutionStartedEventAttributes](api-childworkflowexecutionstartedeventattributes.md) object

Required: No

**childWorkflowExecutionTerminatedEventAttributes**

If the event is of type `ChildWorkflowExecutionTerminated` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [ChildWorkflowExecutionTerminatedEventAttributes](api-childworkflowexecutionterminatedeventattributes.md) object

Required: No

**childWorkflowExecutionTimedOutEventAttributes**

If the event is of type `ChildWorkflowExecutionTimedOut` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [ChildWorkflowExecutionTimedOutEventAttributes](api-childworkflowexecutiontimedouteventattributes.md) object

Required: No

**completeWorkflowExecutionFailedEventAttributes**

If the event is of type `CompleteWorkflowExecutionFailed` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [CompleteWorkflowExecutionFailedEventAttributes](api-completeworkflowexecutionfailedeventattributes.md) object

Required: No

**continueAsNewWorkflowExecutionFailedEventAttributes**

If the event is of type `ContinueAsNewWorkflowExecutionFailed` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [ContinueAsNewWorkflowExecutionFailedEventAttributes](api-continueasnewworkflowexecutionfailedeventattributes.md) object

Required: No

**decisionTaskCompletedEventAttributes**

If the event is of type `DecisionTaskCompleted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [DecisionTaskCompletedEventAttributes](api-decisiontaskcompletedeventattributes.md) object

Required: No

**decisionTaskScheduledEventAttributes**

If the event is of type `DecisionTaskScheduled` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [DecisionTaskScheduledEventAttributes](api-decisiontaskscheduledeventattributes.md) object

Required: No

**decisionTaskStartedEventAttributes**

If the event is of type `DecisionTaskStarted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [DecisionTaskStartedEventAttributes](api-decisiontaskstartedeventattributes.md) object

Required: No

**decisionTaskTimedOutEventAttributes**

If the event is of type `DecisionTaskTimedOut` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [DecisionTaskTimedOutEventAttributes](api-decisiontasktimedouteventattributes.md) object

Required: No

**externalWorkflowExecutionCancelRequestedEventAttributes**

If the event is of type `ExternalWorkflowExecutionCancelRequested` then this member is set and
provides detailed information about the event. It isn't set for other event types.

Type: [ExternalWorkflowExecutionCancelRequestedEventAttributes](api-externalworkflowexecutioncancelrequestedeventattributes.md) object

Required: No

**externalWorkflowExecutionSignaledEventAttributes**

If the event is of type `ExternalWorkflowExecutionSignaled` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [ExternalWorkflowExecutionSignaledEventAttributes](api-externalworkflowexecutionsignaledeventattributes.md) object

Required: No

**failWorkflowExecutionFailedEventAttributes**

If the event is of type `FailWorkflowExecutionFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [FailWorkflowExecutionFailedEventAttributes](api-failworkflowexecutionfailedeventattributes.md) object

Required: No

**lambdaFunctionCompletedEventAttributes**

Provides the details of the `LambdaFunctionCompleted` event. It isn't set
for other event types.

Type: [LambdaFunctionCompletedEventAttributes](api-lambdafunctioncompletedeventattributes.md) object

Required: No

**lambdaFunctionFailedEventAttributes**

Provides the details of the `LambdaFunctionFailed` event. It isn't set for
other event types.

Type: [LambdaFunctionFailedEventAttributes](api-lambdafunctionfailedeventattributes.md) object

Required: No

**lambdaFunctionScheduledEventAttributes**

Provides the details of the `LambdaFunctionScheduled` event. It isn't set
for other event types.

Type: [LambdaFunctionScheduledEventAttributes](api-lambdafunctionscheduledeventattributes.md) object

Required: No

**lambdaFunctionStartedEventAttributes**

Provides the details of the `LambdaFunctionStarted` event. It isn't set for
other event types.

Type: [LambdaFunctionStartedEventAttributes](api-lambdafunctionstartedeventattributes.md) object

Required: No

**lambdaFunctionTimedOutEventAttributes**

Provides the details of the `LambdaFunctionTimedOut` event. It isn't set for
other event types.

Type: [LambdaFunctionTimedOutEventAttributes](api-lambdafunctiontimedouteventattributes.md) object

Required: No

**markerRecordedEventAttributes**

If the event is of type `MarkerRecorded` then this member is set and provides detailed information
about the event. It isn't set for other event types.

Type: [MarkerRecordedEventAttributes](api-markerrecordedeventattributes.md) object

Required: No

**recordMarkerFailedEventAttributes**

If the event is of type `DecisionTaskFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [RecordMarkerFailedEventAttributes](api-recordmarkerfailedeventattributes.md) object

Required: No

**requestCancelActivityTaskFailedEventAttributes**

If the event is of type `RequestCancelActivityTaskFailed` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [RequestCancelActivityTaskFailedEventAttributes](api-requestcancelactivitytaskfailedeventattributes.md) object

Required: No

**requestCancelExternalWorkflowExecutionFailedEventAttributes**

If the event is of type `RequestCancelExternalWorkflowExecutionFailed` then this member is set and
provides detailed information about the event. It isn't set for other event types.

Type: [RequestCancelExternalWorkflowExecutionFailedEventAttributes](api-requestcancelexternalworkflowexecutionfailedeventattributes.md) object

Required: No

**requestCancelExternalWorkflowExecutionInitiatedEventAttributes**

If the event is of type `RequestCancelExternalWorkflowExecutionInitiated` then this member is set and
provides detailed information about the event. It isn't set for other event types.

Type: [RequestCancelExternalWorkflowExecutionInitiatedEventAttributes](api-requestcancelexternalworkflowexecutioninitiatedeventattributes.md) object

Required: No

**scheduleActivityTaskFailedEventAttributes**

If the event is of type `ScheduleActivityTaskFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [ScheduleActivityTaskFailedEventAttributes](api-scheduleactivitytaskfailedeventattributes.md) object

Required: No

**scheduleLambdaFunctionFailedEventAttributes**

Provides the details of the `ScheduleLambdaFunctionFailed` event. It isn't
set for other event types.

Type: [ScheduleLambdaFunctionFailedEventAttributes](api-schedulelambdafunctionfailedeventattributes.md) object

Required: No

**signalExternalWorkflowExecutionFailedEventAttributes**

If the event is of type `SignalExternalWorkflowExecutionFailed` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [SignalExternalWorkflowExecutionFailedEventAttributes](api-signalexternalworkflowexecutionfailedeventattributes.md) object

Required: No

**signalExternalWorkflowExecutionInitiatedEventAttributes**

If the event is of type `SignalExternalWorkflowExecutionInitiated` then this member is set and
provides detailed information about the event. It isn't set for other event types.

Type: [SignalExternalWorkflowExecutionInitiatedEventAttributes](api-signalexternalworkflowexecutioninitiatedeventattributes.md) object

Required: No

**startChildWorkflowExecutionFailedEventAttributes**

If the event is of type `StartChildWorkflowExecutionFailed` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [StartChildWorkflowExecutionFailedEventAttributes](api-startchildworkflowexecutionfailedeventattributes.md) object

Required: No

**startChildWorkflowExecutionInitiatedEventAttributes**

If the event is of type `StartChildWorkflowExecutionInitiated` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [StartChildWorkflowExecutionInitiatedEventAttributes](api-startchildworkflowexecutioninitiatedeventattributes.md) object

Required: No

**startLambdaFunctionFailedEventAttributes**

Provides the details of the `StartLambdaFunctionFailed` event. It isn't set
for other event types.

Type: [StartLambdaFunctionFailedEventAttributes](api-startlambdafunctionfailedeventattributes.md) object

Required: No

**startTimerFailedEventAttributes**

If the event is of type `StartTimerFailed` then this member is set and provides detailed information
about the event. It isn't set for other event types.

Type: [StartTimerFailedEventAttributes](api-starttimerfailedeventattributes.md) object

Required: No

**timerCanceledEventAttributes**

If the event is of type `TimerCanceled` then this member is set and provides detailed information
about the event. It isn't set for other event types.

Type: [TimerCanceledEventAttributes](api-timercanceledeventattributes.md) object

Required: No

**timerFiredEventAttributes**

If the event is of type `TimerFired` then this member is set and provides detailed information about
the event. It isn't set for other event types.

Type: [TimerFiredEventAttributes](api-timerfiredeventattributes.md) object

Required: No

**timerStartedEventAttributes**

If the event is of type `TimerStarted` then this member is set and provides detailed information
about the event. It isn't set for other event types.

Type: [TimerStartedEventAttributes](api-timerstartedeventattributes.md) object

Required: No

**workflowExecutionCanceledEventAttributes**

If the event is of type `WorkflowExecutionCanceled` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionCanceledEventAttributes](api-workflowexecutioncanceledeventattributes.md) object

Required: No

**workflowExecutionCancelRequestedEventAttributes**

If the event is of type `WorkflowExecutionCancelRequested` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [WorkflowExecutionCancelRequestedEventAttributes](api-workflowexecutioncancelrequestedeventattributes.md) object

Required: No

**workflowExecutionCompletedEventAttributes**

If the event is of type `WorkflowExecutionCompleted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionCompletedEventAttributes](api-workflowexecutioncompletedeventattributes.md) object

Required: No

**workflowExecutionContinuedAsNewEventAttributes**

If the event is of type `WorkflowExecutionContinuedAsNew` then this member is set and provides
detailed information about the event. It isn't set for other event types.

Type: [WorkflowExecutionContinuedAsNewEventAttributes](api-workflowexecutioncontinuedasneweventattributes.md) object

Required: No

**workflowExecutionFailedEventAttributes**

If the event is of type `WorkflowExecutionFailed` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionFailedEventAttributes](api-workflowexecutionfailedeventattributes.md) object

Required: No

**workflowExecutionSignaledEventAttributes**

If the event is of type `WorkflowExecutionSignaled` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionSignaledEventAttributes](api-workflowexecutionsignaledeventattributes.md) object

Required: No

**workflowExecutionStartedEventAttributes**

If the event is of type `WorkflowExecutionStarted` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionStartedEventAttributes](api-workflowexecutionstartedeventattributes.md) object

Required: No

**workflowExecutionTerminatedEventAttributes**

If the event is of type `WorkflowExecutionTerminated` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionTerminatedEventAttributes](api-workflowexecutionterminatedeventattributes.md) object

Required: No

**workflowExecutionTimedOutEventAttributes**

If the event is of type `WorkflowExecutionTimedOut` then this member is set and provides detailed
information about the event. It isn't set for other event types.

Type: [WorkflowExecutionTimedOutEventAttributes](api-workflowexecutiontimedouteventattributes.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/historyevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/historyevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/historyevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailWorkflowExecutionFailedEventAttributes

LambdaFunctionCompletedEventAttributes

All content copied from https://docs.aws.amazon.com/.
