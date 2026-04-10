This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::DomainConfiguration AuthorizerConfig

An object that specifies the authorization service for a domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowAuthorizerOverride" : Boolean,
  "DefaultAuthorizerName" : String
}

```

### YAML

```yaml

  AllowAuthorizerOverride: Boolean
  DefaultAuthorizerName: String

```

## Properties

`AllowAuthorizerOverride`

A Boolean that specifies whether the domain configuration's authorization service can be overridden.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultAuthorizerName`

The name of the authorization service for a domain configuration.

_Required_: No

_Type_: String

_Pattern_: `^[\w=,@-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::DomainConfiguration

ClientCertificateConfig

All content copied from https://docs.aws.amazon.com/.
