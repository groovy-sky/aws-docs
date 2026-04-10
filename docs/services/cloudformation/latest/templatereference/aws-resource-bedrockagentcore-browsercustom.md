This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::BrowserCustom

AgentCore Browser tool provides a fast, secure, cloud-based browser runtime to enable
AI agents to interact with websites at scale.

For more information about using the custom browser, see [Interact with web\
applications using Amazon Bedrock AgentCore Browser](../../../bedrock-agentcore/latest/devguide/browser-tool.md).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::BrowserCustom",
  "Properties" : {
      "BrowserSigning" : BrowserSigning,
      "Description" : String,
      "ExecutionRoleArn" : String,
      "Name" : String,
      "NetworkConfiguration" : BrowserNetworkConfiguration,
      "RecordingConfig" : RecordingConfig,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::BrowserCustom
Properties:
  BrowserSigning:
    BrowserSigning
  Description: String
  ExecutionRoleArn: String
  Name: String
  NetworkConfiguration:
    BrowserNetworkConfiguration
  RecordingConfig:
    RecordingConfig
  Tags:
    Key: Value

```

## Properties

`BrowserSigning`

The browser signing configuration that enables cryptographic agent identification using HTTP message signatures for web bot authentication.

_Required_: No

_Type_: [BrowserSigning](aws-properties-bedrockagentcore-browsercustom-browsersigning.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the browser.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the IAM role that provides permissions for the browser to access AWS services.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws(?:-cn|-us-gov|-iso(?:-[bef])?)?):iam::[0-9]{12}:role/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the browser. The name must be unique within your account.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfiguration`

The network configuration for the browser. This configuration specifies the network mode for the browser.

_Required_: Yes

_Type_: [BrowserNetworkConfiguration](aws-properties-bedrockagentcore-browsercustom-browsernetworkconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecordingConfig`

The recording configuration for the browser. When enabled, browser sessions are recorded and stored in the specified Amazon S3 location.

_Required_: No

_Type_: [RecordingConfig](aws-properties-bedrockagentcore-browsercustom-recordingconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A map of tag keys and values to assign to the browser. Tags enable you to categorize your resources in different ways, for example, by purpose, owner, or environment.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the custom browser. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:browser-custom/MyBrowser-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BrowserArn`

The ARN for the custom browser.

`BrowserId`

The ID for the custom browser.

`CreatedAt`

The time at which the custom browser was created.

`FailureReason`

The reason for failure if the browser is in a failed state.

`LastUpdatedAt`

The time at which the custom browser was last updated.

`Status`

The status of the custom browser.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BrowserNetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
