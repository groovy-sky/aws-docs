This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index Search

Provides information about how a custom index field is used during a search.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Displayable" : Boolean,
  "Facetable" : Boolean,
  "Searchable" : Boolean,
  "Sortable" : Boolean
}

```

### YAML

```yaml

  Displayable: Boolean
  Facetable: Boolean
  Searchable: Boolean
  Sortable: Boolean

```

## Properties

`Displayable`

Determines whether the field is returned in the query response. The default is
`true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Facetable`

Indicates that the field can be used to create search facets, a count of results for
each value in the field. The default is `false` .

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Searchable`

Determines whether the field is used in the search. If the `Searchable`
field is `true`, you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is `true` for
string fields and `false` for number and date fields.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sortable`

Determines whether the field can be used to sort the results of a query. The default
is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Relevance

ServerSideEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
