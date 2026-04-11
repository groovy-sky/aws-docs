# Integration subtype reference

The following
[integration subtypes](../../../apigatewayv2/latest/api-reference/apis-apiid-integrations-integrationid.md#apis-apiid-integrations-integrationid-prop-integration-integrationsubtype) are supported for HTTP APIs.

###### Integration subtypes

- [EventBridge-PutEvents 1.0](#EventBridge-PutEvents)

- [SQS-SendMessage 1.0](#SQS-SendMessage)

- [SQS-ReceiveMessage 1.0](#SQS-ReceiveMessage)

- [SQS-DeleteMessage 1.0](#SQS-DeleteMessage)

- [SQS-PurgeQueue 1.0](#SQS-PurgeQueue)

- [AppConfig-GetConfiguration 1.0](#AppConfig-GetConfiguration)

- [Kinesis-PutRecord 1.0](#Kinesis-PutRecord)

- [StepFunctions-StartExecution 1.0](#StepFunctions-StartExecution)

- [StepFunctions-StartSyncExecution 1.0](#StepFunctions-StartSyncExecution)

- [StepFunctions-StopExecution 1.0](#StepFunctions-StopExecution)

## EventBridge-PutEvents 1.0

Sends custom events to Amazon EventBridge so that they can be matched to rules.

ParameterRequiredDetailTrueDetailTypeTrueSourceTrueTimeFalseEventBusNameFalseResourcesFalseRegionFalseTraceHeaderFalse

To learn more, see [PutEvents](../../../../reference/eventbridge/latest/apireference/api-putevents.md) in the
_Amazon EventBridge API Reference_.

## SQS-SendMessage 1.0

Delivers a message to the specified queue.

ParameterRequiredQueueUrlTrueMessageBodyTrueDelaySecondsFalseMessageAttributesFalseMessageDeduplicationIdFalseMessageGroupIdFalseMessageSystemAttributesFalseRegionFalse

To learn more, see [SendMessage](../../../../reference/awssimplequeueservice/latest/apireference/api-sendmessage.md) in the
_Amazon Simple Queue Service API Reference_.

## SQS-ReceiveMessage 1.0

Retrieves one or more messages (up to 10), from the specified queue.

ParameterRequiredQueueUrlTrueAttributeNamesFalseMaxNumberOfMessagesFalseMessageAttributeNamesFalseReceiveRequestAttemptIdFalseVisibilityTimeoutFalseWaitTimeSecondsFalseRegionFalse

To learn more, see [ReceiveMessage](../../../../reference/awssimplequeueservice/latest/apireference/api-receivemessage.md) in
the _Amazon Simple Queue Service API Reference_.

## SQS-DeleteMessage 1.0

Deletes the specified message from the specified queue.

ParameterRequiredReceiptHandleTrueQueueUrlTrueRegionFalse

To learn more, see [DeleteMessage](../../../../reference/awssimplequeueservice/latest/apireference/api-deletemessage.md) in the
_Amazon Simple Queue Service API Reference_.

## SQS-PurgeQueue 1.0

Deletes all messages in the specified queue.

ParameterRequiredQueueUrlTrueRegionFalse

To learn more, see [PurgeQueue](../../../../reference/awssimplequeueservice/latest/apireference/api-purgequeue.md) in the
_Amazon Simple Queue Service API Reference_.

## AppConfig-GetConfiguration 1.0

Receive information about a configuration.

ParameterRequiredApplicationTrueEnvironmentTrueConfigurationTrueClientIdTrueClientConfigurationVersionFalseRegionFalse

To learn more, see [GetConfiguration](../../../../reference/appconfig/2019-10-09/apireference/api-getconfiguration.md) in
the _AWS AppConfig API Reference_.

## Kinesis-PutRecord 1.0

Writes a single data record into an Amazon Kinesis data stream.

ParameterRequiredStreamNameTrueDataTruePartitionKeyTrueSequenceNumberForOrderingFalseExplicitHashKeyFalseRegionFalse

To learn more, see [PutRecord](../../../../reference/kinesis/latest/apireference/api-putrecord.md) in the
_Amazon Kinesis Data Streams API Reference_.

## StepFunctions-StartExecution 1.0

Starts a state machine execution.

ParameterRequiredStateMachineArnTrueNameFalseInputFalseRegionFalse

To learn more, see [StartExecution](../../../../reference/step-functions/latest/apireference/api-startexecution.md) in
the _AWS Step Functions API Reference_.

## StepFunctions-StartSyncExecution 1.0

Starts a synchronous state machine execution.

ParameterRequiredStateMachineArnTrueNameFalseInputFalseRegionFalseTraceHeaderFalse

To learn more, see [StartSyncExecution](../../../../reference/step-functions/latest/apireference/api-startsyncexecution.md) in the
_AWS Step Functions API Reference_.

## StepFunctions-StopExecution 1.0

Stops an execution.

ParameterRequiredExecutionArnTrueCauseFalseErrorFalseRegionFalse

To learn more, see [StopExecution](../../../../reference/step-functions/latest/apireference/api-stopexecution.md) in the
_AWS Step Functions API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS service integrations

Private integrations

All content copied from https://docs.aws.amazon.com/.
