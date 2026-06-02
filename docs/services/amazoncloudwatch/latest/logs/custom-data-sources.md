---
title: "Custom sources"
---

# Custom sources

For custom application logs, you can define your own data source categorization
using one of the following methods:

- **Log group tags** – Add tags to your log
groups using the keys `cw:datasource:name` and
`cw:datasource:type` to specify the data source name and type
for all logs ingested in the log group.

- **Pipeline configuration** – Configure data
source information through log processing pipelines when ingesting your
application logs.

###### Note

Custom data source names cannot start with "aws" or "amazon" to avoid
conflicts with AWS service logs.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Supported third-party sources for data sources

Analyzing log data with CloudWatch Logs Insights

All content copied from https://docs.aws.amazon.com/.
