---
title: "ExportSummary"
---

# ExportSummary

Summary information about an export task.

## Contents

###### Note

In the following list, the required parameters are described first.

**ExportArn**

The Amazon Resource Name (ARN) of the export.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**ExportStatus**

Export can be in one of the following states: IN\_PROGRESS, COMPLETED, or
FAILED.

Type: String

Valid Values: `IN_PROGRESS | COMPLETED | FAILED`

Required: No

**ExportType**

The type of export that was performed. Valid values are `FULL_EXPORT` or
`INCREMENTAL_EXPORT`.

Type: String

Valid Values: `FULL_EXPORT | INCREMENTAL_EXPORT`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/ExportSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/ExportSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/ExportSummary)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportDescription

FailureException

All content copied from https://docs.aws.amazon.com/.
