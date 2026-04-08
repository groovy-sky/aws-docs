# ModifyCertificates

Override the system-default Secure Sockets Layer/Transport Layer Security (SSL/TLS)
certificate for Amazon RDS for new DB instances, or remove the override.

By using this operation, you can specify an RDS-approved SSL/TLS certificate for new DB
instances that is different from the default certificate provided by RDS. You can also
use this operation to remove the override, so that new DB instances use the default
certificate provided by RDS.

You might need to override the default certificate in the following situations:

- You already migrated your applications to support the latest certificate authority (CA) certificate, but the new CA certificate is not yet
the RDS default CA certificate for the specified AWS Region.

- RDS has already moved to a new default CA certificate for the specified AWS
Region, but you are still in the process of supporting the new CA certificate.
In this case, you temporarily need additional time to finish your application
changes.

For more information about rotating your SSL/TLS certificate for RDS DB engines, see
[Rotating Your SSL/TLS Certificate](../../../../services/amazonrds/latest/userguide/usingwithrds-ssl-certificate-rotation.md) in the _Amazon RDS User Guide_.

For more information about rotating your SSL/TLS certificate for Aurora DB engines, see
[Rotating Your SSL/TLS Certificate](../../../../services/amazonrds/latest/aurorauserguide/usingwithrds-ssl-certificate-rotation.md) in the _Amazon Aurora User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**CertificateIdentifier**

The new default certificate identifier to override the current one with.

To determine the valid values, use the `describe-certificates`
AWS CLI
command or the `DescribeCertificates` API operation.

Type: String

Required: No

**RemoveCustomerOverride**

Specifies whether to remove the override for the default certificate.
If the override is removed, the default certificate is the system
default.

Type: Boolean

Required: No

## Response Elements

The following element is returned by the service.

**Certificate**

A CA certificate for an AWS account.

For more information, see [Using SSL/TLS to encrypt a connection to a DB \
instance](../../../../services/amazonrds/latest/userguide/usingwithrds-ssl.md) in the _Amazon RDS User Guide_ and
[Using SSL/TLS to encrypt a connection to a DB cluster](../../../../services/amazonrds/latest/aurorauserguide/usingwithrds-ssl.md) in the _Amazon Aurora_
_User Guide_.

Type: [Certificate](api-certificate.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CertificateNotFound**

`CertificateIdentifier` doesn't refer to an
existing certificate.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifycertificates.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifycertificates.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModifyActivityStream

ModifyCurrentDBClusterCapacity

All content copied from https://docs.aws.amazon.com/.
