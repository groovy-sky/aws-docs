This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::QueryDefinition QueryParameter

This structure defines a query parameter for a saved CloudWatch Logs Insights query
definition. Query parameters are supported only for Logs Insights QL queries. They are
placeholder variables that you can reference in a query string using the
`{{parameterName}}` syntax. Each parameter can include a default value and a
description.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : String,
  "Description" : String,
  "Name" : String
}

```

### YAML

```yaml

  DefaultValue: String
  Description: String
  Name: String

```

## Properties

`DefaultValue`

The default value to use for this query parameter if no value is supplied at execution
time.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the query parameter that explains its purpose or expected values.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the query parameter. A query parameter name must start with a letter or
underscore, and contain only letters, digits, and underscores.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z0-9_]*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Logs::QueryDefinition

AWS::Logs::ResourcePolicy

All content copied from https://docs.aws.amazon.com/.
