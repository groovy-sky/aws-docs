# CloudWatchLogOptionsSpecification

Options for sending VPN tunnel logs to CloudWatch.

## Contents

**BgpLogEnabled**

Specifies whether to enable BGP logging for the VPN connection. Default value is `False`.

Valid values: `True` \| `False`

Type: Boolean

Required: No

**BgpLogGroupArn**

The Amazon Resource Name (ARN) of the CloudWatch log group where BGP logs will be sent.

Type: String

Required: No

**BgpLogOutputFormat**

The desired output format for BGP logs to be sent to CloudWatch. Default format is `json`.

Valid values: `json` \| `text`

Type: String

Required: No

**LogEnabled**

Enable or disable VPN tunnel logging feature. Default value is `False`.

Valid values: `True` \| `False`

Type: Boolean

Required: No

**LogGroupArn**

The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.

Type: String

Required: No

**LogOutputFormat**

Set log format. Default format is `json`.

Valid values: `json` \| `text`

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cloudwatchlogoptionsspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cloudwatchlogoptionsspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cloudwatchlogoptionsspecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CloudWatchLogOptions

CoipAddressUsage
