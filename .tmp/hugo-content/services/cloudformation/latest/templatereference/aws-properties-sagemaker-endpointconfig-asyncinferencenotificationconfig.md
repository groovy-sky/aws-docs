This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig AsyncInferenceNotificationConfig

Specifies the configuration for notifications of inference results for asynchronous inference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorTopic" : String,
  "IncludeInferenceResponseIn" : [ String, ... ],
  "SuccessTopic" : String
}

```

### YAML

```yaml

  ErrorTopic: String
  IncludeInferenceResponseIn:
    - String
  SuccessTopic: String

```

## Properties

`ErrorTopic`

Amazon SNS topic to post a notification to when an inference fails. If no topic is provided, no notification is
sent on failure.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IncludeInferenceResponseIn`

The Amazon SNS topics where you want the inference response to be included.

###### Note

The inference response is included only if the response size is less than or equal
to 128 KB.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SuccessTopic`

Amazon SNS topic to post a notification to when an inference completes successfully. If no topic is provided, no
notification is sent on success.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AsyncInferenceConfig

AsyncInferenceOutputConfig

All content copied from https://docs.aws.amazon.com/.
