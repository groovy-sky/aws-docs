This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SalesforceSourceProperties

The properties that are applied when Salesforce is being used as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataTransferApi" : String,
  "EnableDynamicFieldUpdate" : Boolean,
  "IncludeDeletedRecords" : Boolean,
  "Object" : String
}

```

### YAML

```yaml

  DataTransferApi: String
  EnableDynamicFieldUpdate: Boolean
  IncludeDeletedRecords: Boolean
  Object: String

```

## Properties

`DataTransferApi`

Specifies which Salesforce API is used by Amazon AppFlow when your flow transfers
data from Salesforce.

AUTOMATIC

The default. Amazon AppFlow selects which API to use based on the number of
records that your flow transfers from Salesforce. If your flow transfers fewer than
1,000,000 records, Amazon AppFlow uses Salesforce REST API. If your flow transfers
1,000,000 records or more, Amazon AppFlow uses Salesforce Bulk API 2.0.

Each of these Salesforce APIs structures data differently. If Amazon AppFlow
selects the API automatically, be aware that, for recurring flows, the data output might
vary from one flow run to the next. For example, if a flow runs daily, it might use REST
API on one day to transfer 900,000 records, and it might use Bulk API 2.0 on the next
day to transfer 1,100,000 records. For each of these flow runs, the respective
Salesforce API formats the data differently. Some of the differences include how dates
are formatted and null values are represented. Also, Bulk API 2.0 doesn't transfer
Salesforce compound fields.

By choosing this option, you optimize flow performance for both small and large data
transfers, but the tradeoff is inconsistent formatting in the output.

BULKV2

Amazon AppFlow uses only Salesforce Bulk API 2.0. This API runs asynchronous
data transfers, and it's optimal for large sets of data. By choosing this option, you
ensure that your flow writes consistent output, but you optimize performance only for
large data transfers.

Note that Bulk API 2.0 does not transfer Salesforce compound fields.

REST\_SYNC

Amazon AppFlow uses only Salesforce REST API. By choosing this option, you
ensure that your flow writes consistent output, but you decrease performance for large
data transfers that are better suited for Bulk API 2.0. In some cases, if your flow
attempts to transfer a vary large set of data, it might fail wituh a timed out
error.

_Required_: No

_Type_: [String](aws-properties-appflow-flow-datatransferapi.md)

_Allowed values_: `AUTOMATIC | BULKV2 | REST_SYNC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDynamicFieldUpdate`

The flag that enables dynamic fetching of new (recently added) fields in the Salesforce
objects while running a flow.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeDeletedRecords`

Indicates whether Amazon AppFlow includes deleted files in the flow run.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Object`

The object specified in the Salesforce flow source.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SalesforceSourceProperties](../../../../reference/appflow/1-0/apireference/api-salesforcesourceproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceDestinationProperties

SAPODataDestinationProperties

All content copied from https://docs.aws.amazon.com/.
