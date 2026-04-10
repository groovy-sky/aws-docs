This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model ModelAccessConfig

The access configuration file to control access to the ML model. You can explicitly accept the model
end-user license agreement (EULA) within the `ModelAccessConfig`.

- If you are a Jumpstart user, see the [End-user license agreements](../../../sagemaker/latest/dg/jumpstart-foundation-models-choose.md#jumpstart-foundation-models-choose-eula) section for more details on accepting the EULA.

- If you are an AutoML user, see the _Optional Parameters_ section of
_Create an AutoML job to fine-tune text generation models using the_
_API_ for details on [How to set the EULA acceptance when fine-tuning a model using the AutoML\
API](../../../sagemaker/latest/dg/autopilot-create-experiment-finetune-llms.md#autopilot-llms-finetuning-api-optional-params).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AcceptEula" : Boolean
}

```

### YAML

```yaml

  AcceptEula: Boolean

```

## Properties

`AcceptEula`

Specifies agreement to the model end-user license agreement (EULA). The
`AcceptEula` value must be explicitly defined as `True` in order
to accept the EULA that this model requires. You are responsible for reviewing and
complying with any applicable license terms and making sure they are acceptable for your
use case before downloading or using a model.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceExecutionConfig

ModelDataSource

All content copied from https://docs.aws.amazon.com/.
