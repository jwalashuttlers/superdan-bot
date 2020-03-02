## Bot Commands

`/add` => Adds the money to a pool account. Checks if the user is in `ADMIN_GROUPS`. Gives 403 otherwise.
`/book <int>` => Confirms the booking of <n> courts. Deducts the amount from each user's wallet _equally_ for the court booking. 

### Scheduled commands

- **At 9:30AM on Mon-Fri**: Asks for head count for people playing on that day
- **At 2:30PM on Mon-Fri**: Confirms the head count and _suggests_ the number of court bookings.
(A general rule of 3 people/court is applied). Displays the total amount required to book the suggested number. For phase 1 the
booking is manual. Once done `/book` is invoked.
- **At 5:00PM** every alternate day: Checks if any user's wallet balance is negative. Sends a reminder to user to add money.

## Tasks

- [ ] Telegram Bot API token
