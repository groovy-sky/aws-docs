This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProduct ProvisioningArtifactProperties

Information about a provisioning artifact (also known as a version) for a product.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "DisableTemplateValidation" : Boolean,
  "Info" : Json,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Description: String
  DisableTemplateValidation: Boolean
  Info: Json
  Name: String
  Type: String

```

## Properties

`Description`

The description of the provisioning artifact, including how it differs from the previous provisioning artifact.

_Required_: No

_Type_: String

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableTemplateValidation`

If set to true, AWS Service Catalog stops validating the specified provisioning artifact even if it is invalid.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Info`

Specify the template source with one of the following options, but not both.
Keys accepted: \[ `LoadTemplateFromURL`, `ImportFromPhysicalId` \]

The URL of the AWS CloudFormation template in Amazon S3 in JSON format.
Specify the URL in JSON format as follows:

`"LoadTemplateFromURL": "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/..."`

`ImportFromPhysicalId`: The physical id of the resource that contains the
template. Currently only supports AWS CloudFormation stack arn. Specify the physical id in JSON
format as follows: `ImportFromPhysicalId: “arn:aws:cloudformation:[us-east-1]:[accountId]:stack/[StackName]/[resourceId]`

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the provisioning artifact (for example, v1 v2beta). No spaces are allowed.

_Required_: No

_Type_: String

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of provisioning artifact.

- `CLOUD_FORMATION_TEMPLATE` \- AWS CloudFormation template

- `TERRAFORM_OPEN_SOURCE` \- Terraform Open Source configuration file

- `TERRAFORM_CLOUD` \- Terraform Cloud configuration file

- `EXTERNAL` \- External configuration file

_Required_: No

_Type_: String

_Allowed values_: `CLOUD_FORMATION_TEMPLATE | MARKETPLACE_AMI | MARKETPLACE_CAR | TERRAFORM_OPEN_SOURCE | EXTERNAL | TERRAFORM_CLOUD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ProvisioningArtifactProperties](../../../servicecatalog/latest/dg/api-provisioningartifactproperties.md) in the _AWS Service Catalog_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionParameters

SourceConnection

All content copied from https://docs.aws.amazon.com/.
