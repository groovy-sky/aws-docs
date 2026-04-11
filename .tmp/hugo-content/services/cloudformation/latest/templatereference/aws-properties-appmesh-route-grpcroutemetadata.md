This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route GrpcRouteMetadata

An object that represents the match metadata for the route.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Invert" : Boolean,
  "Match" : GrpcRouteMetadataMatchMethod,
  "Name" : String
}

```

### YAML

```yaml

  Invert: Boolean
  Match:
    GrpcRouteMetadataMatchMethod
  Name: String

```

## Properties

`Invert`

Specify `True` to match anything except the match criteria. The default value is `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

An object that represents the data to match from the request.

_Required_: No

_Type_: [GrpcRouteMetadataMatchMethod](aws-properties-appmesh-route-grpcroutemetadatamatchmethod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the route.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcRouteMatch

GrpcRouteMetadataMatchMethod

All content copied from https://docs.aws.amazon.com/.
