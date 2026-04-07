This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::SecurityConfig

Specifies a security configuration for OpenSearch Serverless. For more information, see [SAML\
authentication for Amazon OpenSearch Serverless](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchServerless::SecurityConfig",
  "Properties" : {
      "Description" : String,
      "IamFederationOptions" : IamFederationConfigOptions,
      "IamIdentityCenterOptions" : IamIdentityCenterConfigOptions,
      "Name" : String,
      "SamlOptions" : SamlConfigOptions,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchServerless::SecurityConfig
Properties:
  Description: String
  IamFederationOptions:
    IamFederationConfigOptions
  IamIdentityCenterOptions:
    IamIdentityCenterConfigOptions
  Name: String
  SamlOptions:
    SamlConfigOptions
  Type: String

```

## Properties

`Description`

The description of the security configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamFederationOptions`

Describes IAM federation options in the form of a key-value map. Contains
configuration details about how OpenSearch Serverless integrates with external identity
providers through federation.

_Required_: No

_Type_: [IamFederationConfigOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchserverless-securityconfig-iamfederationconfigoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamIdentityCenterOptions`

Describes IAM Identity Center options in the form of a key-value map.

_Required_: No

_Type_: [IamIdentityCenterConfigOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchserverless-securityconfig-iamidentitycenterconfigoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the security configuration.

_Required_: No

_Type_: String

_Pattern_: `^[a-z][a-z0-9-]{2,31}$`

_Minimum_: `3`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SamlOptions`

SAML options for the security configuration in the form of a key-value map.

_Required_: No

_Type_: [SamlConfigOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchserverless-securityconfig-samlconfigoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of security configuration. Currently the only option is
`saml`.

_Required_: No

_Type_: String

_Allowed values_: `saml | iamidentitycenter | iamfederation`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the name of the ID of the security configuration. For
more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`GetAtt` returns a value for a specified attribute of this type. For more
information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html). The following are the available attributes and sample return
values.

`IamIdentityCenterOptions.ApplicationArn`

Property description not available.

`IamIdentityCenterOptions.ApplicationDescription`

Property description not available.

`IamIdentityCenterOptions.ApplicationName`

Property description not available.

`Id`

The unique identifier of the security configuration. For example,
`saml/123456789012/myprovider`.

## Examples

### Create a security configuration that specifies a YAML provider

The following example specifies an OpenSearch Serverless SAML provider named
`my-provider` with a custom group attribute
`ALLGroups`.

#### JSON

```json

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Description":"OpenSearch Serverless security policy template",
   "Resources":{
      "TestSecurityConfig":{
         "Type":"AWS::OpenSearchServerless::SecurityConfig",
         "Properties":{
            "Name":"my-provider",
            "Type":"saml",
            "Description":"Serverless SAML configuration",
            "SamlOptions":{
               "Metadata":"<?xml version=\"1.0\"
                encoding=\"UTF-8\"?><md:EntityDescriptor
                entityID=\"http://www.okta.com/foobar\"
                xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\"><md:IDPSSODescriptor
                WantAuthnRequestsSigned=\"false\"
                protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:KeyDescriptor
                use=\"signing\"><ds:KeyInfo
                xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:X509Data><ds:X509Certificate>Mfoobar</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat><md:SingleSignOnService
                Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\"
                Location=\"https://trial-1234567.okta.com/app/trial-1234567_saml2_1/foobar/sso/saml\"/><md:SingleSignOnService
                Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect\"
                Location=\"https://trial-1234567.okta.com/app/trial-1234567_saml2_1/foobar/sso/saml\"/></md:IDPSSODescriptor></md:EntityDescriptor>",
               "UserAttribute":"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier",
               "GroupAttribute":"ALLGroups",
               "SessionTimeout":120
            }
         }
      }
   }
}
```

#### YAML

```yaml

Description: OpenSearch Serverless security policy template
Resources:
  TestSecurityConfig:
      Type: 'AWS::OpenSearchService::Domain'
      Properties:
        Name: my-provider
        Type: saml
        Description: Serverless SAML configuration
        SamlOptions:
          Metadata: >-
             <?xml
             version="1.0" encoding="UTF-8"?><md:EntityDescriptor
             entityID="http://www.okta.com/foobar"
             xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata"><md:IDPSSODescriptor
             WantAuthnRequestsSigned="false"
             protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><md:KeyDescriptor
             use="signing"><ds:KeyInfo
             xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><ds:X509Data><ds:X509Certificate>Mfoobar</ds:X509Certificate></ds:X509Data></ds:KeyInfo></md:KeyDescriptor><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat><md:SingleSignOnService
             Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
             Location="https://trial-1234567.okta.com/app/trial-1234567_saml2_1/foobar/sso/saml"/><md:SingleSignOnService
             Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
             Location="https://trial-1234567.okta.com/app/trial-1234567_saml2_1/foobar/sso/saml"/></md:IDPSSODescriptor></md:EntityDescriptor>
             UserAttribute: 'http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier'
             GroupAttribute: ALLGroups SessionTimeout: 120
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::OpenSearchServerless::LifecyclePolicy

IamFederationConfigOptions
