This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route GrpcRouteMatch

An object that represents the criteria for determining a request match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Metadata" : [ GrpcRouteMetadata, ... ],
  "MethodName" : String,
  "Port" : Integer,
  "ServiceName" : String
}

```

### YAML

```yaml

  Metadata:
    - GrpcRouteMetadata
  MethodName: String
  Port: Integer
  ServiceName: String

```

## Properties

`Metadata`

An object that represents the data to match from the request.

_Required_: No

_Type_: Array of [GrpcRouteMetadata](aws-properties-appmesh-route-grpcroutemetadata.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MethodName`

The method name to match from the request. If you specify a name, you must also specify
a `serviceName`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number to match on.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The fully qualified domain name for the service to match from the request.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GrpcRouteAction

GrpcRouteMetadata

All content copied from https://docs.aws.amazon.com/.
