This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain Idp

The SAML Identity Provider's information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EntityId" : String,
  "MetadataContent" : String
}

```

### YAML

```yaml

  EntityId: String
  MetadataContent: String

```

## Properties

`EntityId`

The unique entity ID of the application in the SAML identity provider.

_Required_: Yes

_Type_: String

_Minimum_: `8`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetadataContent`

The metadata of the SAML application, in XML format.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1048576`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdentityCenterOptions

JWTOptions

All content copied from https://docs.aws.amazon.com/.
