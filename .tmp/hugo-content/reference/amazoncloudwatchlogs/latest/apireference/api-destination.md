# Destination

Represents a cross-account destination that receives subscription log events.

## Contents

**accessPolicy**

An IAM policy document that governs which AWS accounts can create
subscription filters against this destination.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**arn**

The ARN of this destination.

Type: String

Required: No

**creationTime**

The creation time of the destination, expressed as the number of milliseconds after Jan
1, 1970 00:00:00 UTC.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**destinationName**

The name of the destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**roleArn**

A role for impersonation, used when delivering log events to the target.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**targetArn**

The Amazon Resource Name (ARN) of the physical target where the log events are
delivered (for example, a Kinesis stream).

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/destination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/destination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/destination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeliverySource

DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
