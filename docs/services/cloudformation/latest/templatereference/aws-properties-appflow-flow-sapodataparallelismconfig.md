This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SAPODataParallelismConfig

Sets the number of _concurrent processes_ that transfer OData records
from your SAP instance. A concurrent process is query that retrieves a batch of records as
part of a flow run. Amazon AppFlow can run multiple concurrent processes in parallel to
transfer data faster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "maxParallelism" : Integer
}

```

### YAML

```yaml

  maxParallelism: Integer

```

## Properties

`maxParallelism`

The maximum number of processes that Amazon AppFlow runs at the same time when it
retrieves your data from your SAP application.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SAPODataPaginationConfig

SAPODataSourceProperties

All content copied from https://docs.aws.amazon.com/.
