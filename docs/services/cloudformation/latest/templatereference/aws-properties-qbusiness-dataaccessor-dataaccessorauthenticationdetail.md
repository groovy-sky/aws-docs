This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor DataAccessorAuthenticationDetail

Contains the authentication configuration details for a data accessor. This structure
defines how the ISV authenticates when accessing data through the data accessor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationConfiguration" : DataAccessorAuthenticationConfiguration,
  "AuthenticationType" : String,
  "ExternalIds" : [ String, ... ]
}

```

### YAML

```yaml

  AuthenticationConfiguration:
    DataAccessorAuthenticationConfiguration
  AuthenticationType: String
  ExternalIds:
    - String

```

## Properties

`AuthenticationConfiguration`

The specific authentication configuration based on the authentication type.

_Required_: No

_Type_: [DataAccessorAuthenticationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-dataaccessor-dataaccessorauthenticationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

The type of authentication to use for the data accessor. This determines how the ISV
authenticates when accessing data. You can use one of two authentication types:

- `AWS_IAM_IDC_TTI` \- Authentication using IAM Identity Center
Trusted Token Issuer (TTI). This authentication type allows the ISV to use a
trusted token issuer to generate tokens for accessing the data.

- `AWS_IAM_IDC_AUTH_CODE` \- Authentication using IAM Identity Center authorization code flow. This authentication type uses the
standard OAuth 2.0 authorization code flow for authentication.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_IAM_IDC_TTI | AWS_IAM_IDC_AUTH_CODE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalIds`

A list of external identifiers associated with this authentication configuration.
These are used to correlate the data accessor with external systems.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1000 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataAccessorAuthenticationConfiguration

DataAccessorIdcTrustedTokenIssuerConfiguration
