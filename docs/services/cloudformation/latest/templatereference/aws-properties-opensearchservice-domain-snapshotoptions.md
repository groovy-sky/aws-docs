This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain SnapshotOptions

**DEPRECATED**. This setting is only relevant to domains
running legacy Elasticsearch OSS versions earlier than 5.3. It does not apply to OpenSearch
domains.

The automated snapshot configuration for the OpenSearch Service domain indexes.

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

The hour in UTC during which the service takes an automated daily snapshot of the indexes
in the OpenSearch Service domain. For example, if you specify 0, OpenSearch Service takes an
automated snapshot everyday between midnight and 1 am. You can specify a value between 0 and
23.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceSoftwareOptions

SoftwareUpdateOptions
