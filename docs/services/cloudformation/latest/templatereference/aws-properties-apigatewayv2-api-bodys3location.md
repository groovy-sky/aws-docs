This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Api BodyS3Location

The `BodyS3Location` property specifies an S3 location from which to
import an OpenAPI definition. Supported only for HTTP APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Etag" : String,
  "Key" : String,
  "Version" : String
}

```

### YAML

```yaml

  Bucket: String
  Etag: String
  Key: String
  Version: String

```

## Properties

`Bucket`

The S3 bucket that contains the OpenAPI definition to import. Required if you specify a `BodyS3Location` for an API.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Etag`

The Etag of the S3 object.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key of the S3 object. Required if you specify a `BodyS3Location` for an API.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the S3 object.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGatewayV2::Api

Cors
