---
title: "Focus an application in Connect Customer agent workspace"
---

# Focus an application in Connect Customer agent workspace

Brings the application into focus in the Connect Customer agent workspace for the given
application instance ID.

**Signature**

```typescript

focusApp(instanceId: string): Promise<AppFocusResult>
```

**Usage**

```typescript

const applicationFocusResult: AppFocusResult = await appControllerClient.focusApp(appInstanceId);
```

**Input**

**Parameter****Type****Description**appInstanceId _Required_stringThe instance ID of the application

**Output - AppFocusResult**

**Parameter****Type****Description**instanceIdstringThe AmazonResourceName(ARN) of the applicationresult"queued" \| "completed" \| "failed"Indicates if the request is queued, completed or failed

**Permissions required:**

```typescript

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

closeApp()

getApp()

All content copied from https://docs.aws.amazon.com/.
