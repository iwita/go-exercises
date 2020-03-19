### Initiallize the CLI using Cobra
spf13/cobra package autogenerates code for the cli commands

### Use Bolt as a Database
Bolt is a pure Go key/value store. 
Simple, fast and reliable database for projects that don't require a full database server such as Postgres or MySQL.

Used by Heroku, Shopify

Buckets contain key/value pairs
Key: a byte slice
Value: a byte slice

High Read - Low Write scenario -> Read or Write at a given time. Works well for many reads