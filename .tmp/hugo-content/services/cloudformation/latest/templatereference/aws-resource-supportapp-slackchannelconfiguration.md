This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SupportApp::SlackChannelConfiguration

You can use the `AWS::SupportApp::SlackChannelConfiguration` resource to
specify your AWS account when you configure the AWS Support App. This resource includes the following information:

- The Slack channel name and ID

- The team ID in Slack

- The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role

- Whether you want the AWS Support App to notify you when your support
cases are created, updated, resolved, or reopened

- The case severity that you want to get notified for

For more information, see the following topics in the _AWS Support User Guide_:

- [AWS Support App in Slack](../../../awssupport/latest/user/aws-support-app-for-slack.md)

- [Creating AWS Support App in Slack resources with AWS CloudFormation](../../../awssupport/latest/user/creating-resources-with-cloudformation.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SupportApp::SlackChannelConfiguration",
  "Properties" : {
      "ChannelId" : String,
      "ChannelName" : String,
      "ChannelRoleArn" : String,
      "NotifyOnAddCorrespondenceToCase" : Boolean,
      "NotifyOnCaseSeverity" : String,
      "NotifyOnCreateOrReopenCase" : Boolean,
      "NotifyOnResolveCase" : Boolean,
      "TeamId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SupportApp::SlackChannelConfiguration
Properties:
  ChannelId: String
  ChannelName: String
  ChannelRoleArn: String
  NotifyOnAddCorrespondenceToCase: Boolean
  NotifyOnCaseSeverity: String
  NotifyOnCreateOrReopenCase: Boolean
  NotifyOnResolveCase: Boolean
  TeamId: String

```

## Properties

`ChannelId`

The channel ID in Slack. This ID identifies a channel within a Slack workspace.

_Required_: Yes

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelName`

The channel name in Slack. This is the channel where you invite the AWS Support App.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelRoleArn`

The Amazon Resource Name (ARN) of the IAM role for this Slack channel
configuration. The AWS Support App uses this role to perform AWS Support and Service Quotas actions on your behalf.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:iam::[0-9]{12}:role\/(.+)$`

_Minimum_: `31`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotifyOnAddCorrespondenceToCase`

Whether to get notified when a correspondence is added to your support cases.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotifyOnCaseSeverity`

The case severity for your support cases that you want to receive notifications. You
can specify `none`, `all`, or `high`.

_Required_: Yes

_Type_: String

_Allowed values_: `none | all | high`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotifyOnCreateOrReopenCase`

Whether to get notified when your support cases are created or reopened

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotifyOnResolveCase`

Whether to get notified when your support cases are resolved.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeamId`

The team ID in Slack. This ID uniquely identifies a Slack workspace.

_Required_: Yes

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

The following examples can help you create AWS Support App resources using
CloudFormation.

### Example

Use this template to create a Slack channel configuration and IAM role for your account or your organization. You must replace
the Slack workspace and channel IDs.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Example stack to create a Slack channel configuration",
    "Resources": {
        "SlackChannelConfiguration": {
            "Type": "AWS::SupportApp::SlackChannelConfiguration",
            "Properties": {
                "TeamId": "T012ABCDEFG",
                "ChannelId": "C01234A5BCD",
                "ChannelName": "cloudformationtemplatechannel",
                "NotifyOnCreateOrReopenCase": true,
                "NotifyOnAddCorrespondenceToCase": false,
                "NotifyOnResolveCase": true,
                "NotifyOnCaseSeverity": "high",
                "ChannelRoleArn": {"Fn::GetAtt" : ["AWSSupportSlackAppCFNRole", "Arn"] }
            }
        },
        "AWSSupportSlackAppCFNRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "supportapp.amazonaws.com"
                                ]
                            },
                            "Action": [
                                "sts:AssumeRole"
                            ]
                        }
                    ]
                },
                "ManagedPolicyArns": [
                    "arn:aws:iam::aws:policy/AWSSupportAppFullAccess"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Example stack to create a Slack channel configuration
Resources:
  SlackChannelConfiguration:
    Type: AWS::SupportApp::SlackChannelConfiguration
    Properties:
      TeamId: T012ABCDEFG
      ChannelId: C01234A5BCD
      ChannelName: cfntemplatechannel
      NotifyOnCreateOrReopenCase: true
      NotifyOnAddCorrespondenceToCase: false
      NotifyOnResolveCase: true
      NotifyOnCaseSeverity: high
      ChannelRoleArn: !GetAtt 'AWSSupportSlackAppCFNRole.Arn'
  AWSSupportSlackAppCFNRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - supportapp.amazonaws.com
            Action:
              - sts:AssumeRole
      ManagedPolicyArns:
        - arn:aws:iam::aws:policy/AWSSupportAppFullAccess

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SupportApp::AccountAlias

AWS::SupportApp::SlackWorkspaceConfiguration

All content copied from https://docs.aws.amazon.com/.
