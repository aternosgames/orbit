# What we need
The network requires a central point of contact for various data from e.g. players, servers, leaderboards, etc. . We want to avoid that every component of the network has direct access to the database. The above mentioned contact point should also be able to do some processing including the data, e.g. calculating the level by XP, so that this does not always have to be a tedious job for the developers. This point of contact for the data should also be easy to maintain, if anything in the processing of the data has to be changed, this should not happen in every component of the network

# Idea
The idea of Orbit is a central contact point to provide data and to handle this via RestAPI and websockets. So the database logic is outsourced and easy to maintain. The database in the background could be RethinkDB. Orbit authorizes the client and grants indirect access to data from the database.

### Authentication
How the authentication works is at the moment not exactly determined. However, the focus is on JWT.