This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceExplorer2::View SearchFilter

A search filter defines which resources can be part of a search query result
set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilterString" : String
}

```

### YAML

```yaml

  FilterString:
    String

```

## Properties

`FilterString`

The string that contains the search keywords, prefixes, and operators to control the
results that can be returned by a Search operation.

For information about the supported syntax, see [Search query\
reference](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html) in the _AWS Resource Explorer User_
_Guide_.

###### Important

This query string in the context of this operation supports only [filter prefixes](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-filters) with optional [operators](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-operators). It doesn't support free-form text. For example, the string
`region:us* service:ec2 -tag:stage=prod` includes all Amazon EC2 resources in any AWS Region that begin with the
letters `us` and are _not_ tagged with a key
`Stage` that has the value `prod`.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IncludedProperty

Next
