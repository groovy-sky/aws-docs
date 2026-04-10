This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain SAMLOptions

Container for information about the SAML configuration for OpenSearch
Dashboards.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "Idp" : Idp,
  "MasterBackendRole" : String,
  "MasterUserName" : String,
  "RolesKey" : String,
  "SessionTimeoutMinutes" : Integer,
  "SubjectKey" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  Idp:
    Idp
  MasterBackendRole: String
  MasterUserName: String
  RolesKey: String
  SessionTimeoutMinutes: Integer
  SubjectKey: String

```

## Properties

`Enabled`

True to enable SAML authentication for a domain.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Idp`

The SAML Identity Provider's information.

_Required_: No

_Type_: [Idp](aws-properties-opensearchservice-domain-idp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterBackendRole`

The backend role that the SAML master user is mapped to.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUserName`

The SAML master user name, which is stored in the domain's internal user
database.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RolesKey`

Element of the SAML assertion to use for backend roles. Default is
`roles`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionTimeoutMinutes`

The duration, in minutes, after which a user session becomes inactive. Acceptable
values are between 1 and 1440, and the default value is 60.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubjectKey`

Element of the SAML assertion to use for the user name. Default is
`NameID`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3VectorsEngine

ServerlessVectorAcceleration

All content copied from https://docs.aws.amazon.com/.
