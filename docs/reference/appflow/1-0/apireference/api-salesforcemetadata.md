# SalesforceMetadata

The connector metadata specific to Salesforce.

## Contents

**dataTransferApis**

The Salesforce APIs that you can have Amazon AppFlow use when your flows transfers
data to or from Salesforce.

Type: Array of strings

Valid Values: `AUTOMATIC | BULKV2 | REST_SYNC`

Required: No

**oauth2GrantTypesSupported**

The OAuth 2.0 grant types that Amazon AppFlow can use when it requests an access
token from Salesforce. Amazon AppFlow requires an access token each time it attempts to
access your Salesforce records.

AUTHORIZATION\_CODE

Amazon AppFlow passes an authorization code when it requests the access token
from Salesforce. Amazon AppFlow receives the authorization code from Salesforce
after you log in to your Salesforce account and authorize Amazon AppFlow to access
your records.

JWT\_BEARER

Amazon AppFlow passes a JSON web token (JWT) when it requests the access token
from Salesforce. You provide the JWT to Amazon AppFlow when you define the
connection to your Salesforce account. When you use this grant type, you don't need to
log in to your Salesforce account to authorize Amazon AppFlow to access your
records.

###### Note

The CLIENT\_CREDENTIALS value is not supported for Salesforce.

Type: Array of strings

Valid Values: `CLIENT_CREDENTIALS | AUTHORIZATION_CODE | JWT_BEARER`

Required: No

**oAuthScopes**

The desired authorization scope for the Salesforce account.

Type: Array of strings

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/salesforcemetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/salesforcemetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/salesforcemetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceDestinationProperties

SalesforceSourceProperties

All content copied from https://docs.aws.amazon.com/.
