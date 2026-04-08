# SAPODataConnectorProfileProperties

The connector-specific profile properties required when using SAPOData.

## Contents

**applicationHostUrl**

The location of the SAPOData resource.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

Required: Yes

**applicationServicePath**

The application path to catalog service.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**clientNumber**

The client number for the client creating the connection.

Type: String

Length Constraints: Fixed length of 3.

Pattern: `^\d{3}$`

Required: Yes

**portNumber**

The port number of the SAPOData instance.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 65535.

Required: Yes

**disableSSO**

If you set this parameter to `true`, Amazon AppFlow bypasses the single
sign-on (SSO) settings in your SAP account when it accesses your SAP OData instance.

Whether you need this option depends on the types of credentials that you applied to your
SAP OData connection profile. If your profile uses basic authentication credentials, SAP SSO
can prevent Amazon AppFlow from connecting to your account with your username and
password. In this case, bypassing SSO makes it possible for Amazon AppFlow to connect
successfully. However, if your profile uses OAuth credentials, this parameter has no
affect.

Type: Boolean

Required: No

**logonLanguage**

The logon language of SAPOData instance.

Type: String

Length Constraints: Maximum length of 2.

Pattern: `^[a-zA-Z0-9_]*$`

Required: No

**oAuthProperties**

The SAPOData OAuth properties required for OAuth type authentication.

Type: [OAuthProperties](api-oauthproperties.md) object

Required: No

**privateLinkServiceName**

The SAPOData Private Link service name to be used for private data transfers.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `^$|com.amazonaws.vpce.[\w/!:@#.\-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/sapodataconnectorprofileproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/sapodataconnectorprofileproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/sapodataconnectorprofileproperties.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SAPODataConnectorProfileCredentials

SAPODataDestinationProperties
