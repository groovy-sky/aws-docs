# FrameworkMetadata

The metadata of a framework, such as the name, ID, or description.

## Contents

**complianceType**

The compliance standard that's associated with the framework. For example, this could
be PCI DSS or HIPAA.

Type: String

Length Constraints: Maximum length of 100.

Pattern: `^[\w\W\s\S]*$`

Required: No

**description**

The description of the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `^[\w\W\s\S]*$`

Required: No

**logo**

The logo that's associated with the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[\w,\s-]+\.[A-Za-z]+$`

Required: No

**name**

The name of the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/FrameworkMetadata)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/FrameworkMetadata)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/FrameworkMetadata)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Framework

Insights
