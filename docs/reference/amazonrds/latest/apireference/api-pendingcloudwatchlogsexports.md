---
title: "PendingCloudwatchLogsExports"
---

# PendingCloudwatchLogsExports

A list of the log types whose configuration is still pending. In other words, these log types are in the process of being activated or deactivated.

## Contents

###### Note

In the following list, the required parameters are described first.

**LogTypesToDisable.member.N**

Log types that are in the process of being enabled. After they are enabled, these log types are exported to CloudWatch Logs.

Type: Array of strings

Required: No

**LogTypesToEnable.member.N**

Log types that are in the process of being deactivated. After they are deactivated, these log types aren't exported to CloudWatch Logs.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/PendingCloudwatchLogsExports)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/PendingCloudwatchLogsExports)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/PendingCloudwatchLogsExports)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Parameter

PendingMaintenanceAction

All content copied from https://docs.aws.amazon.com/.
