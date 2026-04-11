This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route QueryParameter

An object that represents the query parameter in the request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Match" : HttpQueryParameterMatch,
  "Name" : String
}

```

### YAML

```yaml

  Match:
    HttpQueryParameterMatch
  Name: String

```

## Properties

`Match`

The query parameter to match on.

_Required_: No

_Type_: [HttpQueryParameterMatch](aws-properties-appmesh-route-httpqueryparametermatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the query parameter that will be matched on.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MatchRange

RouteSpec

All content copied from https://docs.aws.amazon.com/.
