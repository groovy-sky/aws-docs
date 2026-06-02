---
title: "Get the current network connection status of the agent in Connect Customer agent workspace"
---

# Get the current network connection status of the agent in Connect Customer agent workspace

Returns the current network connection health status of the agent's
connection to Connect Customer backend services.

```typescript

async getNetworkConnectionStatus(): Promise<NetworkConnectionStatusChanged>

```

**Output - NetworkConnectionStatusChanged**

**Parameter****Type****Description**statusNetworkConnectionStatusThe connection health status. One of
`"connected"`, `"connecting"`,
`"disconnected"`, or `"failed"`.timestampnumberEpoch milliseconds when the status was reported.

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

onNextAvailabilityStateChanged()

onNetworkConnectionStatusChanged()

All content copied from https://docs.aws.amazon.com/.
