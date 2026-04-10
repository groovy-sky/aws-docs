This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration HttpRequest

The `HttpRequest` property type specifies Property description not available. for an [AWS::MediaTailor::PlaybackConfiguration](aws-resource-mediatailor-playbackconfiguration.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Body" : String,
  "CompressRequest" : String,
  "Headers" : {Key: Value, ...},
  "HttpMethod" : String
}

```

### YAML

```yaml

  Body: String
  CompressRequest: String
  Headers:
    Key: Value
  HttpMethod: String

```

## Properties

`Body`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompressRequest`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `NONE | GZIP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Headers`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpMethod`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `GET | POST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsConfiguration

LivePreRollConfiguration

All content copied from https://docs.aws.amazon.com/.
