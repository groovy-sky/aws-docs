This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualGateway VirtualGatewayBackendDefaults

An object that represents the default properties for a backend.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientPolicy" : VirtualGatewayClientPolicy
}

```

### YAML

```yaml

  ClientPolicy:
    VirtualGatewayClientPolicy

```

## Properties

`ClientPolicy`

A reference to an object that represents a client policy.

_Required_: No

_Type_: [VirtualGatewayClientPolicy](aws-properties-appmesh-virtualgateway-virtualgatewayclientpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VirtualGatewayAccessLog

VirtualGatewayClientPolicy

All content copied from https://docs.aws.amazon.com/.
