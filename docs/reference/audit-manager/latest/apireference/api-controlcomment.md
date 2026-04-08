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

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/controlcomment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/controlcomment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/controlcomment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Control

ControlDomainInsights
