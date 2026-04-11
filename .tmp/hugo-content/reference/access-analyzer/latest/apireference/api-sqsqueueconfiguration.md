# SqsQueueConfiguration

The proposed access control configuration for an Amazon SQS queue. You can propose a
configuration for a new Amazon SQS queue or an existing Amazon SQS queue that you own by specifying
the Amazon SQS policy. If the configuration is for an existing Amazon SQS queue and you do not
specify the Amazon SQS policy, the access preview uses the existing Amazon SQS policy for the queue.
If the access preview is for a new resource and you do not specify the policy, the access
preview assumes an Amazon SQS queue without a policy. To propose deletion of an existing Amazon SQS
queue policy, you can specify an empty string for the Amazon SQS policy. For more information
about Amazon SQS policy limits, see [Quotas related\
to policies](../../../../services/awssimplequeueservice/latest/sqsdeveloperguide/quotas-policies.md).

## Contents

**queuePolicy**

The proposed resource policy for the Amazon SQS queue.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/sqsqueueconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/sqsqueueconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/sqsqueueconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Span

StatusReason
