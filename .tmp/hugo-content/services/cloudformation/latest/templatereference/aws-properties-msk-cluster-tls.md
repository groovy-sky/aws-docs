This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster Tls

Details for client authentication using TLS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateAuthorityArnList" : [ String, ... ],
  "Enabled" : Boolean
}

```

### YAML

```yaml

  CertificateAuthorityArnList:
    - String
  Enabled: Boolean

```

## Properties

`CertificateAuthorityArnList`

List of AWS Private CA ARNs.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

TLS authentication is enabled or not.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageInfo

Unauthenticated

All content copied from https://docs.aws.amazon.com/.
