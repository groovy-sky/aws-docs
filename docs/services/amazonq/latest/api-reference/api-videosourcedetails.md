# VideoSourceDetails

Details about a video source, including its identifier, format, and time information.

## Contents

**endTimeMilliseconds**

The ending timestamp in milliseconds for the relevant video segment.

Type: Long

Required: No

**mediaId**

Unique identifier for the video media file.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**mediaMimeType**

The MIME type of the video file (e.g., video/mp4, video/avi).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**startTimeMilliseconds**

The starting timestamp in milliseconds for the relevant video segment.

Type: Long

Required: No

**videoExtractionType**

The type of video extraction performed on the content.

Type: String

Valid Values: `TRANSCRIPT | SUMMARY`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/VideoSourceDetails)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/VideoSourceDetails)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/VideoSourceDetails)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VideoExtractionConfiguration

WebExperience
