This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::Application

Creates an OAuth 2.0 customer managed application in IAM Identity Center for the given
application provider.

###### Note

This API does not support creating SAML 2.0 customer managed applications or AWS
managed applications. To learn how to create an AWS managed application, see the
application user guide. You can create a SAML 2.0 customer managed application in
the AWS Management Console only. See [Setting\
up customer managed SAML 2.0 applications](../../../singlesignon/latest/userguide/customermanagedapps-saml2-setup.md). For more information on these
application types, see [AWS managed\
applications](../../../singlesignon/latest/userguide/awsapps.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSO::Application",
  "Properties" : {
      "ApplicationProviderArn" : String,
      "Description" : String,
      "InstanceArn" : String,
      "Name" : String,
      "PortalOptions" : PortalOptionsConfiguration,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSO::Application
Properties:
  ApplicationProviderArn: String
  Description: String
  InstanceArn: String
  Name: String
  PortalOptions:
    PortalOptionsConfiguration
  Status: String
  Tags:
    - Tag

```

## Properties

`ApplicationProviderArn`

The ARN of the application provider for this application.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]{1,5}){0,3}:sso::aws:applicationProvider/[a-zA-Z0-9-/]+$`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The ARN of the instance of IAM Identity Center that is configured with this application.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]{1,5}){0,3}:sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the application.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w+=,.@-]+$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalOptions`

A structure that describes the options for the access portal associated with this
application.

_Required_: No

_Type_: [PortalOptionsConfiguration](aws-properties-sso-application-portaloptionsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the application in this instance of IAM Identity Center.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies tags to be attached to the application.

_Required_: No

_Type_: Array of [Tag](aws-properties-sso-application-tag.md)

_Maximum_: `75`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, combined by all fields with the delimiter
`|`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`ApplicationArn`

The ARN of the application.

## Examples

### Creating an application in IAM Identity Center

The following example creates a new custom application with an Application URL
sign-in option.

#### JSON

```json

{
  "Type" : "AWS::SSO::Application",
  "Properties" : {
    "ApplicationProviderArn" : "arn:sso::aws:applicationProvider/example",
    "Description" : "This is a sample application",
    "InstanceArn" : "arn:aws:sso:::instance/ssoins-instanceId",
    "Name" : "Application",
    "PortalOptions" : {
      "SignInOptions" : {
        "ApplicationUrl" : "http://www.example.com",
        "Origin" : "APPLICATION"
      },
      "Visibility" : "ENABLED"
    },
    "Status" : "ENABLED",
    "Tags": [
      {
        "Key": "tagKey",
        "Value": "tagValue"
      }
    ]
  }
}
```

#### YAML

```yaml

Type: AWS::SSO::Application
Properties:
  ApplicationProviderArn: arn:sso::aws:applicationProvider/example
  Description: This is a sample application
  InstanceArn: arn:aws:sso:::instance/ssoins-instanceId
  Name: Application
  PortalOptions:
    SignInOptions:
      ApplicationUrl: http://www.example.com
      Origin: APPLICATION
    Visibility: ENABLED
  Status: ENABLED
  Tags:
  - Key: tagKey
    Value: tagValue
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IAM Identity Center

PortalOptionsConfiguration

All content copied from https://docs.aws.amazon.com/.
