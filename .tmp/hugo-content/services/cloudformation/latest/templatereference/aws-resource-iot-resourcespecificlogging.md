This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ResourceSpecificLogging

Configure resource-specific logging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::ResourceSpecificLogging",
  "Properties" : {
      "LogLevel" : String,
      "TargetName" : String,
      "TargetType" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::ResourceSpecificLogging
Properties:
  LogLevel: String
  TargetName: String
  TargetType: String

```

## Properties

`LogLevel`

The default log level.Valid Values: `DEBUG | INFO | ERROR | WARN | DISABLED`

_Required_: Yes

_Type_: String

_Allowed values_: `ERROR | WARN | INFO | DEBUG | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetName`

The target name.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9.:\s_\-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetType`

The target type. Valid Values: `DEFAULT | THING_GROUP`

_Required_: Yes

_Type_: String

_Allowed values_: `THING_GROUP | CLIENT_ID | SOURCE_IP | PRINCIPAL_ID | EVENT_TYPE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource-specific log ID. For example:

`{"Ref": "MyResourceLog-12345" }`

### Fn::GetAtt

`TargetId`

The target Id.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoT::RoleAlias

All content copied from https://docs.aws.amazon.com/.
