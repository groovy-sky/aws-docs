This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPNConnection CloudwatchLogOptionsSpecification

Options for sending VPN tunnel logs to CloudWatch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BgpLogEnabled" : Boolean,
  "BgpLogGroupArn" : String,
  "BgpLogOutputFormat" : String,
  "LogEnabled" : Boolean,
  "LogGroupArn" : String,
  "LogOutputFormat" : String
}

```

### YAML

```yaml

  BgpLogEnabled: Boolean
  BgpLogGroupArn: String
  BgpLogOutputFormat: String
  LogEnabled: Boolean
  LogGroupArn: String
  LogOutputFormat: String

```

## Properties

`BgpLogEnabled`

Specifies whether to enable BGP logging for the VPN connection. Default value is `False`.

Valid values: `True` \| `False`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BgpLogGroupArn`

The Amazon Resource Name (ARN) of the CloudWatch log group where BGP logs will be sent.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BgpLogOutputFormat`

The desired output format for BGP logs to be sent to CloudWatch. Default format is `json`.

Valid values: `json` \| `text`

_Required_: No

_Type_: String

_Allowed values_: `json | text`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogEnabled`

Enable or disable VPN tunnel logging feature. Default value is `False`.

Valid values: `True` \| `False`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupArn`

The Amazon Resource Name (ARN) of the CloudWatch log group to send logs to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogOutputFormat`

Set log format. Default format is `json`.

Valid values: `json` \| `text`

_Required_: No

_Type_: String

_Allowed values_: `json | text`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPNConnection

IKEVersionsRequestListValue

All content copied from https://docs.aws.amazon.com/.
