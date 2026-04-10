This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource SqlConfiguration

Provides information that configures Amazon Kendra to use a SQL database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueryIdentifiersEnclosingOption" : String
}

```

### YAML

```yaml

  QueryIdentifiersEnclosingOption: String

```

## Properties

`QueryIdentifiersEnclosingOption`

Determines whether Amazon Kendra encloses SQL identifiers for tables and column names
in double quotes (") when making a database query. You can set the value to
`DOUBLE_QUOTES` or `NONE`.

By default, Amazon Kendra passes SQL identifiers the way that they are entered into
the data source configuration. It does not change the case of identifiers or enclose
them in quotes.

PostgreSQL internally converts uppercase characters to lower case characters in
identifiers unless they are quoted. Choosing this option encloses identifiers in quotes
so that PostgreSQL does not convert the character's case.

For MySQL databases, you must enable the ansi\_quotes option when you set this field
to `DOUBLE_QUOTES`.

_Required_: No

_Type_: String

_Allowed values_: `DOUBLE_QUOTES | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SharePointConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
