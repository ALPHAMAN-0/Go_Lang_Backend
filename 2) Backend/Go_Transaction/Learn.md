## Learning Path: Database Transactions in Go with sqlc

If you want to learn this topic step by step, follow this order:

1. Understand the transaction idea.
	- Learn that a transaction is one unit of work.
	- It should either fully succeed or fully fail.

2. Learn the ACID properties.
	- Atomicity: all operations happen together or not at all.
	- Consistency: the database stays valid.
	- Isolation: concurrent operations do not interfere incorrectly.
	- Durability: committed data should not be lost.

3. Learn the basic Go database flow.
	- Open a database connection.
	- Run simple queries.
	- Understand `BEGIN`, `COMMIT`, and `ROLLBACK`.

4. Learn how sqlc generates database code.
	- Write SQL queries in query files.
	- Generate typed Go code from those queries.
	- Use the generated query methods in your application.

5. Learn how to add a transaction layer.
	- Create a `Store` struct to hold query logic.
	- Add a transaction execution helper.
	- Make sure all related database actions share the same transaction.

6. Practice with a transfer example.
	- Create two accounts.
	- Deduct money from one account.
	- Add money to the other account.
	- Record the transfer and entries in the same transaction.

7. Learn how to handle errors safely.
	- Roll back the transaction when any step fails.
	- Commit only after every step succeeds.
	- Return clear errors to the caller.

8. Test concurrency.
	- Use goroutines to run multiple transfers at the same time.
	- Use channels to collect results.
	- Check that balances stay correct under load.

9. Verify edge cases.
	- Test insufficient balance.
	- Test duplicate requests.
	- Test race conditions and deadlocks.

10. Review and improve.
	 - Read the code again after testing.
	 - Simplify repeated logic.
	 - Add comments only where the flow is hard to follow.

## Short Summary

The main learning goal is to understand how Go can use sqlc and database transactions to keep multi-step operations safe, consistent, and testable.
