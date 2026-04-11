This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function ImageConfig

Configuration values that override the container image Dockerfile settings. For more information, see [Container image\
settings](../../../lambda/latest/dg/images-create.md#images-parms).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Command" : [ String, ... ],
  "EntryPoint" : [ String, ... ],
  "WorkingDirectory" : String
}

```

### YAML

```yaml

  Command:
    - String
  EntryPoint:
    - String
  WorkingDirectory: String

```

## Properties

`Command`

Specifies parameters that you want to pass in with ENTRYPOINT. You can specify a maximum of 1,500 parameters
in the list.

_Required_: No

_Type_: Array of String

_Maximum_: `1500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntryPoint`

Specifies the entry point to their application, which is typically the location of the runtime
executable. You can specify a maximum of 1,500 string entries in the list.

_Required_: No

_Type_: Array of String

_Maximum_: `1500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkingDirectory`

Specifies the working directory. The length of the directory string cannot exceed 1,000 characters.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunctionScalingConfig

LambdaManagedInstancesCapacityProviderConfig

All content copied from https://docs.aws.amazon.com/.
