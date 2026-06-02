---
title: "expand"
---

# expand

Use `expand` to take a field containing a JSON array and
create a separate log event for each element in the array. All other
fields from the original log event are duplicated in each new event.

**Syntax**

```nohighlight

expand fieldName
```

**Example**

If a log event contains
`items = ["apple","banana","cherry"]` and
`host = "web-01"`, then `expand items` produces
three log events: `{items: "apple", host: "web-01"}`,
`{items: "banana", host: "web-01"}`, and
`{items: "cherry", host: "web-01"}`.

```nohighlight

expand items
| stats count(*) by items, host
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

relevantfields

sort

All content copied from https://docs.aws.amazon.com/.
