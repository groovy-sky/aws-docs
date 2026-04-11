This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::GatewayRoute HttpGatewayRouteRewrite

An object representing the gateway route to rewrite.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hostname" : GatewayRouteHostnameRewrite,
  "Path" : HttpGatewayRoutePathRewrite,
  "Prefix" : HttpGatewayRoutePrefixRewrite
}

```

### YAML

```yaml

  Hostname:
    GatewayRouteHostnameRewrite
  Path:
    HttpGatewayRoutePathRewrite
  Prefix:
    HttpGatewayRoutePrefixRewrite

```

## Properties

`Hostname`

The host name to rewrite.

_Required_: No

_Type_: [GatewayRouteHostnameRewrite](aws-properties-appmesh-gatewayroute-gatewayroutehostnamerewrite.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to rewrite.

_Required_: No

_Type_: [HttpGatewayRoutePathRewrite](aws-properties-appmesh-gatewayroute-httpgatewayroutepathrewrite.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The specified beginning characters to rewrite.

_Required_: No

_Type_: [HttpGatewayRoutePrefixRewrite](aws-properties-appmesh-gatewayroute-httpgatewayrouteprefixrewrite.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpGatewayRoutePrefixRewrite

HttpPathMatch

All content copied from https://docs.aws.amazon.com/.
