# ControlComment

A comment that's posted by a user on a control. This includes the author's name, the
comment text, and a timestamp.

## Contents

**authorName**

The name of the user who authored the comment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^[a-zA-Z0-9-_()\s\+=,.@]+$`

Required: No

**commentBody**

The body text of a control comment.

Type: String

Length Constraints: Maximum length of 500.

Pattern: `^[\w\W\s\S]*$`

Required: No

**postedDate**

The time when the comment was posted.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/ControlComment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/ControlComment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/ControlComment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Control

ControlDomainInsights
