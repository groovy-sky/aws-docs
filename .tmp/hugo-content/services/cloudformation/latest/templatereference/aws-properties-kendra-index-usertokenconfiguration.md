This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index UserTokenConfiguration

Provides the configuration information for a token.

###### Important

If you're using an Amazon Kendra Gen AI Enterprise Edition index and you try to use
`UserTokenConfigurations` to configure user context policy, Amazon Kendra returns
a `ValidationException` error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonTokenTypeConfiguration" : JsonTokenTypeConfiguration,
  "JwtTokenTypeConfiguration" : JwtTokenTypeConfiguration
}

```

### YAML

```yaml

  JsonTokenTypeConfiguration:
    JsonTokenTypeConfiguration
  JwtTokenTypeConfiguration:
    JwtTokenTypeConfiguration

```

## Properties

`JsonTokenTypeConfiguration`

Information about the JSON token type configuration.

_Required_: No

_Type_: [JsonTokenTypeConfiguration](aws-properties-kendra-index-jsontokentypeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwtTokenTypeConfiguration`

Information about the JWT token type configuration.

_Required_: No

_Type_: [JwtTokenTypeConfiguration](aws-properties-kendra-index-jwttokentypeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ValueImportanceItem

All content copied from https://docs.aws.amazon.com/.
