This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint VariantProperty

Specifies a production variant property type for an Endpoint.

If you are updating an Endpoint with the [RetainAllVariantProperties](../../../sagemaker/latest/dg/api-updateendpoint.md#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) option set to `true`, the `VarientProperty` objects listed
in [ExcludeRetainedVariantProperties](../../../sagemaker/latest/dg/api-updateendpoint.md#SageMaker-UpdateEndpoint-request-ExcludeRetainedVariantProperties) override the existing variant properties of the Endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VariantPropertyType" : String
}

```

### YAML

```yaml

  VariantPropertyType: String

```

## Properties

`VariantPropertyType`

The type of variant property. The supported values are:

- `DesiredInstanceCount`: Overrides the existing variant instance counts using the [InitialInstanceCount](../../../sagemaker/latest/dg/api-productionvariant.md#SageMaker-Type-ProductionVariant-InitialInstanceCount) values in the [ProductionVariants](../../../sagemaker/latest/dg/api-createendpointconfig.md#SageMaker-CreateEndpointConfig-request-ProductionVariants).

- `DesiredWeight`: Overrides the existing variant weights using the [InitialVariantWeight](../../../sagemaker/latest/dg/api-productionvariant.md#SageMaker-Type-ProductionVariant-InitialVariantWeight) values in the [ProductionVariants](../../../sagemaker/latest/dg/api-createendpointconfig.md#SageMaker-CreateEndpointConfig-request-ProductionVariants).

- `DataCaptureConfig`: (Not currently supported.)

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrafficRoutingConfig

AWS::SageMaker::EndpointConfig

All content copied from https://docs.aws.amazon.com/.
