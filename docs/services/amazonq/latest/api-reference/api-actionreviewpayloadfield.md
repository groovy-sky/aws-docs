# ActionReviewPayloadField

A user input field in an plugin action review payload.

## Contents

**allowedFormat**

The expected data format for the action review input field value. For example, in PTO
request, `from` and `to` would be of `datetime` allowed
format.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**allowedValues**

Information about the field values that an end user can use to provide to
Amazon Q Business for Amazon Q Business to perform the requested plugin action.

Type: Array of [ActionReviewPayloadFieldAllowedValue](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_ActionReviewPayloadFieldAllowedValue.html) objects

Required: No

**arrayItemJsonSchema**

Use to create a custom form with array fields (fields with nested objects inside an
array).

Type: JSON value

Required: No

**displayDescription**

The field level description of each action review input field. This could be an
explanation of the field. In the Amazon Q Business web experience, these descriptions could
be used to display as tool tips to help users understand the field.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**displayName**

The name of the field.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**displayOrder**

The display order of fields in a payload.

Type: Integer

Required: No

**required**

Information about whether the field is required.

Type: Boolean

Required: No

**type**

The type of field.

Type: String

Valid Values: `STRING | NUMBER | ARRAY | BOOLEAN`

Required: No

**value**

The field value.

Type: JSON value

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/ActionReviewPayloadField)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/ActionReviewPayloadField)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/ActionReviewPayloadField)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ActionReviewEvent

ActionReviewPayloadFieldAllowedValue
