This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource DataSourceErrorInfo

Error information for the data source creation or update.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Message" : String,
  "Type" : String
}

```

### YAML

```yaml

  Message: String
  Type: String

```

## Properties

`Message`

Error message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Error type.

_Required_: No

_Type_: String

_Allowed values_: `ACCESS_DENIED | COPY_SOURCE_NOT_FOUND | TIMEOUT | ENGINE_VERSION_NOT_SUPPORTED | UNKNOWN_HOST | GENERIC_SQL_FAILURE | CONFLICT | UNKNOWN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceCredentials

DataSourceParameters

All content copied from https://docs.aws.amazon.com/.
