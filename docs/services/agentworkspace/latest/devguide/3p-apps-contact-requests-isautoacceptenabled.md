---
title: "Check whether auto-accept is enabled for the given contact in Connect Customer agent workspace"
---

# Check whether auto-accept is enabled for the given contact in Connect Customer agent workspace

Returns whether auto-accept is enabled for the given contact. When auto-accept
is enabled, an incoming contact is automatically accepted on the agent's behalf
without requiring an explicit [accept()](3p-apps-contact-requests-accept.md) call.

**Signature**

```

isAutoAcceptEnabled(contactId: string): Promise<boolean>

```

**Usage**

```

const enabled: boolean = await contactClient.isAutoAcceptEnabled(contactId);

```

**Input**

**Parameter****Type****Description**contactId _Required_stringThe id of the contact to check.

**Permissions required:**

```

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

disconnectSelf()

onConnecting()

All content copied from https://docs.aws.amazon.com/.
