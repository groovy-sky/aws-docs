# ExecutionTimeFilter

Used to filter the workflow executions in visibility APIs by various time-based rules. Each parameter, if
specified, defines a rule that must be satisfied by each returned query result. The parameter values are in the [Unix Time format](https://en.wikipedia.org/wiki/Unix_time). For example:
`"oldestDate": 1325376070.`

## Contents

**oldestDate**

Specifies the oldest start or close date and time to return.

Type: Timestamp

Required: Yes

**latestDate**

Specifies the latest start or close date and time to return.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/executiontimefilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/executiontimefilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/executiontimefilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainInfo

ExternalWorkflowExecutionCancelRequestedEventAttributes

All content copied from https://docs.aws.amazon.com/.
