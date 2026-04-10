This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::SlackChannelConfiguration

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](../../../chatbot/latest/adminguide/service-rename.md)

`Type` attribute values remain unchanged.

The `AWS::Chatbot::SlackChannelConfiguration` resource configures a Slack channel to allow users to use Amazon Q Developer with CloudFormation templates.

This resource requires some setup to be done in the Amazon Q Developer in chat applications console. To provide the required Slack workspace ID, you must perform the initial authorization flow with
Slack in the Amazon Q Developer in chat applications console, then copy and paste the workspace ID from the console.
For more details, see steps 1-3 in [Tutorial: Get started with Slack](../../../chatbot/latest/adminguide/slack-setup.md#slack-client-setup) in the _Amazon Q Developer in chat applications User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Chatbot::SlackChannelConfiguration",
  "Properties" : {
      "ConfigurationName" : String,
      "CustomizationResourceArns" : [ String, ... ],
      "GuardrailPolicies" : [ String, ... ],
      "IamRoleArn" : String,
      "LoggingLevel" : String,
      "SlackChannelId" : String,
      "SlackWorkspaceId" : String,
      "SnsTopicArns" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "UserRoleRequired" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Chatbot::SlackChannelConfiguration
Properties:
  ConfigurationName: String
  CustomizationResourceArns:
    - String
  GuardrailPolicies:
    - String
  IamRoleArn: String
  LoggingLevel: String
  SlackChannelId: String
  SlackWorkspaceId: String
  SnsTopicArns:
    - String
  Tags:
    - Tag
  UserRoleRequired: Boolean

```

## Properties

`ConfigurationName`

The name of the configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomizationResourceArns`

Links a list of resource ARNs (for example, custom action ARNs) to a Slack channel configuration for Amazon Q Developer.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuardrailPolicies`

The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

The ARN of the IAM role that defines the permissions for Amazon Q Developer.

This is a user-defined role that Amazon Q Developer will assume. This is not the service-linked role. For more information, see [IAM Policies for Amazon Q Developerin chat applications](../../../chatbot/latest/adminguide/chatbot-iam-policies.md).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.-]{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingLevel`

Specifies the logging level for this configuration. This property affects the log entries pushed to Amazon CloudWatch Logs.

Logging levels include `ERROR`, `INFO`, or `NONE`.

_Required_: No

_Type_: String

_Pattern_: `^(ERROR|INFO|NONE)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlackChannelId`

The ID of the Slack channel.

To get the ID, open Slack, right click on the channel name in the left pane, then choose Copy Link. The channel ID is the character string at the end of the URL. For example, `ABCBBLZZZ`.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlackWorkspaceId`

The ID of the Slack workspace authorized with Amazon Q Developer.

To get the workspace ID, you must perform the initial authorization flow with Slack in the Amazon Q Developer in chat applications console. Then you can copy and paste the workspace ID from the console.
For more details, see steps 1-3 in [Tutorial: Get started with Slack](../../../chatbot/latest/adminguide/slack-setup.md#slack-client-setup) in the _Amazon Q Developer in chat applications User Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Z]{1,255}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnsTopicArns`

The ARNs of the SNS topics that deliver notifications to Amazon Q Developer.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-chatbot-slackchannelconfiguration-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserRoleRequired`

Enables use of a user role requirement in your chat configuration.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

### Fn::GetAtt

`Arn`

The ARN of the resource.

## Remarks

Common troubleshooting scenarios:

- _I don't have a workspace ID._

If you don't have a workspace ID, you must perform the initial authorization flow in the Amazon Q Developer in chat applications console. Then you will be able to copy and paste the workspace ID from the console. For more details, see steps 1-3 in [Tutorial: Get started with Slack](../../../chatbot/latest/adminguide/slack-setup.md#slack-client-setup) in the _Amazon Q Developer in chat applications Administrator Guide_.

- _I have already done the initial authorization for my workspace. Do I need to do it again?_

No, you can use your existing workspace. You must log into the Amazon Q Developer in chat applications console to get the workspace ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
