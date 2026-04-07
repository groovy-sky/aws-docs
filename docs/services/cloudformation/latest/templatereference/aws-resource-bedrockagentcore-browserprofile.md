This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::BrowserProfile

Specifies a browser profile for Amazon Bedrock AgentCore. A browser profile stores persistent browser session data, including cookies, local storage, session storage, and browsing history, enabling AI agents to maintain authenticated sessions and reuse browser state across multiple browser sessions.

For more information, see [Manage browser profiles in Amazon Bedrock AgentCore](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/browser-profiles.html).

See the **Properties** section below for descriptions of both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::BrowserProfile",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::BrowserProfile
Properties:
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the browser profile.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the browser profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the browser profile.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the browser profile. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:browser-profile/MyBrowserProfile-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the browser profile was created.

`LastSavedAt`

The timestamp when browser session data was last saved to this profile.

`LastSavedBrowserId`

The identifier of the browser from which data was last saved to this profile.

`LastSavedBrowserSessionId`

The identifier of the browser session from which data was last saved to this profile.

`LastUpdatedAt`

The timestamp when the browser profile was last updated.

`ProfileArn`

The Amazon Resource Name (ARN) of the browser profile.

`ProfileId`

The unique identifier of the browser profile.

`Status`

The current status of the browser profile.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

AWS::BedrockAgentCore::CodeInterpreterCustom
