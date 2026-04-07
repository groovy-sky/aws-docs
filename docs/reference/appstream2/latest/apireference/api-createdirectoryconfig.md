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

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/appstream2/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[CertificateBasedAuthProperties](#API_CreateDirectoryConfig_RequestSyntax)**

The certificate-based authentication properties used to authenticate SAML 2.0 Identity
Provider (IdP) user identities to Active Directory domain-joined streaming instances.
Fallback is turned on by default when certificate-based authentication is **Enabled** . Fallback allows users to log in using their AD
domain password if certificate-based authentication is unsuccessful, or to unlock a
desktop lock screen. **Enabled\_no\_directory\_login\_fallback** enables certificate-based
authentication, but does not allow users to log in using their AD domain password. Users
will be disconnected to re-authenticate using certificates.

Type: [CertificateBasedAuthProperties](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_CertificateBasedAuthProperties.html) object

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

Type: [ServiceAccountCredentials](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ServiceAccountCredentials.html) object

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

Type: [DirectoryConfig](https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/appstream2/latest/APIReference/CommonErrors.html).

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appstream-2016-12-01/CreateDirectoryConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appstream-2016-12-01/CreateDirectoryConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateApplication

CreateEntitlement
