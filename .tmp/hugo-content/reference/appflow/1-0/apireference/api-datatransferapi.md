# DataTransferApi

The API of the connector application that Amazon AppFlow uses to transfer your
data.

## Contents

**Name**

The name of the connector application API.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `[\w/-]+`

Required: No

**Type**

You can specify one of the following types:

AUTOMATIC

The default. Optimizes a flow for datasets that fluctuate in size from small to
large. For each flow run, Amazon AppFlow chooses to use the SYNC or ASYNC API type
based on the amount of data that the run transfers.

SYNC

A synchronous API. This type of API optimizes a flow for small to medium-sized
datasets.

ASYNC

An asynchronous API. This type of API optimizes a flow for large datasets.

Type: String

Valid Values: `SYNC | ASYNC | AUTOMATIC`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/datatransferapi.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/datatransferapi.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/datatransferapi.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatadogSourceProperties

DestinationConnectorProperties

All content copied from https://docs.aws.amazon.com/.
