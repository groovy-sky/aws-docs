This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping SchemaRegistryConfig

Specific configuration settings for a Kafka schema registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessConfigs" : [ SchemaRegistryAccessConfig, ... ],
  "EventRecordFormat" : String,
  "SchemaRegistryURI" : String,
  "SchemaValidationConfigs" : [ SchemaValidationConfig, ... ]
}

```

### YAML

```yaml

  AccessConfigs:
    - SchemaRegistryAccessConfig
  EventRecordFormat: String
  SchemaRegistryURI: String
  SchemaValidationConfigs:
    - SchemaValidationConfig

```

## Properties

`AccessConfigs`

An array of access configuration objects that tell Lambda how to authenticate with your schema registry.

_Required_: No

_Type_: Array of [SchemaRegistryAccessConfig](aws-properties-lambda-eventsourcemapping-schemaregistryaccessconfig.md)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventRecordFormat`

The record format that Lambda delivers to your function after schema validation.

- Choose `JSON` to have Lambda deliver the record to your function as a standard JSON object.

- Choose `SOURCE` to have Lambda deliver the record to your function in its original source format.
Lambda removes all schema metadata, such as the schema ID, before sending the record to your function.

_Required_: No

_Type_: String

_Allowed values_: `JSON | SOURCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaRegistryURI`

The URI for your schema registry. The correct URI format depends on the type of schema registry you're using.

- For AWS Glue schema registries, use the ARN of the registry.

- For Confluent schema registries, use the URL of the registry.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9-/*:_+=.@-]*`

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaValidationConfigs`

An array of schema validation configuration objects, which tell Lambda the message
attributes you want to validate and filter using your schema registry.

_Required_: No

_Type_: Array of [SchemaValidationConfig](aws-properties-lambda-eventsourcemapping-schemavalidationconfig.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SchemaRegistryAccessConfig

SchemaValidationConfig

All content copied from https://docs.aws.amazon.com/.
