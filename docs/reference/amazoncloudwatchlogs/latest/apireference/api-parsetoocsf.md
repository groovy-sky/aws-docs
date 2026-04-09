# ParseToOCSF

This processor converts logs into [Open Cybersecurity Schema\
Framework (OCSF)](https://ocsf.io/) events.

For more information about this processor including examples, see [parseToOCSF](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseToOCSF) in the _CloudWatch Logs User Guide_.

## Contents

**eventSource**

Specify the service or process that produces the log events that will be converted with
this processor.

Type: String

Valid Values: `CloudTrail | Route53Resolver | VPCFlow | EKSAudit | AWSWAF`

Required: Yes

**ocsfVersion**

Specify which version of the OCSF schema to use for the transformed log events.

Type: String

Valid Values: `V1.1 | V1.5`

Required: Yes

**mappingVersion**

The version of the OCSF mapping to use for parsing log data.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10.

Pattern: `^\d+\.\d+(\.\d+)?$`

Required: No

**source**

The path to the field in the log event that you want to parse. If you omit this value, the
whole log message is parsed.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/parsetoocsf.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/parsetoocsf.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/parsetoocsf.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseRoute53

ParseVPC

All content copied from https://docs.aws.amazon.com/.
