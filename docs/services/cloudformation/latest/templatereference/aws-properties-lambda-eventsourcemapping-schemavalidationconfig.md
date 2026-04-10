This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping SchemaValidationConfig

Specific schema validation configuration settings that tell Lambda the message
attributes you want to validate and filter using your schema registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attribute" : String
}

```

### YAML

```yaml

  Attribute: String

```

## Properties

`Attribute`

The attributes you want your schema registry to validate and filter for. If you selected `JSON` as the
`EventRecordFormat`, Lambda also deserializes the selected message attributes.

_Required_: No

_Type_: String

_Allowed values_: `KEY | VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaRegistryConfig

SelfManagedEventSource

All content copied from https://docs.aws.amazon.com/.
