# Certificate

A CA certificate for an AWS account.

For more information, see [Using SSL/TLS to encrypt a connection to a DB \
instance](../../../../services/amazonrds/latest/userguide/usingwithrds-ssl.md) in the _Amazon RDS User Guide_ and
[Using SSL/TLS to encrypt a connection to a DB cluster](../../../../services/amazonrds/latest/aurorauserguide/usingwithrds-ssl.md) in the _Amazon Aurora_
_User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**CertificateArn**

The Amazon Resource Name (ARN) for the certificate.

Type: String

Required: No

**CertificateIdentifier**

The unique key that identifies a certificate.

Type: String

Required: No

**CertificateType**

The type of the certificate.

Type: String

Required: No

**CustomerOverride**

Indicates whether there is an override for the default certificate identifier.

Type: Boolean

Required: No

**CustomerOverrideValidTill**

If there is an override for the default certificate identifier, when the override
expires.

Type: Timestamp

Required: No

**Thumbprint**

The thumbprint of the certificate.

Type: String

Required: No

**ValidFrom**

The starting date from which the certificate is valid.

Type: Timestamp

Required: No

**ValidTill**

The final date that the certificate continues to be valid.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/certificate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/certificate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/certificate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlueGreenDeploymentTask

CertificateDetails

All content copied from https://docs.aws.amazon.com/.
