# Project

API Design:
- /u/username (GET)
- /login
- /register

Protected routes:
- /u/username/projects (GET)
- /u/username/projects/projectName (GET)
  - Languages (Github ID)
- /u/username/create (POST)
  - ProjectName
  - Github Username
  - Github Repo Name

Tasks:
- /u/username/projects/projectName/tasks/new (POST)
- /u/username/projects/projectName/tasks (GET)

Project Fields:
- ID
- Name
- Description
- Github URL (use helper to extract details from URL)

Task Fields:
- ID
- Type (Feature, Bug, Chore, etc.)
- Name
- Status
- Deadline Date (display "No Date" when no date specified)
