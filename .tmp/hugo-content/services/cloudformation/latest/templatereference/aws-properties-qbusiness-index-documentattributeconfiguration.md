This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Index DocumentAttributeConfiguration

Configuration information for document attributes. Document attributes are metadata or
fields associated with your documents. For example, the company department name
associated with each document.

For more information, see [Understanding document attributes](../../../amazonq/latest/business-use-dg/doc-attributes.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Search" : String,
  "Type" : String
}

```

### YAML

```yaml

  Name: String
  Search: String
  Type: String

```

## Properties

`Name`

The name of the document attribute.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Search`

Information about whether the document attribute can be used by an end user to search
for information on their web experience.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of document attribute.

_Required_: No

_Type_: String

_Allowed values_: `STRING | STRING_LIST | NUMBER | DATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QBusiness::Index

IndexCapacityConfiguration

All content copied from https://docs.aws.amazon.com/.
