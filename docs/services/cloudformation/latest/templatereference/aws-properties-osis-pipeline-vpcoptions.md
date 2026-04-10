This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline VpcOptions

Options that specify the subnets and security groups for an OpenSearch Ingestion
VPC endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ],
  "VpcAttachmentOptions" : VpcAttachmentOptions,
  "VpcEndpointManagement" : String
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  VpcAttachmentOptions:
    VpcAttachmentOptions
  VpcEndpointManagement: String

```

## Properties

`SecurityGroupIds`

A list of security groups associated with the VPC endpoint.

_Required_: No

_Type_: Array of String

_Minimum_: `11`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

A list of subnet IDs associated with the VPC endpoint.

_Required_: Yes

_Type_: Array of String

_Minimum_: `15`

_Maximum_: `24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcAttachmentOptions`

Options for attaching a VPC to a pipeline.

_Required_: No

_Type_: [VpcAttachmentOptions](aws-properties-osis-pipeline-vpcattachmentoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcEndpointManagement`

Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.

_Required_: No

_Type_: String

_Allowed values_: `CUSTOMER | SERVICE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcEndpoint

Next

All content copied from https://docs.aws.amazon.com/.
