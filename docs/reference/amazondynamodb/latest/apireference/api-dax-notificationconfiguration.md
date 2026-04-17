---
title: "NotificationConfiguration"
---

# NotificationConfiguration

Describes a notification topic and its status. Notification topics are used for
publishing DAX events to subscribers using Amazon Simple Notification Service
(SNS).

## Contents

###### Note

In the following list, the required parameters are described first.

**TopicArn**

The Amazon Resource Name (ARN) that identifies the topic.

Type: String

Required: No

**TopicStatus**

The current state of the topic. A value of “active” means that notifications will
be sent to the topic. A value of “inactive” means that notifications will not be sent to
the topic.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dax-2017-04-19/NotificationConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dax-2017-04-19/NotificationConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dax-2017-04-19/NotificationConfiguration)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeTypeSpecificValue

Parameter

All content copied from https://docs.aws.amazon.com/.
