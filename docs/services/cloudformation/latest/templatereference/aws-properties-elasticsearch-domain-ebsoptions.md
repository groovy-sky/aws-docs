This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain EBSOptions

The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to
data nodes in the OpenSearch Service domain. For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the _Amazon OpenSearch Service Developer_
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
  "EBSEnabled" : Boolean,
  "Iops" : Integer,
  "VolumeSize" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  EBSEnabled: Boolean
  Iops: Integer
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`EBSEnabled`

Specifies whether Amazon EBS volumes are attached to data nodes in the OpenSearch Service
domain.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Iops`

The number of I/O operations per second (IOPS) that the volume supports. This property
applies only to provisioned IOPS EBS volume types.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSize`

The size (in GiB) of the EBS volume for each data node. The minimum and maximum size of an
EBS volume depends on the EBS volume type and the instance type to which it is attached. For
more information, see [EBS volume size\
limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the _Amazon OpenSearch Service Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeType`

The EBS volume type to use with the OpenSearch Service domain, such as standard, gp2, or
io1. For more information about each type, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the
_Amazon EC2 User Guide for Linux Instances_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DomainEndpointOptions

ElasticsearchClusterConfig
