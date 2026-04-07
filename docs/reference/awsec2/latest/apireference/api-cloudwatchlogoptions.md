# CloudWatchLogOptions

Options for sending VPN tunnel logs to CloudWatch.

## Contents

**bgpLogEnabled**

Indicates whether Border Gateway Protocol (BGP) logging is enabled for the VPN connection. Default value is `False`.

Valid values: `True` \| `False`

Type: Boolean

Required: No

**bgpLogGroupArn**

The Amazon Resource Name (ARN) of the CloudWatch log group for BGP logs.

Type: String

Required: No

**bgpLogOutputFormat**

The output format for BGP logs sent to CloudWatch. Default format is `json`.

Valid values: `json` \| `text`

Type: String

Required: No

**logEnabled**

Status of VPN tunnel logging feature. Default value is `False`.

Valid values: `True` \| `False`

Type: Boolean

Required: No

**logGroupArn**

The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.

Type: String

Required: No

**logOutputFormat**

Configured log format. Default format is `json`.

Valid values: `json` \| `text`

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cloudwatchlogoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cloudwatchlogoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cloudwatchlogoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientVpnRouteStatus

CloudWatchLogOptionsSpecification
