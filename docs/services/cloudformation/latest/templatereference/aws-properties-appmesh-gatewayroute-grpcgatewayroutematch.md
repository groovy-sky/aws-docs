This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute GrpcGatewayRouteMatch

An object that represents the criteria for determining a request match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hostname" : GatewayRouteHostnameMatch,
  "Metadata" : [ GrpcGatewayRouteMetadata, ... ],
  "Port" : Integer,
  "ServiceName" : String
}

```

### YAML

```yaml

  Hostname:
    GatewayRouteHostnameMatch
  Metadata:
    - GrpcGatewayRouteMetadata
  Port: Integer
  ServiceName: String

```

## Properties

`Hostname`

The gateway route host name to be matched on.

_Required_: No

_Type_: [GatewayRouteHostnameMatch](aws-properties-appmesh-gatewayroute-gatewayroutehostnamematch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metadata`

The gateway route metadata to be matched on.

_Required_: No

_Type_: Array of [GrpcGatewayRouteMetadata](aws-properties-appmesh-gatewayroute-grpcgatewayroutemetadata.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The gateway route port to be matched on.

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

GrpcGatewayRouteAction

GrpcGatewayRouteMetadata

All content copied from https://docs.aws.amazon.com/.
