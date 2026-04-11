# WebsiteConfiguration

Specifies website configuration parameters for an Amazon S3 bucket.

## Contents

**ErrorDocument**

The name of the error document for the website.

Type: [ErrorDocument](api-errordocument.md) data type

Required: No

**IndexDocument**

The name of the index document for the website.

Type: [IndexDocument](api-indexdocument.md) data type

Required: No

**RedirectAllRequestsTo**

The redirect behavior for every request to this bucket's website endpoint.

###### Important

If you specify this property, you can't specify any other property.

Type: [RedirectAllRequestsTo](api-redirectallrequeststo.md) data type

Required: No

**RoutingRules**

Rules that define when a redirect is applied and the redirect behavior.

Type: Array of [RoutingRule](api-routingrule.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/websiteconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/websiteconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/websiteconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VersioningConfiguration

Amazon S3 Control

All content copied from https://docs.aws.amazon.com/.
