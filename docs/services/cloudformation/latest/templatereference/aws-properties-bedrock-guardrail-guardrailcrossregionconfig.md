This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail GuardrailCrossRegionConfig

The system-defined guardrail profile that you're using with your guardrail. Guardrail profiles define the destination AWS Regions where guardrail inference requests can be automatically routed. Using guardrail profiles helps maintain guardrail performance and reliability when demand increases.

For more information, see the [Amazon Bedrock User Guide](../../../bedrock/latest/userguide/guardrails-cross-region.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GuardrailProfileArn" : String
}

```

### YAML

```yaml

  GuardrailProfileArn: String

```

## Properties

`GuardrailProfileArn`

The Amazon Resource Name (ARN) of the guardrail profile that your guardrail is using. Guardrail profile availability depends on your current AWS Region. For more information, see the [Amazon Bedrock User Guide](../../../bedrock/latest/userguide/guardrails-cross-region-support.md).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:guardrail-profile/[a-z0-9-]+[.]{1}guardrail[.]{1}v[0-9:]+$`

_Minimum_: `15`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContextualGroundingPolicyConfig

ManagedWordsConfig

All content copied from https://docs.aws.amazon.com/.
