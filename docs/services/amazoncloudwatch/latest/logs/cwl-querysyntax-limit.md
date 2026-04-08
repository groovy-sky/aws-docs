# limit

Use `limit` to specify the number of log events that you want
your query to return. If you omit `limit`, the query will return
as many as 10,000 log events in the results.

For example, the following example returns only the 25 most recent log
events

```nohighlight

fields @timestamp, @message | sort @timestamp desc | limit 25
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

stats

dedup

All content copied from https://docs.aws.amazon.com/.
