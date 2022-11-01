#  Guide on how to run and use application



### For run application:
```
make run
```

## How to use this project?
After you have launched the application, 
you need to go to the address 
http://localhost:8000. There are 3 routes:

* `/download`
* `/search`
* `create-file`

First of all, you need to follow the route `/download` (http://localhost:8000/download) - it will download
file and then saved to the database.

Next, you can follow the route `/search`, and also add to the corresponding search arguments.
The application can search for transactions by the following parameters:
* `transaction_id:`
  * http://localhost:8000/search?transaction=1
  * http://localhost:8000/search?transaction=2,5,10
* `terminal_id:`
    * http://localhost:8000/search?terminal=3506
    * http://localhost:8000/search?terminal_id=3506,3507
* `status:`
    * http://localhost:8000/search?status=accepted
    * http://localhost:8000/search?status=accepted,declined
* `payment_type:`
    * http://localhost:8000/search?payment=cash
    * http://localhost:8000/search?payment=cash,card
* `from/to date_post:`
  * http://localhost:8000/search?from=2022-08-12&to=2022-08-18
* `part of payment_narrative:`
    * http://localhost:8000/search?narrative=Перерахування


All transactions ordered by `transaction_id`.
___
Next, you can follow the route `/create-file`, (with appropriate parameters like in `/search` route) and data what you 
want search will be saved in file. Links for example:
* http://localhost:8000/create-file?terminal_id=3506,3507
* http://localhost:8000/create-file?transaction=2,5,10
* http://localhost:8000/create-file?from=2022-08-12&to=2022-08-18
