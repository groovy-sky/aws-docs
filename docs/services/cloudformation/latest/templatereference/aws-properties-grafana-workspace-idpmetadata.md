This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace IdpMetadata

A structure containing the identity provider (IdP) metadata used to integrate the
identity provider with this workspace. You can specify the metadata either by providing
a URL to its location in the `url` parameter, or by specifying the full
metadata in XML format in the `xml` parameter. Specifying both will cause an
error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Url" : String,
  "Xml" : String
}

```

### YAML

```yaml

  Url: String
  Xml: String

```

## Properties

`Url`

The URL of the location containing the IdP metadata.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Xml`

The full IdP metadata, in XML format.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssertionAttributes

NetworkAccessControl

All content copied from https://docs.aws.amazon.com/.
