This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig AsyncInferenceClientConfig

Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous
inference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConcurrentInvocationsPerInstance" : Integer
}

```

### YAML

```yaml

  MaxConcurrentInvocationsPerInstance: Integer

```

## Properties

`MaxConcurrentInvocationsPerInstance`

The maximum number of concurrent requests sent by the SageMaker client to the model container. If no value is
provided, SageMaker will choose an optimal value for you.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::EndpointConfig

AsyncInferenceConfig

All content copied from https://docs.aws.amazon.com/.
