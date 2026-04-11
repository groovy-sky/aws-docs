This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application InteractiveConfiguration

The configuration to use to enable the different types of interactive use cases in an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LivyEndpointEnabled" : Boolean,
  "StudioEnabled" : Boolean
}

```

### YAML

```yaml

  LivyEndpointEnabled: Boolean
  StudioEnabled: Boolean

```

## Properties

`LivyEndpointEnabled`

Enables an Apache Livy endpoint that you can connect to and run interactive jobs.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`StudioEnabled`

Enables you to connect an application to Amazon EMR Studio to run interactive workloads in a notebook.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InitialCapacityConfigKeyValuePair

LogTypeMapKeyValuePair

All content copied from https://docs.aws.amazon.com/.
