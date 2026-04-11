This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::Endpoint ElasticsearchSettings

Provides information that defines an OpenSearch endpoint. This
information includes the output format of records applied to the endpoint and details of
transaction and control table data information. For more information about the available settings, see
[Extra connection attributes when using OpenSearch as a target for AWS DMS](../../../dms/latest/userguide/chap-target-elasticsearch.md#CHAP_Target.Elasticsearch.Configuration)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndpointUri" : String,
  "ErrorRetryDuration" : Integer,
  "FullLoadErrorPercentage" : Integer,
  "ServiceAccessRoleArn" : String
}

```

### YAML

```yaml

  EndpointUri: String
  ErrorRetryDuration: Integer
  FullLoadErrorPercentage: Integer
  ServiceAccessRoleArn: String

```

## Properties

`EndpointUri`

The endpoint for the OpenSearch cluster. AWS DMS uses HTTPS if a transport
protocol (either HTTP or HTTPS) isn't specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorRetryDuration`

The maximum number of seconds for which DMS retries failed API requests to the
OpenSearch cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FullLoadErrorPercentage`

The maximum percentage of records that can fail to be written before a full load
operation stops.

To avoid early failure, this counter is only effective after 1,000 records
are transferred. OpenSearch also has the concept of error monitoring during the
last 10 minutes of an Observation Window. If transfer of all records fail in the
last 10 minutes, the full load operation stops.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccessRoleArn`

The Amazon Resource Name (ARN) used by the service to access the IAM role.
The role must allow the `iam:PassRole` action.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamoDbSettings

GcpMySQLSettings

All content copied from https://docs.aws.amazon.com/.
