This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping SchemaRegistryAccessConfig

Specific access configuration settings that tell Lambda how to authenticate with your schema registry.

If you're working with an AWS Glue schema registry, don't provide authentication details in this object.
Instead, ensure that your execution role has the required permissions for Lambda to access your cluster.

If you're working with a Confluent schema registry, choose the authentication method in the `Type` field,
and provide the AWS Secrets Manager secret ARN in the `URI` field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "URI" : String
}

```

### YAML

```yaml

  Type: String
  URI: String

```

## Properties

`Type`

The type of authentication Lambda uses to access your schema registry.

_Required_: No

_Type_: String

_Allowed values_: `BASIC_AUTH | CLIENT_CERTIFICATE_TLS_AUTH | SERVER_ROOT_CA_CERTIFICATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`URI`

The URI of the secret (Secrets Manager secret ARN) to authenticate with your schema registry.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\-])+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:(.*)`

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingConfig

SchemaRegistryConfig

All content copied from https://docs.aws.amazon.com/.
