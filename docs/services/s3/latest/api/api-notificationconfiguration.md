# NotificationConfiguration

A container for specifying the notification configuration of the bucket. If this element is empty,
notifications are turned off for the bucket.

## Contents

**EventBridgeConfiguration**

Enables delivery of events to Amazon EventBridge.

Type: [EventBridgeConfiguration](api-eventbridgeconfiguration.md) data type

Required: No

**LambdaFunctionConfigurations**

Describes the AWS Lambda functions to invoke and the events for which to invoke them.

Type: Array of [LambdaFunctionConfiguration](api-lambdafunctionconfiguration.md) data types

Required: No

**QueueConfigurations**

The Amazon Simple Queue Service queues to publish messages to and the events for which to publish
messages.

Type: Array of [QueueConfiguration](api-queueconfiguration.md) data types

Required: No

**TopicConfigurations**

The topic to which notifications are sent and the events for which notifications are
generated.

Type: Array of [TopicConfiguration](api-topicconfiguration.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/notificationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/notificationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/notificationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NoncurrentVersionTransition

NotificationConfigurationDeprecated

All content copied from https://docs.aws.amazon.com/.
