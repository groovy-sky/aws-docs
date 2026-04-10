This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain LogPublishingOption

Specifies whether the OpenSearch Service domain publishes application, search slow logs,
or index slow logs to Amazon CloudWatch. Each option must be an object of name
`SEARCH_SLOW_LOGS`, `ES_APPLICATION_LOGS`,
`INDEX_SLOW_LOGS`, or `AUDIT_LOGS` depending on the type of logs you
want to publish. For the full syntax, see the [examples](../userguide/aws-resource-opensearchservice-domain.md#aws-resource-opensearchservice-domain--examples).

Before you enable log publishing, you need to create a CloudWatch log group and provide
OpenSearch Service the correct permissions to write to it. To learn more, see [Enabling log publishing (AWS CloudFormation)](../../../opensearch-service/latest/developerguide/createdomain-configure-slow-logs.md#createdomain-configure-slow-logs-cfn).

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

Specifies the CloudWatch log group to publish to. Required if you enable log
publishing.

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

JWTOptions

MasterUserOptions

All content copied from https://docs.aws.amazon.com/.
