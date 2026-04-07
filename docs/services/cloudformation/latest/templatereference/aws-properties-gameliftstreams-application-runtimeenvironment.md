This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLiftStreams::Application RuntimeEnvironment

Configuration settings that identify the operating system for an application resource. This can also include a compatibility layer and
other drivers.

A runtime environment can be one of the following:

- For Linux applications

- Ubuntu 22.04 LTS ( `Type=UBUNTU, Version=22_04_LTS`)

- For Windows applications

- Microsoft Windows Server 2022 Base ( `Type=WINDOWS, Version=2022`)

- Proton 9.0-2 ( `Type=PROTON, Version=20250516`)

- Proton 8.0-5 ( `Type=PROTON, Version=20241007`)

- Proton 8.0-2c ( `Type=PROTON, Version=20230704`)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Version" : String
}

```

### YAML

```yaml

  Type: String
  Version: String

```

## Properties

`Type`

The operating system and other drivers. For Proton, this also includes the Proton compatibility layer.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

Versioned container environment for the application operating system.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GameLiftStreams::Application

AWS::GameLiftStreams::StreamGroup
