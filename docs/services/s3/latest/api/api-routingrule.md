# RoutingRule

Specifies the redirect behavior and when a redirect is applied. For more information about routing
rules, see [Configuring\
advanced conditional redirects](../dev/how-to-page-redirect.md#advanced-conditional-redirects) in the _Amazon S3 User Guide_.

## Contents

**Redirect**

Container for redirect information. You can redirect requests to another host, to another page, or
with another protocol. In the event of an error, you can specify a different error code to
return.

Type: [Redirect](api-redirect.md) data type

Required: Yes

**Condition**

A container for describing a condition that must be met for the specified redirect to apply. For
example, 1. If request is for pages in the `/docs` folder, redirect to the
`/documents` folder. 2. If request results in HTTP error 4xx, redirect request to another
host where you might process the error.

Type: [Condition](api-condition.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/routingrule.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/routingrule.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/routingrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreStatus

Rule

All content copied from https://docs.aws.amazon.com/.
