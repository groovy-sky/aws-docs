This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent

Creates an inference component, which is a SageMaker AI hosting object that you can
use to deploy a model to an endpoint. In the inference component settings, you specify the
model, the endpoint, and how the model utilizes the resources that the endpoint hosts. You
can optimize resource utilization by tailoring how the required CPU cores, accelerators,
and memory are allocated. You can deploy multiple inference components to an endpoint,
where each inference component contains one model and the resource utilization needs for
that individual model. After you deploy an inference component, you can directly invoke the
associated model when you use the InvokeEndpoint API action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::InferenceComponent",
  "Properties" : {
      "DeploymentConfig" : InferenceComponentDeploymentConfig,
      "EndpointArn" : String,
      "EndpointName" : String,
      "InferenceComponentName" : String,
      "RuntimeConfig" : InferenceComponentRuntimeConfig,
      "Specification" : InferenceComponentSpecification,
      "Tags" : [ Tag, ... ],
      "VariantName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::InferenceComponent
Properties:
  DeploymentConfig:
    InferenceComponentDeploymentConfig
  EndpointArn: String
  EndpointName: String
  InferenceComponentName: String
  RuntimeConfig:
    InferenceComponentRuntimeConfig
  Specification:
    InferenceComponentSpecification
  Tags:
    - Tag
  VariantName: String

```

## Properties

`DeploymentConfig`

The deployment configuration for an endpoint, which contains the desired deployment
strategy and rollback configurations.

_Required_: No

_Type_: [InferenceComponentDeploymentConfig](aws-properties-sagemaker-inferencecomponent-inferencecomponentdeploymentconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointArn`

The Amazon Resource Name (ARN) of the endpoint that hosts the inference component.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointName`

The name of the endpoint that hosts the inference component.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceComponentName`

The name of the inference component.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeConfig`

Runtime settings for the inference component, including the number of copies to deploy.

_Required_: No

_Type_: [InferenceComponentRuntimeConfig](aws-properties-sagemaker-inferencecomponent-inferencecomponentruntimeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Specification`

The specification for the inference component, including the model and container configuration.

_Required_: Yes

_Type_: [InferenceComponentSpecification](aws-properties-sagemaker-inferencecomponent-inferencecomponentspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-inferencecomponent-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariantName`

The name of the production variant that hosts the inference component.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the inference component.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time when the inference component was created.

`FailureReason`

The reason why the inference component failed, if applicable.

`InferenceComponentArn`

The Amazon Resource Name (ARN) of the inference component.

`InferenceComponentStatus`

The status of the inference component.

`LastModifiedTime`

The time when the inference component was last updated.

`RuntimeConfig.CurrentCopyCount`

The number of runtime copies of the model container that are currently deployed.

`RuntimeConfig.DesiredCopyCount`

The number of runtime copies of the model container that you requested to deploy with
the inference component.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::ImageVersion

Alarm

All content copied from https://docs.aws.amazon.com/.
