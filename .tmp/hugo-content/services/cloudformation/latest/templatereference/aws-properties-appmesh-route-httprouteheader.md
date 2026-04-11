This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route HttpRouteHeader

An object that represents the HTTP header in the request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Invert" : Boolean,
  "Match" : HeaderMatchMethod,
  "Name" : String
}

```

### YAML

```yaml

  Invert: Boolean
  Match:
    HeaderMatchMethod
  Name: String

```

## Properties

`Invert`

Specify `True` to match anything except the match criteria. The default value is `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

The `HeaderMatchMethod` object.

_Required_: No

_Type_: [HeaderMatchMethod](aws-properties-appmesh-route-headermatchmethod.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the HTTP header in the client request that will be matched on.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpRouteAction

HttpRouteMatch

All content copied from https://docs.aws.amazon.com/.
