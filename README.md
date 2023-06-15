# SPQR-API

SQPR-API is an opensource API for querying information about the ancient Romans.

The API is written in Go and available here: <LINK>

The project is split into two separate applications: the api and the scraper. 
The scraper is a simple script that uses Colly to scrap popular websites for 
information about the romans. The API uses different Go technologies such as Gin 
and GORM to create the server. 

## Usage

**Rulers**
Rome had many rulers between the years of x and y, and not merely emporers. 
Rupublican Rome (between x and y) saw the senate in control, and around 
them the two elected consuls of Rome, who had vitto power over descisions
of law, war and the common people. 

This endpoint sees to try and capture the popular rulers of rome that we know of, 
and serves as a good who's who of Roman politics throughout the years. 

Below give the usage for this endpoint and outlines useful filters and sorting options.

```
spqr-api.com/api/rulers

GET /api/rulers/random   :   Get a random ruler
GET /api/rulers          :   Get all ruler data (with default pagination)

Filters

from    : Beginning Range (default 625 BC)
to      : End Range (default AD 476)
dynasty : Emporer dynasty (default all)
items   : Items per paginated page (default 20)
page    : Page returned after pagination (default 1)

Sorting

???
```