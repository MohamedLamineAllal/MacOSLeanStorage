# Project description

We want to build a CLI tool and Daemon using Go Lang (for better memory usage). That clean safe to remove files and directories that consume storage. That are taken by many applications. Notably Arc, google chrome, and many browsers, electron applications, Discord, vscode, cursor, antigravity, OpenAIAtlas ...

## Instructions

### Agentic Workflow

- **Resumability**: At the start of every interaction, the agent MUST read `MEMORY.md` to understand the current phase, progress, and context.
- **Prompt Logging**: The agent MUST log every significant user prompt into `Prompts.log` to maintain a history of directives and intent.
- **Contextual Awareness**: The agent should always reference `MEMORY.md` before proposing new actions to ensure continuity.
- **Incremental Updates**: Every significant step or decision must be recorded in `MEMORY.md`.
- **Response Documentation**: When creating a substantial response, analysis, guide, recommendation, or decision-support note, write it as a Markdown file under `docs` using a clear well named subfolder structure and organization. Keep response files organized by topic, use descriptive filenames, and reference the created file in the final chat response.
- **Git Integration**: All significant changes should be committed with descriptive messages following conventional commits if possible.

### Memory Logic

- `MEMORY.md` serves as the source of truth for the project's state.
- It includes:
    - Current Phase.
    - Summary of completed tasks.
    - Pending tasks for the current phase.
    - Brainstormed items, analysis results, and decisions (as they are generated).

### Execution and context
- Make sure you tackle deeply and efficiently all of the request of the prompt, take your time, do things in multiple steps, as much as it's required. Stop when needed and resume till you finish everything.