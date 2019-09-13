- [IEEE Web application](#ieee-web-application)
  - [Why](#why)
  - [How to start](#how-to-start)
    - [Docker](#docker)
    - [Manually](#manually)
  - [Environment Variables](#environment-variables)
  - [Planned Features](#planned-features)
    - [APP](#app)
    - [WEB](#web)

# IEEE Web application

A IEEE Student Branch Dashboard & Management tool for student branches of the IEEE Organization.
## Why

Members and the community of IEEE Student Branches will be able to advertise and interact with the community itself in an easier way. This application intends to tie the community in a much tighter form and shape a more robust social construct.

Its more important aspect is the fact that the student branch can advertise the SB's events. This means that the community will be more self-aware, more people will know about the SB's events and advertisment will be easier and more effortless.

**In short: We spread knowledge, we make new friends and we meet people we haven't met before in a much more easy way.**

## How to start

### Docker

Docker setups everything **automatically** without you having to install any additional tools that are required or maintain the local instance. Docker also happens to be multiplatform.

1. Install docker https://docs.docker.com/install/
2. Install docker-compose https://docs.docker.com/compose/install/
3. Run
    ```bash
    # Get composer files
    wget https://raw.githubusercontent.com/ionian-uni-ieee/ieeesb-app/dev/build/docker/docker-compose.yml https://raw.githubusercontent.com/ionian-uni-ieee/ieeesb-app/dev/build/docker/docker-compose.prod.ym
    
    # Deploy swarm
    docker swarm init
    docker stack deploy -c docker-compose.yml -c docker-compose.prod.yml ieeesb-app
    ```

### Manually

In this case, we don't provide further support for other platforms than Linux, and you have to figure out the setup on your own.

1. Install **golang** (>= v1.10)
2. Install **nodejs**,**npm** & **yarn**
3. Install **mongodb**
4. Clone repository
    ```bash
    git clone https://github.com/ionian-uni-ieee/ieeesb-app
    ```
5. For **backend**-server run 
   ```bash
   sudo systemctl start mongod
   go run main.go
   ```
6. For **frontend**-server run 
   ```bash
   cd ./web && yarn && yarn start
   ```

## Environment Variables

**Docker** sets all of them up for you. 

But it's **necessary** to setup a `API_DATABASE_USERNAME` and a `API_DATABASE_PASSWORD`.

| Name                      | Description                    |
| ------------------------- | ------------------------------ |
| API_HOST                  | Server host                    |
| API_PORT                  | Server port                    |
| API_DATABASE_HOST         | Database host                  |
| API_DATABASE_PORT         | Database port                  |
| API_DATABASE_NAME         | Database collection/table name |
| **API_DATABASE_USERNAME** | Database username credential   |
| **API_DATABASE_PASSWORD** | Database password credential   |

## Planned Features
After they're all completed, we will pass to an **Pre-Alpha** phase.

### APP
- [x] Manager Users
  - [x] Delete
  - [x] Edit
  - [x] Get Users

- [x] Authorization
  - [x] Get Profile
  - [x] Login
  - [x] Logout
  - [x] Register

- [ ] Contact Tickets
  - [x] Contact
  - [x] Close Ticket
  - [x] Get Tickets
  - [ ] Manager Respond

- [x] Sponsors
  - [x] Add
  - [x] Delete
  - [x] Edit
  - [x] Get Sponsors

- [ ] Events
  - [ ] Add
  - [ ] Delete
  - [ ] Edit
  - [x] Get Events

- [ ] Blog (small priority)
  - [ ] Add article
  - [ ] Delete article
  - [ ] Edit article
  - [ ] Add comment
  - [ ] Edit comment
  - [ ] Delete comment

### WEB

- [ ] Calendar
  - [ ] Show events

- [ ] Upcoming Events
  - [ ] Show upcoming events

- [ ] Contact form

- [ ] Control panel
    - [ ] Events
      - [ ] Create
      - [ ] Delete
      - [ ] Edit
      - [ ] View
    - [ ] Tickets
      - [ ] Respond
      - [ ] Close
      - [ ] View
    - [ ] Sponsors
      - [ ] Add
      - [ ] Delete
      - [ ] Edit
      - [ ] View
    - [ ] Managers
      - [ ] Add
      - [ ] Delete
      - [ ] Edit
      - [ ] View