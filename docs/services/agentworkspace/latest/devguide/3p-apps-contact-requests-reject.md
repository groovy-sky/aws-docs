---
title: "Reject the incoming contact for the given contactId in Connect Customer agent workspace"
---

# Reject the incoming contact for the given contactId in Connect Customer agent workspace

Reject the incoming contact for the given contactId. The contact will not be
connected to the agent and will be routed back to the queue.

**Signature**

```

reject(contactId: string): Promise<void>

```

**Usage**

```

await contactClient.reject(contactId);

```

**Input**

**Parameter****Type****Description**contactId _Required_stringThe id of the incoming contact to reject.

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

transfer()

disconnectSelf()

All content copied from https://docs.aws.amazon.com/.
