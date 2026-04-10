This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction PublishFindingToSnsParams

Parameters to define a mitigation action that publishes findings to Amazon SNS. You can implement your own custom actions in response to the Amazon SNS messages.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicArn" : String
}

```

### YAML

```yaml

  TopicArn: String

```

## Properties

`TopicArn`

The ARN of the topic to which you want to publish the findings.

_Required_: Yes

_Type_: String

_Minimum_: `11`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnableIoTLoggingParams

ReplaceDefaultPolicyVersionParams

All content copied from https://docs.aws.amazon.com/.
