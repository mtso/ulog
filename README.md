# ulog [![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

[![Build Status](https://travis-ci.org/mtso/ulog.svg?branch=master)](https://travis-ci.org/mtso/ulog)

A URI logger to demo Postgresql's data types. An attempt at rolling my own go/postgresql web service after following the tutorial over at: http://www.alexedwards.net/blog/practical-persistence-sql.

## Resulting API

#### `GET /log`
    
Returns all URI's stored in the database in the format:

    log_id=[serial integer] log_timestamp=[2006-01-02T15:04:05.999999Z]
    log_description=[description text]
    log_uri=[uri string]
    
#### `POST /log?uri=[uri][&description=[text]]`

Saves a URI with a description and timestamp. The description is optional. Parameters must be in valid uri encoding. If successful, returns with a response of: 

    Successfully created "[log_description]...", 1 row(s) affected.

## Backlog

- Add tests and hook them up to the Postgresql service in travis-ci. Ref: https://docs.travis-ci.com/user/database-setup/#PostgreSQL