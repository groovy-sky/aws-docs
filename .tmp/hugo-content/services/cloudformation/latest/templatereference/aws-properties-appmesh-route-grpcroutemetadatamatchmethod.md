This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route GrpcRouteMetadataMatchMethod

An object that represents the match method. Specify one of the match values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exact" : String,
  "Prefix" : String,
  "Range" : MatchRange,
  "Regex" : String,
  "Suffix" : String
}

```

### YAML

```yaml

  Exact: String
  Prefix: String
  Range:
    MatchRange
  Regex: String
  Suffix: String

```

## Properties

`Exact`

The value sent by the client must match the specified value exactly.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The value sent by the client must begin with the specified characters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Range`

An object that represents the range of values to match on.

_Required_: No

_Type_: [MatchRange](aws-properties-appmesh-route-matchrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regex`

The value sent by the client must include the specified characters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Suffix`

The value sent by the client must end with the specified characters.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcRouteMetadata

GrpcTimeout

All content copied from https://docs.aws.amazon.com/.
