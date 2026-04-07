This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig AsyncInferenceConfig

Specifies configuration for how an endpoint performs asynchronous inference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientConfig" : AsyncInferenceClientConfig,
  "OutputConfig" : AsyncInferenceOutputConfig
}

```

### YAML

```yaml

  ClientConfig:
    AsyncInferenceClientConfig
  OutputConfig:
    AsyncInferenceOutputConfig

```

## Properties

`ClientConfig`

Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous
inference.

_Required_: No

_Type_: [AsyncInferenceClientConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpointconfig-asyncinferenceclientconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputConfig`

Specifies the configuration for asynchronous inference invocation outputs.

_Required_: Yes

_Type_: [AsyncInferenceOutputConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-endpointconfig-asyncinferenceoutputconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AsyncInferenceClientConfig

AsyncInferenceNotificationConfig
