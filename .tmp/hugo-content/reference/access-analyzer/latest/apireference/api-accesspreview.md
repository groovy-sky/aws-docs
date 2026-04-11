# AccessPreview

Contains information about an access preview.

## Contents

**analyzerArn**

The ARN of the analyzer used to generate the access preview.

Type: String

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**configurations**

A map of resource ARNs for the proposed resource configuration.

Type: String to [Configuration](api-configuration.md) object map

Required: Yes

**createdAt**

The time at which the access preview was created.

Type: Timestamp

Required: Yes

**id**

The unique ID for the access preview.

Type: String

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**status**

The status of the access preview.

- `Creating` \- The access preview creation is in progress.

- `Completed` \- The access preview is complete. You can preview findings
for external access to the resource.

- `Failed` \- The access preview creation has failed.

Type: String

Valid Values: `COMPLETED | CREATING | FAILED`

Required: Yes

**statusReason**

Provides more details about the current status of the access preview.

For example, if the creation of the access preview fails, a `Failed` status
is returned. This failure can be due to an internal issue with the analysis or due to an
invalid resource configuration.

Type: [AccessPreviewStatusReason](api-accesspreviewstatusreason.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/accesspreview.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/accesspreview.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/accesspreview.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Access

AccessPreviewFinding
