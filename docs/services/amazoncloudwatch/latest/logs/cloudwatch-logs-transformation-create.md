# Create and manage log transformers

A log transformer includes one or more _processors_ that are in a
logical pipeline together. Each processor is applied to a log event, one after the other
in the order that they are listed in the transformer configuration.

Some processors are of the _parser_ type. Each transformer must
have at least one parser, and the first processor in a transformer must be a
parser.

Some of the parsers are built-in parsers that are configured for a certain type of
AWS vended log.

Other processor types are string mutators, JSON mutators, and data processors.

You can create transformers for individual log groups, and you can also create
account-level transformers that apply to all or many log groups in your account. If a
log group has a log group-level transformer, that transformer overrides any
account-level transformer that would otherwise apply to that log group. You can have as
many as 20 account-level transformers in a Region in your account.

You must follow these guidelines when you create a transformer:

- If you include a pre-configured parser for a type of AWS vended logs, it
must be the first processor listed in the transformer. You can include only one
such processor in a transformer.

- You can include only one `grok` processor in a transformer.

- You must have at least one parser-type processor in a transformer. You can
include as many as five parser-type processors. This limit of five includes both
built-in parsers and configurable parsers.

- You can have as many as 20 processors in a transformer.

- You can include only one **addKeys** processor in a
transformer.

- You can include only one **copyValue** processor in a
transformer.

- Each transformer can extract up to 200 fields from a log event.

- Each log event **MUST** be below 512KB. Total size of log
events can still go over 512KB.

###### Topics

- [Create an account-level transformer policy](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Transformer-CreateAccountLevel.html)

- [Edit or delete an account-level transformer policy](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Transformer-EditAccountLevel.html)

- [Create a log-group-level log transformer from scratch](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-CreateNew.html)

- [Create a log-group-level transformer by copying an existing one](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Copy.html)

- [Edit a log-group-level transformer](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Edit.html)

- [Delete a log-group-level transformer](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation-Delete.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Transform logs during ingestion

Create an account-level transformer policy
