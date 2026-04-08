# RegistrationOutput

Describes the status of an attempt from Amazon AppFlow to register a resource.

When you run a flow that you've configured to use a metadata catalog, Amazon AppFlow
registers a metadata table and data partitions with that catalog. This operation provides the
status of that registration attempt. The operation also indicates how many related resources
Amazon AppFlow created or updated.

## Contents

**message**

Explains the status of the registration attempt from Amazon AppFlow. If the attempt
fails, the message explains why.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: No

**result**

Indicates the number of resources that Amazon AppFlow created or updated. Possible
resources include metadata tables and data partitions.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: No

**status**

Indicates the status of the registration attempt from Amazon AppFlow.

Type: String

Valid Values: `InProgress | Successful | Error | CancelStarted | Canceled`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/registrationoutput.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/registrationoutput.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/registrationoutput.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RedshiftMetadata

S3DestinationProperties
