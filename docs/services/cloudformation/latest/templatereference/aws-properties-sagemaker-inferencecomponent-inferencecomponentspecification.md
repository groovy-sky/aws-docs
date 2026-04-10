This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent InferenceComponentSpecification

Details about the resources to deploy with this inference component, including the
model, container, and compute resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseInferenceComponentName" : String,
  "ComputeResourceRequirements" : InferenceComponentComputeResourceRequirements,
  "Container" : InferenceComponentContainerSpecification,
  "ModelName" : String,
  "StartupParameters" : InferenceComponentStartupParameters
}

```

### YAML

```yaml

  BaseInferenceComponentName: String
  ComputeResourceRequirements:
    InferenceComponentComputeResourceRequirements
  Container:
    InferenceComponentContainerSpecification
  ModelName: String
  StartupParameters:
    InferenceComponentStartupParameters

```

## Properties

`BaseInferenceComponentName`

The name of an existing inference component that is to contain the inference component
that you're creating with your request.

Specify this parameter only if your request is meant to create an adapter inference
component. An adapter inference component contains the path to an adapter model. The
purpose of the adapter model is to tailor the inference output of a base foundation model,
which is hosted by the base inference component. The adapter inference component uses the
compute resources that you assigned to the base inference component.

When you create an adapter inference component, use the `Container` parameter
to specify the location of the adapter artifacts. In the parameter value, use the
`ArtifactUrl` parameter of the
`InferenceComponentContainerSpecification` data type.

Before you can create an adapter inference component, you must have an existing
inference component that contains the foundation model that you want to adapt.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputeResourceRequirements`

The compute resources allocated to run the model, plus any
adapter models, that you assign to the inference component.

Omit this parameter if your request is meant to create an adapter inference component.
An adapter inference component is loaded by a base inference component, and it uses the
compute resources of the base inference component.

_Required_: No

_Type_: [InferenceComponentComputeResourceRequirements](aws-properties-sagemaker-inferencecomponent-inferencecomponentcomputeresourcerequirements.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Container`

Defines a container that provides the runtime environment for a model that you deploy
with an inference component.

_Required_: No

_Type_: [InferenceComponentContainerSpecification](aws-properties-sagemaker-inferencecomponent-inferencecomponentcontainerspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelName`

The name of an existing SageMaker AI model object in your account that you want to
deploy with the inference component.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartupParameters`

Settings that take effect while the model container starts up.

_Required_: No

_Type_: [InferenceComponentStartupParameters](aws-properties-sagemaker-inferencecomponent-inferencecomponentstartupparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceComponentRuntimeConfig

InferenceComponentStartupParameters

All content copied from https://docs.aws.amazon.com/.
