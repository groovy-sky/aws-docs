This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain VPCOptions

The virtual private cloud (VPC) configuration for the OpenSearch Service domain. For more
information, see [Launching your Amazon OpenSearch\
Service domains using a VPC](../../../opensearch-service/latest/developerguide/vpc.md) in the _Amazon OpenSearch Service Developer_
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
for the VPC. To learn more, see [Security groups for your VPC](../../../vpc/latest/userguide/vpc-securitygroups.md) in
the _Amazon VPC User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

Provide one subnet ID for each Availability Zone that your domain uses. For example, you
must specify three subnet IDs for a three-AZ domain. To learn more, see [VPCs and subnets](../../../vpc/latest/userguide/vpc-subnets.md) in
the _Amazon VPC User Guide_.

If you specify more than one subnet, you must also configure
`ZoneAwarenessEnabled` and `ZoneAwarenessConfig` within [ClusterConfig](../userguide/aws-properties-opensearchservice-domain-clusterconfig.md), otherwise you'll see the error "You must specify exactly one subnet"
during template creation.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Launching your Amazon OpenSearch Service domains within a VPC](../../../opensearch-service/latest/developerguide/vpc.md) in the
_Amazon OpenSearch Service Developer Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WindowStartTime

All content copied from https://docs.aws.amazon.com/.
