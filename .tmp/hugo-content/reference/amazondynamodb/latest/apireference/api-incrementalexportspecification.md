# IncrementalExportSpecification

Optional object containing the parameters specific to an incremental export.

## Contents

###### Note

In the following list, the required parameters are described first.

**ExportFromTime**

Time in the past which provides the inclusive start range for the export table's data,
counted in seconds from the start of the Unix epoch. The incremental export will reflect
the table's state including and after this point in time.

Type: Timestamp

Required: No

**ExportToTime**

Time in the past which provides the exclusive end range for the export table's data,
counted in seconds from the start of the Unix epoch. The incremental export will reflect
the table's state just prior to this point in time. If this is not provided, the latest
time with data available will be used.

Type: Timestamp

Required: No

**ExportViewType**

The view type that was chosen for the export. Valid values are
`NEW_AND_OLD_IMAGES` and `NEW_IMAGES`. The default value is
`NEW_AND_OLD_IMAGES`.

Type: String

Valid Values: `NEW_IMAGE | NEW_AND_OLD_IMAGES`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/incrementalexportspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/incrementalexportspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/incrementalexportspecification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportTableDescription

InputFormatOptions

All content copied from https://docs.aws.amazon.com/.
