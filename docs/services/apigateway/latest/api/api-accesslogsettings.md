# AccessLogSettings

Access log settings, including the access log format and access log destination ARN.

## Contents

**destinationArn**

The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with `amazon-apigateway-`.

Type: String

Required: No

**format**

A single line format of the access logs of data, as specified by selected $context variables. The format must include at least `$context.requestId`.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/accesslogsettings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/accesslogsettings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/accesslogsettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

ApiKey

All content copied from https://docs.aws.amazon.com/.
