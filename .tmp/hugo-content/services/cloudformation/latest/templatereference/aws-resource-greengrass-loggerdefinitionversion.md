This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::LoggerDefinitionVersion

The `AWS::Greengrass::LoggerDefinitionVersion` resource represents a logger
definition version for AWS IoT Greengrass. A logger definition version contains a list of
loggers.

###### Note

To create a logger definition version, you must specify the ID of the logger
definition that you want to associate with the version. For information about creating a
logger definition, see [`AWS::Greengrass::LoggerDefinition`](../userguide/aws-resource-greengrass-loggerdefinition.md).

After you create a logger definition version that contains the loggers you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::LoggerDefinitionVersion",
  "Properties" : {
      "LoggerDefinitionId" : String,
      "Loggers" : [ Logger, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::LoggerDefinitionVersion
Properties:
  LoggerDefinitionId: String
  Loggers:
    - Logger

```

## Properties

`LoggerDefinitionId`

The ID of the logger definition associated with this version. This value is a
GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Loggers`

The loggers in this version.

_Required_: Yes

_Type_: Array of [Logger](aws-properties-greengrass-loggerdefinitionversion-logger.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the logger definition
version, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/loggers/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Logger Definition Version Snippet

The following snippet defines logger definition and logger definition version
resources. The logger definition version references the logger definition and
contains a logger.

For an example of a complete template, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

"TestLoggerDefinition": {
    "Type": "AWS::Greengrass::LoggerDefinition",
    "Properties": {
        "Name": "DemoTestLoggerDefinition"
    }
},
"TestLoggerDefinitionVersion": {
    "Type": "AWS::Greengrass::LoggerDefinitionVersion",
    "Properties": {
        "LoggerDefinitionId": {
            "Ref": "TestLoggerDefinition"
        },
        "Loggers": [
            {
                "Id": "TestLogger1",
                "Type": "FileSystem",
                "Component": "GreengrassSystem",
                "Level": "INFO",
                "Space": "128"
            }
        ]
    }
}
```

#### YAML

```yaml

TestLoggerDefinition:
  Type: 'AWS::Greengrass::LoggerDefinition'
  Properties:
    Name: DemoTestLoggerDefinition
TestLoggerDefinitionVersion:
  Type: 'AWS::Greengrass::LoggerDefinitionVersion'
  Properties:
    LoggerDefinitionId: !Ref TestLoggerDefinition
    Loggers:
      - Id: TestLogger1
        Type: FileSystem
        Component: GreengrassSystem
        Level: INFO
        Space: '128'
```

## See also

- [CreateLoggerDefinitionVersion](../../../../reference/greengrass/v1/apireference/createloggerdefinitionversion-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggerDefinitionVersion

Logger

All content copied from https://docs.aws.amazon.com/.
