# go API exercise

An implementation an API of the specification the drive a UI for engaging with services of an organization.

Aimed to use a few dependencies outside of the Std Lib. I switched to mux as I was having issues with POSTs and the build in net/http

## Persistence
Using file system with one JSON file per organization for simplicity, seems a reasonable choice for the exercise. Typically would use a relational DB for this. The JSON file approach makes migration to a Document DB the easiest next step.

On startup a sample JSON file is read and converted into a Struct called Services which is an Array of Service Struct. This is stored globally in the app. Mixed feelings around this but no better idea struct.

I introduced an mutation, createService to the datastructure to play with Locking and flush to disk considerations.

On mutation, the Services Array is locked and flushed to disk when mutation is complete. No locking done on reads as *I think* that is safe. My understanding is reads will block will the lock in place so potential optimazation to support read locks.

## API
Went REST as I'm most familiar with that approach and enough new concerns in play.

Used the URL path for required parameter, e.g. id of service to get details on. Used Query Path for optional parameters, search, sort on getting services. Seems OK. I did not use mux's feature which look like it has conviences for building an API.

Did not come up with a good documentation experience for a consumer of the API. I have an Insomnia session which perhaps I can export. I looked into swagger for goland but looked like too much effort.

## Testing
Decided to explore testing with the standard library approaches. For integration tests it  would have been more valuable to have used the same language as the consumer on second though as could act as documentation and be a better resource. For unit tests the go build in testing makes sense. 

Only partial coverage but was useful during refactoring.

## Authentication/Autherization
Did not go down this path due to time. JWT would make sense in my opinion, it could store the orgID and userID which could drive decisions.

Only serving one Organization services requirement was not addressed. so this impl could only work for one org.

I was planning on using a JSON file for each organization services but would take time to sort out the initialization and flushing to disk for that approach.

## Potential Next Steps

- Choosing a DB approach over file persistence.
- AuthN/AuthZ
- Supporting multi org
- Covering rest of CRUD operations.





