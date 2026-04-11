This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Guardrail ContentFilterConfig

Contains filter strengths for harmful content. Guardrails support the following content filters to detect and filter harmful user inputs and FM-generated outputs.

- **Hate** – Describes language or a statement that discriminates, criticizes, insults, denounces, or dehumanizes a person or group on the basis of an identity (such as race, ethnicity, gender, religion, sexual orientation, ability, and national origin).

- **Insults** – Describes language or a statement that includes demeaning, humiliating, mocking, insulting, or belittling language. This type of language is also labeled as bullying.

- **Sexual** – Describes language or a statement that indicates sexual interest, activity, or arousal using direct or indirect references to body parts, physical traits, or sex.

- **Violence** – Describes language or a statement that includes glorification of or threats to inflict physical pain, hurt, or injury toward a person, group or thing.

Content filtering depends on the confidence classification of user inputs and FM
responses across each of the four harmful categories. All input and output statements are
classified into one of four confidence levels (NONE, LOW, MEDIUM, HIGH) for each
harmful category. For example, if a statement is classified as
_Hate_ with HIGH confidence, the likelihood of the statement
representing hateful content is high. A single statement can be classified across
multiple categories with varying confidence levels. For example, a single statement
can be classified as _Hate_ with HIGH confidence, _Insults_ with LOW confidence, _Sexual_ with NONE confidence, and _Violence_ with MEDIUM confidence.

For more information, see [Guardrails content filters](../../../bedrock/latest/userguide/guardrails-filters.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputAction" : String,
  "InputEnabled" : Boolean,
  "InputModalities" : [ String, ... ],
  "InputStrength" : String,
  "OutputAction" : String,
  "OutputEnabled" : Boolean,
  "OutputModalities" : [ String, ... ],
  "OutputStrength" : String,
  "Type" : String
}

```

### YAML

```yaml

  InputAction: String
  InputEnabled: Boolean
  InputModalities:
    - String
  InputStrength: String
  OutputAction: String
  OutputEnabled: Boolean
  OutputModalities:
    - String
  OutputStrength: String
  Type: String

```

## Properties

`InputAction`

Specifies the action to take when harmful content is detected. Supported values include:

- `BLOCK` – Block the content and replace it with blocked
messaging.

- `NONE` – Take no action but return detection information in the
trace response.

_Required_: No

_Type_: String

_Allowed values_: `BLOCK | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputEnabled`

Specifies whether to enable guardrail evaluation on the input. When disabled, you aren't
charged for the evaluation. The evaluation doesn't appear in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputModalities`

The input modalities selected for the guardrail content filter configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputStrength`

The strength of the content filter to apply to prompts. As you
increase the filter strength, the likelihood of filtering harmful content increases
and the probability of seeing harmful content in your application reduces.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputAction`

Specifies the action to take when harmful content is detected in the output. Supported values include:

- `BLOCK` – Block the content and replace it with blocked
messaging.

- `NONE` – Take no action but return detection information in the trace
response.

_Required_: No

_Type_: String

_Allowed values_: `BLOCK | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputEnabled`

Specifies whether to enable guardrail evaluation on the output. When disabled, you
aren't charged for the evaluation. The evaluation doesn't appear in the response.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputModalities`

The output modalities selected for the guardrail content filter configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputStrength`

The strength of the content filter to apply to model responses. As you
increase the filter strength, the likelihood of filtering harmful content increases
and the probability of seeing harmful content in your application reduces.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The harmful category that the content filter is applied to.

_Required_: Yes

_Type_: String

_Allowed values_: `SEXUAL | VIOLENCE | HATE | INSULTS | MISCONDUCT | PROMPT_ATTACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutomatedReasoningPolicyConfig

ContentFiltersTierConfig

All content copied from https://docs.aws.amazon.com/.
