This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource ConnectionConfiguration

Provides the configuration information that's required to connect to a
database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseHost" : String,
  "DatabaseName" : String,
  "DatabasePort" : Integer,
  "SecretArn" : String,
  "TableName" : String
}

```

### YAML

```yaml

  DatabaseHost: String
  DatabaseName: String
  DatabasePort: Integer
  SecretArn: String
  TableName: String

```

## Properties

`DatabaseHost`

The name of the host for the database. Can be either a string
(host.subdomain.domain.tld) or an IPv4 or IPv6 address.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of the database containing the document data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabasePort`

The port that the database uses for connections.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that stores
the credentials. The credentials should be a user-password pair. For more information, see [Using a\
Database Data Source](../../../kendra/latest/dg/data-source-database.md). For more information about AWS Secrets Manager, see
[What\
Is AWS Secrets Manager](../../../secretsmanager/latest/userguide/intro.md) in the _AWS Secrets Manager_
user guide.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the table that contains the document data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfluenceSpaceToIndexFieldMapping

CustomDocumentEnrichmentConfiguration

All content copied from https://docs.aws.amazon.com/.
