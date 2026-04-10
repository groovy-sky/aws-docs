This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Project ServiceCatalogProvisioningDetails

Details that you specify to provision a service catalog product. For information about
service catalog, see [What is AWS Service\
Catalog](../../../servicecatalog/latest/adminguide/introduction.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PathId" : String,
  "ProductId" : String,
  "ProvisioningArtifactId" : String,
  "ProvisioningParameters" : [ ProvisioningParameter, ... ]
}

```

### YAML

```yaml

  PathId: String
  ProductId: String
  ProvisioningArtifactId: String
  ProvisioningParameters:
    - ProvisioningParameter

```

## Properties

`PathId`

The path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProductId`

The ID of the product to provision.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisioningArtifactId`

The ID of the provisioning artifact.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisioningParameters`

A list of key value pairs that you specify when you provision a product.

_Required_: No

_Type_: Array of [ProvisioningParameter](aws-properties-sagemaker-project-provisioningparameter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceCatalogProvisionedProductDetails

Tag

All content copied from https://docs.aws.amazon.com/.
