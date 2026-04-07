This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain ElasticsearchClusterConfig

The cluster configuration for the OpenSearch Service domain. You can specify options such
as the instance type and the number of instances. For more information, see [Creating and managing Amazon OpenSearch Service domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html) in the _Amazon_
_OpenSearch Service Developer Guide_.

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
  "ColdStorageOptions" : ColdStorageOptions,
  "DedicatedMasterCount" : Integer,
  "DedicatedMasterEnabled" : Boolean,
  "DedicatedMasterType" : String,
  "InstanceCount" : Integer,
  "InstanceType" : String,
  "WarmCount" : Integer,
  "WarmEnabled" : Boolean,
  "WarmType" : String,
  "ZoneAwarenessConfig" : ZoneAwarenessConfig,
  "ZoneAwarenessEnabled" : Boolean
}

```

### YAML

```yaml

  ColdStorageOptions:
    ColdStorageOptions
  DedicatedMasterCount: Integer
  DedicatedMasterEnabled: Boolean
  DedicatedMasterType: String
  InstanceCount: Integer
  InstanceType: String
  WarmCount: Integer
  WarmEnabled: Boolean
  WarmType: String
  ZoneAwarenessConfig:
    ZoneAwarenessConfig
  ZoneAwarenessEnabled: Boolean

```

## Properties

`ColdStorageOptions`

Specifies cold storage options for the domain.

_Required_: No

_Type_: [ColdStorageOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-coldstorageoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DedicatedMasterCount`

The number of instances to use for the master node. If you specify this property, you must
specify true for the DedicatedMasterEnabled property.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DedicatedMasterEnabled`

Indicates whether to use a dedicated master node for the OpenSearch Service domain. A
dedicated master node is a cluster node that performs cluster management tasks, but doesn't
hold data or respond to data upload requests. Dedicated master nodes offload cluster
management tasks to increase the stability of your search clusters. See [Dedicated master nodes in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-dedicatedmasternodes.html).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DedicatedMasterType`

The hardware configuration of the computer that hosts the dedicated master node, such as
`m3.medium.elasticsearch`. If you specify this property, you must specify true
for the `DedicatedMasterEnabled` property. For valid values, see [Supported\
instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceCount`

The number of data nodes (instances) to use in the OpenSearch Service domain.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type for your data nodes, such as `m3.medium.elasticsearch`. For
valid values, see [Supported\
instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmCount`

The number of warm nodes in the cluster. Required if you enable warm storage.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmEnabled`

Whether to enable warm storage for the cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WarmType`

The instance type for the cluster's warm nodes. Required if you enable warm
storage.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZoneAwarenessConfig`

Specifies zone awareness configuration options. Only use if
`ZoneAwarenessEnabled` is `true`.

_Required_: Conditional

_Type_: [ZoneAwarenessConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-zoneawarenessconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZoneAwarenessEnabled`

Indicates whether to enable zone awareness for the OpenSearch Service domain. When you
enable zone awareness, OpenSearch Service allocates the nodes and replica index shards that
belong to a cluster across two Availability Zones (AZs) in the same region to prevent data
loss and minimize downtime in the event of node or data center failure. Don't enable zone
awareness if your cluster has no replica index shards or is a single-node cluster. For more
information, see [Configuring a\
multi-AZ domain in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EBSOptions

EncryptionAtRestOptions
