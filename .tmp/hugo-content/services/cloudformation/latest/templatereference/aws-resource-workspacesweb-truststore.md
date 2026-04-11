This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::TrustStore

This resource specifies a trust store that can be associated with a web portal. A trust store contains
certificate authority (CA) certificates. Once associated with a web portal, the browser in
a streaming session will recognize certificates that have been issued using any of the CAs
in the trust store. If your organization has internal websites that use certificates issued
by private CAs, you should add the private CA certificate to the trust store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::TrustStore",
  "Properties" : {
      "CertificateList" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::TrustStore
Properties:
  CertificateList:
    - String
  Tags:
    - Tag

```

## Properties

`CertificateList`

A list of CA certificates to be added to the trust store.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the trust store. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesweb-truststore-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that this trust store is associated with.

`TrustStoreArn`

The ARN of the trust store.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
