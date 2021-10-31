## Initial Thoughts

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