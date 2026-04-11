# SalesforceConnectorProfileProperties

The connector-specific profile properties required when using Salesforce.

## Contents

**instanceUrl**

The location of the Salesforce resource.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: No

**isSandboxEnvironment**

Indicates whether the connector profile applies to a sandbox or production environment.

Type: Boolean

Required: No

**usePrivateLinkForMetadataAndAuthorization**

If the connection mode for the connector profile is private, this parameter sets whether
Amazon AppFlow uses the private network to send metadata and authorization calls to
Salesforce. Amazon AppFlow sends private calls through AWS PrivateLink. These
calls travel through AWS infrastructure without being exposed to the public
internet.

Set either of the following values:

true

Amazon AppFlow sends all calls to Salesforce over the private network.

These private calls are:

- Calls to get metadata about your Salesforce records. This metadata describes
your Salesforce objects and their fields.

- Calls to get or refresh access tokens that allow Amazon AppFlow to access
your Salesforce records.

- Calls to transfer your Salesforce records as part of a flow run.

false

The default value. Amazon AppFlow sends some calls to Salesforce privately and
other calls over the public internet.

The public calls are:

- Calls to get metadata about your Salesforce records.

- Calls to get or refresh access tokens.

The private calls are:

- Calls to transfer your Salesforce records as part of a flow run.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/salesforceconnectorprofileproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/salesforceconnectorprofileproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/salesforceconnectorprofileproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SalesforceConnectorProfileCredentials

SalesforceDestinationProperties

All content copied from https://docs.aws.amazon.com/.
