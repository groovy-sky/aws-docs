# KeyNameConstraint

If provided, the generated manifest includes only source bucket objects whose object
keys match the string constraints specified for `MatchAnyPrefix`,
`MatchAnySuffix`, and `MatchAnySubstring`.

## Contents

**MatchAnyPrefix**

If provided, the generated manifest includes objects where the specified string appears
at the start of the object key string. Each KeyNameConstraint filter accepts an array of strings with a length of 1 string.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**MatchAnySubstring**

If provided, the generated manifest includes objects where the specified string appears
anywhere within the object key string. Each KeyNameConstraint filter accepts an array of strings with a length of 1 string.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**MatchAnySuffix**

If provided, the generated manifest includes objects where the specified string appears
at the end of the object key string. Each KeyNameConstraint filter accepts an array of strings with a length of 1 string.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/keynameconstraint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/keynameconstraint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/keynameconstraint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobTimers

LambdaInvokeOperation

All content copied from https://docs.aws.amazon.com/.
