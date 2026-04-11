This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::SigningConfiguration RepositoryFilter

A repository filter used to determine which repositories have their
images automatically signed on push. Each filter consists of a filter type and filter value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filter" : String,
  "FilterType" : String
}

```

### YAML

```yaml

  Filter: String
  FilterType: String

```

## Properties

`Filter`

The filter value used to match repository names. When using
`WILDCARD_MATCH`, the `*` character matches any sequence of characters.

Examples:

- `myapp/*` \- Matches all repositories starting with
`myapp/`

- `*/production` \- Matches all repositories ending with
`/production`

- `*prod*` \- Matches all repositories containing
`prod`

_Required_: Yes

_Type_: String

_Pattern_: `^(?:[a-z0-9*]+(?:[._-][a-z0-9*]+)*/)*[a-z0-9*]+(?:[._-][a-z0-9*]+)*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterType`

The type of filter to apply. Currently, only `WILDCARD_MATCH` is supported,
which uses wildcard patterns to match repository names.

_Required_: Yes

_Type_: String

_Allowed values_: `WILDCARD_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECR::SigningConfiguration

Rule

All content copied from https://docs.aws.amazon.com/.
