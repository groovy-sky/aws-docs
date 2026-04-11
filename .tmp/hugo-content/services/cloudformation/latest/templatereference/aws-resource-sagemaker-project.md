This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Project

Creates a machine learning (ML) project that can contain one or more templates that set
up an ML pipeline from training to deploying an approved model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Project",
  "Properties" : {
      "ProjectDescription" : String,
      "ProjectName" : String,
      "ServiceCatalogProvisionedProductDetails" : ServiceCatalogProvisionedProductDetails,
      "ServiceCatalogProvisioningDetails" : ServiceCatalogProvisioningDetails,
      "Tags" : [ Tag, ... ],
      "TemplateProviderDetails" : [ TemplateProviderDetail, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Project
Properties:
  ProjectDescription: String
  ProjectName: String
  ServiceCatalogProvisionedProductDetails:
    ServiceCatalogProvisionedProductDetails
  ServiceCatalogProvisioningDetails:
    ServiceCatalogProvisioningDetails
  Tags:
    - Tag
  TemplateProviderDetails:
    - TemplateProviderDetail

```

## Properties

`ProjectDescription`

The description of the project.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectName`

The name of the project.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceCatalogProvisionedProductDetails`

Details of a provisioned service catalog product. For information about service catalog,
see [What is AWS Service\
Catalog](../../../servicecatalog/latest/adminguide/introduction.md).

_Required_: No

_Type_: [ServiceCatalogProvisionedProductDetails](aws-properties-sagemaker-project-servicecatalogprovisionedproductdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceCatalogProvisioningDetails`

The product ID and provisioning artifact ID to provision a service catalog. For information, see [What is AWS Service\
Catalog](../../../servicecatalog/latest/adminguide/introduction.md).

_Required_: No

_Type_: [ServiceCatalogProvisioningDetails](aws-properties-sagemaker-project-servicecatalogprovisioningdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of key-value pairs to apply to this resource.

For more information, see [Resource Tag](../userguide/aws-properties-resource-tags.md) and [Using Cost\
Allocation Tags](../../../awsaccountbilling/latest/aboutv2/cost-alloc-tags.md#allocation-what) in the _AWS Billing and Cost Management User_
_Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-project-tag.md)

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateProviderDetails`

An array of template providers associated with the project.

_Required_: No

_Type_: Array of [TemplateProviderDetail](aws-properties-sagemaker-project-templateproviderdetail.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the Project.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time that the project was created.

`ProjectArn`

The Amazon Resource Name (ARN) of the project.

`ProjectId`

The ID of the project. This ID is prepended to all entities associated with this project.

`ProjectStatus`

The status of the project.

## Examples

### SageMaker Project Example

The following example creates a SageMaker Project.

#### JSON

```json

{
   "Description": "AWS SageMaker Project basic template",
   "Resources": {
      "SampleProject": {
         "Type": "AWS::SageMaker::Project",
         "Properties": {
            "ProjectName": "project1",
            "ProjectDescription": "Project Description",
            "ServiceCatalogProvisioningDetails": {
               "ProductId": "prod-53ibyqbj2cgmo",
               "ProvisioningArtifactId": "pa-sm4pjfuzictpe"
            }
         }
      }
   }
}
```

#### YAML

```yaml

---
Description: AWS SageMaker Project basic template

Resources:

  SampleProject:
    Type: AWS::SageMaker::Project
    Properties:
      ProjectName: "SampleProject"
      ProjectDescription: "Project Description"
      ServiceCatalogProvisioningDetails:
        ProductId: "prod-53ibyqbj2cgmo"
        ProvisioningArtifactId: "pa-sm4pjfuzictpe"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

CfnStackParameter

All content copied from https://docs.aws.amazon.com/.
