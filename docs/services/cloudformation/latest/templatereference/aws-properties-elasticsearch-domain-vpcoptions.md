This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain VPCOptions

The virtual private cloud (VPC) configuration for the OpenSearch Service domain. For more
information, see [Launching your Amazon OpenSearch\
Service domains using a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the _Amazon OpenSearch Service Developer_
_Guide_.

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
Elasticsearch. For more information about the service rename, see [New resource\
types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the _Amazon OpenSearch Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`SecurityGroupIds`

The list of security group IDs that are associated with the VPC endpoints for the domain.
If you don't provide a security group ID, OpenSearch Service uses the default security group
for the VPC. To learn more, see [Security groups for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in
the _Amazon VPC User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

Provide one subnet ID for each Availability Zone that your domain uses. For example, you
must specify three subnet IDs for a three Availability Zone domain. To learn more, see [VPCs and subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in
the _Amazon VPC User Guide_.

Required if you're creating your domain inside a VPC.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the
_Amazon OpenSearch Service Developer Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ZoneAwarenessConfig
