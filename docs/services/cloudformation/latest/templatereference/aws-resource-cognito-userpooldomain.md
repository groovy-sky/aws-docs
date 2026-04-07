This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolDomain

The AWS::Cognito::UserPoolDomain resource creates a new domain for a user pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolDomain",
  "Properties" : {
      "CustomDomainConfig" : CustomDomainConfigType,
      "Domain" : String,
      "ManagedLoginVersion" : Integer,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolDomain
Properties:
  CustomDomainConfig:
    CustomDomainConfigType
  Domain: String
  ManagedLoginVersion: Integer
  UserPoolId: String

```

## Properties

`CustomDomainConfig`

The configuration for a custom domain that hosts the sign-up and sign-in pages for
your application. Use this object to specify an SSL certificate that is managed by
ACM.

When you create a custom domain, the passkey RP ID defaults to the custom domain. If
you had a prefix domain active, this will cause passkey integration for your prefix
domain to stop working due to a mismatch in RP ID. To keep the prefix domain passkey
integration working, you can explicitly set RP ID to the prefix domain.

_Required_: No

_Type_: [CustomDomainConfigType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpooldomain-customdomainconfigtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The name of the domain that you want to update. For custom domains, this is the
fully-qualified domain name, for example `auth.example.com`. For prefix
domains, this is the prefix alone, such as `myprefix`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedLoginVersion`

A version number that indicates the state of managed login for your domain. Version
`1` is hosted UI (classic). Version `2` is the newer managed
login with the branding editor. For more information, see [Managed login](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-managed-login.html).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool that is associated with the domain you're updating.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns physicalResourceId, which is “Domain". For
example:

`{ "Ref": "your-test-domain" }`

For the Amazon Cognito user pool domain `your-test-domain`, Ref returns the
name of the user pool domain.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CloudFrontDistribution`

The Amazon CloudFront endpoint that you use as the target of the alias that you set up
with your Domain Name Service (DNS) provider.

## Examples

- [Creating a new custom domain for a user pool](#aws-resource-cognito-userpooldomain--examples--Creating_a_new_custom_domain_for_a_user_pool)

- [Creating a new default domain for a user pool](#aws-resource-cognito-userpooldomain--examples--Creating_a_new_default_domain_for_a_user_pool)

### Creating a new custom domain for a user pool

The following example creates a custom domain, "my-test-user-pool-domain", in
the referenced user pool.

#### JSON

```json

{
   "UserPoolDomain":{
      "Type":"AWS::Cognito::UserPoolDomain",
      "Properties":{
         "UserPoolId":{
            "Ref":"UserPool"
         },
         "Domain":"my-test-user-pool-domain.myapplication.com",
         "ManagedLoginVersion": "2",
         "CustomDomainConfig":{
            "CertificateArn":{
               "Ref":"CertificateArn"
            }
         }
      }
   }
}
```

#### YAML

```yaml

UserPoolDomain:
  Type: AWS::Cognito::UserPoolDomain
  Properties:
    UserPoolId: !Ref UserPool
    Domain: "my-test-user-pool-domain.myapplication.com"
    ManagedLoginVersion: "2"
    CustomDomainConfig:
      CertificateArn: !Ref CertificateArn
```

### Creating a new default domain for a user pool

The following example creates a new default domain,
"my-test-user-pool-domain", in the referenced user pool.

#### JSON

```json

{
   "UserPoolDomain":{
      "Type":"AWS::Cognito::UserPoolDomain",
      "Properties":{
         "UserPoolId":{
            "Ref":"UserPool"
         },
         "Domain":"my-test-user-pool-domain",
         "ManagedLoginVersion": "2"
      }
   }
}
```

#### YAML

```yaml

UserPoolDomain:
  Type: AWS::Cognito::UserPoolDomain
  Properties:
    UserPoolId: !Ref UserPool
    Domain: "my-test-user-pool-domain"
    ManagedLoginVersion: "2"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TokenValidityUnits

CustomDomainConfigType
