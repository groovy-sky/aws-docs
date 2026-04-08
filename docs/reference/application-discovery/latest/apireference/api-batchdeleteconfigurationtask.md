# BatchDeleteConfigurationTask

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

A metadata object that represents the deletion task being executed.

## Contents

**configurationType**

The type of configuration item to delete. Supported types are: SERVER.

Type: String

Valid Values: `SERVER`

Required: No

**deletedConfigurations**

The list of configuration IDs that were successfully deleted by the deletion task.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: No

**deletionWarnings**

A list of configuration IDs that produced warnings regarding their deletion, paired with
a warning message.

Type: Array of [DeletionWarning](api-deletionwarning.md) objects

Required: No

**endTime**

An epoch seconds timestamp (UTC) of when the deletion task was completed or failed.

Type: Timestamp

Required: No

**failedConfigurations**

A list of configuration IDs that failed to delete during the deletion task, each paired
with an error message.

Type: Array of [FailedConfiguration](api-failedconfiguration.md) objects

Required: No

**requestedConfigurations**

The list of configuration IDs that were originally requested to be deleted by the
deletion task.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: No

**startTime**

An epoch seconds timestamp (UTC) of when the deletion task was started.

Type: Timestamp

Required: No

**status**

The current execution status of the deletion task. Valid status are: INITIALIZING \|
VALIDATING \| DELETING \| COMPLETED \| FAILED.

Type: String

Valid Values: `INITIALIZING | VALIDATING | DELETING | COMPLETED | FAILED`

Required: No

**taskId**

The deletion task's unique identifier.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/batchdeleteconfigurationtask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/batchdeleteconfigurationtask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/batchdeleteconfigurationtask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

BatchDeleteAgentError

BatchDeleteImportDataError
