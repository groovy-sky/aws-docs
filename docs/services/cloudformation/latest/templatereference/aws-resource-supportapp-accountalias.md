This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SupportApp::AccountAlias

You can use the `AWS::SupportApp::AccountAlias` resource to specify your
AWS account when you configure the AWS Support App in
Slack. Your alias name appears on the AWS Support App page in the Support Center Console and in messages from the AWS Support App. You
can use this alias to identify the account you've configured with the AWS Support App.

For more information, see [AWS Support App in Slack](https://docs.aws.amazon.com/awssupport/latest/user/aws-support-app-for-slack.html) in the _AWS Support User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SupportApp::AccountAlias",
  "Properties" : {
      "AccountAlias" : String
    }
}

```

### YAML

```yaml

Type: AWS::SupportApp::AccountAlias
Properties:
  AccountAlias: String

```

## Properties

`AccountAlias`

An alias or short name for an AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\- ]+$`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`AccountAliasResourceId`

The `AccountAlias` resource type has an attribute
`AccountAliasResourceId`. You can use this attribute to identify the
resource.

The `AccountAliasResourceId` will be
`AccountAlias_for_accountId`. In this example,
`AccountAlias_for_` is the prefix and `accountId` is your
AWS account number, such as
`AccountAlias_for_123456789012`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Support App

AWS::SupportApp::SlackChannelConfiguration
