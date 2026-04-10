This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::ConfigurationProfile Validators

A validator provides a syntactic or semantic check to ensure the configuration that you
want to deploy functions as intended. To validate your application configuration data, you
provide a schema or an AWS Lambda function that runs against the configuration. The
configuration deployment or update can only proceed when the configuration data is valid.
For more information, see [About validators](../../../appconfig/latest/userguide/appconfig-creating-configuration-profile.md#appconfig-creating-configuration-and-profile-validators) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String,
  "Type" : String
}

```

### YAML

```yaml

  Content: String
  Type: String

```

## Properties

`Content`

Either the JSON Schema content or the Amazon Resource Name (ARN) of an Lambda
function.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `32768`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

AWS AppConfig supports validators of type `JSON_SCHEMA` and
`LAMBDA`

_Required_: No

_Type_: String

_Allowed values_: `JSON_SCHEMA | LAMBDA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tags

AWS::AppConfig::Deployment

All content copied from https://docs.aws.amazon.com/.
