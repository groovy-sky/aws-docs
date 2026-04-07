# DynamoDB throttling resolution guide

This section provides targeted resolution guidance for each specific throttling reason
that DynamoDB may return. Each entry includes suggested resolution approaches based on best
practices and corresponding CloudWatch metrics to monitor.

DynamoDB implements 16 distinct throttling reasons across four main categories. Use the
throttling reasons from your application's exception to navigate directly to the relevant
guidance.

## Key range throughput exceeded (hot partitions)

These throttling reasons occur when individual partitions exceed their throughput
limits, affecting both provisioned and on-demand modes:

- [TableReadKeyRangeThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-key-range-limit-exceeded-mitigation.html#throttling-table-read-keyrange)

- [TableWriteKeyRangeThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-key-range-limit-exceeded-mitigation.html#throttling-table-write-keyrange)

- [IndexReadKeyRangeThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-key-range-limit-exceeded-mitigation.html#throttling-index-read-keyrange)

- [IndexWriteKeyRangeThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-key-range-limit-exceeded-mitigation.html#throttling-index-write-keyrange)

## Provisioned throughput exceeded

These throttling reasons occur when consumption rates exceed provisioned capacity
limits in provisioned mode:

- [TableReadProvisionedThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-provisioned-capacity-exceeded-mitigation.html#throttling-table-read-provisioned)

- [TableWriteProvisionedThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-provisioned-capacity-exceeded-mitigation.html#throttling-table-write-provisioned)

- [IndexReadProvisionedThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-provisioned-capacity-exceeded-mitigation.html#throttling-index-read-provisioned)

- [IndexWriteProvisionedThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-provisioned-capacity-exceeded-mitigation.html#throttling-index-write-provisioned)

## Account limits exceeded

These throttling reasons occur when consumption rates exceed account-level throughput
quotas in your AWS Region:

- [TableReadAccountLimitExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-account-limit-exceeded-mitigation.html#throttling-table-read-account-limit)

- [TableWriteAccountLimitExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-account-limit-exceeded-mitigation.html#throttling-table-write-account-limit)

- [IndexReadAccountLimitExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-account-limit-exceeded-mitigation.html#throttling-index-read-account-limit)

- [IndexWriteAccountLimitExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-account-limit-exceeded-mitigation.html#throttling-index-write-account-limit)

## On-demand maximum throughput exceeded

These throttling reasons occur when consumption rates exceed configured maximum
throughput limits in on-demand mode:

- [TableReadMaxOnDemandThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-ondemand-capacity-exceeded-mitigation.html#throttling-diagnostic-table-read-max-ondemand)

- [TableWriteMaxOnDemandThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-ondemand-capacity-exceeded-mitigation.html#throttling-diagnostic-table-write-max-ondemand)

- [IndexReadMaxOnDemandThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-ondemand-capacity-exceeded-mitigation.html#throttling-diagnostic-index-read-max-ondemand)

- [IndexWriteMaxOnDemandThroughputExceeded](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/throttling-ondemand-capacity-exceeded-mitigation.html#throttling-diagnostic-index-write-max-ondemand)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Diagnosing throttling

1- Key range throughput exceeded (hot partitions)
