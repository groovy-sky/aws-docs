This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule HttpMatch

Describes criteria that can be applied to incoming requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderMatches" : [ HeaderMatch, ... ],
  "Method" : String,
  "PathMatch" : PathMatch
}

```

### YAML

```yaml

  HeaderMatches:
    - HeaderMatch
  Method: String
  PathMatch:
    PathMatch

```

## Properties

`HeaderMatches`

The header matches. Matches incoming requests with rule based on request header value before
applying rule action.

_Required_: No

_Type_: Array of [HeaderMatch](aws-properties-vpclattice-rule-headermatch.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Method`

The HTTP method type.

_Required_: No

_Type_: String

_Allowed values_: `CONNECT | DELETE | GET | HEAD | OPTIONS | POST | PUT | TRACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PathMatch`

The path match.

_Required_: No

_Type_: [PathMatch](aws-properties-vpclattice-rule-pathmatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeaderMatchType

Match

All content copied from https://docs.aws.amazon.com/.
