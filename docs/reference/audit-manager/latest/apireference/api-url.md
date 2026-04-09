# URL

Short for uniform resource locator. A URL is used as a unique identifier to locate a
resource on the internet.

## Contents

**hyperlinkName**

The name or word that's used as a hyperlink to the URL.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `^[\w\W\s\S]*$`

Required: No

**link**

The unique identifier for the internet resource.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8192.

Pattern: `^(https?:\/\/)?(www\.)?[a-zA-Z0-9-_]+([\.]+[a-zA-Z]+)+[\/\w]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/url.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/url.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/url.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAssessmentFrameworkControlSet

ValidationExceptionField

All content copied from https://docs.aws.amazon.com/.
