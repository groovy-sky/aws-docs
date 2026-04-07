This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::OIDCProvider

Creates or updates an IAM entity to describe an identity provider (IdP)
that supports [OpenID Connect (OIDC)](http://openid.net/connect).

The OIDC provider that you create with this operation can be used as a principal in a
role's trust policy. Such a policy establishes a trust relationship between AWS and the OIDC provider.

When you create the IAM OIDC provider, you specify the following:

- The URL of the OIDC identity provider (IdP) to trust

- A list of client IDs (also known as audiences) that identify the application or
applications that are allowed to authenticate using the OIDC provider

- A list of tags that are attached to the specified IAM OIDC provider

- A list of thumbprints of one or more server certificates that the IdP uses

You get all of this information from the OIDC IdP that you want to use to access AWS.

When you update the IAM OIDC provider, you specify the following:

- The URL of the OIDC identity provider (IdP) to trust

- A list of client IDs (also known as audiences) that replaces the existing list of
client IDs associated with the OIDC IdP

- A list of tags that replaces the existing list of tags attached to the specified
IAM OIDC provider

- A list of thumbprints that replaces the existing list of server certificates
thumbprints that the IdP uses

###### Note

The trust for the OIDC provider is derived from the IAM provider that
this operation creates. Therefore, it is best to limit access to the [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html) operation to highly privileged users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::OIDCProvider",
  "Properties" : {
      "ClientIdList" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "ThumbprintList" : [ String, ... ],
      "Url" : String
    }
}

```

### YAML

```yaml

Type: AWS::IAM::OIDCProvider
Properties:
  ClientIdList:
    - String
  Tags:
    - Tag
  ThumbprintList:
    - String
  Url: String

```

## Properties

`ClientIdList`

A list of client IDs (also known as audiences) that are associated with the specified
IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html).

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags that are attached to the specified IAM OIDC provider. The
returned list of tags is sorted by tag key. For more information about tagging, see [Tagging IAM\
resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the _IAM User Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iam-oidcprovider-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThumbprintList`

A list of certificate thumbprints that are associated with the specified IAM OIDC provider resource object. For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html).

This property is optional. If it is not included, IAM will retrieve and
use the top intermediate certificate authority (CA) thumbprint of the OpenID Connect
identity provider server certificate.

_Required_: No

_Type_: Array of String

_Minimum_: `40`

_Maximum_: `40 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL that the IAM OIDC provider resource object is associated with.
For more information, see [CreateOpenIDConnectProvider](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the specified
`AWS::IAM::OIDCProvider` resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IAM::ManagedPolicy

Tag
