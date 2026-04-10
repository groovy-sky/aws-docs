This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::LogStream

The `AWS::Logs::LogStream` resource specifies an Amazon CloudWatch Logs log stream in a specific log group. A log stream represents the sequence of events
coming from an application instance or resource that you are monitoring.

There is no limit on the number of log streams that you can create for a log
group.

You must use the following guidelines when naming a log stream:

- Log stream names must be unique within the log group.

- Log stream names can be between 1 and 512 characters long.

- The ':' (colon) and '\*' (asterisk) characters are not allowed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::LogStream",
  "Properties" : {
      "LogGroupName" : String,
      "LogStreamName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Logs::LogStream
Properties:
  LogGroupName: String
  LogStreamName: String

```

## Properties

`LogGroupName`

The name of the log group where the log stream is created.

_Required_: Yes

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogStreamName`

The name of the log stream. The name must be unique within the log group.

_Required_: No

_Type_: String

_Pattern_: `[^:*]*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name, such as `
            MyAppLogStream`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a Log Stream

The following example creates a log stream named `MyAppLogStream`
in the `exampleLogGroup` log group.

#### JSON

```json

"LogStream": {
  "Type": "AWS::Logs::LogStream",
  "Properties": {
    "LogGroupName" : "exampleLogGroup",
    "LogStreamName": "MyAppLogStream"
  }
}
```

#### YAML

```yaml

LogStream:
  Type: AWS::Logs::LogStream
  Properties:
    LogGroupName: "exampleLogGroup"
    LogStreamName: "MyAppLogStream"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Logs::MetricFilter

All content copied from https://docs.aws.amazon.com/.
