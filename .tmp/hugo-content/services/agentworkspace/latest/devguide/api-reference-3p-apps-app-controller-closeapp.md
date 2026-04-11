# Close an application in Amazon Connect Agent Workspace

Closes the application for the given application instance ID in the Amazon Connect agent
workspace.

**Signature**

```typescript

closeApp(instanceId: string): Promise<void>
```

**Usage**

```typescript

await appControllerClient.closeApp(appInstanceId);
```

**Input**

**Parameter****Type****Description**appInstanceId _Required_stringThe instance ID of the application

**Permissions required:**

```typescript

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppController

focusApp()

All content copied from https://docs.aws.amazon.com/.
