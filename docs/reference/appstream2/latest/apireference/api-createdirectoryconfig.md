# CreateDirectoryConfig

Creates a Directory Config object in WorkSpaces Applications. This object includes the configuration information required to join fleets and image builders to Microsoft Active Directory domains.

## Request Syntax

```nohighlight

{
   "CertificateBasedAuthProperties": {
      "CertificateAuthorityArn": "string",
      "Status": "string"
   },
   "DirectoryName": "string",
   "OrganizationalUnitDistinguishedNames": [ "string" ],
   "ServiceAccountCredentials": {
      "AccountName": "string",
      "AccountPassword": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CertificateBasedAuthProperties](#API_CreateDirectoryConfig_RequestSyntax)**

The certificate-based authentication properties used to authenticate SAML 2.0 Identity
Provider (IdP) user identities to Active Directory domain-joined streaming instances.
Fallback is turned on by default when certificate-based authentication is **Enabled** . Fallback allows users to log in using their AD
domain password if certificate-based authentication is unsuccessful, or to unlock a
desktop lock screen. **Enabled\_no\_directory\_login\_fallback** enables certificate-based
authentication, but does not allow users to log in using their AD domain password. Users
will be disconnected to re-authenticate using certificates.

Type: [CertificateBasedAuthProperties](api-certificatebasedauthproperties.md) object

Required: No

**[DirectoryName](#API_CreateDirectoryConfig_RequestSyntax)**

The fully qualified name of the directory (for example, corp.example.com).

Type: String

Required: Yes

**[OrganizationalUnitDistinguishedNames](#API_CreateDirectoryConfig_RequestSyntax)**

The distinguished names of the organizational units for computer accounts.

Type: Array of strings

Length Constraints: Maximum length of 2000.

Required: Yes

**[ServiceAccountCredentials](#API_CreateDirectoryConfig_RequestSyntax)**

The credentials for the service account used by the fleet or image builder to connect to the directory.

Type: [ServiceAccountCredentials](api-serviceaccountcredentials.md) object

Required: No

## Response Syntax

```nohighlight

{
   "DirectoryConfig": {
      "CertificateBasedAuthProperties": {
         "CertificateAuthorityArn": "string",
         "Status": "string"
      },
      "CreatedTime": number,
      "DirectoryName": "string",
      "OrganizationalUnitDistinguishedNames": [ "string" ],
      "ServiceAccountCredentials": {
         "AccountName": "string",
         "AccountPassword": "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DirectoryConfig](#API_CreateDirectoryConfig_ResponseSyntax)**

Information about the directory configuration.

Type: [DirectoryConfig](api-directoryconfig.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidAccountStatusException**

The resource cannot be created because your AWS account is suspended. For assistance, contact AWS Support.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidRoleException**

The specified role is invalid.

**Message**

The error message in the exception.

HTTP Status Code: 400

**LimitExceededException**

The requested limit exceeds the permitted limit for an account.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceAlreadyExistsException**

The specified resource already exists.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for Python](../../../../services/goto/boto3/appstream-2016-12-01/createdirectoryconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/createdirectoryconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateApplication

CreateEntitlement
