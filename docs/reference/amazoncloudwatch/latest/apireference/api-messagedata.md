# MessageData

A message returned by the `GetMetricData` API, including a code and a
description.

If a cross-Region `GetMetricData` operation fails with a code of
`Forbidden` and a value of `Authentication too complex to retrieve
                cross region data`, you can correct the problem by running the
`GetMetricData` operation in the same Region where the metric data
is.

## Contents

**Code**

The error code or status code associated with the message.

Type: String

Required: No

**Value**

The message text.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/messagedata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/messagedata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/messagedata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedRuleState

Metric

All content copied from https://docs.aws.amazon.com/.
