This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile SnowflakeConnectorProfileProperties

The connector-specific profile properties required when using Snowflake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountName" : String,
  "BucketName" : String,
  "BucketPrefix" : String,
  "PrivateLinkServiceName" : String,
  "Region" : String,
  "Stage" : String,
  "Warehouse" : String
}

```

### YAML

```yaml

  AccountName: String
  BucketName: String
  BucketPrefix: String
  PrivateLinkServiceName: String
  Region: String
  Stage: String
  Warehouse: String

```

## Properties

`AccountName`

The name of the account.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

The name of the Amazon S3 bucket associated with Snowflake.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The bucket path that refers to the Amazon S3 bucket associated with Snowflake.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateLinkServiceName`

The Snowflake Private Link service name to be used for private data transfers.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The AWS Region of the Snowflake account.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snowflake account. This is written in the following format: <
Database>< Schema><Stage Name>.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Warehouse`

The name of the Snowflake warehouse.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\w/!@#+=.-]*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SnowflakeConnectorProfileProperties](../../../../reference/appflow/1-0/apireference/api-snowflakeconnectorprofileproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeConnectorProfileCredentials

TrendmicroConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
