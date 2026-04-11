# BatchCreateCategoryInputCategory

The category object to be created.

## Contents

**title**

The name of the category.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 30.

Pattern: `[a-zA-Z0-9_]+( [a-zA-Z0-9_]+)*`

Required: Yes

**color**

The color to be associated with a category. The color must be a hexadecimal value of
either 3 or 6 digits.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 7.

Pattern: `#([A-Fa-f0-9]{3}|[A-Fa-f0-9]{6})`

Required: No

**id**

The unique identifier to be associated with a category. If you don't include a value, the
category is automatically assigned a unique identifier.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/batchcreatecategoryinputcategory.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/batchcreatecategoryinputcategory.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/batchcreatecategoryinputcategory.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeFilter

Card

All content copied from https://docs.aws.amazon.com/.
