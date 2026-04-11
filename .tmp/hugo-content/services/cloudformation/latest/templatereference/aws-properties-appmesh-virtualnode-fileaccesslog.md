This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::VirtualNode FileAccessLog

An object that represents an access log file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Format" : LoggingFormat,
  "Path" : String
}

```

### YAML

```yaml

  Format:
    LoggingFormat
  Path: String

```

## Properties

`Format`

The specified format for the logs. The format is either `json_format` or
`text_format`.

_Required_: No

_Type_: [LoggingFormat](aws-properties-appmesh-virtualnode-loggingformat.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The file path to write access logs to. You can use `/dev/stdout` to send
access logs to standard out and configure your Envoy container to use a log driver, such as
`awslogs`, to export the access logs to a log storage service such as Amazon
CloudWatch Logs. You can also specify a path in the Envoy container's file system to write
the files to disk.

###### Note

The Envoy process must have write permissions to the path that you specify here.
Otherwise, Envoy fails to bootstrap properly.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Duration

GrpcTimeout

All content copied from https://docs.aws.amazon.com/.
