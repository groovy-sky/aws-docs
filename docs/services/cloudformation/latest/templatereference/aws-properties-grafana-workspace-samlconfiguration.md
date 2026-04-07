This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace SamlConfiguration

A structure containing information about how this workspace works with SAML.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedOrganizations" : [ String, ... ],
  "AssertionAttributes" : AssertionAttributes,
  "IdpMetadata" : IdpMetadata,
  "LoginValidityDuration" : Number,
  "RoleValues" : RoleValues
}

```

### YAML

```yaml

  AllowedOrganizations:
    - String
  AssertionAttributes:
    AssertionAttributes
  IdpMetadata:
    IdpMetadata
  LoginValidityDuration: Number
  RoleValues:
    RoleValues

```

## Properties

`AllowedOrganizations`

Lists which organizations defined in the SAML assertion are allowed to use the Amazon Managed Grafana workspace. If this is empty, all organizations in the assertion attribute
have access.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssertionAttributes`

A structure that defines which attributes in the SAML assertion are to be used to
define information about the users authenticated by that IdP to use the
workspace.

_Required_: No

_Type_: [AssertionAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-grafana-workspace-assertionattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdpMetadata`

A structure containing the identity provider (IdP) metadata used to integrate the
identity provider with this workspace.

_Required_: Yes

_Type_: [IdpMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-grafana-workspace-idpmetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginValidityDuration`

How long a sign-on session by a SAML user is valid, before the user has to sign on
again.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleValues`

A structure containing arrays that map group names in the SAML assertion to the
Grafana `Admin` and `Editor` roles in the workspace.

_Required_: No

_Type_: [RoleValues](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-grafana-workspace-rolevalues.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RoleValues

VpcConfiguration
