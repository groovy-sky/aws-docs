This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::App ResourceMapping

Defines a resource mapping.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EksSourceName" : String,
  "LogicalStackName" : String,
  "MappingType" : String,
  "PhysicalResourceId" : PhysicalResourceId,
  "ResourceName" : String,
  "TerraformSourceName" : String
}

```

### YAML

```yaml

  EksSourceName: String
  LogicalStackName: String
  MappingType: String
  PhysicalResourceId:
    PhysicalResourceId
  ResourceName: String
  TerraformSourceName: String

```

## Properties

`EksSourceName`

Name of the Amazon Elastic Kubernetes Service cluster and namespace that this resource is mapped to when the `mappingType` is
`EKS`.

###### Note

This parameter accepts values in "eks-cluster/namespace" format.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalStackName`

Name of the CloudFormation stack this resource is mapped to when the `mappingType` is `CfnStack`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MappingType`

Specifies the type of resource mapping.

_Required_: Yes

_Type_: String

_Pattern_: `CfnStack|Resource|Terraform|EKS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhysicalResourceId`

Identifier of the physical resource.

_Required_: Yes

_Type_: [PhysicalResourceId](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resiliencehub-app-physicalresourceid.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceName`

Name of the resource that this resource is mapped to when the `mappingType` is `Resource`.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9_\-]{1,59}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerraformSourceName`

Name of the Terraform source that this resource is mapped to when the `mappingType` is `Terraform`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PhysicalResourceId

AWS::ResilienceHub::ResiliencyPolicy
