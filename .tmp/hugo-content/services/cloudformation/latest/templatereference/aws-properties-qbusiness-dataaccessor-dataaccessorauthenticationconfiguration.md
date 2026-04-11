This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor DataAccessorAuthenticationConfiguration

A union type that contains the specific authentication configuration based on the
authentication type selected.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdcTrustedTokenIssuerConfiguration" : DataAccessorIdcTrustedTokenIssuerConfiguration
}

```

### YAML

```yaml

  IdcTrustedTokenIssuerConfiguration:
    DataAccessorIdcTrustedTokenIssuerConfiguration

```

## Properties

`IdcTrustedTokenIssuerConfiguration`

Configuration for IAM Identity Center Trusted Token Issuer (TTI) authentication used
when the authentication type is `AWS_IAM_IDC_TTI`.

_Required_: Yes

_Type_: [DataAccessorIdcTrustedTokenIssuerConfiguration](aws-properties-qbusiness-dataaccessor-dataaccessoridctrustedtokenissuerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeFilter

DataAccessorAuthenticationDetail

All content copied from https://docs.aws.amazon.com/.
