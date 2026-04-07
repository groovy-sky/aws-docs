This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::GameServerGroup LaunchTemplate

**This data type is used with the GameLift FleetIQ and game server groups.**

An Amazon EC2 launch template that contains configuration settings and game server code to
be deployed to all instances in a game server group. The launch template is specified
when creating a new game server group with `GameServerGroup`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateId" : String,
  "LaunchTemplateName" : String,
  "Version" : String
}

```

### YAML

```yaml

  LaunchTemplateId: String
  LaunchTemplateName: String
  Version: String

```

## Properties

`LaunchTemplateId`

A unique identifier for an existing Amazon EC2 launch template.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplateName`

A readable identifier for an existing Amazon EC2 launch template.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9\(\)\.\-/_]+`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the Amazon EC2 launch template to use. If no version is specified, the
default version will be used. With Amazon EC2, you can specify a default version for a launch
template. If none is set, the default is the first version created.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceDefinition

Tag
