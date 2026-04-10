This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin APISchema

Contains details about the OpenAPI schema for a custom plugin. For more information,
see [custom plugin OpenAPI schemas](../../../amazonq/latest/qbusiness-ug/custom-plugin.md#plugins-api-schema). You can either include
the schema directly in the payload field or you can upload it to an S3 bucket and
specify the S3 bucket location in the `s3` field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Payload" : String,
  "S3" : S3
}

```

### YAML

```yaml

  Payload: String
  S3:
    S3

```

## Properties

`Payload`

The JSON or YAML-formatted payload defining the OpenAPI schema for a custom plugin.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Contains details about the S3 object containing the OpenAPI schema for a custom
plugin. The schema could be in either JSON or YAML format.

_Required_: No

_Type_: [S3](aws-properties-qbusiness-plugin-s3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QBusiness::Plugin

BasicAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
