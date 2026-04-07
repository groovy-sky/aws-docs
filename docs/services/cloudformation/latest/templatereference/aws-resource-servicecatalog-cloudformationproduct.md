This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProduct

Specifies a product.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::CloudFormationProduct",
  "Properties" : {
      "AcceptLanguage" : String,
      "Description" : String,
      "Distributor" : String,
      "Name" : String,
      "Owner" : String,
      "ProductType" : String,
      "ProvisioningArtifactParameters" : [ ProvisioningArtifactProperties, ... ],
      "ReplaceProvisioningArtifacts" : Boolean,
      "SourceConnection" : SourceConnection,
      "SupportDescription" : String,
      "SupportEmail" : String,
      "SupportUrl" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::CloudFormationProduct
Properties:
  AcceptLanguage: String
  Description: String
  Distributor: String
  Name: String
  Owner: String
  ProductType: String
  ProvisioningArtifactParameters:
    - ProvisioningArtifactProperties
  ReplaceProvisioningArtifacts: Boolean
  SourceConnection:
    SourceConnection
  SupportDescription: String
  SupportEmail: String
  SupportUrl: String
  Tags:
    - Tag

```

## Properties

`AcceptLanguage`

The language code.

- `jp` \- Japanese

- `zh` \- Chinese

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the product.

_Required_: No

_Type_: String

_Maximum_: `8191`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Distributor`

The distributor of the product.

_Required_: No

_Type_: String

_Maximum_: `8191`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the product.

_Required_: Yes

_Type_: String

_Maximum_: `8191`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Owner`

The owner of the product.

_Required_: Yes

_Type_: String

_Maximum_: `8191`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductType`

The type of product.

_Required_: No

_Type_: String

_Allowed values_: `CLOUD_FORMATION_TEMPLATE | MARKETPLACE | TERRAFORM_OPEN_SOURCE | EXTERNAL | TERRAFORM_CLOUD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningArtifactParameters`

The configuration of the provisioning artifact (also known as a version).

_Required_: No

_Type_: Array of [ProvisioningArtifactProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicecatalog-cloudformationproduct-provisioningartifactproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplaceProvisioningArtifacts`

This property is turned off by default. If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more)
and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be
changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.

If turned on, provisioning artifacts will be given a new unique identifier when you update the product or provisioning artifacts.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceConnection`

A top level `ProductViewDetail` response containing details about the product’s connection.
AWS Service Catalog returns this field for the `CreateProduct`, `UpdateProduct`,
`DescribeProductAsAdmin`, and `SearchProductAsAdmin` APIs.
This response contains the same fields as the `ConnectionParameters` request, with the
addition of the `LastSync` response.

_Required_: No

_Type_: [SourceConnection](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicecatalog-cloudformationproduct-sourceconnection.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportDescription`

The support information about the product.

_Required_: No

_Type_: String

_Maximum_: `8191`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportEmail`

The contact email for product support.

_Required_: No

_Type_: String

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportUrl`

The contact URL for product support.

`^https?:\/\// `/ is the pattern used to validate SupportUrl.

_Required_: No

_Type_: String

_Maximum_: `2083`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-servicecatalog-cloudformationproduct-tag.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the provisioning artifact, such as
`pa-3mc34fbybfmgp`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ProductName`

The name of the product.

`ProvisioningArtifactIds`

The IDs of the provisioning artifacts.

`ProvisioningArtifactNames`

The names of the provisioning artifacts.

## See also

- [CreateProduct](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_CreateProduct.html) in
the _AWS Service Catalog API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ServiceCatalog::AcceptedPortfolioShare

CodeStarParameters
