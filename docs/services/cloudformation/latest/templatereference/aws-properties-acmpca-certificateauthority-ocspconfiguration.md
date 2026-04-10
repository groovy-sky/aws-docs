This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ACMPCA::CertificateAuthority OcspConfiguration

Contains information to enable and configure Online Certificate Status Protocol (OCSP)
for validating certificate revocation status.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "OcspCustomCname" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  OcspCustomCname: String

```

## Properties

`Enabled`

Flag enabling use of the Online Certificate Status Protocol (OCSP) for validating
certificate revocation status.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OcspCustomCname`

By default, AWS Private CA injects an Amazon domain into certificates being
validated by the Online Certificate Status Protocol (OCSP). A customer can alternatively
use this object to define a CNAME specifying a customized OCSP domain.

###### Note

The content of a Canonical Name (CNAME) record must conform to [RFC2396](https://www.ietf.org/rfc/rfc2396.txt) restrictions on the
use of special characters in URIs. Additionally, the value of the CNAME must not
include a protocol prefix such as "http://" or "https://".

_Required_: No

_Type_: String

_Pattern_: `[-a-zA-Z0-9;/?:@&=+$,%_.!~*()']*`

_Minimum_: `0`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyUsage

OtherName

All content copied from https://docs.aws.amazon.com/.
