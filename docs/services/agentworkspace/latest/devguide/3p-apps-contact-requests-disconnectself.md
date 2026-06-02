---
title: "Disconnect the agent from the given contact in Connect Customer agent workspace"
---

# Disconnect the agent from the given contact in Connect Customer agent workspace

Disconnects the currently logged-in agent from the given contact. Other
participants on the contact (such as the customer) remain connected.

**Signature**

```

disconnectSelf(contactId: string): Promise<void>

```

**Usage**

```

await contactClient.disconnectSelf(contactId);

```

**Input**

**Parameter****Type****Description**contactId _Required_stringThe id of the contact to disconnect the agent from.

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

reject()

isAutoAcceptEnabled()

All content copied from https://docs.aws.amazon.com/.
