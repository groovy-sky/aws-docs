This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority CrlDistributionPointExtensionConfiguration

Contains configuration information for the default behavior of the CRL Distribution Point (CDP) extension in certificates issued by your CA. This extension
contains a link to download the CRL, so you can check whether a certificate has been revoked. To choose whether you want this extension
omitted or not in certificates issued by your CA, you can set the **OmitExtension** parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OmitExtension" : Boolean
}

```

### YAML

```yaml

  OmitExtension: Boolean

```

## Properties

`OmitExtension`

Configures whether the CRL Distribution Point extension should be populated with the default URL to the CRL. If set to `true`, then the CDP extension will
not be present in any certificates issued by that CA unless otherwise specified through CSR or API passthrough.

###### Note

Only set this if you have another way to distribute the CRL Distribution Points for certificates issued by your CA, such as the Matter Distributed Compliance Ledger.

This configuration cannot be enabled with a custom CNAME set.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CrlConfiguration

CsrExtensions

All content copied from https://docs.aws.amazon.com/.
