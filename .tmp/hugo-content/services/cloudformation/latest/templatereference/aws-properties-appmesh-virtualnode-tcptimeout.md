This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode TcpTimeout

An object that represents types of timeouts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Idle" : Duration
}

```

### YAML

```yaml

  Idle:
    Duration

```

## Properties

`Idle`

An object that represents an idle timeout. An idle timeout bounds the amount of time that a connection may be idle. The default value is none.

_Required_: No

_Type_: [Duration](aws-properties-appmesh-virtualnode-duration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TlsValidationContext

All content copied from https://docs.aws.amazon.com/.
