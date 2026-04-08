# BillingModeSummary

Contains the details for the read/write capacity mode. This page talks about
`PROVISIONED` and `PAY_PER_REQUEST` billing modes. For more
information about these modes, see [Read/write capacity mode](../../../../services/dynamodb/latest/developerguide/howitworks-readwritecapacitymode.md).

###### Note

You may need to switch to on-demand mode at least once in order to return a
`BillingModeSummary` response.

## Contents

###### Note

In the following list, the required parameters are described first.

**BillingMode**

Controls how you are charged for read and write throughput and how you manage
capacity. This setting can be changed later.

- `PROVISIONED` \- Sets the read/write capacity mode to
`PROVISIONED`. We recommend using `PROVISIONED` for
predictable workloads.

- `PAY_PER_REQUEST` \- Sets the read/write capacity mode to
`PAY_PER_REQUEST`. We recommend using
`PAY_PER_REQUEST` for unpredictable workloads.

Type: String

Valid Values: `PROVISIONED | PAY_PER_REQUEST`

Required: No

**LastUpdateToPayPerRequestDateTime**

Represents the time when `PAY_PER_REQUEST` was last set as the read/write
capacity mode.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/billingmodesummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/billingmodesummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/billingmodesummary.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BatchStatementResponse

CancellationReason
