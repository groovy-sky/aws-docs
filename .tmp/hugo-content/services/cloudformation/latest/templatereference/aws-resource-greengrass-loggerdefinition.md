This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::LoggerDefinition

The `AWS::Greengrass::LoggerDefinition` resource represents a logger
definition for AWS IoT Greengrass. Logger definitions are used to organize your logger
definition versions.

Logger definitions can reference multiple logger definition versions. All logger
definition versions must be associated with a logger definition. Each logger definition
version can contain one or more loggers.

###### Note

When you create a logger definition, you can optionally include an initial logger
definition version. To associate a logger definition version later, create an [`AWS::Greengrass::LoggerDefinitionVersion`](../userguide/aws-resource-greengrass-loggerdefinitionversion.md) resource and
specify the ID of this logger definition.

After you create the logger definition version that contains the loggers you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::LoggerDefinition",
  "Properties" : {
      "InitialVersion" : LoggerDefinitionVersion,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::LoggerDefinition
Properties:
  InitialVersion:
    LoggerDefinitionVersion
  Name: String
  Tags:
    - Tag

```

## Properties

`InitialVersion`

The logger definition version to include when the logger definition is created. A logger
definition version contains a list of [`logger`](../userguide/aws-properties-greengrass-loggerdefinition-logger.md) property types.

###### Note

To associate a logger definition version after the logger definition is created,
create an [`AWS::Greengrass::LoggerDefinitionVersion`](../userguide/aws-resource-greengrass-loggerdefinitionversion.md) resource and specify the ID of this logger definition.

_Required_: No

_Type_: [LoggerDefinitionVersion](aws-properties-greengrass-loggerdefinition-loggerdefinitionversion.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the logger definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Application-specific metadata to attach to the logger definition. You can use tags in
IAM policies to control access to AWS IoT Greengrass resources. You
can also use tags to categorize your resources. For more information, see [Tagging Your\
AWS IoT Greengrass Resources](../../../greengrass/v1/developerguide/tagging.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses
the following format, which is different from most `Tags` implementations in
CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the logger definition, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `LoggerDefinition`, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/loggers/1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`Id`

The ID of the `LoggerDefinition`, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`LatestVersionArn`

The ARN of the last `LoggerDefinitionVersion` that was added to the
`LoggerDefinition`, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/loggers/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

`Name`

The name of the `LoggerDefinition`, such as `MyLoggerDefinition`.

## Examples

### Logger Definition Snippet

The following snippet defines a logger definition resource with an initial version
that contains a logger.

For an example of a complete template, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

"TestLoggerDefinition": {
    "Type": "AWS::Greengrass::LoggerDefinition",
    "Properties": {
        "Name": "DemoTestLoggerDefinition",
        "InitialVersion": {
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
}
```

#### YAML

```yaml

TestLoggerDefinition:
  Type: 'AWS::Greengrass::LoggerDefinition'
  Properties:
    Name: DemoTestLoggerDefinition
    InitialVersion:
      Loggers:
        - Id: TestLogger1
          Type: FileSystem
          Component: GreengrassSystem
          Level: INFO
          Space: '128'
```

## See also

- [CreateLoggerDefinition](../../../../reference/greengrass/v1/apireference/createloggerdefinition-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Greengrass::GroupVersion

Logger

All content copied from https://docs.aws.amazon.com/.
