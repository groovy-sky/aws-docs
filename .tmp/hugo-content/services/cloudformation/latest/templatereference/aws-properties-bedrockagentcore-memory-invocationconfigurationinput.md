This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory InvocationConfigurationInput

The memory invocation configuration input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PayloadDeliveryBucketName" : String,
  "TopicArn" : String
}

```

### YAML

```yaml

  PayloadDeliveryBucketName: String
  TopicArn: String

```

## Properties

`PayloadDeliveryBucketName`

The message invocation configuration information for the bucket name.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicArn`

The memory trigger condition topic Amazon Resource Name (ARN).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws(?:-cn|-us-gov|-iso(?:-[bef])?)?):[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EpisodicReflectionConfigurationInput

KinesisResource

All content copied from https://docs.aws.amazon.com/.
