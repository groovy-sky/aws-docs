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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/audiosourcedetails.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/audiosourcedetails.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/audiosourcedetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioExtractionConfiguration

AuthChallengeRequest

All content copied from https://docs.aws.amazon.com/.
