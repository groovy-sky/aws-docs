# FpgaImageState

Describes the state of the bitstream generation process for an Amazon FPGA image (AFI).

## Contents

**code**

The state. The following are the possible values:

- `pending` \- AFI bitstream generation is in progress.

- `available` \- The AFI is available for use.

- `failed` \- AFI bitstream generation failed.

- `unavailable` \- The AFI is no longer available for use.

Type: String

Valid Values: `pending | failed | available | unavailable`

Required: No

**message**

If the state is `failed`, this is the error message.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/fpgaimagestate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/fpgaimagestate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/fpgaimagestate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FpgaImageAttribute

FpgaInfo
