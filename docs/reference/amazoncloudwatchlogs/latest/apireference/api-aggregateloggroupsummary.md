---
title: "AggregateLogGroupSummary"
---

# AggregateLogGroupSummary

Contains an aggregate summary of log groups grouped by data source characteristics,
including the count of log groups and their grouping identifiers.

## Contents

**groupingIdentifiers**

An array of key-value pairs that identify the data source characteristics used to group
the log groups.

The size and content of this array depends on the `groupBy` parameter specified
in the request.

Type: Array of [GroupingIdentifier](api-groupingidentifier.md) objects

Required: No

**logGroupCount**

The number of log groups in this aggregate summary group.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/AggregateLogGroupSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/AggregateLogGroupSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/AggregateLogGroupSummary)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddKeys

Anomaly

All content copied from https://docs.aws.amazon.com/.
