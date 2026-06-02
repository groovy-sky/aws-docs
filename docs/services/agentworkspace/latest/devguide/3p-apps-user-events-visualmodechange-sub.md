---
title: "Subscribe a callback function when an Connect Customer agent workspace user's visual mode changes"
---

# Subscribe a callback function when an Connect Customer agent workspace user's visual mode changes

Subscribes a callback function to-be-invoked whenever the user's visual mode
changes in the Connect Customer agent workspace.

**Signature**

```typescript

onVisualModeChange(handler: VisualModeChangedHandler)

```

**Usage**

```typescript

const handler: VisualModeChangedHandler = async (data: VisualModeChanged) => {
    console.log("Visual mode changed! " + data.visualMode);
};

settingsClient.onVisualModeChange(handler);

// VisualModeChanged Structure
{
  visualMode: VisualMode;
  previous: {
    visualMode: VisualMode | null;
  };
}

```

**Permissions required:**

```typescript

*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

setVisualMode()

offVisualModeChange()

All content copied from https://docs.aws.amazon.com/.
