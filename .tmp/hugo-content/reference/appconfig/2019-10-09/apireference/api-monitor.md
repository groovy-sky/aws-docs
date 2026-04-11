# Monitor

Amazon CloudWatch alarms to monitor during the deployment process.

## Contents

**AlarmArn**

Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**AlarmRoleArn**

ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor
`AlarmArn`.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^((arn):(aws|aws-cn|aws-iso|aws-iso-[a-z]{1}|aws-us-gov|aws-eusc):(iam)::\d{12}:role[/].*)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/monitor.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/monitor.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/monitor.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InvalidConfigurationDetail

Parameter

All content copied from https://docs.aws.amazon.com/.
