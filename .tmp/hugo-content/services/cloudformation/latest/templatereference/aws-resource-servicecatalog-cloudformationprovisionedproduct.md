This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProvisionedProduct

Provisions the specified product.

A provisioned product is a resourced instance of a product. For example, provisioning
a product based on a AWS CloudFormation template launches a AWS CloudFormation stack and its
underlying resources. You can check the status of this request using [DescribeRecord](../../../servicecatalog/latest/dg/api-describerecord.md).

If the request contains a tag key with an empty list of values, there is a tag
conflict for that key. Do not include conflicted keys as tags, or this causes the error
"Parameter validation failed: Missing required parameter in
Tags\[ _N_\]: _Value_".

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::CloudFormationProvisionedProduct",
  "Properties" : {
      "AcceptLanguage" : String,
      "NotificationArns" : [ String, ... ],
      "PathId" : String,
      "PathName" : String,
      "ProductId" : String,
      "ProductName" : String,
      "ProvisionedProductName" : String,
      "ProvisioningArtifactId" : String,
      "ProvisioningArtifactName" : String,
      "ProvisioningParameters" : [ ProvisioningParameter, ... ],
      "ProvisioningPreferences" : ProvisioningPreferences,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::CloudFormationProvisionedProduct
Properties:
  AcceptLanguage: String
  NotificationArns:
    - String
  PathId: String
  PathName: String
  ProductId: String
  ProductName: String
  ProvisionedProductName: String
  ProvisioningArtifactId: String
  ProvisioningArtifactName: String
  ProvisioningParameters:
    - ProvisioningParameter
  ProvisioningPreferences:
    ProvisioningPreferences
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

_Allowed values_: `en | jp | zh`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationArns`

Passed to AWS CloudFormation. The SNS topic ARNs to which to publish stack-related
events.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PathId`

The path identifier of the product. This value is optional if the product has a
default path, and required if the product has more than one path. To list the paths for a
product, use [ListLaunchPaths](../../../servicecatalog/latest/dg/api-listlaunchpaths.md).

###### Note

You must provide the name or ID, but not both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PathName`

The name of the path. This value is optional if the product has a
default path, and required if the product has more than one path. To list the paths for a
product, use [ListLaunchPaths](../../../servicecatalog/latest/dg/api-listlaunchpaths.md).

###### Note

You must provide the name or ID, but not both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductId`

The product identifier.

###### Note

You must specify either the ID or the name of the product,
but not both.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductName`

The name of the Service Catalog product.

Each time a stack is created or updated, if `ProductName` is provided it will
successfully resolve to `ProductId` as long as only one product exists in the
account or Region with that `ProductName`.

###### Note

You must specify either
the name or the ID of the product, but not both.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedProductName`

A user-friendly name for the provisioned product. This value must be
unique for the AWS account and cannot be updated after the product is provisioned.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisioningArtifactId`

The identifier of the provisioning artifact (also known as a version).

###### Note

You must specify either the ID or the name of the provisioning artifact, but not both.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningArtifactName`

The name of the provisioning artifact (also known as a version) for the product. This
name must be unique for the product.

###### Note

You must specify either the name or the ID of the provisioning artifact, but not both. You must also specify either the name or the ID of the product, but not both.

_Required_: Conditional

_Type_: String

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningParameters`

Parameters specified by the administrator that are required for provisioning the
product.

_Required_: No

_Type_: Array of [ProvisioningParameter](aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningPreferences`

StackSet preferences that are required for provisioning the product or updating a provisioned product.

_Required_: No

_Type_: [ProvisioningPreferences](aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags.

###### Note

Requires the provisioned product to have an [ResourceUpdateConstraint](../userguide/aws-resource-servicecatalog-resourceupdateconstraint.md) resource with
`TagUpdatesOnProvisionedProduct` set to `ALLOWED` to allow tag
updates. If `RESOURCE_UPDATE` constraint is not present, tags updates are ignored.

_Required_: No

_Type_: Array of [Tag](aws-properties-servicecatalog-cloudformationprovisionedproduct-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the provisioned product ID, such as
`pp-hfyszaotincww`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CloudformationStackArn`

Property description not available.

`ProvisionedProductId`

The ID of the provisioned product.

`RecordId`

The ID of the record, such as `rec-rjeatvy434trk`.

## Examples

### GetAtt Example

#### YAML

```yaml

      AWSTemplateFormatVersion: '2010-09-09'
      Description: Serverless Stack
      Resources:
         SimpleLambda:
           Type: AWS::ServiceCatalog::CloudFormationProvisionedProduct
           Properties:
            ProductName: Basic Lambda
            ProvisioningArtifactName: '1.0'

         SimpleApiGateway:
           Type: AWS::ServiceCatalog::CloudFormationProvisionedProduct
           Properties:
            ProductName: API Gateway
            ProvisioningArtifactName: '1.0'
            ProvisioningParameters:
               -
                  Key: LambdaArn
                  Value: !GetAtt [SimpleLambda, Outputs.SCLambdaArn]

```

## See also

- [ProvisionProduct](../../../servicecatalog/latest/dg/api-provisionproduct.md) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ProvisioningParameter

All content copied from https://docs.aws.amazon.com/.
