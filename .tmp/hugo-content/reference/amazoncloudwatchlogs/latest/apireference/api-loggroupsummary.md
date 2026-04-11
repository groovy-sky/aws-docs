# LogGroupSummary

This structure contains information about one log group in your account.

## Contents

**logGroupArn**

The Amazon Resource Name (ARN) of the log group.

Type: String

Required: No

**logGroupClass**

The log group class for this log group. For details about the features supported by each
log group class, see [Log classes](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-log-classes.md)

Type: String

Valid Values: `STANDARD | INFREQUENT_ACCESS | DELIVERY`

Required: No

**logGroupName**

The name of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/loggroupsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/loggroupsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/loggroupsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogGroupField

LogStream

All content copied from https://docs.aws.amazon.com/.
