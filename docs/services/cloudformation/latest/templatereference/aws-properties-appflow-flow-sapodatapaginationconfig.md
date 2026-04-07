This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SAPODataPaginationConfig

Sets the page size for each _concurrent process_ that transfers OData
records from your SAP instance. A concurrent process is query that retrieves a batch of
records as part of a flow run. Amazon AppFlow can run multiple concurrent processes in
parallel to transfer data faster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "maxPageSize" : Integer
}

```

### YAML

```yaml

  maxPageSize: Integer

```

## Properties

`maxPageSize`

The maximum number of records that Amazon AppFlow receives in each page of the
response from your SAP application. For transfers of OData records, the maximum page size is
3,000. For transfers of data that comes from an ODP provider, the maximum page size
is 10,000.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SAPODataDestinationProperties

SAPODataParallelismConfig
