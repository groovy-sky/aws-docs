This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginAccessControl OriginAccessControlConfig

Creates a new origin access control in CloudFront. After you create an origin access
control, you can add it to an origin in a CloudFront distribution so that CloudFront sends
authenticated (signed) requests to the origin.

This makes it possible to block public access to the origin, allowing viewers (users) to
access the origin's content only through CloudFront.

For more information about using a CloudFront origin access control, see [Restricting access to an AWS origin](../../../amazoncloudfront/latest/developerguide/private-content-restricting-access-to-origin.md) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Name" : String,
  "OriginAccessControlOriginType" : String,
  "SigningBehavior" : String,
  "SigningProtocol" : String
}

```

### YAML

```yaml

  Description: String
  Name: String
  OriginAccessControlOriginType: String
  SigningBehavior: String
  SigningProtocol: String

```

## Properties

`Description`

A description of the origin access control.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name to identify the origin access control. You can specify up to 64 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginAccessControlOriginType`

The type of origin that this origin access control is for.

_Required_: Yes

_Type_: String

_Pattern_: `^(s3|mediastore|lambda|mediapackagev2)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningBehavior`

Specifies which requests CloudFront signs (adds authentication information to). Specify
`always` for the most common use case. For more information, see [origin access control advanced settings](../../../amazoncloudfront/latest/developerguide/private-content-restricting-access-to-s3.md#oac-advanced-settings) in the
_Amazon CloudFront Developer Guide_.

This field can have one of the following values:

- `always` – CloudFront signs all origin requests, overwriting the
`Authorization` header from the viewer request if one
exists.

- `never` – CloudFront doesn't sign any origin requests. This value turns
off origin access control for all origins in all distributions that use this
origin access control.

- `no-override` – If the viewer request doesn't contain the
`Authorization` header, then CloudFront signs the origin request. If
the viewer request contains the `Authorization` header, then CloudFront
doesn't sign the origin request and instead passes along the
`Authorization` header from the viewer request. **WARNING: To pass along the `Authorization` header**
**from the viewer request, you _must_ add the**
**`Authorization` header to a [cache policy](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md) for all cache behaviors that**
**use origins associated with this origin access control.**

_Required_: Yes

_Type_: String

_Pattern_: `^(never|no-override|always)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningProtocol`

The signing protocol of the origin access control, which determines how CloudFront signs
(authenticates) requests. The only valid value is `sigv4`.

_Required_: Yes

_Type_: String

_Pattern_: `^(sigv4)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFront::OriginAccessControl

AWS::CloudFront::OriginRequestPolicy
