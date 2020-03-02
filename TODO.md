## Bot Commands

- `/payin` => Adds the money to a pool account. Checks if the user is in `ADMIN_GROUPS`. Gives 403 otherwise.
- `/deduct amount: <float> description: <string>` => Amount to be deducted. For ex. book court, shuttlecock expense etc.

### Scheduled commands

- At 9:30AM on Mon-Fri: Asks for head count for people playing on that day
- At 2:30PM on Mon-Fri: Confirms the head count and _suggests_ the number of court bookings.
(A general rule of 3 people/court is applied). Displays the total amount required to book the suggested number. For phase 1 the
booking is manual. Once done `/book` is invoked.
- At 5:00PM every alternate day: Checks if any user's wallet balance is negative. Sends a reminder to user to add money.

## Tasks

- [x] Telegram Bot API token
- [ ] Basic framework for bot using [telebot](https://github.com/tucnak/telebot)
- [ ] Write methods for bot commands
- [ ] Use `gocron` for scheduled commands
- [ ] Config parsing
