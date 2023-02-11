## [MongoDB] Example 2 - Tickers



### `session.Clone()`
In the code, session.Clone() is used to create a new session for each call to storeDataInDB and deleteOldDataFromDB. This is done to ensure that each call has its own independent session, which is necessary for concurrency safety.

MongoDB sessions are not safe to use concurrently, so it's important to create a new session for each operation that needs to be performed concurrently. The Clone method is used to create a new session that shares the same underlying socket connection to the MongoDB server as the original session, but with its own unique session ID and context. This allows multiple concurrent operations to share the same socket connection and reduce the overhead of creating a new socket connection for each operation.

By using session.Clone() instead of just session, the original session remains intact and can be used by other goroutines while the cloned session is used in the current goroutine. This ensures that each call to storeDataInDB and deleteOldDataFromDB uses its own session, avoiding any concurrent use issues and ensuring that each call has a consistent view of the data.