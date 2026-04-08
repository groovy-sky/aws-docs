# ImportTask

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

An array of information related to the import task request that includes status
information, times, IDs, the Amazon S3 Object URL for the import file, and more.

## Contents

**applicationImportFailure**

The total number of application records in the import file that failed to be
imported.

Type: Integer

Required: No

**applicationImportSuccess**

The total number of application records in the import file that were successfully
imported.

Type: Integer

Required: No

**clientRequestToken**

A unique token used to prevent the same import request from occurring more than once. If
you didn't provide a token, a token was automatically generated when the import task request
was sent.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Required: No

**errorsAndFailedEntriesZip**

A link to a compressed archive folder (in the ZIP format) that contains an error log and a
file of failed records. You can use these two files to quickly identify records that failed,
why they failed, and correct those records. Afterward, you can upload the corrected file to
your Amazon S3 bucket and create another import task request.

This field also includes authorization information so you can confirm the authenticity of
the compressed archive before you download it.

If some records failed to be imported we recommend that you correct the records in the
failed entries file and then imports that failed entries file. This prevents you from having
to correct and update the larger original file and attempt importing it again.

Type: String

Required: No

**fileClassification**

The type of file detected by the import task.

Type: String

Valid Values: `MODELIZEIT_EXPORT | RVTOOLS_EXPORT | VMWARE_NSX_EXPORT | IMPORT_TEMPLATE`

Required: No

**importCompletionTime**

The time that the import task request finished, presented in the Unix time stamp
format.

Type: Timestamp

Required: No

**importDeletedTime**

The time that the import task request was deleted, presented in the Unix time stamp
format.

Type: Timestamp

Required: No

**importRequestTime**

The time that the import task request was made, presented in the Unix time stamp
format.

Type: Timestamp

Required: No

**importTaskId**

The unique ID for a specific import task. These IDs aren't globally unique, but they are
unique within an AWS account.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `^import-task-[a-fA-F0-9]{32}$`

Required: No

**importUrl**

The URL for your import file that you've uploaded to Amazon S3.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4000.

Pattern: `\S+://\S+/[\s\S]*\S[\s\S]*`

Required: No

**name**

A descriptive name for an import task. You can use this name to filter future requests
related to this import task, such as identifying applications and servers that were included
in this import task. We recommend that you use a meaningful name for each import task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\s\S]*\S[\s\S]*`

Required: No

**serverImportFailure**

The total number of server records in the import file that failed to be imported.

Type: Integer

Required: No

**serverImportSuccess**

The total number of server records in the import file that were successfully
imported.

Type: Integer

Required: No

**status**

The status of the import task. An import can have the status of
`IMPORT_COMPLETE` and still have some records fail to import from the overall
request. More information can be found in the downloadable archive defined in the
`errorsAndFailedEntriesZip` field, or in the Migration Hub management
console.

Type: String

Valid Values: `IMPORT_IN_PROGRESS | IMPORT_COMPLETE | IMPORT_COMPLETE_WITH_ERRORS | IMPORT_FAILED | IMPORT_FAILED_SERVER_LIMIT_EXCEEDED | IMPORT_FAILED_RECORD_LIMIT_EXCEEDED | IMPORT_FAILED_UNSUPPORTED_FILE_TYPE | DELETE_IN_PROGRESS | DELETE_COMPLETE | DELETE_FAILED | DELETE_FAILED_LIMIT_EXCEEDED | INTERNAL_ERROR`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/importtask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/importtask.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/importtask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Filter

ImportTaskFilter
