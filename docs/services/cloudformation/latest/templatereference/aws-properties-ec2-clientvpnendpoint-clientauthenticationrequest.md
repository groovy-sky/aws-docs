This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint ClientAuthenticationRequest

Describes the authentication method to be used by a Client VPN endpoint. For more information, see [Authentication](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/authentication-authrization.html#client-authentication)
in the _AWS Client VPN Administrator Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActiveDirectory" : DirectoryServiceAuthenticationRequest,
  "FederatedAuthentication" : FederatedAuthenticationRequest,
  "MutualAuthentication" : CertificateAuthenticationRequest,
  "Type" : String
}

```

### YAML

```yaml

  ActiveDirectory:
    DirectoryServiceAuthenticationRequest
  FederatedAuthentication:
    FederatedAuthenticationRequest
  MutualAuthentication:
    CertificateAuthenticationRequest
  Type: String

```

## Properties

`ActiveDirectory`

Information about the Active Directory to be used, if applicable. You must provide this information if **Type** is `directory-service-authentication`.

_Required_: No

_Type_: [DirectoryServiceAuthenticationRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-clientvpnendpoint-directoryserviceauthenticationrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FederatedAuthentication`

Information about the IAM SAML identity provider, if applicable.

_Required_: No

_Type_: [FederatedAuthenticationRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-clientvpnendpoint-federatedauthenticationrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MutualAuthentication`

Information about the authentication certificates to be used, if applicable. You must provide this information if **Type** is `certificate-authentication`.

_Required_: No

_Type_: [CertificateAuthenticationRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-clientvpnendpoint-certificateauthenticationrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of client authentication to be used.

_Required_: Yes

_Type_: String

_Allowed values_: `certificate-authentication | directory-service-authentication | federated-authentication`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CertificateAuthenticationRequest

ClientConnectOptions
