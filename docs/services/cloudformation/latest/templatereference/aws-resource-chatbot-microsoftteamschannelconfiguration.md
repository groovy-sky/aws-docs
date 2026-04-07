This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::MicrosoftTeamsChannelConfiguration

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](https://docs.aws.amazon.com/chatbot/latest/adminguide/service-rename.html)

`Type` attribute values remain unchanged.

The `AWS::Chatbot::MicrosoftTeamsChannelConfiguration` resource configures a Microsoft Teams channel to allow users to use Amazon Q Developer with CloudFormation templates.

This resource requires some setup to be done in the Amazon Q Developer in chat applications console. To provide the required Microsoft Teams team and tenant IDs, you must perform the initial authorization flow with
Microsoft Teams in the Amazon Q Developer in chat applications console, then copy and paste the IDs from the console.
For more details, see steps 1-3 in [Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the _Amazon Q Developer in chat applications Administrator Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Chatbot::MicrosoftTeamsChannelConfiguration",
  "Properties" : {
      "ConfigurationName" : String,
      "CustomizationResourceArns" : [ String, ... ],
      "GuardrailPolicies" : [ String, ... ],
      "IamRoleArn" : String,
      "LoggingLevel" : String,
      "SnsTopicArns" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TeamId" : String,
      "TeamsChannelId" : String,
      "TeamsChannelName" : String,
      "TeamsTenantId" : String,
      "UserRoleRequired" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Chatbot::MicrosoftTeamsChannelConfiguration
Properties:
  ConfigurationName: String
  CustomizationResourceArns:
    - String
  GuardrailPolicies:
    - String
  IamRoleArn: String
  LoggingLevel: String
  SnsTopicArns:
    - String
  Tags:
    - Tag
  TeamId: String
  TeamsChannelId: String
  TeamsChannelName: String
  TeamsTenantId: String
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

Links a list of resource ARNs (for example, custom action ARNs) to a Microsoft Teams channel configuration for Amazon Q Developer.

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

This is a user-defined role that Amazon Q Developer will assume. This is not the service-linked role. For more information, see [IAM Policies for Amazon Q Developer in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html).

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

`SnsTopicArns`

The ARNs of the SNS topics that deliver notifications to Amazon Q Developer.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the configuration.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-chatbot-microsoftteamschannelconfiguration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeamId`

The ID of the Microsoft Team authorized with Amazon Q Developer.

To get the team ID, you must perform the initial authorization flow with Microsoft Teams in the Amazon Q Developer in chat applications console. Then you can copy and paste the team ID from the console.
For more details, see steps 1-3 in [Tutorial: Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html) in the _Amazon Q Developer in chat applications Administrator Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TeamsChannelId`

The ID of the Microsoft Teams channel.

To get the channel ID, open Microsoft Teams, right click on the channel name in the left pane, then choose **Copy**. An example of the channel ID syntax is: `19%3ab6ef35dc342d56ba5654e6fc6d25a071%40thread.tacv2`.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9-_=+/.,])*%3[aA]([a-zA-Z0-9-_=+/.,])*%40([a-zA-Z0-9-_=+/.,])*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeamsChannelName`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^(.*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeamsTenantId`

The ID of the Microsoft Teams tenant.

To get the tenant ID, you must perform the initial authorization flow with Microsoft Teams in the Amazon Q Developer in chat applications console. Then you can copy and paste the tenant ID from the console.
For more details, see steps 1-3 in [Tutorial: Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html) in the _Amazon Q Developer in chat applications Administrator Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

- _I don't have a teams, tenant, or channel ID._

If you don't have one or any of the aforementioned IDs, you must perform the initial authorization flow in the Amazon Q Developer in chat applications console. Then you will be able to copy and paste the IDs from the console. For more details, see steps 1-3 in [Tutorial: Get started with Microsoft Teams](https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup) in the _Amazon Q Developer in chat applications Administrator Guide_.

- _I have already done the initial authorization for my workspace. Do I need to do it again?_

No, you can use your existing workspace. You must log into the Amazon Q Developer in chat applications console to get the workspace ID.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
