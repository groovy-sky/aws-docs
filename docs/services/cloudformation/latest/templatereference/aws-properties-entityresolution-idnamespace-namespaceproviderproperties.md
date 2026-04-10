This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdNamespace NamespaceProviderProperties

An object containing `providerConfiguration` and
`providerServiceArn`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProviderConfiguration" : {Key: Value, ...},
  "ProviderServiceArn" : String
}

```

### YAML

```yaml

  ProviderConfiguration:
    Key: Value
  ProviderServiceArn: String

```

## Properties

`ProviderConfiguration`

An object which defines any additional configurations required by the provider
service.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderServiceArn`

The Amazon Resource Name (ARN) of the provider service.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):(entityresolution):([a-z]{2}-[a-z]{1,10}-[0-9])::providerservice/([a-zA-Z0-9_-]{1,255})/([a-zA-Z0-9_-]{1,255})$`

_Minimum_: `20`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdNamespaceInputSource

NamespaceRuleBasedProperties

All content copied from https://docs.aws.amazon.com/.
