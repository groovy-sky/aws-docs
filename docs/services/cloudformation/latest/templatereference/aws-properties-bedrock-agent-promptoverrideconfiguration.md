This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent PromptOverrideConfiguration

Contains configurations to override prompts in different parts of an agent sequence. For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OverrideLambda" : String,
  "PromptConfigurations" : [ PromptConfiguration, ... ]
}

```

### YAML

```yaml

  OverrideLambda: String
  PromptConfigurations:
    - PromptConfiguration

```

## Properties

`OverrideLambda`

The ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence. If you specify this field, at least one of the `promptConfigurations` must contain a `parserMode` value that is set to `OVERRIDDEN`. For more information, see [Parser Lambda function in Amazon Bedrock Agents](https://docs.aws.amazon.com/bedrock/latest/userguide/lambda-parser.html).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:lambda:[a-z0-9-]{1,20}:\d{12}:function:[a-zA-Z0-9-_\.]+(:(\$LATEST|[a-zA-Z0-9-_]+))?$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptConfigurations`

Contains configurations to override a prompt template in one part of an agent sequence. For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html).

_Required_: Yes

_Type_: Array of [PromptConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-promptconfiguration.html)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PromptConfiguration

S3Identifier
