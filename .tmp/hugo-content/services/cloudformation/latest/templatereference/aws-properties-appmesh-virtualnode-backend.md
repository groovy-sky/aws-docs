This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode Backend

An object that represents the backends that a virtual node is expected to send outbound
traffic to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VirtualService" : VirtualServiceBackend
}

```

### YAML

```yaml

  VirtualService:
    VirtualServiceBackend

```

## Properties

`VirtualService`

Specifies a virtual service to use as a backend.

_Required_: No

_Type_: [VirtualServiceBackend](aws-properties-appmesh-virtualnode-virtualservicebackend.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsCloudMapServiceDiscovery

BackendDefaults

All content copied from https://docs.aws.amazon.com/.
