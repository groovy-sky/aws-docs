This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SupportApp::SlackWorkspaceConfiguration

You can use the `AWS::SupportApp::SlackWorkspaceConfiguration` resource to
specify your Slack workspace configuration. This resource configures your AWS account so that you can use the specified Slack workspace in the
AWS Support App. This resource includes the following information:

- The team ID for the Slack workspace

- The version ID of the resource to use with AWS CloudFormation

For more information, see the following topics in the _AWS Support User Guide_:

- [AWS Support App in Slack](../../../awssupport/latest/user/aws-support-app-for-slack.md)

- [Creating AWS Support App in Slack resources with AWS CloudFormation](../../../awssupport/latest/user/creating-resources-with-cloudformation.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SupportApp::SlackWorkspaceConfiguration",
  "Properties" : {
      "TeamId" : String,
      "VersionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SupportApp::SlackWorkspaceConfiguration
Properties:
  TeamId: String
  VersionId: String

```

## Properties

`TeamId`

The team ID in Slack. This ID uniquely identifies a Slack workspace, such as
`T012ABCDEFG`.

_Required_: Yes

_Type_: String

_Pattern_: `^\S+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VersionId`

An identifier used to update an existing Slack workspace configuration in AWS CloudFormation, such as `100`.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Slack workspace, such as
`T012ABCDEFG`.

For the AWS Support App Slack workspace configuration, `Ref`
returns the value of the Slack workspace ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

The following examples can help you create AWS Support App resources using
CloudFormation.

### Example

Use this template to create a Slack workspace configuration for your account
or your organization. You must replace the Slack workspace ID.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Example stack to create a Slack workspace configuration",
    "Resources": {
        "SlackWorkspaceConfiguration": {
            "Type": "AWS::SupportApp::SlackWorkspaceConfiguration",
            "Properties": {
                "TeamId": "T012ABCDEFG",
                "VersionId": "1"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Example stack to create a Slack workspace configuration
Resources:
  SlackWorkspaceConfiguration:
    Type: AWS::SupportApp::SlackWorkspaceConfiguration
    Properties:
      TeamId: T012ABCDEFG
      VersionId: '1'

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SupportApp::SlackChannelConfiguration

Next

All content copied from https://docs.aws.amazon.com/.
