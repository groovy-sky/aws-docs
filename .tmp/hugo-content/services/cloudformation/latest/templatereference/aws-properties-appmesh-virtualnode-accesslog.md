This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode AccessLog

An object that represents the access logging information for a virtual node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "File" : FileAccessLog
}

```

### YAML

```yaml

  File:
    FileAccessLog

```

## Properties

`File`

The file object to send virtual node access logs to.

_Required_: No

_Type_: [FileAccessLog](aws-properties-appmesh-virtualnode-fileaccesslog.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppMesh::VirtualNode

AwsCloudMapInstanceAttribute

All content copied from https://docs.aws.amazon.com/.
