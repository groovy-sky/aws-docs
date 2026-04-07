This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AIOps::InvestigationGroup ChatbotNotificationChannel

Use this structure to integrate CloudWatch investigations with chat applications. This
structure is a string array. For the first string, specify the ARN of an Amazon SNS
topic. For the array of strings, specify the ARNs of one or more chat applications configurations that you want to associate with that topic. For more
information about these configuration ARNs, see [Getting started with Amazon Q in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/getting-started.html) and [Resource type defined by AWS Chatbot](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awschatbot.html#awschatbot-resources-for-iam-policies).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChatConfigurationArns" : [ String, ... ],
  "SNSTopicArn" : String
}

```

### YAML

```yaml

  ChatConfigurationArns:
    - String
  SNSTopicArn: String

```

## Properties

`ChatConfigurationArns`

Returns the Amazon Resource Name (ARN) of any third-party chat integrations configured
for the account.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SNSTopicArn`

Returns the ARN of an Amazon SNS topic used for third-party chat integrations.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AIOps::InvestigationGroup

CrossAccountConfiguration
