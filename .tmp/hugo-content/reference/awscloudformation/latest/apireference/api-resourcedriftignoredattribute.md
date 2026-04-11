# ResourceDriftIgnoredAttribute

The `ResourceDriftIgnoredAttribute` data type.

## Contents

**Path**

Path of the resource attribute for which drift was ignored.

Type: String

Required: No

**Reason**

Reason why drift was ignored for the attribute, can have 2 possible values:

- `WRITE_ONLY_PROPERTY` \- Property is not included in read response for the
resource’s live state.

- `MANAGED_BY_AWS` \- Property is managed by an AWS service and is expected to be
dynamically modified.

Type: String

Valid Values: `MANAGED_BY_AWS | WRITE_ONLY_PROPERTY`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/resourcedriftignoredattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/resourcedriftignoredattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/resourcedriftignoredattribute.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceDetail

ResourceIdentifierSummary

All content copied from https://docs.aws.amazon.com/.
