# ImportTaskFilter

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

A name-values pair of elements you can use to filter the results when querying your import
tasks. Currently, wildcards are not supported for filters.

###### Note

When filtering by import status, all other filter values are ignored.

## Contents

**name**

The name, status, or import task ID for a specific import task.

Type: String

Valid Values: `IMPORT_TASK_ID | STATUS | NAME | FILE_CLASSIFICATION`

Required: No

**values**

An array of strings that you can provide to match against a specific name, status, or
import task ID to filter the results for your import task queries.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/importtaskfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/importtaskfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/importtaskfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportTask

NeighborConnectionDetail

All content copied from https://docs.aws.amazon.com/.
