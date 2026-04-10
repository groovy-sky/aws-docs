This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain SnapshotOptions

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](../userguide/aws-resource-opensearchservice-domain.md) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
Elasticsearch. For more information about the service rename, see [New resource\
types](../../../opensearch-service/latest/developerguide/rename.md#rename-resource) in the _Amazon OpenSearch Service Developer_
_Guide_.

**DEPRECATED**. For domains running Elasticsearch 5.3 and
later, OpenSearch Service takes hourly automated snapshots, making this setting irrelevant. For domains
running earlier versions of Elasticsearch, OpenSearch Service takes daily automated snapshots.

The automated snapshot configuration for the OpenSearch Service domain indices.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutomatedSnapshotStartHour" : Integer
}

```

### YAML

```yaml

  AutomatedSnapshotStartHour: Integer

```

## Properties

`AutomatedSnapshotStartHour`

The hour in UTC during which the service takes an automated daily snapshot of the indices
in the OpenSearch Service domain. For example, if you specify 0, OpenSearch Service takes an automated snapshot
everyday between midnight and 1 am. You can specify a value between 0 and 23.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeToNodeEncryptionOptions

Tag

All content copied from https://docs.aws.amazon.com/.
