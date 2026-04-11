# Handling exceptions

An AWS Certificate Manager command might fail for several reasons. For information about each exception,
see the table below.

## Private certificate exception handling

The following exceptions can occur when you attempt to renew a private PKI certificate
issued by AWS Private CA.

###### Note

AWS Private CA is not supported in the China (Beijing) Region and the China (Ningxia)
Region.

ACM failure code

Comment

`PCA_ACCESS_DENIED`

The private CA has not granted ACM permissions. This triggers a
AWS Private CA `AccessDeniedException` failure code.

To remedy the problem, grant the necessary permissions to the ACM service
principal using the AWS Private CA [CreatePermission](../../../../reference/privateca/latest/apireference/api-createpermission.md) operation.

`PCA_INVALID_DURATION`

The validity period of the requested certificate exceeds the validity period
of the issuing private CA. This triggers a AWS Private CA
`ValidationException` failure code.

To remedy the problem, [install a\
new CA certificate](../../../privateca/latest/userguide/pcacertinstall.md) with an appropriate validity period.

`PCA_INVALID_STATE`

The private CA being called is not in the correct state to perform the
requested ACM operation. This triggers a AWS Private CA
`InvalidStateException` failure code.

Resolve the issue as follows:

- If the CA has the status `CREATING`, wait for creation to
finish and then install the CA certificate.

- If the CA has status `PENDING_CERTIFICATE`, install the CA
certificate.

- If the CA has status `DISABLED`, update it to
`ACTIVE` status.

- If the CA has status `DELETED`, restore it.

- If the CA has status `EXPIRED`, install a new
certificate

- If the CA has status `FAILED`, and you cannot resolve the
issue, contact [Support](https://console.aws.amazon.com/support/home).

`PCA_LIMIT_EXCEEDED`

The private CA has reached an issuance quota. This triggers a AWS Private CA
`LimitExceededException` failure code. Try repeating your request
before proceeding with this help.

If the error persists, contact [Support](https://console.aws.amazon.com/support/home) to request a quota increase.

`PCA_REQUEST_FAILED`

A network or system error occurred. This triggers a AWS Private CA
`RequestFailedException` failure code. Try repeating your request
before proceeding with this help.

If the error persists, contact [Support](https://console.aws.amazon.com/support/home).

`PCA_RESOURCE_NOT_FOUND`

The private CA has been permanently deleted. This triggers a AWS Private CA
`ResourceNotFoundException` failure code. Verify that you used the
correct ARN. If that fails, you won't be able to use this CA.

To remedy the problem, [create a new\
CA](../../../privateca/latest/userguide/pcacreateca.md).

`SLR_NOT_FOUND`In order to renew a certificate signed by a private CA that resides in another
account, ACM requires a Service Linked Role (SLR) on the account where the
certificate resides. If you need to recreate a deleted SLR, see [Creating the SLR for ACM](acm-slr.md#create-slr).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Problems with the ACM service-linked role (SLR)

Quotas

All content copied from https://docs.aws.amazon.com/.
