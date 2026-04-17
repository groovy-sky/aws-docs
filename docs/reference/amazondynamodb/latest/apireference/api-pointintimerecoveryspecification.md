---
title: "PointInTimeRecoverySpecification"
---

# PointInTimeRecoverySpecification

Represents the settings used to enable point in time recovery.

## Contents

###### Note

In the following list, the required parameters are described first.

**PointInTimeRecoveryEnabled**

Indicates whether point in time recovery is enabled (true) or disabled (false) on the
table.

Type: Boolean

Required: Yes

**RecoveryPeriodInDays**

The number of preceding days for which continuous backups are taken and maintained.
Your table data is only recoverable to any point-in-time from within the configured
recovery period. This parameter is optional. If no value is provided, the value will
default to 35.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 35.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/PointInTimeRecoverySpecification)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/PointInTimeRecoverySpecification)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/PointInTimeRecoverySpecification)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PointInTimeRecoveryDescription

Projection

All content copied from https://docs.aws.amazon.com/.
