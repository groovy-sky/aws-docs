This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin CustomPluginConfiguration

Configuration information required to create a custom plugin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiSchema" : APISchema,
  "ApiSchemaType" : String,
  "Description" : String
}

```

### YAML

```yaml

  ApiSchema:
    APISchema
  ApiSchemaType: String
  Description: String

```

## Properties

`ApiSchema`

Contains either details about the S3 object containing the OpenAPI schema for the
action group or the JSON or YAML-formatted payload defining the schema.

_Required_: Yes

_Type_: [APISchema](aws-properties-qbusiness-plugin-apischema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApiSchemaType`

The type of OpenAPI schema to use.

_Required_: Yes

_Type_: String

_Allowed values_: `OPEN_API_V3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for your custom plugin configuration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthConfiguration

OAuth2ClientCredentialConfiguration

All content copied from https://docs.aws.amazon.com/.
