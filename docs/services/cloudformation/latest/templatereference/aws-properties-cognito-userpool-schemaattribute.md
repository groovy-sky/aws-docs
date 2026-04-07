This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool SchemaAttribute

A list of the user attributes and their properties in your user pool. The attribute
schema contains standard attributes, custom attributes with a `custom:`
prefix, and developer attributes with a `dev:` prefix. For more information,
see [User pool\
attributes](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html).

Developer-only `dev:` attributes are a legacy feature of user pools, and
are read-only to all app clients. You can create and update developer-only attributes
only with IAM-authenticated API operations. Use app client read/write permissions
instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeDataType" : String,
  "DeveloperOnlyAttribute" : Boolean,
  "Mutable" : Boolean,
  "Name" : String,
  "NumberAttributeConstraints" : NumberAttributeConstraints,
  "Required" : Boolean,
  "StringAttributeConstraints" : StringAttributeConstraints
}

```

### YAML

```yaml

  AttributeDataType: String
  DeveloperOnlyAttribute: Boolean
  Mutable: Boolean
  Name: String
  NumberAttributeConstraints:
    NumberAttributeConstraints
  Required: Boolean
  StringAttributeConstraints:
    StringAttributeConstraints

```

## Properties

`AttributeDataType`

The data format of the values for your attribute. When you choose an
`AttributeDataType`, Amazon Cognito validates the input against the data type. A
custom attribute value in your user's ID token is always a string, for example
`"custom:isMember" : "true"` or `"custom:YearsAsMember" :
                "12"`.

_Required_: No

_Type_: String

_Allowed values_: `String | Number | DateTime | Boolean`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeveloperOnlyAttribute`

###### Note

You should use [WriteAttributes](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolClientType.html#CognitoUserPools-Type-UserPoolClientType-WriteAttributes) in the user pool client to control how attributes can
be mutated for new use cases instead of using
`DeveloperOnlyAttribute`.

Specifies whether the attribute type is developer only. This attribute can only be
modified by an administrator. Users won't be able to modify this attribute using their
access token. For example, `DeveloperOnlyAttribute` can be modified using
AdminUpdateUserAttributes but can't be updated using UpdateUserAttributes.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mutable`

Specifies whether the value of the attribute can be changed.

Any user pool attribute whose value you map from an IdP attribute must be mutable,
with a parameter value of `true`. Amazon Cognito updates mapped attributes when users
sign in to your application through an IdP. If an attribute is immutable, Amazon Cognito throws
an error when it attempts to update the attribute. For more information, see [Specifying Identity Provider Attribute Mappings for Your User\
Pool](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.html).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of your user pool attribute. When you create or update a user pool, adding a
schema attribute creates a custom or developer-only attribute. When you add an attribute
with a `Name` value of `MyAttribute`, Amazon Cognito creates the custom
attribute `custom:MyAttribute`. When `DeveloperOnlyAttribute` is
`true`, Amazon Cognito creates your attribute as `dev:MyAttribute`. In
an operation that describes a user pool, Amazon Cognito returns this value as `value`
for standard attributes, `custom:value` for custom attributes, and
`dev:value` for developer-only attributes..

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberAttributeConstraints`

Specifies the constraints for an attribute of the number type.

_Required_: No

_Type_: [NumberAttributeConstraints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-numberattributeconstraints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Required`

Specifies whether a user pool attribute is required. If the attribute is required and
the user doesn't provide a value, registration or sign-in will fail.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringAttributeConstraints`

Specifies the constraints for an attribute of the string type.

_Required_: No

_Type_: [StringAttributeConstraints](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpool-stringattributeconstraints.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecoveryOption

SignInPolicy
