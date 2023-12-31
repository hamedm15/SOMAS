# Self-Organising Multi-Agent System
An Inter-dependent Collective Action Problem

## Overview

Built upon common participation. In SOMAS World, a dynamic electronic environment, multiple fictitious entities termed "agents" venture into a pursuit of survival, energy and point acquisition. These agents confront interdependent collective action challenges amidst varying social relationship scenarios, governed by intricate movement, allocation and interaction strategies. Systems of this nature present interesting implications in computer science and collective intelligence. Independently acting agents equipped with decision-making mechanisms mirror the dynamics of real-world social dilemmas and collective action, offering insight into the capabilities and limitations of such systems. SOMAS World aims to simulate these system’s principles. Through observation of the agents, we can address emerging behaviours, decision-making and adapting social dynamics in a test of survival and resource collection. In reference to self-organizational systems, this project combines the aspects of planning, decision-making, and collaboration to optimize each agent’s overall performance.

## Official Documents
### ['In-Depth Report'](Report.pdf)
- [Rules and Implentation](./docs/Rules%20and%20Implentation.md)
- [SOMAS Base Platform User Manual](https://imperiallondon.sharepoint.com/sites/elec70071-202310/Shared%20Documents/General/basePlatformSOMAS_User_Manual.pdf?CT=1698662908166&OR=ItemsView)
- [Introduction to SOMAS Base Platform](https://imperiallondon.sharepoint.com/sites/elec70071-202310/Shared%20Documents/General/basePlatformSOMAS.pdf?CT=1699098039015&OR=ItemsView)
- [SOMAS ACW Rules](https://imperiallondon.sharepoint.com/sites/elec70071-202310/Class%20Materials/Coursework/SOMAS%20ACW%202023.pdf?CT=1699098083591&OR=ItemsView)

## Running code
See [Setup & Rules](./docs/SETUP.md) for requirements.

```bash
# Approach 1
go run . # Linux and macOS: Use `sudo go run .` if you encounter any "Permission denied" errors.

# Approach 2
go build # build step
./SOMAS2023 # SOMAS2023.exe if you're on Windows. Use `sudo` on Linux and macOS as Approach 1 if required.
```

### Parameters & Help
```bash
go run . --help
```

## Structure

### [`docs`](docs)
Important documents pertaining to codebase organisation, code conventions and project management. Read before writing code.
The rules can be found here [Rules and Implementation](./docs/Rules%20and%20Implementation.md)

### [`internal`](internal)
Internal SOMAS2020 packages. Most development occurs here, including client and server code.

- [`clients`](internal/clients)
Individual team code goes into the respective folders in this directory.

- [`common`](internal/common)
Common utilities, or system-wide code such as game specification etc.

- [`server`](internal/server)
Self-explanatory.

