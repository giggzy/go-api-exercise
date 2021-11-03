# go API exercise

An implementation an API of the specification the drive a UI for engaging with services of an organization.

Uses a sample JSON file to initialize a global services struct.
When the struct is mutated it is serialized to disk and preferentially read on next server startup achieving persistance.
Only create service is implemented.
Only one organization is supported currently.
Search on service name, sort on service name and serviceCount are supported.

## Questions and snags

- Moved to mux as hit snags on POSTs using build in http lib.
  - Don't understand some basics
- Tried to implement multiple orgs each having a file but started getting messy
- No AuthN or authZ due to time

## TODO

- decompose into modules/files, handlers, util, routes
  - routes is so simple that it seems overkill
  - handlers are getting fat so maybe [X]
- cross cutting concerns are leading to code duplication, DRY
  - missing decorators, but can achieve same with wrapping functions inside functions, worry about readability and community conventions
- Stuck on "merging" structs into a JSON, would like a meta block to help frontend dev and myself in development loop
- CRUD
  - would like to at least implement "update"
    - with file based persistence locking is my responsiblity, see some approaches online but not confident
    - when to persist?
      - on every mutation? on server shutdown? Lean towards every mutation

## Initial Thoughts (Historical)

- Use built-in web server
- Don't use another framework unless it gets messy
- Persistence
  - use file system to start
  - SQLite is an option?
  - Leaning towards ReadOnly
- API Design
  - REST
  - GET will get pretty far
  - CALLS
    - services_list
    - service_detail
      - needs versions
- Search
  - start with simple such as prefix
  - Frontend may want to send ajax calls
  - Only name of of service? Should Description come into play? Could be a future improvement
- Pagination
  - offsets
  - Funny that the designer indexes from 0, first time seeing that in my career.
- Testing
  - Explore, there is a built-in for go
  - Use Insomnia to exercise API as part of development loop
- VCS
  - Local git first, then upload to github
  - should this be private? .. checking for public repos with peoples attempts?
    - yes, some public but not specifically for this takehome

### Optional

- Full CRUD
  - lean to no, without DB (i.e. files) it's annoying, and not pleasant with DB
- Tests
  - lean to yes
- AuthN/AuthZ
  - at least make skeletons

## API Responses

Thinking requests are GETs with path and query string driving.

```json
{
    "services": [
        {
            "url": "https://example.com/a_Service",
            "versionCount": 1,
            "description": "a blah, blah, blah",
            "id": "1",
            "name": "a_Service"
        },

            "url": "https://example.com/b_Service",
            "versionCount": 2,
            "description": "b blah, blah, blah",
            "id": "2",
            "name": "b_Service"
        },
        {
            "url": "https://example.com/c_Service",
            "versionCount": 3,
            "description": "c blah, blah, blah",
            "id": "3",
            "name": "c_Service"
        }
    ]
}
```

## Resources

[The Spec](https://docs.google.com/document/d/1GcqaLwUv2MC7CmXs7ZCrTrfOwkSOiZHRLoWgLzTr9Vc/)
