# AudioSourceDetails

Details about an audio source, including its identifier, format, and time information.

## Contents

**audioExtractionType**

The type of audio extraction performed on the content.

Type: String

Valid Values: `TRANSCRIPT | SUMMARY`

Required: No

**endTimeMilliseconds**

The ending timestamp in milliseconds for the relevant audio segment.

Type: Long

Required: No

**mediaId**

Unique identifier for the audio media file.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**mediaMimeType**

The MIME type of the audio file (e.g., audio/mp3, audio/wav).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**startTimeMilliseconds**

The starting timestamp in milliseconds for the relevant audio segment.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/AudioSourceDetails)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/AudioSourceDetails)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/AudioSourceDetails)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AudioExtractionConfiguration

AuthChallengeRequest
