This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space SpaceIdleSettings

Settings related to idle shutdown of Studio applications in a space.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleTimeoutInMinutes" : Integer
}

```

### YAML

```yaml

  IdleTimeoutInMinutes: Integer

```

## Properties

`IdleTimeoutInMinutes`

The time that SageMaker waits after the application becomes idle before shutting it
down.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `525600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpaceCodeEditorAppSettings

SpaceJupyterLabAppSettings

All content copied from https://docs.aws.amazon.com/.
