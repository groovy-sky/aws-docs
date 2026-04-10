This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain LogPublishingOption

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](../userguide/aws-resource-opensearchservice-domain.md) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
Elasticsearch. For more information about the service rename, see [New resource\
types](../../../opensearch-service/latest/developerguide/rename.md#rename-resource) in the _Amazon OpenSearch Service Developer_
_Guide_.

Specifies whether the OpenSearch Service domain publishes the Elasticsearch application,
search slow logs, or index slow logs to Amazon CloudWatch. Each option must be an object of
name `SEARCH_SLOW_LOGS`, `ES_APPLICATION_LOGS`,
`INDEX_SLOW_LOGS`, or `AUDIT_LOGS` depending on the type of logs you
want to publish.

If you enable a slow log, you still have to enable the _collection_ of
slow logs using the Configuration API. To learn more, see [Enabling log publishing (AWS CLI)](../../../opensearch-service/latest/developerguide/createdomain-configure-slow-logs.md#createdomain-configure-slow-logs-cli).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsLogGroupArn" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  CloudWatchLogsLogGroupArn: String
  Enabled: Boolean

```

## Properties

`CloudWatchLogsLogGroupArn`

Specifies the CloudWatch log group to publish to. Required if you enable log publishing
for the domain.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

If `true`, enables the publishing of logs to CloudWatch.

Default: `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Monitoring OpenSearch logs with Amazon CloudWatch Logs](../../../opensearch-service/latest/developerguide/createdomain-configure-slow-logs.md) and [LogPublishingOptions](../../../opensearch-service/latest/developerguide/configuration-api.md#configuration-api-datatypes-logpublishingoptions) in the _Amazon OpenSearch Service Developer_
_Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionAtRestOptions

MasterUserOptions

All content copied from https://docs.aws.amazon.com/.
