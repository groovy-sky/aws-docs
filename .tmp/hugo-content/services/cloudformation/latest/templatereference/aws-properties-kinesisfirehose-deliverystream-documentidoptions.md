This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DocumentIdOptions

Indicates the method for setting up document ID. The supported methods are Firehose generated document ID and OpenSearch Service generated document ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultDocumentIdFormat" : String
}

```

### YAML

```yaml

  DefaultDocumentIdFormat: String

```

## Properties

`DefaultDocumentIdFormat`

When the `FIREHOSE_DEFAULT` option is chosen, Firehose generates
a unique document ID for each record based on a unique internal identifier. The generated
document ID is stable across multiple delivery attempts, which helps prevent the same
record from being indexed multiple times with different document IDs.

When the `NO_DOCUMENT_ID` option is chosen, Firehose does not
include any document IDs in the requests it sends to the Amazon OpenSearch Service. This
causes the Amazon OpenSearch Service domain to generate document IDs. In case of multiple
delivery attempts, this may cause the same record to be indexed more than once with
different document IDs. This option enables write-heavy operations, such as the ingestion
of logs and observability data, to consume less resources in the Amazon OpenSearch Service
domain, resulting in improved performance.

_Required_: Yes

_Type_: String

_Allowed values_: `FIREHOSE_DEFAULT | NO_DOCUMENT_ID`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DirectPutSourceConfiguration

DynamicPartitioningConfiguration

All content copied from https://docs.aws.amazon.com/.
